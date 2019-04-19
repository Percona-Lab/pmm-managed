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
	"fmt"
	"net"
	"net/url"
	"sort"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/api/agentpb"

	"github.com/percona/pmm-managed/models"
)

func mongodbExporterConfig(service *models.Service, exporter *models.Agent) *agentpb.SetStateRequest_AgentProcess {
	tdp := templateDelimsPair(
		pointer.GetString(service.Address),
		pointer.GetString(exporter.Username),
		pointer.GetString(exporter.Password),
		pointer.GetString(exporter.MetricsURL),
	)

	args := []string{
		"--collect.database",
		"--collect.collection",
		"--collect.topmetrics",
		"--web.listen-address=:" + tdp.left + " .listen_port " + tdp.right,
	}

	if pointer.GetString(exporter.MetricsURL) != "" {
		args = append(args, "--web.telemetry-path="+*exporter.MetricsURL)
	}

	sort.Strings(args)

	connString := mongoDSN(service, exporter)

	return &agentpb.SetStateRequest_AgentProcess{
		Type:               agentpb.Type_MONGODB_EXPORTER,
		TemplateLeftDelim:  tdp.left,
		TemplateRightDelim: tdp.right,
		Args:               args,
		Env: []string{
			fmt.Sprintf("MONGODB_URI=%s", connString),
		},
	}
}

// qanMongoDBProfilerAgentConfig returns desired configuration of qan-mysql-perfschema internal agent.
func qanMongoDBProfilerAgentConfig(service *models.Service, exporter *models.Agent) *agentpb.SetStateRequest_BuiltinAgent {
	return &agentpb.SetStateRequest_BuiltinAgent{
		Type: agentpb.Type_QAN_MONGODB_PROFILER_AGENT,
		Dsn:  mongoDSN(service, exporter),
	}
}

func mongoDSN(service *models.Service, exporter *models.Agent) string {

	host := pointer.GetString(service.Address)
	port := pointer.GetUint16(service.Port)

	connURL := url.URL{
		Scheme: "mongodb",
		Host:   net.JoinHostPort(host, strconv.Itoa(int(port))),
	}

	username := pointer.GetString(exporter.Username)
	password := pointer.GetString(exporter.Password)
	if username != "" {
		connURL.User = url.UserPassword(username, password)
	}

	return connURL.String()
}
