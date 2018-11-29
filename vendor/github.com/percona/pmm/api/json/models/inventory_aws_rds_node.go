// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryAWSRDSNode AWSRDSNode represents AWS RDS Node.
// swagger:model inventoryAWSRDSNode
type InventoryAWSRDSNode struct {

	// Hostname. Unique in combination with region.
	Hostname string `json:"hostname,omitempty"`

	// Unique Node identifier.
	ID int64 `json:"id,omitempty"`

	// Unique user-defined Node name.
	Name string `json:"name,omitempty"`

	// AWS region. Unique in combination with hostname.
	Region string `json:"region,omitempty"`
}

// Validate validates this inventory AWS RDS node
func (m *InventoryAWSRDSNode) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryAWSRDSNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryAWSRDSNode) UnmarshalBinary(b []byte) error {
	var res InventoryAWSRDSNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
