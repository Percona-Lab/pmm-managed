package models

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type agentNodeViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *agentNodeViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("agent_nodes").
func (v *agentNodeViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *agentNodeViewType) Columns() []string {
	return []string{"agent_id", "node_id"}
}

// NewStruct makes a new struct for that view or table.
func (v *agentNodeViewType) NewStruct() reform.Struct {
	return new(AgentNode)
}

// AgentNodeView represents agent_nodes view or table in SQL database.
var AgentNodeView = &agentNodeViewType{
	s: parse.StructInfo{Type: "AgentNode", SQLSchema: "", SQLName: "agent_nodes", Fields: []parse.FieldInfo{{Name: "AgentID", PKType: "", Column: "agent_id"}, {Name: "NodeID", PKType: "", Column: "node_id"}}, PKFieldIndex: -1},
	z: new(AgentNode).Values(),
}

// String returns a string representation of this struct or record.
func (s AgentNode) String() string {
	res := make([]string, 2)
	res[0] = "AgentID: " + reform.Inspect(s.AgentID, true)
	res[1] = "NodeID: " + reform.Inspect(s.NodeID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *AgentNode) Values() []interface{} {
	return []interface{}{
		s.AgentID,
		s.NodeID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *AgentNode) Pointers() []interface{} {
	return []interface{}{
		&s.AgentID,
		&s.NodeID,
	}
}

// View returns View object for that struct.
func (s *AgentNode) View() reform.View {
	return AgentNodeView
}

// check interfaces
var (
	_ reform.View   = AgentNodeView
	_ reform.Struct = new(AgentNode)
	_ fmt.Stringer  = new(AgentNode)
)

func init() {
	parse.AssertUpToDate(&AgentNodeView.s, new(AgentNode))
}
