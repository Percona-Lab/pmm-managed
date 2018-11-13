// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryAddVirtualMachineNodeRequest inventory add virtual machine node request
// swagger:model inventoryAddVirtualMachineNodeRequest
type InventoryAddVirtualMachineNodeRequest struct {

	// Hostname. Is not unique.
	Hostname string `json:"hostname,omitempty"`

	// Unique user-defined Node name.
	Name string `json:"name,omitempty"`
}

// Validate validates this inventory add virtual machine node request
func (m *InventoryAddVirtualMachineNodeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryAddVirtualMachineNodeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryAddVirtualMachineNodeRequest) UnmarshalBinary(b []byte) error {
	var res InventoryAddVirtualMachineNodeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
