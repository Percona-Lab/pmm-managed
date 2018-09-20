// pmm-managed
// Copyright (C) 2018 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package logs

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	servicelib "github.com/percona/kardianos-service"

	"github.com/percona/pmm-managed/utils/logger"
	"gopkg.in/yaml.v2"
)

// File represents log file content.
type File struct {
	Name string
	Data []byte
	Err  error
}

type Log struct {
	FilePath string
	UnitName string
	Command  string
}

var DefaultLogs = []Log{
	{"/var/log/consul.log", "consul", ""},
	{"/var/log/createdb.log", "", ""},
	{"/var/log/cron.log", "crond", ""},
	{"/var/log/dashboard-upgrade.log", "", ""},
	{"/var/log/grafana/grafana.log", "", ""},
	{"/var/log/mysql.log", "", ""},
	{"/var/log/mysqld.log", "mysqld", ""},
	{"/var/log/nginx.log", "nginx", ""},
	{"/var/log/nginx/access.log", "", ""},
	{"/var/log/nginx/error.log", "", ""},
	{"/var/log/node_exporter.log", "node_exporter", ""},
	{"/var/log/orchestrator.log", "orchestrator", ""},
	{"/var/log/pmm-manage.log", "pmm-manage", ""},
	{"/var/log/pmm-managed.log", "pmm-managed", ""},
	{"/var/log/prometheus1.log", "prometheus1", ""},
	{"/var/log/prometheus.log", "prometheus", ""},
	{"/var/log/qan-api.log", "percona-qan-api", ""},
	{"/var/log/supervisor/supervisord.log", "", ""},
	{"/etc/prometheus.yml", "expand", "/usr/bin/cat"},
	{"/etc/supervisord.d/pmm.ini", "expand", "/usr/bin/cat"},
	{"/etc/nginx/conf.d/pmm.conf", "expand", "/usr/bin/cat"},
	{"prometheus_targets.txt", "expand", "/usr/bin/curl -s http://localhost/prometheus/targets"},
	{"consul_nodes.txt", "expand", "/usr/bin/curl -s http://localhost/v1/internal/ui/nodes?dc=dc1"},
	{"qan-api_instances.txt", "expand", "/usr/bin/curl -s http://localhost/qan-api/instances"},
	{"managed_RDS-Aurora.txt", "expand", "/usr/bin/curl -s http://localhost/managed/RDS"},
	{"pmm-version.txt", "expand", "head -1 /srv/update/main.yml"},
	{"supervisorctl_status.log", "expand", "supervisorctl status"},
	{"systemctl_status.log", "expand", "systemctl -l status"},
	{"pt-summary.log", "expand", "pt-summary"},
}

// Logs is responsible for interactions with logs.
type Logs struct {
	n              int
	logs           []Log
	journalctlPath string
	ctx            context.Context
}

type manageConfig struct {
	Users []struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"users"`
}

// Fetch PMM credential
func getCredential(ctx context.Context) string {
	u := ""
	f, err := os.Open("/srv/update/pmm-manage.yml")
	if err != nil {
		return u
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return u
	}

	var config manageConfig
	if err = yaml.Unmarshal(b, &config); err != nil {
		return u
	}
	if len(config.Users) > 0 && config.Users[0].Username != "" {
		u = (strings.Join([]string{config.Users[0].Username, config.Users[0].Password}, ":"))
	}
	return u
}

// New creates a new Logs service.
// n is a number of last lines of log to read.
func New(ctx context.Context, logs []Log, n int) *Logs {
	l := &Logs{
		n:    n,
		logs: logs,
		ctx:  ctx,
	}

	// PMM Server Docker image contails journalctl, so we can't use exec.LookPath("journalctl") alone for detection.
	// TODO Probably, that check should be moved to supervisor service.
	//      Or the whole logs service should be merged with it.
	if servicelib.Platform() == "linux-systemd" {
		l.journalctlPath, _ = exec.LookPath("journalctl")
	}

	return l
}

// Zip creates .zip archive with all logs.
func (l *Logs) Zip(ctx context.Context, w io.Writer) error {
	zw := zip.NewWriter(w)

	now := time.Now().UTC()
	for _, log := range l.logs {
		name, content, err := l.readLog(ctx, &log)
		if name == "" {
			continue
		}

		if err != nil {
			logger.Get(l.ctx).WithField("component", "logs").Error(err)

			// do not let a single error break the whole archive
			if len(content) > 0 {
				content = append(content, "\n\n"...)
			}
			content = append(content, []byte(err.Error())...)
		}

		f, err := zw.CreateHeader(&zip.FileHeader{
			Name:     name,
			Method:   zip.Deflate,
			Modified: now,
		})
		if err != nil {
			return err
		}
		if _, err = f.Write(content); err != nil {
			return err
		}
	}

	// make sure to check the error on Close
	return zw.Close()
}

// Files returns list of logs and their content.
func (l *Logs) Files(ctx context.Context) []File {
	files := make([]File, len(l.logs))

	for i, log := range l.logs {
		var file File
		file.Name, file.Data, file.Err = l.readLog(ctx, &log)
		files[i] = file
	}

	return files
}

// readLog reads last l.n lines from defined Log configuration.
func (l *Logs) readLog(ctx context.Context, log *Log) (name string, data []byte, err error) {
	if log.UnitName == "expand" {
		s := strings.Split(log.Command, "//")
		credential := getCredential(ctx)
		if len(s) > 1 && len(credential) > 1 {
			log.Command = fmt.Sprintf("%s//%s@%s", s[0], credential, s[1])
		}
		name = filepath.Base(log.FilePath)
		data, err = l.collectExec(ctx, log.FilePath, log.Command)
		return
	}

	if log.UnitName != "" && l.journalctlPath != "" {
		name = log.UnitName
		data, err = l.journalctlN(ctx, log.UnitName)
		return
	}

	if log.FilePath != "" {
		name = filepath.Base(log.FilePath)
		data, err = l.tailN(ctx, log.FilePath)
		return
	}

	return
}

// journalctlN reads last l.n lines from systemd unit u using `journalctl` command.
func (l *Logs) journalctlN(ctx context.Context, u string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, l.journalctlPath, "-n", strconv.Itoa(l.n), "-u", u)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	b, err := cmd.Output()
	if err != nil {
		return b, fmt.Errorf("%s: %s: %s", strings.Join(cmd.Args, " "), err, stderr.String())
	}
	return b, nil
}

// tailN reads last l.n lines from log file at given path using `tail` command.
func (l *Logs) tailN(ctx context.Context, path string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "/usr/bin/tail", "-n", strconv.Itoa(l.n), path)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	b, err := cmd.Output()
	if err != nil {
		return b, fmt.Errorf("%s: %s: %s", strings.Join(cmd.Args, " "), err, stderr.String())
	}
	return b, nil
}

// collect output from various commands
func (l *Logs) collectExec(ctx context.Context, path string, command string) ([]byte, error) {
	cmd := &exec.Cmd{}
	if filepath.Dir(path) != "." {
		cmd = exec.CommandContext(ctx, command, path)
	} else {
		command := strings.Split(command, " ")
		cmd = exec.CommandContext(ctx, command[0], command[1:]...)
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	b, err := cmd.Output()
	if err != nil {
		return b, fmt.Errorf("%s: %s: %s", strings.Join(cmd.Args, " "), err, stderr.String())
	}
	return b, nil
}
