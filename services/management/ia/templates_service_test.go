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

package ia

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testBadTemplates     = "../../../testdata/ia/bad/*.yml"
	testBuiltinTemplates = "../../../testdata/ia/builtin/*.yml"
	testUser2Templates   = "../../../testdata/ia/user2/*.yml"
	testUserTemplates    = "../../../testdata/ia/user/*.yml"
	testMissingTemplates = "/no/such/path/*.yml"
)

func TestCollect(t *testing.T) {
	t.Parallel()

	t.Run("bad and missing template paths", func(t *testing.T) {
		t.Parallel()

		svc := NewTemplatesService()
		svc.builtinTemplatesPath = testMissingTemplates
		svc.userTemplatesPath = testBadTemplates
		svc.collectRuleTemplates(context.Background())

		require.Empty(t, svc.rules)
	})

	t.Run("valid template paths", func(t *testing.T) {
		t.Parallel()

		svc := NewTemplatesService()
		svc.builtinTemplatesPath = testBuiltinTemplates
		svc.userTemplatesPath = testUserTemplates
		svc.collectRuleTemplates(context.Background())

		require.NotEmpty(t, svc.rules)
		require.Len(t, svc.rules, 2)
		assert.Contains(t, svc.rules, "builtin_rule")
		assert.Contains(t, svc.rules, "user_rule")

		// check whether map was cleared and updated on a subsequent call
		svc.userTemplatesPath = testUser2Templates
		svc.collectRuleTemplates(context.Background())

		require.NotEmpty(t, svc.rules)
		require.Len(t, svc.rules, 2)
		assert.NotContains(t, svc.rules, "user_rule")
		assert.Contains(t, svc.rules, "builtin_rule")
		assert.Contains(t, svc.rules, "user2_rule")
	})
}
