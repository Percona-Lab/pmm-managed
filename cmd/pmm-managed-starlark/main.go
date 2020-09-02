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

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/percona-platform/saas/pkg/starlark"
	"github.com/percona/pmm-managed/services/checks"

	"github.com/percona/pmm/api/agentpb"
	"github.com/percona/pmm/version"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("stdlog: ")

	kingpin.Version(version.FullInfo())
	kingpin.HelpFlag.Short('h')

	kingpin.Parse()

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000-07:00",
	})
	if on, _ := strconv.ParseBool(os.Getenv("PMM_DEBUG")); on {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if on, _ := strconv.ParseBool(os.Getenv("PMM_TRACE")); on {
		logrus.SetLevel(logrus.TraceLevel)
	}

	gobDecoder := gob.NewDecoder(os.Stdin)
	var data checks.StarlarkScriptData
	gobDecoder.Decode(&data)

	env, err := starlark.NewEnv(data.CheckName, data.Script, data.Funcs)
	if err != nil {
		os.Exit(1)
	}

	scriptInput, err := agentpb.UnmarshalActionQueryResult(data.ScriptInput)
	if err != nil {
		os.Exit(1)
	}

	results, err := env.Run(data.CheckName, scriptInput, data.PrintFn)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print(results)
}
