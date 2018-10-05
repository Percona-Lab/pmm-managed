// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIScrapeConfigsGetResponse api scrape configs get response
// swagger:model apiScrapeConfigsGetResponse
type APIScrapeConfigsGetResponse struct {

	// scrape config
	ScrapeConfig *APIScrapeConfig `json:"scrape_config,omitempty"`

	// Scrape targets health for this scrape job
	ScrapeTargetsHealth []*APIScrapeTargetHealth `json:"scrape_targets_health"`
}

// Validate validates this api scrape configs get response
func (m *APIScrapeConfigsGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScrapeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScrapeTargetsHealth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIScrapeConfigsGetResponse) validateScrapeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ScrapeConfig) { // not required
		return nil
	}

	if m.ScrapeConfig != nil {
		if err := m.ScrapeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scrape_config")
			}
			return err
		}
	}

	return nil
}

func (m *APIScrapeConfigsGetResponse) validateScrapeTargetsHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.ScrapeTargetsHealth) { // not required
		return nil
	}

	for i := 0; i < len(m.ScrapeTargetsHealth); i++ {
		if swag.IsZero(m.ScrapeTargetsHealth[i]) { // not required
			continue
		}

		if m.ScrapeTargetsHealth[i] != nil {
			if err := m.ScrapeTargetsHealth[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scrape_targets_health" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIScrapeConfigsGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIScrapeConfigsGetResponse) UnmarshalBinary(b []byte) error {
	var res APIScrapeConfigsGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
