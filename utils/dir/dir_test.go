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

package dir

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDataDir(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name      string
		path      string
		username  string
		groupname string
		perm      os.FileMode
		err       string
	}{{
		name:      "valid params",
		path:      "/tmp/testdir",
		username:  "pmm",
		groupname: "pmm",
		perm:      os.FileMode(0o775),
		err:       ``,
	}, {
		name:      "invalid path",
		path:      "",
		username:  "pmm",
		groupname: "pmm",
		perm:      os.FileMode(0o775),
		err:       `cannot create path "": mkdir : no such file or directory`,
	}, {
		name:      "unknown user",
		path:      "/tmp/testdir",
		username:  "$",
		groupname: "pmm",
		perm:      os.FileMode(0o775),
		err:       `user: unknown user $`,
	}, {
		name:      "unknown group",
		path:      "/tmp/testdir",
		username:  "pmm",
		groupname: "$",
		perm:      os.FileMode(0o775),
		err:       `group: unknown group $`,
	}}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			defer os.Remove(tc.path) //nolint:errcheck

			err := CreateDataDir(tc.path, tc.username, tc.groupname, tc.perm)
			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
				return
			}

			assert.NoError(t, err)
			stat, err := os.Stat(tc.path)
			require.NoError(t, err)
			assert.True(t, stat.IsDir())
			assert.Equal(t, tc.perm, stat.Mode().Perm())
		})
	}
}
