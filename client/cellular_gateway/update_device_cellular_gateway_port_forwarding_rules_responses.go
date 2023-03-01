// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateDeviceCellularGatewayPortForwardingRulesReader is a Reader for the UpdateDeviceCellularGatewayPortForwardingRules structure.
type UpdateDeviceCellularGatewayPortForwardingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceCellularGatewayPortForwardingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceCellularGatewayPortForwardingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDeviceCellularGatewayPortForwardingRulesOK creates a UpdateDeviceCellularGatewayPortForwardingRulesOK with default headers values
func NewUpdateDeviceCellularGatewayPortForwardingRulesOK() *UpdateDeviceCellularGatewayPortForwardingRulesOK {
	return &UpdateDeviceCellularGatewayPortForwardingRulesOK{}
}

/*
UpdateDeviceCellularGatewayPortForwardingRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceCellularGatewayPortForwardingRulesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update device cellular gateway port forwarding rules o k response has a 2xx status code
func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device cellular gateway port forwarding rules o k response has a 3xx status code
func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device cellular gateway port forwarding rules o k response has a 4xx status code
func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device cellular gateway port forwarding rules o k response has a 5xx status code
func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device cellular gateway port forwarding rules o k response a status code equal to that given
func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device cellular gateway port forwarding rules o k response
func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) Code() int {
	return 200
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/cellularGateway/portForwardingRules][%d] updateDeviceCellularGatewayPortForwardingRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/cellularGateway/portForwardingRules][%d] updateDeviceCellularGatewayPortForwardingRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceCellularGatewayPortForwardingRulesBody update device cellular gateway port forwarding rules body
// Example: {"rules":[{"access":"any","lanIp":"172.31.128.5","localPort":"4","name":"test","protocol":"tcp","publicPort":"11-12","uplink":"both"},{"access":"restricted","allowedIps":["10.10.10.10","10.10.10.11"],"lanIp":"172.31.128.5","localPort":"5","name":"test 2","protocol":"tcp","publicPort":"99","uplink":"both"}]}
swagger:model UpdateDeviceCellularGatewayPortForwardingRulesBody
*/
type UpdateDeviceCellularGatewayPortForwardingRulesBody struct {

	// An array of port forwarding params
	Rules []*UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0 `json:"rules"`
}

// Validate validates this update device cellular gateway port forwarding rules body
func (o *UpdateDeviceCellularGatewayPortForwardingRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesBody) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(o.Rules) { // not required
		return nil
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCellularGatewayPortForwardingRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCellularGatewayPortForwardingRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update device cellular gateway port forwarding rules body based on the context it is used
func (o *UpdateDeviceCellularGatewayPortForwardingRulesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {
			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCellularGatewayPortForwardingRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCellularGatewayPortForwardingRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayPortForwardingRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayPortForwardingRulesBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularGatewayPortForwardingRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0 update device cellular gateway port forwarding rules params body rules items0
swagger:model UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0
*/
type UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0 struct {

	// `any` or `restricted`. Specify the right to make inbound connections on the specified ports or port ranges. If `restricted`, a list of allowed IPs is mandatory.
	// Required: true
	Access *string `json:"access"`

	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.
	AllowedIps []string `json:"allowedIps"`

	// The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	// Required: true
	LanIP *string `json:"lanIp"`

	// A port or port ranges that will receive the forwarded traffic from the WAN
	// Required: true
	LocalPort *string `json:"localPort"`

	// A descriptive name for the rule
	Name string `json:"name,omitempty"`

	// TCP or UDP
	// Required: true
	Protocol *string `json:"protocol"`

	// A port or port ranges that will be forwarded to the host on the LAN
	// Required: true
	PublicPort *string `json:"publicPort"`
}

// Validate validates this update device cellular gateway port forwarding rules params body rules items0
func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLanIP(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocalPort(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePublicPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", o.Access); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) validateLanIP(formats strfmt.Registry) error {

	if err := validate.Required("lanIp", "body", o.LanIP); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) validateLocalPort(formats strfmt.Registry) error {

	if err := validate.Required("localPort", "body", o.LocalPort); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", o.Protocol); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) validatePublicPort(formats strfmt.Registry) error {

	if err := validate.Required("publicPort", "body", o.PublicPort); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device cellular gateway port forwarding rules params body rules items0 based on context it is used
func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularGatewayPortForwardingRulesParamsBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
