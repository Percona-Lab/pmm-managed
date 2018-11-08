// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryMySqldExporter MySQLdExporter represents mysqld_exporter Agent configuration.
// swagger:model inventoryMySQLdExporter
type InventoryMySqldExporter struct {

	// Unique agent identifier.
	ID int64 `json:"id,omitempty"`

	// HTTP listen port for exposing metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// MySQL password for extracting metrics.
	Password string `json:"password,omitempty"`

	// Node identifier where agent runs.
	RunsOnNodeID int64 `json:"runs_on_node_id,omitempty"`

	// Service identifier for extracting metrics.
	ServiceID int64 `json:"service_id,omitempty"`

	// MySQL username for extracting metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this inventory my sqld exporter
func (m *InventoryMySqldExporter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryMySqldExporter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryMySqldExporter) UnmarshalBinary(b []byte) error {
	var res InventoryMySqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
