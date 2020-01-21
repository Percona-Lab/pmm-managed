// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package server implements pmm-managed Server API.
package server

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/percona/pmm/api/serverpb"
	"github.com/percona/pmm/utils/pdeathsig"
	"github.com/percona/pmm/version"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

// Server represents service for checking PMM Server status and changing settings.
type Server struct {
	db               *reform.DB
	prometheus       prometheusService
	supervisord      supervisordService
	telemetryService telemetryService
	checker          *AWSInstanceChecker
	l                *logrus.Entry

	pmmUpdateAuthFileM sync.Mutex
	pmmUpdateAuthFile  string

	envRW                  sync.RWMutex
	envDisableUpdates      bool
	envDisableTelemetry    bool
	envMetricsResolutionHR time.Duration
	envMetricsResolutionMR time.Duration
	envMetricsResolutionLR time.Duration
	envDataRetention       time.Duration

	sshKeyM sync.Mutex

	// To make this testeable. To run API tests we need to write the rules file but on dev envs
	// there is no /srv/prometheus/rules/ directory
	alertManagerFile string
}

type pmmUpdateAuth struct {
	AuthToken string `json:"auth_token"`
}

// NewServer returns new server for Server service.
func NewServer(db *reform.DB, prometheus prometheusService, supervisord supervisordService,
	telemetryService telemetryService, checker *AWSInstanceChecker, alertManagerFile string) (*Server, error) {
	path := os.TempDir()
	if _, err := os.Stat(path); err != nil {
		return nil, errors.WithStack(err)
	}
	path = filepath.Join(path, "pmm-update.json")

	s := &Server{
		db:                db,
		prometheus:        prometheus,
		supervisord:       supervisord,
		telemetryService:  telemetryService,
		checker:           checker,
		l:                 logrus.WithField("component", "server"),
		pmmUpdateAuthFile: path,
		alertManagerFile:  alertManagerFile,
	}
	return s, nil
}

// UpdateSettingsFromEnv updates settings in the database with environment variables values.
// It returns only validation or database errors; invalid environment variables are logged and skipped.
func (s *Server) UpdateSettingsFromEnv(env []string) error {
	s.envRW.Lock()
	defer s.envRW.Unlock()

	return s.db.InTransaction(func(tx *reform.TX) error {
		settings, err := models.GetSettings(tx.Querier)
		if err != nil {
			return err
		}

		for _, e := range env {
			p := strings.SplitN(e, "=", 2)
			if len(p) != 2 {
				s.l.Warnf("Failed to parse environment variable %q.", e)
				continue
			}

			k, v := strings.ToUpper(p[0]), strings.ToLower(p[1])
			var err error
			var value interface{}
			switch k {
			case "DISABLE_UPDATES":
				var b bool
				b, err = strconv.ParseBool(v)
				if err == nil {
					value = b
					s.envDisableUpdates = b
				}

			case "DISABLE_TELEMETRY":
				var b bool
				b, err = strconv.ParseBool(v)
				if err == nil {
					value = b
					s.envDisableTelemetry = b
					settings.Telemetry.Disabled = b
				}

			case "METRICS_RESOLUTION", "METRICS_RESOLUTION_HR":
				var d time.Duration
				d, err = time.ParseDuration(v)
				if err == nil {
					value = d
					s.envMetricsResolutionHR = d
					settings.MetricsResolutions.HR = d
				}

			case "METRICS_RESOLUTION_MR":
				var d time.Duration
				d, err = time.ParseDuration(v)
				if err == nil {
					value = d
					s.envMetricsResolutionMR = d
					settings.MetricsResolutions.MR = d
				}

			case "METRICS_RESOLUTION_LR":
				var d time.Duration
				d, err = time.ParseDuration(v)
				if err == nil {
					value = d
					s.envMetricsResolutionLR = d
					settings.MetricsResolutions.LR = d
				}

			case "DATA_RETENTION":
				var d time.Duration
				d, err = time.ParseDuration(v)
				if err == nil {
					value = d
					s.envDataRetention = d
					settings.DataRetention = d
				}

			default:
				s.l.Tracef("Skipping %q.", e)
				continue
			}

			if err == nil {
				s.l.Infof("Environment variable %q parsed: %v.", e, value)
			} else {
				s.l.Warnf("Failed to parse environment variable %q: %s.", e, err)
			}
		}

		return models.SaveSettings(tx.Querier, settings)
	})
}

// Version returns PMM Server version.
func (s *Server) Version(ctx context.Context, req *serverpb.VersionRequest) (*serverpb.VersionResponse, error) {
	// for API testing of authentication, panic handling, etc.
	if req.Dummy != "" {
		switch {
		case strings.HasPrefix(req.Dummy, "panic-"):
			switch req.Dummy {
			case "panic-error":
				panic(errors.New("panic-error"))
			case "panic-fmterror":
				panic(fmt.Errorf("panic-fmterror"))
			default:
				panic(req.Dummy)
			}

		case strings.HasPrefix(req.Dummy, "grpccode-"):
			code, err := strconv.Atoi(strings.TrimPrefix(req.Dummy, "grpccode-"))
			if err != nil {
				return nil, err
			}
			grpcCode := codes.Code(code)
			return nil, status.Errorf(grpcCode, "gRPC code %d (%s)", grpcCode, grpcCode)
		}
	}

	res := &serverpb.VersionResponse{
		// always return something in this field:
		// it is used by PMM 1.x's pmm-client for compatibility checking
		Version: version.Version,

		Managed: &serverpb.VersionInfo{
			Version:     version.Version,
			FullVersion: version.FullCommit,
		},

		DistributionMethod: s.telemetryService.DistributionMethod(),
	}
	if t, err := version.Time(); err == nil {
		ts, _ := ptypes.TimestampProto(t)
		res.Managed.Timestamp = ts
	}

	if v := s.supervisord.InstalledPMMVersion(); v != nil {
		res.Version = v.Version
		res.Server = &serverpb.VersionInfo{
			Version:     v.Version,
			FullVersion: v.FullVersion,
		}
		if v.BuildTime != nil {
			res.Server.Timestamp, _ = ptypes.TimestampProto(*v.BuildTime)
		}
	}

	return res, nil
}

// Readiness returns an error when some PMM Server component is not ready yet or is being restarted.
// It can be used as for Docker health check or Kubernetes readiness probe.
func (s *Server) Readiness(ctx context.Context, req *serverpb.ReadinessRequest) (*serverpb.ReadinessResponse, error) {
	// TODO https://jira.percona.com/browse/PMM-1962

	if err := s.prometheus.Check(ctx); err != nil {
		return nil, err
	}
	return &serverpb.ReadinessResponse{}, nil
}

// CheckUpdates checks PMM Server updates availability.
func (s *Server) CheckUpdates(ctx context.Context, req *serverpb.CheckUpdatesRequest) (*serverpb.CheckUpdatesResponse, error) {
	if req.Force {
		if err := s.supervisord.ForceCheckUpdates(); err != nil {
			return nil, err
		}
	}

	v, lastCheck := s.supervisord.LastCheckUpdatesResult()
	if v == nil {
		return nil, status.Error(codes.Unavailable, "failed to check for updates")
	}

	res := &serverpb.CheckUpdatesResponse{
		Installed: &serverpb.VersionInfo{
			Version:     v.Installed.Version,
			FullVersion: v.Installed.FullVersion,
		},
		Latest: &serverpb.VersionInfo{
			Version:     v.Latest.Version,
			FullVersion: v.Latest.FullVersion,
		},
		UpdateAvailable: v.UpdateAvailable,
		LatestNewsUrl:   v.LatestNewsURL,
	}
	res.LastCheck, _ = ptypes.TimestampProto(lastCheck)
	if v.Installed.BuildTime != nil {
		t := v.Installed.BuildTime.UTC().Truncate(24 * time.Hour) // return only date
		res.Installed.Timestamp, _ = ptypes.TimestampProto(t)
	}
	if v.Latest.BuildTime != nil {
		t := v.Latest.BuildTime.UTC().Truncate(24 * time.Hour) // return only date
		res.Latest.Timestamp, _ = ptypes.TimestampProto(t)
	}
	return res, nil
}

// StartUpdate starts PMM Server update.
func (s *Server) StartUpdate(ctx context.Context, req *serverpb.StartUpdateRequest) (*serverpb.StartUpdateResponse, error) {
	s.envRW.RLock()
	updatesDisabled := s.envDisableUpdates
	s.envRW.RUnlock()

	if updatesDisabled {
		return nil, status.Error(codes.FailedPrecondition, "Updates are disabled via DISABLE_UPDATES environment variable.")
	}

	offset, err := s.supervisord.StartUpdate()
	if err != nil {
		return nil, err
	}

	authToken := uuid.New().String()
	if err = s.writeUpdateAuthToken(authToken); err != nil {
		return nil, err
	}

	return &serverpb.StartUpdateResponse{
		AuthToken: authToken,
		LogOffset: offset,
	}, nil
}

// UpdateStatus returns PMM Server update status.
func (s *Server) UpdateStatus(ctx context.Context, req *serverpb.UpdateStatusRequest) (*serverpb.UpdateStatusResponse, error) {
	token, err := s.readUpdateAuthToken()
	if err != nil {
		return nil, err
	}
	if subtle.ConstantTimeCompare([]byte(req.AuthToken), []byte(token)) == 0 {
		return nil, status.Error(codes.PermissionDenied, "Invalid authentication token.")
	}

	// wait up to 30 seconds for new log lines
	var lines []string
	var newOffset uint32
	var done bool
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for ctx.Err() == nil {
		done = !s.supervisord.UpdateRunning()
		if done {
			// give supervisord a second to flush logs to file
			time.Sleep(time.Second)
		}

		lines, newOffset, err = s.supervisord.UpdateLog(req.LogOffset)
		if err != nil {
			s.l.Warn(err)
		}

		if len(lines) != 0 || done {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}

	return &serverpb.UpdateStatusResponse{
		LogLines:  lines,
		LogOffset: newOffset,
		Done:      done,
	}, nil
}

// writeUpdateAuthToken writes authentication token for getting update status and logs to the file.
//
// We can't rely on Grafana for authentication or on PostgreSQL for storage as their configuration
// is being changed during update.
func (s *Server) writeUpdateAuthToken(token string) error {
	s.pmmUpdateAuthFileM.Lock()
	defer s.pmmUpdateAuthFileM.Unlock()

	a := &pmmUpdateAuth{
		AuthToken: token,
	}
	f, err := os.OpenFile(s.pmmUpdateAuthFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600|os.ModeExclusive)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			s.l.Error(err)
		}
	}()

	return errors.WithStack(json.NewEncoder(f).Encode(a))
}

// readUpdateAuthToken reads authentication token for getting update status and logs from the file.
func (s *Server) readUpdateAuthToken() (string, error) {
	s.pmmUpdateAuthFileM.Lock()
	defer s.pmmUpdateAuthFileM.Unlock()

	f, err := os.OpenFile(s.pmmUpdateAuthFile, os.O_RDONLY, os.ModeExclusive)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			s.l.Error(err)
		}
	}()

	var a pmmUpdateAuth
	err = json.NewDecoder(f).Decode(&a)
	return a.AuthToken, errors.WithStack(err)
}

func convertSettings(s *models.Settings) *serverpb.Settings {
	return &serverpb.Settings{
		TelemetryEnabled: !s.Telemetry.Disabled,
		MetricsResolutions: &serverpb.MetricsResolutions{
			Hr: ptypes.DurationProto(s.MetricsResolutions.HR),
			Mr: ptypes.DurationProto(s.MetricsResolutions.MR),
			Lr: ptypes.DurationProto(s.MetricsResolutions.LR),
		},
		DataRetention:       ptypes.DurationProto(s.DataRetention),
		AwsPartitions:       s.AWSPartitions,
		SshKey:              s.SSHKey,
		AlertManagerAddress: s.AlertManagerAddress,
	}
}

// GetSettings returns current PMM Server settings.
func (s *Server) GetSettings(ctx context.Context, req *serverpb.GetSettingsRequest) (*serverpb.GetSettingsResponse, error) {
	s.envRW.RLock()
	defer s.envRW.RUnlock()

	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}
	res := &serverpb.GetSettingsResponse{
		Settings: convertSettings(settings),
	}
	res.Settings.UpdatesDisabled = s.envDisableUpdates
	alertManagerRules, err := loadAlertManagerRules(s.alertManagerFile)
	if err != nil {
		s.l.Warnf("cannot load alert manager rules: %s", err)
	}
	res.Settings.AlertManagerRules = string(alertManagerRules)
	return res, nil
}

// ChangeSettings changes PMM Server settings.
func (s *Server) ChangeSettings(ctx context.Context, req *serverpb.ChangeSettingsRequest) (*serverpb.ChangeSettingsResponse, error) {
	s.envRW.RLock()
	defer s.envRW.RUnlock()

	if req.EnableTelemetry && req.DisableTelemetry {
		return nil, status.Error(codes.InvalidArgument, "Both enable_telemetry and disable_telemetry are present.")
	}

	var settings *models.Settings
	err := s.db.InTransaction(func(tx *reform.TX) error {
		var e error
		if settings, e = models.GetSettings(tx); e != nil {
			return e
		}

		if s.envDisableTelemetry && (req.EnableTelemetry || req.DisableTelemetry) {
			return status.Error(codes.FailedPrecondition, "Telemetry is disabled via DISABLE_TELEMETRY environment variable.")
		}
		if req.EnableTelemetry {
			settings.Telemetry.Disabled = false
		}
		if req.DisableTelemetry {
			settings.Telemetry.Disabled = true
		}

		// absent or zero value means "do not change"
		res := req.GetMetricsResolutions()
		if hr, e := ptypes.Duration(res.GetHr()); e == nil && hr != 0 {
			if s.envMetricsResolutionHR != 0 {
				return status.Error(codes.FailedPrecondition,
					"High resolution for metrics is set via METRICS_RESOLUTION_HR (or METRICS_RESOLUTION) environment variable.")
			}
			settings.MetricsResolutions.HR = hr
		}
		if mr, e := ptypes.Duration(res.GetMr()); e == nil && mr != 0 {
			if s.envMetricsResolutionMR != 0 {
				return status.Error(codes.FailedPrecondition,
					"Medium resolution for metrics is set via METRICS_RESOLUTION_MR environment variable.")
			}
			settings.MetricsResolutions.MR = mr
		}
		if lr, e := ptypes.Duration(res.GetLr()); e == nil && lr != 0 {
			if s.envMetricsResolutionLR != 0 {
				return status.Error(codes.FailedPrecondition,
					"Low resolution for metrics is set via METRICS_RESOLUTION_LR environment variable.")
			}
			settings.MetricsResolutions.LR = lr
		}

		// absent or zero value means "do not change"
		if dr, e := ptypes.Duration(req.GetDataRetention()); e == nil && dr != 0 {
			if s.envDataRetention != 0 {
				return status.Error(codes.FailedPrecondition, "Data retention for queries is set via DATA_RETENTION environment variable.")
			}
			settings.DataRetention = dr
		}

		// absent or zero value means "do not change"
		if p := req.GetAwsPartitions(); len(p) > 0 {
			settings.AWSPartitions = p
		}

		if req.GetSshKey() != "" {
			if e = s.validateSSHKey(ctx, req.SshKey); e != nil {
				return e
			}

			if e = s.writeSSHKey(req.SshKey); e != nil {
				return e
			}

			settings.SSHKey = req.SshKey
		}

		if req.AlertManagerRules != "" && !req.RemoveAlertManagerRules {
			if e := validateAlertManagerRulesFile(req.AlertManagerRules); e != nil {
				return e
			}
			if e := ioutil.WriteFile(s.alertManagerFile, []byte(req.AlertManagerRules), os.ModePerm); e != nil {
				return e
			}
			if e := os.Chmod(s.alertManagerFile, 0644); e != nil {
				s.l.Warn(e)
			}
		}

		if req.RemoveAlertManagerRules {
			if e := os.Remove(s.alertManagerFile); e != nil {
				s.l.Warn(e)
			}
		}

		if req.AlertManagerAddress != "" && !req.RemoveAlertManagerAddress {
			settings.AlertManagerAddress = req.AlertManagerAddress
		}

		if req.RemoveAlertManagerAddress {
			settings.AlertManagerAddress = ""
		}

		return models.SaveSettings(tx, settings)
	})
	if err != nil {
		return nil, err
	}

	err = s.supervisord.UpdateConfiguration(settings)
	s.prometheus.RequestConfigurationUpdate()
	if err != nil {
		return nil, err
	}

	res := &serverpb.ChangeSettingsResponse{
		Settings: convertSettings(settings),
	}
	res.Settings.UpdatesDisabled = s.envDisableUpdates

	alertManagerRules, err := loadAlertManagerRules(s.alertManagerFile)
	if err != nil {
		return nil, err
	}

	res.Settings.AlertManagerRules = string(alertManagerRules)
	return res, nil
}

func loadAlertManagerRules(filename string) ([]byte, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, nil
	}
	return ioutil.ReadFile(filename)
}

func validateAlertManagerRulesFile(rules string) error {
	tmpfile, err := ioutil.TempFile("", "alter_mgr_test_rules.*.yml")
	if err != nil {
		return err
	}

	defer os.Remove(tmpfile.Name()) //nolint:errcheck

	if _, err := tmpfile.WriteString(rules); err != nil {
		return err
	}

	if err := tmpfile.Close(); err != nil {
		return err
	}

	return exec.Command("promtool", "check", "rules", tmpfile.Name()).Run()
}

func (s *Server) validateSSHKey(ctx context.Context, sshKey string) error {
	tempFile, err := ioutil.TempFile("", "temp_keys_*")
	if err != nil {
		return errors.WithStack(err)
	}
	tempFile.Close()                 //nolint:errcheck
	defer os.Remove(tempFile.Name()) //nolint:errcheck

	if err = ioutil.WriteFile(tempFile.Name(), []byte(sshKey), os.FileMode(0600)); err != nil {
		return errors.WithStack(err)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(timeoutCtx, "ssh-keygen", "-l", "-f", tempFile.Name()) //nolint:gosec
	pdeathsig.Set(cmd, unix.SIGKILL)

	if err = cmd.Run(); err != nil {
		if e, ok := err.(*exec.ExitError); ok && e.ExitCode() != 0 {
			return status.Errorf(codes.InvalidArgument, "Invalid ssh key")
		}
		return errors.WithStack(err)
	}

	return nil
}

func (s *Server) writeSSHKey(sshKey string) error {
	s.sshKeyM.Lock()
	defer s.sshKeyM.Unlock()

	const username = "admin"
	usr, err := user.Lookup(username)
	if err != nil {
		return errors.WithStack(err)
	}
	sshDirPath := path.Join(usr.HomeDir, ".ssh")
	if err = os.MkdirAll(sshDirPath, os.FileMode(0700)); err != nil {
		return errors.WithStack(err)
	}

	uid, err := strconv.Atoi(usr.Uid)
	if err != nil {
		return errors.WithStack(err)
	}
	gid, err := strconv.Atoi(usr.Gid)
	if err != nil {
		return errors.WithStack(err)
	}
	if err = os.Chown(sshDirPath, uid, gid); err != nil {
		return errors.WithStack(err)
	}
	keysPath := path.Join(sshDirPath, "authorized_keys")
	if err = ioutil.WriteFile(keysPath, []byte(sshKey), os.FileMode(0600)); err != nil {
		return errors.WithStack(err)
	}
	if err = os.Chown(keysPath, uid, gid); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// AWSInstanceCheck checks AWS EC2 instance ID.
func (s *Server) AWSInstanceCheck(ctx context.Context, req *serverpb.AWSInstanceCheckRequest) (*serverpb.AWSInstanceCheckResponse, error) {
	if err := s.checker.check(req.InstanceId); err != nil {
		return nil, err
	}
	return &serverpb.AWSInstanceCheckResponse{}, nil
}

// check interfaces
var (
	_ serverpb.ServerServer = (*Server)(nil)
)
