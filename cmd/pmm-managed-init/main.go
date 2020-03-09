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
	"github.com/percona/pmm-managed/models"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/percona/pmm-managed/utils/envvars"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000-07:00",
	})

	envSettings, errs, warns := envvars.ParseEnvVars(os.Environ())
	for _, warn := range warns {
		logrus.Warnf("Configuration warning: %s.", warn)
	}
	for _, err := range errs {
		logrus.Errorf("Configuration error: %s.", err)
	}

	params := models.ChangeSettingsParams{
		DisableTelemetry: envSettings.DisableTelemetry,
		MetricsResolutions: models.MetricsResolutions{
			HR: envSettings.MetricsResolutions.HR,
			MR: envSettings.MetricsResolutions.MR,
			LR: envSettings.MetricsResolutions.LR,
		},
		DataRetention: envSettings.DataRetention,
	}
	err := models.ValidateSettings(params)
	if err != nil {
		logrus.Errorf("Configuration error: %s.", err)
		os.Exit(1)
	}

	if len(errs) > 0 {
		os.Exit(1)
	}
}
