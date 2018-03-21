// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APILogsAllResponse api logs all response
// swagger:model apiLogsAllResponse

type APILogsAllResponse struct {

	// logs
	Logs map[string]APILog `json:"logs,omitempty"`
}

/* polymorph apiLogsAllResponse logs false */

// Validate validates this api logs all response
func (m *APILogsAllResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APILogsAllResponse) validateLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	if err := validate.Required("logs", "body", m.Logs); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APILogsAllResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APILogsAllResponse) UnmarshalBinary(b []byte) error {
	var res APILogsAllResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
