// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Tenant tenant
// swagger:model Tenant
type Tenant struct {

	// Cloud Instances owned by the Tenant
	// Required: true
	CloudInstances []*CloudInstanceReference `json:"cloudInstances"`

	// Date of Tenant creation
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// Indicates if the tenant is enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// IBM Customer Number
	Icn string `json:"icn,omitempty"`

	// Peering Network Information (optional)
	PeeringNetworks []*PeeringNetwork `json:"peeringNetworks,omitempty"`

	// Tenant SSH Keys
	SSHKeys []*SSHKey `json:"sshKeys"`

	// Tenant ID
	// Required: true
	TenantID *string `json:"tenantID"`
}

// Validate validates this tenant
func (m *Tenant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeeringNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tenant) validateCloudInstances(formats strfmt.Registry) error {

	if err := validate.Required("cloudInstances", "body", m.CloudInstances); err != nil {
		return err
	}

	for i := 0; i < len(m.CloudInstances); i++ {
		if swag.IsZero(m.CloudInstances[i]) { // not required
			continue
		}

		if m.CloudInstances[i] != nil {
			if err := m.CloudInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tenant) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Tenant) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Tenant) validatePeeringNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.PeeringNetworks) { // not required
		return nil
	}

	for i := 0; i < len(m.PeeringNetworks); i++ {
		if swag.IsZero(m.PeeringNetworks[i]) { // not required
			continue
		}

		if m.PeeringNetworks[i] != nil {
			if err := m.PeeringNetworks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peeringNetworks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tenant) validateSSHKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.SSHKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.SSHKeys); i++ {
		if swag.IsZero(m.SSHKeys[i]) { // not required
			continue
		}

		if m.SSHKeys[i] != nil {
			if err := m.SSHKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sshKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tenant) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantID", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tenant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tenant) UnmarshalBinary(b []byte) error {
	var res Tenant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
