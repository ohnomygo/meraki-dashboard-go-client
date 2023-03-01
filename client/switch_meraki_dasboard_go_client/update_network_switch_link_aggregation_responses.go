// Code generated by go-swagger; DO NOT EDIT.

package switch_meraki_dasboard_go_client

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

// UpdateNetworkSwitchLinkAggregationReader is a Reader for the UpdateNetworkSwitchLinkAggregation structure.
type UpdateNetworkSwitchLinkAggregationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchLinkAggregationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchLinkAggregationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSwitchLinkAggregationOK creates a UpdateNetworkSwitchLinkAggregationOK with default headers values
func NewUpdateNetworkSwitchLinkAggregationOK() *UpdateNetworkSwitchLinkAggregationOK {
	return &UpdateNetworkSwitchLinkAggregationOK{}
}

/*
UpdateNetworkSwitchLinkAggregationOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchLinkAggregationOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network switch link aggregation o k response has a 2xx status code
func (o *UpdateNetworkSwitchLinkAggregationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network switch link aggregation o k response has a 3xx status code
func (o *UpdateNetworkSwitchLinkAggregationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network switch link aggregation o k response has a 4xx status code
func (o *UpdateNetworkSwitchLinkAggregationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network switch link aggregation o k response has a 5xx status code
func (o *UpdateNetworkSwitchLinkAggregationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network switch link aggregation o k response a status code equal to that given
func (o *UpdateNetworkSwitchLinkAggregationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network switch link aggregation o k response
func (o *UpdateNetworkSwitchLinkAggregationOK) Code() int {
	return 200
}

func (o *UpdateNetworkSwitchLinkAggregationOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/linkAggregations/{linkAggregationId}][%d] updateNetworkSwitchLinkAggregationOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchLinkAggregationOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/linkAggregations/{linkAggregationId}][%d] updateNetworkSwitchLinkAggregationOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchLinkAggregationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchLinkAggregationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkSwitchLinkAggregationBody update network switch link aggregation body
// Example: {"id":"NDU2N18yXzM=","switchPorts":[{"portId":"1","serial":"Q234-ABCD-0001"},{"portId":"2","serial":"Q234-ABCD-0002"},{"portId":"3","serial":"Q234-ABCD-0003"},{"portId":"4","serial":"Q234-ABCD-0004"},{"portId":"5","serial":"Q234-ABCD-0005"},{"portId":"6","serial":"Q234-ABCD-0006"},{"portId":"7","serial":"Q234-ABCD-0007"},{"portId":"8","serial":"Q234-ABCD-0008"}]}
swagger:model UpdateNetworkSwitchLinkAggregationBody
*/
type UpdateNetworkSwitchLinkAggregationBody struct {

	// Array of switch or stack ports for updating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchPorts []*UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0 `json:"switchPorts"`

	// Array of switch profile ports for updating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchProfilePorts []*UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0 `json:"switchProfilePorts"`
}

// Validate validates this update network switch link aggregation body
func (o *UpdateNetworkSwitchLinkAggregationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSwitchPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSwitchProfilePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchLinkAggregationBody) validateSwitchPorts(formats strfmt.Registry) error {
	if swag.IsZero(o.SwitchPorts) { // not required
		return nil
	}

	for i := 0; i < len(o.SwitchPorts); i++ {
		if swag.IsZero(o.SwitchPorts[i]) { // not required
			continue
		}

		if o.SwitchPorts[i] != nil {
			if err := o.SwitchPorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchLinkAggregation" + "." + "switchPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchLinkAggregation" + "." + "switchPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkSwitchLinkAggregationBody) validateSwitchProfilePorts(formats strfmt.Registry) error {
	if swag.IsZero(o.SwitchProfilePorts) { // not required
		return nil
	}

	for i := 0; i < len(o.SwitchProfilePorts); i++ {
		if swag.IsZero(o.SwitchProfilePorts[i]) { // not required
			continue
		}

		if o.SwitchProfilePorts[i] != nil {
			if err := o.SwitchProfilePorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchLinkAggregation" + "." + "switchProfilePorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchLinkAggregation" + "." + "switchProfilePorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network switch link aggregation body based on the context it is used
func (o *UpdateNetworkSwitchLinkAggregationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSwitchPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSwitchProfilePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchLinkAggregationBody) contextValidateSwitchPorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SwitchPorts); i++ {

		if o.SwitchPorts[i] != nil {
			if err := o.SwitchPorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchLinkAggregation" + "." + "switchPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchLinkAggregation" + "." + "switchPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkSwitchLinkAggregationBody) contextValidateSwitchProfilePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SwitchProfilePorts); i++ {

		if o.SwitchProfilePorts[i] != nil {
			if err := o.SwitchProfilePorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchLinkAggregation" + "." + "switchProfilePorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchLinkAggregation" + "." + "switchProfilePorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchLinkAggregationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchLinkAggregationBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchLinkAggregationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0 update network switch link aggregation params body switch ports items0
swagger:model UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0
*/
type UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0 struct {

	// Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	// Required: true
	PortID *string `json:"portId"`

	// Serial number of the switch.
	// Required: true
	Serial *string `json:"serial"`
}

// Validate validates this update network switch link aggregation params body switch ports items0
func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePortID(formats); err != nil {
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

func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) validatePortID(formats strfmt.Registry) error {

	if err := validate.Required("portId", "body", o.PortID); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", o.Serial); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch link aggregation params body switch ports items0 based on context it is used
func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0 update network switch link aggregation params body switch profile ports items0
swagger:model UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0
*/
type UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0 struct {

	// Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	// Required: true
	PortID *string `json:"portId"`

	// Profile identifier.
	// Required: true
	Profile *string `json:"profile"`
}

// Validate validates this update network switch link aggregation params body switch profile ports items0
func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePortID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) validatePortID(formats strfmt.Registry) error {

	if err := validate.Required("portId", "body", o.PortID); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) validateProfile(formats strfmt.Registry) error {

	if err := validate.Required("profile", "body", o.Profile); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch link aggregation params body switch profile ports items0 based on context it is used
func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
