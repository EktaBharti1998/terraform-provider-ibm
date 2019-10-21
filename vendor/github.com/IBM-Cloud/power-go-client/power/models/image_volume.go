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

// ImageVolume image volume
// swagger:model ImageVolume
type ImageVolume struct {

	// Indicates if the volume is boot capable
	// Required: true
	Bootable *bool `json:"bootable"`

	// Volume Name
	// Required: true
	Name *string `json:"name"`

	// Volume Size
	// Required: true
	Size *float64 `json:"size"`
}

// Validate validates this image volume
func (m *ImageVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageVolume) validateBootable(formats strfmt.Registry) error {

	if err := validate.Required("bootable", "body", m.Bootable); err != nil {
		return err
	}

	return nil
}

func (m *ImageVolume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ImageVolume) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageVolume) UnmarshalBinary(b []byte) error {
	var res ImageVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
