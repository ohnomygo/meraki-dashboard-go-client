// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateOrganizationSnmpReader is a Reader for the UpdateOrganizationSnmp structure.
type UpdateOrganizationSnmpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationSnmpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationSnmpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationSnmpOK creates a UpdateOrganizationSnmpOK with default headers values
func NewUpdateOrganizationSnmpOK() *UpdateOrganizationSnmpOK {
	return &UpdateOrganizationSnmpOK{}
}

/*
UpdateOrganizationSnmpOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationSnmpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update organization snmp o k response has a 2xx status code
func (o *UpdateOrganizationSnmpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization snmp o k response has a 3xx status code
func (o *UpdateOrganizationSnmpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization snmp o k response has a 4xx status code
func (o *UpdateOrganizationSnmpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization snmp o k response has a 5xx status code
func (o *UpdateOrganizationSnmpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization snmp o k response a status code equal to that given
func (o *UpdateOrganizationSnmpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization snmp o k response
func (o *UpdateOrganizationSnmpOK) Code() int {
	return 200
}

func (o *UpdateOrganizationSnmpOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/snmp][%d] updateOrganizationSnmpOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationSnmpOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/snmp][%d] updateOrganizationSnmpOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationSnmpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationSnmpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateOrganizationSnmpBody update organization snmp body
// Example: {"peerIps":["123.123.123.1"],"v2cEnabled":false,"v3AuthMode":"SHA","v3Enabled":true,"v3PrivMode":"AES128"}
swagger:model UpdateOrganizationSnmpBody
*/
type UpdateOrganizationSnmpBody struct {

	// The list of IPv4 addresses that are allowed to access the SNMP server.
	PeerIps []string `json:"peerIps"`

	// Boolean indicating whether SNMP version 2c is enabled for the organization.
	V2cEnabled bool `json:"v2cEnabled,omitempty"`

	// The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.
	// Enum: [MD5 SHA]
	V3AuthMode string `json:"v3AuthMode,omitempty"`

	// The SNMP version 3 authentication password. Must be at least 8 characters if specified.
	V3AuthPass string `json:"v3AuthPass,omitempty"`

	// Boolean indicating whether SNMP version 3 is enabled for the organization.
	V3Enabled bool `json:"v3Enabled,omitempty"`

	// The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
	// Enum: [AES128 DES]
	V3PrivMode string `json:"v3PrivMode,omitempty"`

	// The SNMP version 3 privacy password. Must be at least 8 characters if specified.
	V3PrivPass string `json:"v3PrivPass,omitempty"`
}

// Validate validates this update organization snmp body
func (o *UpdateOrganizationSnmpBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateV3AuthMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateV3PrivMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateOrganizationSnmpBodyTypeV3AuthModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MD5","SHA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateOrganizationSnmpBodyTypeV3AuthModePropEnum = append(updateOrganizationSnmpBodyTypeV3AuthModePropEnum, v)
	}
}

const (

	// UpdateOrganizationSnmpBodyV3AuthModeMD5 captures enum value "MD5"
	UpdateOrganizationSnmpBodyV3AuthModeMD5 string = "MD5"

	// UpdateOrganizationSnmpBodyV3AuthModeSHA captures enum value "SHA"
	UpdateOrganizationSnmpBodyV3AuthModeSHA string = "SHA"
)

// prop value enum
func (o *UpdateOrganizationSnmpBody) validateV3AuthModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateOrganizationSnmpBodyTypeV3AuthModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateOrganizationSnmpBody) validateV3AuthMode(formats strfmt.Registry) error {
	if swag.IsZero(o.V3AuthMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateV3AuthModeEnum("updateOrganizationSnmp"+"."+"v3AuthMode", "body", o.V3AuthMode); err != nil {
		return err
	}

	return nil
}

var updateOrganizationSnmpBodyTypeV3PrivModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AES128","DES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateOrganizationSnmpBodyTypeV3PrivModePropEnum = append(updateOrganizationSnmpBodyTypeV3PrivModePropEnum, v)
	}
}

const (

	// UpdateOrganizationSnmpBodyV3PrivModeAES128 captures enum value "AES128"
	UpdateOrganizationSnmpBodyV3PrivModeAES128 string = "AES128"

	// UpdateOrganizationSnmpBodyV3PrivModeDES captures enum value "DES"
	UpdateOrganizationSnmpBodyV3PrivModeDES string = "DES"
)

// prop value enum
func (o *UpdateOrganizationSnmpBody) validateV3PrivModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateOrganizationSnmpBodyTypeV3PrivModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateOrganizationSnmpBody) validateV3PrivMode(formats strfmt.Registry) error {
	if swag.IsZero(o.V3PrivMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateV3PrivModeEnum("updateOrganizationSnmp"+"."+"v3PrivMode", "body", o.V3PrivMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update organization snmp body based on context it is used
func (o *UpdateOrganizationSnmpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationSnmpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationSnmpBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationSnmpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
