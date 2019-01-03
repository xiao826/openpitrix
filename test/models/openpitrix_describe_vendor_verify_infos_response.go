// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeVendorVerifyInfosResponse openpitrix describe vendor verify infos response
// swagger:model openpitrixDescribeVendorVerifyInfosResponse
type OpenpitrixDescribeVendorVerifyInfosResponse struct {

	// total count
	TotalCount int64 `json:"total_count,omitempty"`

	// vendor verify info set
	VendorVerifyInfoSet OpenpitrixDescribeVendorVerifyInfosResponseVendorVerifyInfoSet `json:"vendor_verify_info_set"`
}

// Validate validates this openpitrix describe vendor verify infos response
func (m *OpenpitrixDescribeVendorVerifyInfosResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeVendorVerifyInfosResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeVendorVerifyInfosResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeVendorVerifyInfosResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}