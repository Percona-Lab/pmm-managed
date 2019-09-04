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
	"sort"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/api/agentpb"

	"github.com/percona/pmm-managed/models"
)

func nodeExporterConfig(node *models.Node, exporter *models.Agent) *agentpb.SetStateRequest_AgentProcess {
	tdp := templateDelimsPair(
		pointer.GetString(exporter.MetricsURL),
	)

	args := []string{
		"--collector.textfile.directory.lr=/usr/local/percona/pmm2/collectors/textfile-collector/low-resolution",
		"--collector.textfile.directory.mr=/usr/local/percona/pmm2/collectors/textfile-collector/medium-resolution",
		"--collector.textfile.directory.hr=/usr/local/percona/pmm2/collectors/textfile-collector/high-resolution",

		"--web.disable-exporter-metrics", // we enable them as a part of HR metrics

		"--web.listen-address=:" + tdp.left + " .listen_port " + tdp.right,
	}

	// do not tweak collectors on macOS as many (but not) of them are Linux-specific
	if node.Distro != "darwin" {
		args = append(args,
			// LR
			"--collector.bonding=true",
			"--collector.entropy=true",
			"--collector.textfile.lr=true",
			"--collector.uname=true",

			// MR
			"--collector.textfile.mr=true",

			// HR
			"--collector.diskstats=true",
			"--collector.filefd=true",
			"--collector.filesystem=true",
			"--collector.loadavg=true",
			"--collector.meminfo=true",
			"--collector.meminfo_numa=true",
			"--collector.netdev=true",
			"--collector.netstat=true",
			"--collector.standard.process=true",
			"--collector.standard.go=true",
			"--collector.stat=true",
			"--collector.textfile.hr=true",
			"--collector.time=true",
			"--collector.vmstat=true",

			// disabled
			"--collector.arp=false",
			"--collector.bcache=false",
			"--collector.buddyinfo=false",
			"--collector.conntrack=false",
			"--collector.cpu=false",
			"--collector.drbd=false",
			"--collector.edac=false",
			"--collector.hwmon=false",
			"--collector.infiniband=false",
			"--collector.interrupts=false",
			"--collector.ipvs=false",
			"--collector.ksmd=false",
			"--collector.logind=false",
			"--collector.mdadm=false",
			"--collector.mountstats=false",
			"--collector.netclass=false",
			"--collector.nfs=false",
			"--collector.nfsd=false",
			"--collector.ntp=false",
			"--collector.processes=false",
			"--collector.qdisc=false",
			"--collector.runit=false",
			"--collector.sockstat=false",
			"--collector.supervisord=false",
			"--collector.systemd=false",
			"--collector.tcpstat=false",
			"--collector.timex=false",
			"--collector.wifi=false",
			"--collector.xfs=false",
			"--collector.zfs=false",

			// add more netstat fields
			"--collector.netstat.fields=^(.*_(InErrors|InErrs|InCsumErrors)"+
				"|Tcp_(ActiveOpens|PassiveOpens|RetransSegs|CurrEstab|AttemptFails|OutSegs|InSegs|EstabResets|OutRsts|OutSegs)|Tcp_Rto(Algorithm|Min|Max)"+
				"|Udp_(RcvbufErrors|SndbufErrors)|Udp(6?|Lite6?)_(InDatagrams|OutDatagrams|RcvbufErrors|SndbufErrors|NoPorts)"+
				"|Icmp6?_(OutEchoReps|OutEchos|InEchos|InEchoReps|InAddrMaskReps|InAddrMasks|OutAddrMaskReps|OutAddrMasks|InTimestampReps|InTimestamps"+
				"|OutTimestampReps|OutTimestamps|OutErrors|InDestUnreachs|OutDestUnreachs|InTimeExcds|InRedirects|OutRedirects|InMsgs|OutMsgs)"+
				"|IcmpMsg_(InType3|OutType3)|Ip(6|Ext)_(InOctets|OutOctets)|Ip_Forwarding|TcpExt_(Listen.*|Syncookies.*|TCPTimeouts))$",

			// add more vmstat fileds
			"--collector.vmstat.fields=^(pg(steal_(kswapd|direct)|refill|alloc)_(movable|normal|dma3?2?)"+
				"|nr_(dirty.*|slab.*|vmscan.*|isolated.*|free.*|shmem.*|i?n?active.*|anon_transparent_.*|writeback.*|unstable"+
				"|unevictable|mlock|mapped|bounce|page_table_pages|kernel_stack)|drop_slab|slabs_scanned|pgd?e?activate"+
				"|pgpg(in|out)|pswp(in|out)|pgm?a?j?fault)$",
		)
	}

	if pointer.GetString(exporter.MetricsURL) != "" {
		args = append(args, "--web.telemetry-path="+*exporter.MetricsURL)
	}

	sort.Strings(args)

	return &agentpb.SetStateRequest_AgentProcess{
		Type:               agentpb.Type_NODE_EXPORTER,
		TemplateLeftDelim:  tdp.left,
		TemplateRightDelim: tdp.right,
		Args:               args,
		Env: []string{
			fmt.Sprintf("HTTP_AUTH=pmm:%s", exporter.AgentID),
		},
	}
}
