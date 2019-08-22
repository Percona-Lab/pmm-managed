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

package prometheus

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services/prometheus/internal/prometheus/config"
	sd_config "github.com/percona/pmm-managed/services/prometheus/internal/prometheus/discovery/config"
	"github.com/percona/pmm-managed/services/prometheus/internal/prometheus/discovery/targetgroup"
)

const addressLabel = model.LabelName(model.AddressLabel)

// scrapeTimeout returns default scrape timeout for given scrape interval.
func scrapeTimeout(interval time.Duration) model.Duration {
	switch {
	case interval <= 2*time.Second:
		return model.Duration(time.Second)
	case interval <= 10*time.Second:
		return model.Duration(interval - time.Second)
	default:
		return model.Duration(10 * time.Second)
	}
}

func scrapeConfigForPrometheus(interval time.Duration) *config.ScrapeConfig {
	return &config.ScrapeConfig{
		JobName:        "prometheus",
		ScrapeInterval: model.Duration(interval),
		ScrapeTimeout:  scrapeTimeout(interval),
		MetricsPath:    "/prometheus/metrics",
		ServiceDiscoveryConfig: sd_config.ServiceDiscoveryConfig{
			StaticConfigs: []*targetgroup.Group{{
				Targets: []model.LabelSet{{addressLabel: "127.0.0.1:9090"}},
				Labels:  model.LabelSet{"instance": "pmm-server"},
			}},
		},
	}
}

func scrapeConfigForGrafana(interval time.Duration) *config.ScrapeConfig {
	return &config.ScrapeConfig{
		JobName:        "grafana",
		ScrapeInterval: model.Duration(interval),
		ScrapeTimeout:  scrapeTimeout(interval),
		MetricsPath:    "/metrics",
		ServiceDiscoveryConfig: sd_config.ServiceDiscoveryConfig{
			StaticConfigs: []*targetgroup.Group{{
				Targets: []model.LabelSet{{addressLabel: "127.0.0.1:3000"}},
				Labels:  model.LabelSet{"instance": "pmm-server"},
			}},
		},
	}
}

func scrapeConfigForPMMManaged(interval time.Duration) *config.ScrapeConfig {
	return &config.ScrapeConfig{
		JobName:        "pmm-managed",
		ScrapeInterval: model.Duration(interval),
		ScrapeTimeout:  scrapeTimeout(interval),
		MetricsPath:    "/debug/metrics",
		ServiceDiscoveryConfig: sd_config.ServiceDiscoveryConfig{
			StaticConfigs: []*targetgroup.Group{{
				Targets: []model.LabelSet{{addressLabel: "127.0.0.1:7773"}},
				Labels:  model.LabelSet{"instance": "pmm-server"},
			}},
		},
	}
}

func mergeLabels(node *models.Node, service *models.Service, agent *models.Agent) (model.LabelSet, error) {
	res := make(model.LabelSet, 16)

	labels, err := node.UnifiedLabels()
	if err != nil {
		return nil, err
	}
	for name, value := range labels {
		res[model.LabelName(name)] = model.LabelValue(value)
	}

	if service != nil {
		labels, err = service.UnifiedLabels()
		if err != nil {
			return nil, err
		}
		for name, value := range labels {
			res[model.LabelName(name)] = model.LabelValue(value)
		}
	}

	labels, err = agent.UnifiedLabels()
	if err != nil {
		return nil, err
	}
	for name, value := range labels {
		res[model.LabelName(name)] = model.LabelValue(value)
	}

	res[model.LabelName("instance")] = model.LabelValue(agent.AgentID)

	if err = res.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to merge labels")
	}
	return res, nil
}

func jobName(agent *models.Agent, interval time.Duration) string {
	return fmt.Sprintf("%s%s_%s", agent.AgentType, strings.Replace(agent.AgentID, "/", "_", -1), interval)
}

// scraperConfigForStandardExporter returns scrape config for standard exporter: /metrics endpoint, high resolution.
// If listen port is not known yet, it returns (nil, nil).
func scraperConfigForStandardExporter(interval time.Duration, node *models.Node, service *models.Service, agent *models.Agent, collect []string) (*config.ScrapeConfig, error) {
	labels, err := mergeLabels(node, service, agent)
	if err != nil {
		return nil, err
	}

	cfg := &config.ScrapeConfig{
		JobName:        jobName(agent, interval),
		ScrapeInterval: model.Duration(interval),
		ScrapeTimeout:  scrapeTimeout(interval),
		MetricsPath:    "/metrics",
	}

	if len(collect) > 0 {
		cfg.Params = url.Values{
			"collect[]": collect,
		}
	}

	port := pointer.GetUint16(agent.ListenPort)
	if port == 0 {
		return nil, nil
	}
	hostport := net.JoinHostPort(node.Address, strconv.Itoa(int(port)))
	target := model.LabelSet{addressLabel: model.LabelValue(hostport)}
	if err = target.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to set targets")
	}

	cfg.ServiceDiscoveryConfig = sd_config.ServiceDiscoveryConfig{
		StaticConfigs: []*targetgroup.Group{{
			Targets: []model.LabelSet{target},
			Labels:  labels,
		}},
	}

	return cfg, nil
}

func scraperConfigsForNodeExporter(s *models.MetricsResolutions, node *models.Node, agent *models.Agent) ([]*config.ScrapeConfig, error) {
	hrc := []string{
		"diskstats",
		"filefd",
		"filesystem",
		"loadavg",
		"meminfo",
		"meminfo_numa",
		"netdev",
		"netstat",
		"stat",
		"time",
		"vmstat",
		"textfile.hr",
		"standard.process",
		"standard.go",
	}
	hr, err := scraperConfigForStandardExporter(s.HR, node, nil, agent, hrc)
	if err != nil {
		return nil, err
	}

	mrc := []string{
		"textfile.mr",
	}
	mr, err := scraperConfigForStandardExporter(s.MR, node, nil, agent, mrc)
	if err != nil {
		return nil, err
	}

	lrc := []string{
		"bonding",
		"entropy",
		"filesystem",
		"uname",
		"textfile.lr",
	}
	lr, err := scraperConfigForStandardExporter(s.LR, node, nil, agent, lrc)
	if err != nil {
		return nil, err
	}

	var r []*config.ScrapeConfig
	if hr != nil {
		r = append(r, hr)
	}
	if mr != nil {
		r = append(r, mr)
	}
	if lr != nil {
		r = append(r, lr)
	}
	return r, nil
}

// scraperConfigsForMySQLdExporter returns scrape config for mysqld_exporter.
// If listen port is not known yet, it returns (nil, nil).
func scraperConfigsForMySQLdExporter(s *models.MetricsResolutions, node *models.Node, service *models.Service, agent *models.Agent) ([]*config.ScrapeConfig, error) {
	hrc := []string{
		"global_status",
		"info_schema.innodb_metrics",
		"custom_query.hr",
		"standard.process",
		"standard.go",
	}
	hr, err := scraperConfigForStandardExporter(s.HR, node, service, agent, hrc)
	if err != nil {
		return nil, err
	}

	mrc := []string{
		"engine_innodb_status",
		"info_schema.innodb_cmp",
		"info_schema.innodb_cmpmem",
		"info_schema.processlist",
		"info_schema.query_response_time",
		"perf_schema.eventswaits",
		"perf_schema.file_events",
		"perf_schema.tablelocks",
		"slave_status",
		"custom_query.mr",
	}
	mr, err := scraperConfigForStandardExporter(s.MR, node, service, agent, mrc)
	if err != nil {
		return nil, err
	}

	lrc := []string{
		"auto_increment.columns",
		"binlog_size",
		"engine_tokudb_status",
		"global_variables",
		"heartbeat",
		"info_schema.clientstats",
		"info_schema.innodb_tablespaces",
		"info_schema.tables",
		"info_schema.tablestats",
		"info_schema.userstats",
		"perf_schema.eventsstatements",
		"perf_schema.file_instances",
		"perf_schema.indexiowaits",
		"perf_schema.tableiowaits",
		"perf_schema.tablestats",
		"custom_query.lr",
	}
	lr, err := scraperConfigForStandardExporter(s.LR, node, service, agent, lrc)
	if err != nil {
		return nil, err
	}

	var r []*config.ScrapeConfig
	if hr != nil {
		r = append(r, hr)
	}
	if mr != nil {
		r = append(r, mr)
	}
	if lr != nil {
		r = append(r, lr)
	}
	return r, nil
}

func scraperConfigsForMongoDBExporter(s *models.MetricsResolutions, node *models.Node, service *models.Service, agent *models.Agent) ([]*config.ScrapeConfig, error) {
	hrc := []string{
		"collection",
		"database",
	}
	hr, err := scraperConfigForStandardExporter(s.HR, node, service, agent, hrc)
	if err != nil {
		return nil, err
	}

	var r []*config.ScrapeConfig
	if hr != nil {
		r = append(r, hr)
	}
	return r, nil
}

func scraperConfigsForPostgresExporter(s *models.MetricsResolutions, node *models.Node, service *models.Service, agent *models.Agent) ([]*config.ScrapeConfig, error) {
	hrc := []string{
		"exporter",
		"custom_query.hr",
		"standard.process",
		"standard.go",
	}
	hr, err := scraperConfigForStandardExporter(s.HR, node, service, agent, hrc)
	if err != nil {
		return nil, err
	}

	mrc := []string{
		"custom_query.mr",
	}
	mr, err := scraperConfigForStandardExporter(s.MR, node, service, agent, mrc)
	if err != nil {
		return nil, err
	}

	lrc := []string{
		"custom_query.lr",
	}
	lr, err := scraperConfigForStandardExporter(s.LR, node, service, agent, lrc)
	if err != nil {
		return nil, err
	}

	var r []*config.ScrapeConfig
	if hr != nil {
		r = append(r, hr)
	}
	if mr != nil {
		r = append(r, mr)
	}
	if lr != nil {
		r = append(r, lr)
	}
	return r, nil
}

func scraperConfigsForProxySQLExporter(s *models.MetricsResolutions, node *models.Node, service *models.Service, agent *models.Agent) ([]*config.ScrapeConfig, error) {
	hrc := []string{
		"mysql_connection_pool",
		"mysql_status",
		"standard.process",
		"standard.go",
	}

	hr, err := scraperConfigForStandardExporter(s.HR, node, service, agent, hrc)
	if err != nil {
		return nil, err
	}

	var r []*config.ScrapeConfig
	if hr != nil {
		r = append(r, hr)
	}
	return r, nil
}
