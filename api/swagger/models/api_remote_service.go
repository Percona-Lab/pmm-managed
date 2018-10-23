// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// APIRemoteService api remote service
// swagger:model apiRemoteService
type APIRemoteService struct {

	// address
	Address string `json:"address,omitempty"`

	// engine
	Engine string `json:"engine,omitempty"`

	// engine version
	EngineVersion string `json:"engine_version,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`
}

// Validate validates this api remote service
func (m *APIRemoteService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIRemoteService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIRemoteService) UnmarshalBinary(b []byte) error {
	var res APIRemoteService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
