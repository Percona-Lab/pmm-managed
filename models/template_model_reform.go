// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type templateTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *templateTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("rule_templates").
func (v *templateTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *templateTableType) Columns() []string {
	return []string{"name", "version", "summary", "tiers", "expr", "params", "for", "severity", "labels", "annotations", "source", "created_at", "updated_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *templateTableType) NewStruct() reform.Struct {
	return new(Template)
}

// NewRecord makes a new record for that table.
func (v *templateTableType) NewRecord() reform.Record {
	return new(Template)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *templateTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// TemplateTable represents rule_templates view or table in SQL database.
var TemplateTable = &templateTableType{
	s: parse.StructInfo{Type: "Template", SQLSchema: "", SQLName: "rule_templates", Fields: []parse.FieldInfo{{Name: "Name", Type: "string", Column: "name"}, {Name: "Version", Type: "uint32", Column: "version"}, {Name: "Summary", Type: "string", Column: "summary"}, {Name: "Tiers", Type: "Tiers", Column: "tiers"}, {Name: "Expr", Type: "string", Column: "expr"}, {Name: "Params", Type: "Params", Column: "params"}, {Name: "For", Type: "Duration", Column: "for"}, {Name: "Severity", Type: "Severity", Column: "severity"}, {Name: "Labels", Type: "Map", Column: "labels"}, {Name: "Annotations", Type: "Map", Column: "annotations"}, {Name: "Source", Type: "string", Column: "source"}, {Name: "CreatedAt", Type: "time.Time", Column: "created_at"}, {Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"}}, PKFieldIndex: 0},
	z: new(Template).Values(),
}

// String returns a string representation of this struct or record.
func (s Template) String() string {
	res := make([]string, 13)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	res[1] = "Version: " + reform.Inspect(s.Version, true)
	res[2] = "Summary: " + reform.Inspect(s.Summary, true)
	res[3] = "Tiers: " + reform.Inspect(s.Tiers, true)
	res[4] = "Expr: " + reform.Inspect(s.Expr, true)
	res[5] = "Params: " + reform.Inspect(s.Params, true)
	res[6] = "For: " + reform.Inspect(s.For, true)
	res[7] = "Severity: " + reform.Inspect(s.Severity, true)
	res[8] = "Labels: " + reform.Inspect(s.Labels, true)
	res[9] = "Annotations: " + reform.Inspect(s.Annotations, true)
	res[10] = "Source: " + reform.Inspect(s.Source, true)
	res[11] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[12] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Template) Values() []interface{} {
	return []interface{}{
		s.Name,
		s.Version,
		s.Summary,
		s.Tiers,
		s.Expr,
		s.Params,
		s.For,
		s.Severity,
		s.Labels,
		s.Annotations,
		s.Source,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Template) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
		&s.Version,
		&s.Summary,
		&s.Tiers,
		&s.Expr,
		&s.Params,
		&s.For,
		&s.Severity,
		&s.Labels,
		&s.Annotations,
		&s.Source,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Template) View() reform.View {
	return TemplateTable
}

// Table returns Table object for that record.
func (s *Template) Table() reform.Table {
	return TemplateTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Template) PKValue() interface{} {
	return s.Name
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Template) PKPointer() interface{} {
	return &s.Name
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Template) HasPK() bool {
	return s.Name != TemplateTable.z[TemplateTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Template) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Name = string(i64)
	} else {
		s.Name = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = TemplateTable
	_ reform.Struct = (*Template)(nil)
	_ reform.Table  = TemplateTable
	_ reform.Record = (*Template)(nil)
	_ fmt.Stringer  = (*Template)(nil)
)

func init() {
	parse.AssertUpToDate(&TemplateTable.s, new(Template))
}
