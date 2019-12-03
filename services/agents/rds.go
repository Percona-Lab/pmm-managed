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

package agents

import (
	"sort"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/api/agentpb"
	"github.com/percona/pmm/api/inventorypb"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
	"gopkg.in/yaml.v2"

	"github.com/percona/pmm-managed/models"
)

// rdsInstance represents a single RDS instance information from configuration file.
type rdsInstance struct {
	Region       string         `yaml:"region"`
	Instance     string         `yaml:"instance"`
	AWSAccessKey string         `yaml:"aws_access_key,omitempty"`
	AWSSecretKey string         `yaml:"aws_secret_key,omitempty"`
	Labels       model.LabelSet `yaml:"labels,omitempty"`
}

// Config contains configuration file information.
type rdsExporterConfigFile struct {
	Instances []rdsInstance `yaml:"instances"`
}

func mergeLabels(node *models.Node, agent *models.Agent) (model.LabelSet, error) {
	res := make(model.LabelSet, 16)

	labels, err := node.UnifiedLabels()
	if err != nil {
		return nil, err
	}
	for name, value := range labels {
		res[model.LabelName(name)] = model.LabelValue(value)
	}

	labels, err = agent.UnifiedLabels()
	if err != nil {
		return nil, err
	}
	for name, value := range labels {
		res[model.LabelName(name)] = model.LabelValue(value)
	}

	delete(res, model.LabelName("region"))

	if err = res.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to merge labels")
	}
	return res, nil
}

// rdsExporterConfig returns desired configuration of rds_exporter process.
func rdsExporterConfig(pairs map[*models.Node]*models.Agent, redactMode redactMode) *agentpb.SetStateRequest_AgentProcess {
	var config rdsExporterConfigFile
	var words []string
	for node, exporter := range pairs {
		labels, _ := mergeLabels(node, exporter) //TODO: add labels for service. Should we do it?
		config.Instances = append(config.Instances, rdsInstance{
			Region:       pointer.GetString(node.Region),
			Instance:     node.Address,
			AWSAccessKey: pointer.GetString(exporter.AWSAccessKey),
			AWSSecretKey: pointer.GetString(exporter.AWSSecretKey),
			Labels:       labels,
		})

		if redactMode != exposeSecrets {
			words = redactWords(exporter)
		}
	}

	tdp := templateDelimsPair()

	args := []string{
		"--web.listen-address=:" + tdp.left + " .listen_port " + tdp.right,
		"--config.file=" + tdp.left + " .TextFiles.config " + tdp.right,
	}
	sort.Strings(args)

	b, _ := yaml.Marshal(config)

	return &agentpb.SetStateRequest_AgentProcess{
		Type:               inventorypb.AgentType_RDS_EXPORTER,
		TemplateLeftDelim:  tdp.left,
		TemplateRightDelim: tdp.right,
		Args:               args,
		TextFiles: map[string]string{
			"config": "---\n" + string(b),
		},
		RedactWords: words,
	}
}
