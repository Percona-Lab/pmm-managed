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
	"os"

	"github.com/sirupsen/logrus"

	"github.com/percona/pmm-managed/utils/validators"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000-07:00",
	})

	_, errs, warns := validators.ValidateEnvVars(os.Environ())
	for _, warn := range warns {
		logrus.Warnf("Configuration warning: %s.", warn)
	}
	for _, err := range errs {
		logrus.Errorf("Configuration error: %s.", err)
	}

	if len(errs) > 0 {
		os.Exit(1)
	}
}
