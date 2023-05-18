// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// UpdateNetworkSwitchAlternateManagementInterfaceReader is a Reader for the UpdateNetworkSwitchAlternateManagementInterface structure.
type UpdateNetworkSwitchAlternateManagementInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchAlternateManagementInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSwitchAlternateManagementInterfaceOK creates a UpdateNetworkSwitchAlternateManagementInterfaceOK with default headers values
func NewUpdateNetworkSwitchAlternateManagementInterfaceOK() *UpdateNetworkSwitchAlternateManagementInterfaceOK {
	return &UpdateNetworkSwitchAlternateManagementInterfaceOK{}
}

/*
UpdateNetworkSwitchAlternateManagementInterfaceOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchAlternateManagementInterfaceOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network switch alternate management interface o k response has a 2xx status code
func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network switch alternate management interface o k response has a 3xx status code
func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network switch alternate management interface o k response has a 4xx status code
func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network switch alternate management interface o k response has a 5xx status code
func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network switch alternate management interface o k response a status code equal to that given
func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network switch alternate management interface o k response
func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) Code() int {
	return 200
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/alternateManagementInterface][%d] updateNetworkSwitchAlternateManagementInterfaceOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/alternateManagementInterface][%d] updateNetworkSwitchAlternateManagementInterfaceOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkSwitchAlternateManagementInterfaceBody update network switch alternate management interface body
// Example: {"enabled":true,"protocols":["radius","snmp","syslog"],"switches":[{"alternateManagementIp":"1.2.3.4","gateway":"1.2.3.5","serial":"Q234-ABCD-5678","subnetMask":"255.255.255.0"}],"vlanId":100}
swagger:model UpdateNetworkSwitchAlternateManagementInterfaceBody
*/
type UpdateNetworkSwitchAlternateManagementInterfaceBody struct {

	// Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Enabled bool `json:"enabled,omitempty"`

	// Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Protocols []string `json:"protocols"`

	// Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	Switches []*UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0 `json:"switches"`

	// Alternate management VLAN, must be between 1 and 4094
	VlanID int64 `json:"vlanId,omitempty"`
}

// Validate validates this update network switch alternate management interface body
func (o *UpdateNetworkSwitchAlternateManagementInterfaceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProtocols(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSwitches(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkSwitchAlternateManagementInterfaceBodyProtocolsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["radius","snmp","syslog"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchAlternateManagementInterfaceBodyProtocolsItemsEnum = append(updateNetworkSwitchAlternateManagementInterfaceBodyProtocolsItemsEnum, v)
	}
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceBody) validateProtocolsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchAlternateManagementInterfaceBodyProtocolsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceBody) validateProtocols(formats strfmt.Registry) error {
	if swag.IsZero(o.Protocols) { // not required
		return nil
	}

	for i := 0; i < len(o.Protocols); i++ {

		// value enum
		if err := o.validateProtocolsItemsEnum("updateNetworkSwitchAlternateManagementInterface"+"."+"protocols"+"."+strconv.Itoa(i), "body", o.Protocols[i]); err != nil {
			return err
		}

	}

	return nil
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceBody) validateSwitches(formats strfmt.Registry) error {
	if swag.IsZero(o.Switches) { // not required
		return nil
	}

	for i := 0; i < len(o.Switches); i++ {
		if swag.IsZero(o.Switches[i]) { // not required
			continue
		}

		if o.Switches[i] != nil {
			if err := o.Switches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchAlternateManagementInterface" + "." + "switches" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchAlternateManagementInterface" + "." + "switches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network switch alternate management interface body based on the context it is used
func (o *UpdateNetworkSwitchAlternateManagementInterfaceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSwitches(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceBody) contextValidateSwitches(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Switches); i++ {

		if o.Switches[i] != nil {
			if err := o.Switches[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchAlternateManagementInterface" + "." + "switches" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchAlternateManagementInterface" + "." + "switches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAlternateManagementInterfaceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAlternateManagementInterfaceBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAlternateManagementInterfaceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0 update network switch alternate management interface params body switches items0
swagger:model UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0
*/
type UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0 struct {

	// Switch alternative management IP. To remove a prior IP setting, provide an empty string
	// Required: true
	AlternateManagementIP *string `json:"alternateManagementIp"`

	// Switch gateway must be in IP format. Only and must be specified for Polaris switches
	Gateway string `json:"gateway,omitempty"`

	// Switch serial number
	// Required: true
	Serial *string `json:"serial"`

	// Switch subnet mask must be in IP format. Only and must be specified for Polaris switches
	SubnetMask string `json:"subnetMask,omitempty"`
}

// Validate validates this update network switch alternate management interface params body switches items0
func (o *UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlternateManagementIP(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0) validateAlternateManagementIP(formats strfmt.Registry) error {

	if err := validate.Required("alternateManagementIp", "body", o.AlternateManagementIP); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", o.Serial); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch alternate management interface params body switches items0 based on context it is used
func (o *UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAlternateManagementInterfaceParamsBodySwitchesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
