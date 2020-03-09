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

// Package prometheus contains business logic of working with Prometheus.
package prometheus

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"reflect"
	"regexp"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/utils/pdeathsig"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
	"gopkg.in/yaml.v2"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services/prometheus/internal/config"
)

const updateBatchDelay = 3 * time.Second

var checkFailedRE = regexp.MustCompile(`FAILED: parsing YAML file \S+: (.+)\n`)

// Service is responsible for interactions with Prometheus.
// It assumes the following:
//   * Prometheus APIs (including lifecycle) are accessible;
//   * Prometheus configuration and rule files are accessible;
//   * promtool is available.
type Service struct {
	configPath     string
	baseConfigPath string
	promtoolPath   string
	db             *reform.DB
	baseURL        *url.URL
	client         *http.Client

	l    *logrus.Entry
	sema chan struct{}
}

// NewService creates new service.
func NewService(configPath, baseConfigPath, promtoolPath string, db *reform.DB, baseURL string) (*Service, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Service{
		configPath:     configPath,
		baseConfigPath: baseConfigPath,
		promtoolPath:   promtoolPath,
		db:             db,
		baseURL:        u,
		client:         new(http.Client),
		l:              logrus.WithField("component", "prometheus"),
		sema:           make(chan struct{}, 1),
	}, nil
}

// Run runs Prometheus configuration update loop until ctx is canceled.
func (svc *Service) Run(ctx context.Context) {
	svc.l.Info("Starting...")
	defer svc.l.Info("Done.")

	for {
		select {
		case <-ctx.Done():
			return

		case <-svc.sema:
			// batch several update requests together by delaying the first one
			sleepCtx, sleepCancel := context.WithTimeout(ctx, updateBatchDelay)
			<-sleepCtx.Done()
			sleepCancel()

			if ctx.Err() != nil {
				return
			}

			if err := svc.updateConfiguration(); err != nil {
				svc.l.Errorf("Failed to update configuration, will retry: %+v.", err)
				svc.RequestConfigurationUpdate()
			}
		}
	}
}

// reload asks Prometheus to reload configuration.
func (svc *Service) reload() error {
	u := *svc.baseURL
	u.Path = path.Join(u.Path, "-", "reload")
	resp, err := svc.client.Post(u.String(), "", nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.Errorf("%d: %s", resp.StatusCode, b)
}

func (svc *Service) loadBaseConfig() *config.Config {
	var cfg config.Config

	buf, err := ioutil.ReadFile(svc.baseConfigPath)
	if err != nil {
		if !os.IsNotExist(err) {
			svc.l.Errorf("Failed to load base prometheus config %s: %s", svc.baseConfigPath, err)
		}
		return &cfg
	}

	if err := yaml.Unmarshal(buf, &cfg); err != nil {
		svc.l.Errorf("Failed to parse base prometheus config %s: %s.", svc.baseConfigPath, err)
		return &config.Config{}
	}

	return &cfg
}

// addScrapeConfigs adds Prometheus scrape configs to cfg for all Agents.
func (svc *Service) addScrapeConfigs(cfg *config.Config, q *reform.Querier, s *models.MetricsResolutions) error {
	agents, err := q.SelectAllFrom(models.AgentTable, "WHERE NOT disabled AND listen_port IS NOT NULL ORDER BY agent_type, agent_id")
	if err != nil {
		return errors.WithStack(err)
	}

	var rdsParams []*scrapeConfigParams
	for _, str := range agents {
		agent := str.(*models.Agent)

		if agent.AgentType == models.PMMAgentType {
			// TODO https://jira.percona.com/browse/PMM-4087
			continue
		}

		// sanity check
		if (agent.NodeID != nil) && (agent.ServiceID != nil) {
			svc.l.Panicf("Both agent.NodeID and agent.ServiceID are present: %s", agent)
		}

		// find Service for this Agent
		var paramsService *models.Service
		if agent.ServiceID != nil {
			paramsService, err = models.FindServiceByID(q, pointer.GetString(agent.ServiceID))
			if err != nil {
				return err
			}
		}

		// find Node for this Agent or Service
		var paramsNode *models.Node
		switch {
		case agent.NodeID != nil:
			paramsNode, err = models.FindNodeByID(q, pointer.GetString(agent.NodeID))
		case paramsService != nil:
			paramsNode, err = models.FindNodeByID(q, paramsService.NodeID)
		}
		if err != nil {
			return err
		}

		// find Node address where pmm-agent runs
		var paramsHost string
		pmmAgent, err := models.FindAgentByID(q, *agent.PMMAgentID)
		if err != nil {
			return errors.WithStack(err)
		}
		pmmAgentNode := &models.Node{NodeID: pointer.GetString(pmmAgent.RunsOnNodeID)}
		if err = q.Reload(pmmAgentNode); err != nil {
			return errors.WithStack(err)
		}
		paramsHost = pmmAgentNode.Address

		var scfgs []*config.ScrapeConfig
		switch agent.AgentType {
		case models.NodeExporterType:
			scfgs, err = scrapeConfigsForNodeExporter(s, &scrapeConfigParams{
				host:    paramsHost,
				node:    paramsNode,
				service: nil,
				agent:   agent,
			})

		case models.MySQLdExporterType:
			scfgs, err = scrapeConfigsForMySQLdExporter(s, &scrapeConfigParams{
				host:    paramsHost,
				node:    paramsNode,
				service: paramsService,
				agent:   agent,
			})

		case models.MongoDBExporterType:
			scfgs, err = scrapeConfigsForMongoDBExporter(s, &scrapeConfigParams{
				host:    paramsHost,
				node:    paramsNode,
				service: paramsService,
				agent:   agent,
			})

		case models.PostgresExporterType:
			scfgs, err = scrapeConfigsForPostgresExporter(s, &scrapeConfigParams{
				host:    paramsHost,
				node:    paramsNode,
				service: paramsService,
				agent:   agent,
			})

		case models.ProxySQLExporterType:
			scfgs, err = scrapeConfigsForProxySQLExporter(s, &scrapeConfigParams{
				host:    paramsHost,
				node:    paramsNode,
				service: paramsService,
				agent:   agent,
			})

		case models.QANMySQLPerfSchemaAgentType, models.QANMySQLSlowlogAgentType:
			continue
		case models.QANMongoDBProfilerAgentType:
			continue
		case models.QANPostgreSQLPgStatementsAgentType:
			continue

		case models.RDSExporterType:
			rdsParams = append(rdsParams, &scrapeConfigParams{
				host:    paramsHost,
				node:    paramsNode,
				service: paramsService,
				agent:   agent,
			})
			continue

		default:
			svc.l.Warnf("Skipping scrape config for %s.", agent)
			continue
		}

		if err != nil {
			svc.l.Warnf("Failed to add %s %q, skipping: %s.", agent.AgentType, agent.AgentID, err)
		}
		cfg.ScrapeConfigs = append(cfg.ScrapeConfigs, scfgs...)
	}

	scfgs, err := scrapeConfigsForRDSExporter(s, rdsParams)
	if err != nil {
		svc.l.Warnf("Failed to add rds_exporter scrape configs: %s.", err)
	}
	cfg.ScrapeConfigs = append(cfg.ScrapeConfigs, scfgs...)

	return nil
}

// marshalConfig marshals Prometheus configuration.
func (svc *Service) marshalConfig() ([]byte, error) {
	cfg := svc.loadBaseConfig()

	e := svc.db.InTransaction(func(tx *reform.TX) error {
		settings, err := models.GetSettings(tx)
		if err != nil {
			return err
		}
		s := settings.MetricsResolutions

		if cfg.GlobalConfig.ScrapeInterval == 0 {
			cfg.GlobalConfig.ScrapeInterval = model.Duration(s.LR)
		}
		if cfg.GlobalConfig.ScrapeTimeout == 0 {
			cfg.GlobalConfig.ScrapeTimeout = scrapeTimeout(s.LR)
		}
		if cfg.GlobalConfig.EvaluationInterval == 0 {
			cfg.GlobalConfig.EvaluationInterval = model.Duration(s.LR)
		}

		cfg.RuleFiles = append(cfg.RuleFiles, "/srv/prometheus/rules/*.rules.yml")

		cfg.ScrapeConfigs = append(cfg.ScrapeConfigs,
			scrapeConfigForPrometheus(s.HR),
			scrapeConfigForGrafana(s.MR),
			scrapeConfigForPMMManaged(s.MR),
			scrapeConfigForQANAPI2(s.MR),
		)

		if settings.AlertManagerURL != "" {
			u, err := url.Parse(settings.AlertManagerURL)
			if err == nil && (u.Opaque != "" || u.Host == "") {
				err = errors.Errorf("parsed incorrectly as %#v", u)
			}

			if err == nil {
				var httpClientConfig config.HTTPClientConfig
				if username := u.User.Username(); username != "" {
					password, _ := u.User.Password()
					httpClientConfig = config.HTTPClientConfig{
						BasicAuth: &config.BasicAuth{
							Username: u.User.Username(),
							Password: password,
						},
					}
				}

				cfg.AlertingConfig.AlertmanagerConfigs = append(cfg.AlertingConfig.AlertmanagerConfigs, &config.AlertmanagerConfig{
					ServiceDiscoveryConfig: config.ServiceDiscoveryConfig{
						StaticConfigs: []*config.Group{{
							Targets: []string{u.Host},
						}},
					},
					HTTPClientConfig: httpClientConfig,
					Scheme:           u.Scheme,
					PathPrefix:       u.Path,
					APIVersion:       config.AlertmanagerAPIVersionV2,
				})
			} else {
				svc.l.Errorf("Failed to parse Alert Manager URL %q: %s.", settings.AlertManagerURL, err)
			}
		}

		return svc.addScrapeConfigs(cfg, tx.Querier, &s)
	})
	if e != nil {
		return nil, e
	}

	// TODO Add comments to each cfg.ScrapeConfigs element.
	// https://jira.percona.com/browse/PMM-3601

	b, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal Prometheus configuration file")
	}

	b = append([]byte("# Managed by pmm-managed. DO NOT EDIT.\n---\n"), b...)
	return b, nil
}

// saveConfigAndReload saves given Prometheus configuration to file and reloads Prometheus.
// If configuration can't be reloaded for some reason, old file is restored, and configuration is reloaded again.
func (svc *Service) saveConfigAndReload(cfg []byte) error {
	// read existing content
	oldCfg, err := ioutil.ReadFile(svc.configPath)
	if err != nil {
		return errors.WithStack(err)
	}

	// compare with new config
	if reflect.DeepEqual(cfg, oldCfg) {
		svc.l.Infof("Configuration not changed, doing nothing.")
		return nil
	}

	fi, err := os.Stat(svc.configPath)
	if err != nil {
		return errors.WithStack(err)
	}

	// restore old content and reload in case of error
	var restore bool
	defer func() {
		if restore {
			if err = ioutil.WriteFile(svc.configPath, oldCfg, fi.Mode()); err != nil {
				svc.l.Error(err)
			}
			if err = svc.reload(); err != nil {
				svc.l.Error(err)
			}
		}
	}()

	// write new content to temporary file, check it
	f, err := ioutil.TempFile("", "pmm-managed-config-")
	if err != nil {
		return errors.WithStack(err)
	}
	if _, err = f.Write(cfg); err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()
	args := []string{"check", "config", f.Name()}
	cmd := exec.Command(svc.promtoolPath, args...) //nolint:gosec
	pdeathsig.Set(cmd, unix.SIGKILL)
	b, err := cmd.CombinedOutput()
	if err != nil {
		svc.l.Errorf("%s", b)

		// return typed error if possible
		s := string(b)
		if m := checkFailedRE.FindStringSubmatch(s); len(m) == 2 {
			return status.Error(codes.Aborted, m[1])
		}
		return errors.Wrap(err, s)
	}
	svc.l.Debugf("%s", b)

	// write to permanent location and reload
	restore = true
	if err = ioutil.WriteFile(svc.configPath, cfg, fi.Mode()); err != nil {
		return errors.WithStack(err)
	}
	if err = svc.reload(); err != nil {
		return err
	}
	svc.l.Infof("Configuration reloaded.")
	restore = false
	return nil
}

// updateConfiguration updates Prometheus configuration.
func (svc *Service) updateConfiguration() error {
	start := time.Now()
	defer func() {
		if dur := time.Since(start); dur > time.Second {
			svc.l.Warnf("updateConfiguration took %s.", dur)
		}
	}()

	cfg, err := svc.marshalConfig()
	if err != nil {
		return err
	}
	return svc.saveConfigAndReload(cfg)
}

// RequestConfigurationUpdate requests Prometheus configuration update.
func (svc *Service) RequestConfigurationUpdate() {
	select {
	case svc.sema <- struct{}{}:
	default:
	}
}

// Check verifies that Prometheus works.
func (svc *Service) Check(ctx context.Context) error {
	// check Prometheus /version API and log version
	u := *svc.baseURL
	u.Path = path.Join(u.Path, "version")
	resp, err := svc.client.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck
	b, err := ioutil.ReadAll(resp.Body)
	svc.l.Debugf("Prometheus: %s", b)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.Errorf("expected 200, got %d", resp.StatusCode)
	}

	// check promtool version
	b, err = exec.CommandContext(ctx, svc.promtoolPath, "--version").CombinedOutput() //nolint:gosec
	if err != nil {
		return errors.Wrap(err, string(b))
	}
	svc.l.Debugf("%s", b)
	return nil
}
