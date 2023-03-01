// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateDeviceCellularSimsReader is a Reader for the UpdateDeviceCellularSims structure.
type UpdateDeviceCellularSimsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceCellularSimsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceCellularSimsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDeviceCellularSimsOK creates a UpdateDeviceCellularSimsOK with default headers values
func NewUpdateDeviceCellularSimsOK() *UpdateDeviceCellularSimsOK {
	return &UpdateDeviceCellularSimsOK{}
}

/*
UpdateDeviceCellularSimsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceCellularSimsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update device cellular sims o k response has a 2xx status code
func (o *UpdateDeviceCellularSimsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device cellular sims o k response has a 3xx status code
func (o *UpdateDeviceCellularSimsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device cellular sims o k response has a 4xx status code
func (o *UpdateDeviceCellularSimsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device cellular sims o k response has a 5xx status code
func (o *UpdateDeviceCellularSimsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device cellular sims o k response a status code equal to that given
func (o *UpdateDeviceCellularSimsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device cellular sims o k response
func (o *UpdateDeviceCellularSimsOK) Code() int {
	return 200
}

func (o *UpdateDeviceCellularSimsOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/cellular/sims][%d] updateDeviceCellularSimsOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCellularSimsOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/cellular/sims][%d] updateDeviceCellularSimsOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCellularSimsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceCellularSimsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceCellularSimsBody update device cellular sims body
// Example: {"apns":[{"allowedIpTypes":["ipv4","ipv6"],"authentication":{"password":"secret","type":"pap","username":"milesmeraki"},"name":"internet"}],"isPrimary":true,"slot":"sim1"}
swagger:model UpdateDeviceCellularSimsBody
*/
type UpdateDeviceCellularSimsBody struct {

	// sim failover
	SimFailover *UpdateDeviceCellularSimsParamsBodySimFailover `json:"simFailover,omitempty"`

	// List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.
	Sims []*UpdateDeviceCellularSimsParamsBodySimsItems0 `json:"sims"`
}

// Validate validates this update device cellular sims body
func (o *UpdateDeviceCellularSimsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSimFailover(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSims(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularSimsBody) validateSimFailover(formats strfmt.Registry) error {
	if swag.IsZero(o.SimFailover) { // not required
		return nil
	}

	if o.SimFailover != nil {
		if err := o.SimFailover.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceCellularSims" + "." + "simFailover")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceCellularSims" + "." + "simFailover")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceCellularSimsBody) validateSims(formats strfmt.Registry) error {
	if swag.IsZero(o.Sims) { // not required
		return nil
	}

	for i := 0; i < len(o.Sims); i++ {
		if swag.IsZero(o.Sims[i]) { // not required
			continue
		}

		if o.Sims[i] != nil {
			if err := o.Sims[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCellularSims" + "." + "sims" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCellularSims" + "." + "sims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update device cellular sims body based on the context it is used
func (o *UpdateDeviceCellularSimsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSimFailover(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSims(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularSimsBody) contextValidateSimFailover(ctx context.Context, formats strfmt.Registry) error {

	if o.SimFailover != nil {
		if err := o.SimFailover.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceCellularSims" + "." + "simFailover")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceCellularSims" + "." + "simFailover")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceCellularSimsBody) contextValidateSims(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Sims); i++ {

		if o.Sims[i] != nil {
			if err := o.Sims[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCellularSims" + "." + "sims" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCellularSims" + "." + "sims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularSimsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCellularSimsParamsBodySimFailover SIM Failover settings.
swagger:model UpdateDeviceCellularSimsParamsBodySimFailover
*/
type UpdateDeviceCellularSimsParamsBodySimFailover struct {

	// Failover to secondary SIM (optional)
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update device cellular sims params body sim failover
func (o *UpdateDeviceCellularSimsParamsBodySimFailover) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device cellular sims params body sim failover based on context it is used
func (o *UpdateDeviceCellularSimsParamsBodySimFailover) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsParamsBodySimFailover) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsParamsBodySimFailover) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularSimsParamsBodySimFailover
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCellularSimsParamsBodySimsItems0 update device cellular sims params body sims items0
swagger:model UpdateDeviceCellularSimsParamsBodySimsItems0
*/
type UpdateDeviceCellularSimsParamsBodySimsItems0 struct {

	// APN configurations. If empty, the default APN will be used.
	Apns []*UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0 `json:"apns"`

	// If true, this SIM is used for boot. Must be true on single-sim devices.
	IsPrimary *bool `json:"isPrimary,omitempty"`

	// SIM slot being configured. Must be 'sim1' on single-sim devices.
	// Enum: [sim1 sim2]
	Slot string `json:"slot,omitempty"`
}

// Validate validates this update device cellular sims params body sims items0
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateApns(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularSimsParamsBodySimsItems0) validateApns(formats strfmt.Registry) error {
	if swag.IsZero(o.Apns) { // not required
		return nil
	}

	for i := 0; i < len(o.Apns); i++ {
		if swag.IsZero(o.Apns[i]) { // not required
			continue
		}

		if o.Apns[i] != nil {
			if err := o.Apns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var updateDeviceCellularSimsParamsBodySimsItems0TypeSlotPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sim1","sim2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceCellularSimsParamsBodySimsItems0TypeSlotPropEnum = append(updateDeviceCellularSimsParamsBodySimsItems0TypeSlotPropEnum, v)
	}
}

const (

	// UpdateDeviceCellularSimsParamsBodySimsItems0SlotSim1 captures enum value "sim1"
	UpdateDeviceCellularSimsParamsBodySimsItems0SlotSim1 string = "sim1"

	// UpdateDeviceCellularSimsParamsBodySimsItems0SlotSim2 captures enum value "sim2"
	UpdateDeviceCellularSimsParamsBodySimsItems0SlotSim2 string = "sim2"
)

// prop value enum
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0) validateSlotEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateDeviceCellularSimsParamsBodySimsItems0TypeSlotPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceCellularSimsParamsBodySimsItems0) validateSlot(formats strfmt.Registry) error {
	if swag.IsZero(o.Slot) { // not required
		return nil
	}

	// value enum
	if err := o.validateSlotEnum("slot", "body", o.Slot); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update device cellular sims params body sims items0 based on the context it is used
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateApns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularSimsParamsBodySimsItems0) contextValidateApns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Apns); i++ {

		if o.Apns[i] != nil {
			if err := o.Apns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularSimsParamsBodySimsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0 update device cellular sims params body sims items0 apns items0
swagger:model UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0
*/
type UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0 struct {

	// IP versions to support (permitted values include 'ipv4', 'ipv6').
	// Required: true
	AllowedIPTypes []string `json:"allowedIpTypes"`

	// authentication
	Authentication *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication `json:"authentication,omitempty"`

	// APN name.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this update device cellular sims params body sims items0 apns items0
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllowedIPTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0) validateAllowedIPTypes(formats strfmt.Registry) error {

	if err := validate.Required("allowedIpTypes", "body", o.AllowedIPTypes); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(o.Authentication) { // not required
		return nil
	}

	if o.Authentication != nil {
		if err := o.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update device cellular sims params body sims items0 apns items0 based on the context it is used
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if o.Authentication != nil {
		if err := o.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication APN authentication configurations.
swagger:model UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication
*/
type UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication struct {

	// APN password, if type is set (if APN password is not supplied, the password is left unchanged).
	Password string `json:"password,omitempty"`

	// APN auth type.
	// Enum: [chap none pap]
	Type *string `json:"type,omitempty"`

	// APN username, if type is set.
	Username string `json:"username,omitempty"`
}

// Validate validates this update device cellular sims params body sims items0 apns items0 authentication
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chap","none","pap"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypeTypePropEnum = append(updateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypeTypePropEnum, v)
	}
}

const (

	// UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypeChap captures enum value "chap"
	UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypeChap string = "chap"

	// UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypeNone captures enum value "none"
	UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypeNone string = "none"

	// UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypePap captures enum value "pap"
	UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypePap string = "pap"
)

// prop value enum
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateDeviceCellularSimsParamsBodySimsItems0ApnsItems0AuthenticationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("authentication"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device cellular sims params body sims items0 apns items0 authentication based on context it is used
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularSimsParamsBodySimsItems0ApnsItems0Authentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
