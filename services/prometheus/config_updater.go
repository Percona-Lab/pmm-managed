// pmm-managed
// Copyright (C) 2017 Percona LLC
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

package prometheus

import (
	"fmt"

	"github.com/prometheus/common/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/percona/pmm-managed/services/prometheus/internal"
)

func convertScrapeConfig(cfg *ScrapeConfig) (*internal.ScrapeConfig, error) {
	var err error
	var interval, timeout model.Duration
	if cfg.ScrapeInterval != "" {
		interval, err = model.ParseDuration(cfg.ScrapeInterval)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "interval: %s", err)
		}
	}
	if cfg.ScrapeTimeout != "" {
		timeout, err = model.ParseDuration(cfg.ScrapeTimeout)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "timeout: %s", err)
		}
	}

	var basicAuth *internal.BasicAuth
	if cfg.BasicAuth != nil {
		basicAuth = &internal.BasicAuth{
			Username: cfg.BasicAuth.Username,
			Password: cfg.BasicAuth.Password,
		}
	}

	tg := make([]*internal.TargetGroup, len(cfg.StaticConfigs))
	for i, sc := range cfg.StaticConfigs {
		tg[i] = new(internal.TargetGroup)

		for _, t := range sc.Targets {
			ls := model.LabelSet{model.AddressLabel: model.LabelValue(t)}
			if err = ls.Validate(); err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "static_configs.targets: %s", err)
			}
			tg[i].Targets = append(tg[i].Targets, ls)
		}

		ls := make(model.LabelSet)
		for _, lp := range sc.Labels {
			ls[model.LabelName(lp.Name)] = model.LabelValue(lp.Value)
		}
		if err = ls.Validate(); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "static_configs.labels: %s", err)
		}
		tg[i].Labels = ls
	}

	return &internal.ScrapeConfig{
		JobName:        cfg.JobName,
		ScrapeInterval: interval,
		ScrapeTimeout:  timeout,
		MetricsPath:    cfg.MetricsPath,
		Scheme:         cfg.Scheme,
		HTTPClientConfig: internal.HTTPClientConfig{
			BasicAuth: basicAuth,
			TLSConfig: internal.TLSConfig{
				InsecureSkipVerify: cfg.TLSConfig.InsecureSkipVerify,
			},
		},
		ServiceDiscoveryConfig: internal.ServiceDiscoveryConfig{
			StaticConfigs: tg,
		},
	}, nil
}

// configUpdater implements Prometheus configuration updating logic:
// it changes both sources while keeping them in sync.
// Input-output is done in Service.
type configUpdater struct {
	consulData []ScrapeConfig
	fileData   []*internal.ScrapeConfig
}

func (cu *configUpdater) addScrapeConfig(scrapeConfig *ScrapeConfig) error {
	cfg, err := convertScrapeConfig(scrapeConfig)
	if err != nil {
		return err
	}

	for _, sc := range cu.consulData {
		if sc.JobName == cfg.JobName {
			return status.Errorf(codes.AlreadyExists, "scrape config with job name %q already exist", cfg.JobName)
		}
	}

	for _, sc := range cu.fileData {
		if sc.JobName == cfg.JobName {
			return status.Errorf(codes.FailedPrecondition, "scrape config with job name %q is built-in", cfg.JobName)
		}
	}

	cu.consulData = append(cu.consulData, *scrapeConfig)
	cu.fileData = append(cu.fileData, cfg)
	return nil
}

func (cu *configUpdater) removeScrapeConfig(jobName string) error {
	consulDataI := -1
	for i, sc := range cu.consulData {
		if sc.JobName == jobName {
			consulDataI = i
			break
		}
	}
	if consulDataI < 0 {
		return status.Errorf(codes.NotFound, "scrape config with job name %q not found", jobName)
	}

	fileDataI := -1
	for i, sc := range cu.fileData {
		if sc.JobName == jobName {
			fileDataI = i
			break
		}
	}
	if fileDataI < 0 {
		return status.Errorf(codes.FailedPrecondition, "scrape config with job name %q not found in configuration file", jobName)
	}

	cu.consulData = append(cu.consulData[:consulDataI], cu.consulData[consulDataI+1:]...)
	cu.fileData = append(cu.fileData[:fileDataI], cu.fileData[fileDataI+1:]...)
	return nil
}

func (cu *configUpdater) addStaticTargets(jobName string, targets []string) error {
	consulDataI := -1
	for i, sc := range cu.consulData {
		if sc.JobName == jobName {
			consulDataI = i
			break
		}
	}
	if consulDataI < 0 {
		return status.Errorf(codes.NotFound, "scrape config with job name %q not found", jobName)
	}

	var staticConfig StaticConfig
	switch len(cu.consulData[consulDataI].StaticConfigs) {
	case 0:
		// nothing
	case 1:
		staticConfig = cu.consulData[consulDataI].StaticConfigs[0]
	default:
		msg := fmt.Sprintf(
			"scrape config with job name %q has %d static configs, that is not supported yet",
			jobName, len(cu.consulData[consulDataI].StaticConfigs),
		)
		return status.Error(codes.Unimplemented, msg)
	}
	for _, add := range targets {
		var found bool
		for _, t := range staticConfig.Targets {
			if t == add {
				found = true
				break
			}
		}
		if found {
			continue
		}
		staticConfig.Targets = append(staticConfig.Targets, add)
	}

	scrapeConfig := cu.consulData[consulDataI]
	scrapeConfig.StaticConfigs = []StaticConfig{staticConfig}
	cfg, err := convertScrapeConfig(&scrapeConfig)
	if err != nil {
		return err
	}

	fileDataI := -1
	for i, sc := range cu.fileData {
		if sc.JobName == jobName {
			fileDataI = i
			break
		}
	}
	if fileDataI < 0 {
		return status.Errorf(codes.FailedPrecondition, "scrape config with job name %q not found in configuration file", jobName)
	}

	cu.consulData[consulDataI] = scrapeConfig
	cu.fileData[fileDataI] = cfg
	return nil
}

func (cu *configUpdater) removeStaticTargets(jobName string, targets []string) error {
	consulDataI := -1
	for i, sc := range cu.consulData {
		if sc.JobName == jobName {
			consulDataI = i
			break
		}
	}
	if consulDataI < 0 {
		return status.Errorf(codes.NotFound, "scrape config with job name %q not found", jobName)
	}

	var staticConfig StaticConfig
	switch len(cu.consulData[consulDataI].StaticConfigs) {
	case 0:
		// nothing
	case 1:
		staticConfig = cu.consulData[consulDataI].StaticConfigs[0]
	default:
		msg := fmt.Sprintf(
			"scrape config with job name %q has %d static configs, that is not supported yet",
			jobName, len(cu.consulData[consulDataI].StaticConfigs),
		)
		return status.Error(codes.Unimplemented, msg)
	}
	for _, remove := range targets {
		for i, t := range staticConfig.Targets {
			if t == remove {
				staticConfig.Targets = append(staticConfig.Targets[:i], staticConfig.Targets[i+1:]...)
				break
			}
		}
	}

	scrapeConfig := cu.consulData[consulDataI]
	if len(staticConfig.Targets) > 0 {
		scrapeConfig.StaticConfigs = []StaticConfig{staticConfig}
	} else {
		scrapeConfig.StaticConfigs = nil
	}
	cfg, err := convertScrapeConfig(&scrapeConfig)
	if err != nil {
		return err
	}

	fileDataI := -1
	for i, sc := range cu.fileData {
		if sc.JobName == jobName {
			fileDataI = i
			break
		}
	}
	if fileDataI < 0 {
		return status.Errorf(codes.FailedPrecondition, "scrape config with job name %q not found in configuration file", jobName)
	}

	cu.consulData[consulDataI] = scrapeConfig
	cu.fileData[fileDataI] = cfg
	return nil
}
