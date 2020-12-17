// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type ruleTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *ruleTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("ia_rules").
func (v *ruleTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *ruleTableType) Columns() []string {
	return []string{
		"template_name",
		"id",
		"summary",
		"disabled",
		"params",
		"for",
		"severity",
		"custom_labels",
		"filters",
		"channel_ids",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *ruleTableType) NewStruct() reform.Struct {
	return new(Rule)
}

// NewRecord makes a new record for that table.
func (v *ruleTableType) NewRecord() reform.Record {
	return new(Rule)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *ruleTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// RuleTable represents ia_rules view or table in SQL database.
var RuleTable = &ruleTableType{
	s: parse.StructInfo{
		Type:    "Rule",
		SQLName: "ia_rules",
		Fields: []parse.FieldInfo{
			{Name: "TemplateName", Type: "string", Column: "template_name"},
			{Name: "ID", Type: "string", Column: "id"},
			{Name: "Summary", Type: "string", Column: "summary"},
			{Name: "Disabled", Type: "bool", Column: "disabled"},
			{Name: "Params", Type: "RuleParams", Column: "params"},
			{Name: "For", Type: "time.Duration", Column: "for"},
			{Name: "Severity", Type: "Severity", Column: "severity"},
			{Name: "CustomLabels", Type: "[]uint8", Column: "custom_labels"},
			{Name: "Filters", Type: "Filters", Column: "filters"},
			{Name: "ChannelIDs", Type: "ChannelIDs", Column: "channel_ids"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 1,
	},
	z: new(Rule).Values(),
}

// String returns a string representation of this struct or record.
func (s Rule) String() string {
	res := make([]string, 12)
	res[0] = "TemplateName: " + reform.Inspect(s.TemplateName, true)
	res[1] = "ID: " + reform.Inspect(s.ID, true)
	res[2] = "Summary: " + reform.Inspect(s.Summary, true)
	res[3] = "Disabled: " + reform.Inspect(s.Disabled, true)
	res[4] = "Params: " + reform.Inspect(s.Params, true)
	res[5] = "For: " + reform.Inspect(s.For, true)
	res[6] = "Severity: " + reform.Inspect(s.Severity, true)
	res[7] = "CustomLabels: " + reform.Inspect(s.CustomLabels, true)
	res[8] = "Filters: " + reform.Inspect(s.Filters, true)
	res[9] = "ChannelIDs: " + reform.Inspect(s.ChannelIDs, true)
	res[10] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[11] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Rule) Values() []interface{} {
	return []interface{}{
		s.TemplateName,
		s.ID,
		s.Summary,
		s.Disabled,
		s.Params,
		s.For,
		s.Severity,
		s.CustomLabels,
		s.Filters,
		s.ChannelIDs,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Rule) Pointers() []interface{} {
	return []interface{}{
		&s.TemplateName,
		&s.ID,
		&s.Summary,
		&s.Disabled,
		&s.Params,
		&s.For,
		&s.Severity,
		&s.CustomLabels,
		&s.Filters,
		&s.ChannelIDs,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Rule) View() reform.View {
	return RuleTable
}

// Table returns Table object for that record.
func (s *Rule) Table() reform.Table {
	return RuleTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Rule) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Rule) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Rule) HasPK() bool {
	return s.ID != RuleTable.z[RuleTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Rule) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = RuleTable
	_ reform.Struct = (*Rule)(nil)
	_ reform.Table  = RuleTable
	_ reform.Record = (*Rule)(nil)
	_ fmt.Stringer  = (*Rule)(nil)
)

func init() {
	parse.AssertUpToDate(&RuleTable.s, new(Rule))
}
