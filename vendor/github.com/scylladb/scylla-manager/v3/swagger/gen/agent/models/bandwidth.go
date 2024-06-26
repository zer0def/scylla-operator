// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Bandwidth bandwidth rate
//
// # Rate at witch to rate limit bandwidth
//
// swagger:model Bandwidth
type Bandwidth struct {

	// String representation of the bandwidth rate limit (eg. 100k, 1M, ...).
	Rate string `json:"rate,omitempty"`
}

// Validate validates this bandwidth
func (m *Bandwidth) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Bandwidth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bandwidth) UnmarshalBinary(b []byte) error {
	var res Bandwidth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
