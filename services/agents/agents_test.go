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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNoDuplicateFlags(t *testing.T, flags []string) {
	s := make(map[string]struct{})
	for _, f := range flags {
		name := strings.Split(f, "=")[0]
		name = strings.TrimPrefix(name, "--no-") // --no-<name> disables --<name>
		name = strings.TrimPrefix(name, "--")
		if _, present := s[name]; present {
			assert.Failf(t, "flag (or no- form) is already present", "%q", name)
		}
		s[name] = struct{}{}
	}
}
