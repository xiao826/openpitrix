// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixValidatePackageRequest openpitrix validate package request
// swagger:model openpitrixValidatePackageRequest
type OpenpitrixValidatePackageRequest struct {

	// version package
	VersionPackage strfmt.Base64 `json:"version_package,omitempty"`

	// optional: vmbased/helm
	VersionType string `json:"version_type,omitempty"`
}

// Validate validates this openpitrix validate package request
func (m *OpenpitrixValidatePackageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixValidatePackageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixValidatePackageRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixValidatePackageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
