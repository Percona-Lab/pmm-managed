// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type jobResultTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *jobResultTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("job_results").
func (v *jobResultTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *jobResultTableType) Columns() []string {
	return []string{
		"id",
		"pmm_agent_id",
		"done",
		"error",
		"result",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *jobResultTableType) NewStruct() reform.Struct {
	return new(JobResult)
}

// NewRecord makes a new record for that table.
func (v *jobResultTableType) NewRecord() reform.Record {
	return new(JobResult)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *jobResultTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// JobResultTable represents job_results view or table in SQL database.
var JobResultTable = &jobResultTableType{
	s: parse.StructInfo{
		Type:    "JobResult",
		SQLName: "job_results",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "string", Column: "id"},
			{Name: "PMMAgentID", Type: "string", Column: "pmm_agent_id"},
			{Name: "Done", Type: "bool", Column: "done"},
			{Name: "Error", Type: "string", Column: "error"},
			{Name: "Result", Type: "[]uint8", Column: "result"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(JobResult).Values(),
}

// String returns a string representation of this struct or record.
func (s JobResult) String() string {
	res := make([]string, 7)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "PMMAgentID: " + reform.Inspect(s.PMMAgentID, true)
	res[2] = "Done: " + reform.Inspect(s.Done, true)
	res[3] = "Error: " + reform.Inspect(s.Error, true)
	res[4] = "Result: " + reform.Inspect(s.Result, true)
	res[5] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[6] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *JobResult) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.PMMAgentID,
		s.Done,
		s.Error,
		s.Result,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *JobResult) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.PMMAgentID,
		&s.Done,
		&s.Error,
		&s.Result,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *JobResult) View() reform.View {
	return JobResultTable
}

// Table returns Table object for that record.
func (s *JobResult) Table() reform.Table {
	return JobResultTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *JobResult) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *JobResult) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *JobResult) HasPK() bool {
	return s.ID != JobResultTable.z[JobResultTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *JobResult) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = JobResultTable
	_ reform.Struct = (*JobResult)(nil)
	_ reform.Table  = JobResultTable
	_ reform.Record = (*JobResult)(nil)
	_ fmt.Stringer  = (*JobResult)(nil)
)

func init() {
	parse.AssertUpToDate(&JobResultTable.s, new(JobResult))
}
