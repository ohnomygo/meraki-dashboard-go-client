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

// CreateNetworkSwitchLinkAggregationReader is a Reader for the CreateNetworkSwitchLinkAggregation structure.
type CreateNetworkSwitchLinkAggregationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchLinkAggregationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchLinkAggregationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkSwitchLinkAggregationCreated creates a CreateNetworkSwitchLinkAggregationCreated with default headers values
func NewCreateNetworkSwitchLinkAggregationCreated() *CreateNetworkSwitchLinkAggregationCreated {
	return &CreateNetworkSwitchLinkAggregationCreated{}
}

/*
CreateNetworkSwitchLinkAggregationCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchLinkAggregationCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network switch link aggregation created response has a 2xx status code
func (o *CreateNetworkSwitchLinkAggregationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network switch link aggregation created response has a 3xx status code
func (o *CreateNetworkSwitchLinkAggregationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network switch link aggregation created response has a 4xx status code
func (o *CreateNetworkSwitchLinkAggregationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network switch link aggregation created response has a 5xx status code
func (o *CreateNetworkSwitchLinkAggregationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network switch link aggregation created response a status code equal to that given
func (o *CreateNetworkSwitchLinkAggregationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network switch link aggregation created response
func (o *CreateNetworkSwitchLinkAggregationCreated) Code() int {
	return 201
}

func (o *CreateNetworkSwitchLinkAggregationCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/linkAggregations][%d] createNetworkSwitchLinkAggregationCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchLinkAggregationCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/linkAggregations][%d] createNetworkSwitchLinkAggregationCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchLinkAggregationCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchLinkAggregationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkSwitchLinkAggregationBody create network switch link aggregation body
// Example: {"switchPorts":[{"portId":"1","serial":"Q234-ABCD-0001"},{"portId":"2","serial":"Q234-ABCD-0002"},{"portId":"3","serial":"Q234-ABCD-0003"},{"portId":"4","serial":"Q234-ABCD-0004"},{"portId":"5","serial":"Q234-ABCD-0005"},{"portId":"6","serial":"Q234-ABCD-0006"},{"portId":"7","serial":"Q234-ABCD-0007"},{"portId":"8","serial":"Q234-ABCD-0008"}]}
swagger:model CreateNetworkSwitchLinkAggregationBody
*/
type CreateNetworkSwitchLinkAggregationBody struct {

	// Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchPorts []*CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0 `json:"switchPorts"`

	// Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchProfilePorts []*CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0 `json:"switchProfilePorts"`
}

// Validate validates this create network switch link aggregation body
func (o *CreateNetworkSwitchLinkAggregationBody) Validate(formats strfmt.Registry) error {
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

func (o *CreateNetworkSwitchLinkAggregationBody) validateSwitchPorts(formats strfmt.Registry) error {
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
					return ve.ValidateName("createNetworkSwitchLinkAggregation" + "." + "switchPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkSwitchLinkAggregation" + "." + "switchPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkSwitchLinkAggregationBody) validateSwitchProfilePorts(formats strfmt.Registry) error {
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
					return ve.ValidateName("createNetworkSwitchLinkAggregation" + "." + "switchProfilePorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkSwitchLinkAggregation" + "." + "switchProfilePorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create network switch link aggregation body based on the context it is used
func (o *CreateNetworkSwitchLinkAggregationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *CreateNetworkSwitchLinkAggregationBody) contextValidateSwitchPorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SwitchPorts); i++ {

		if o.SwitchPorts[i] != nil {
			if err := o.SwitchPorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkSwitchLinkAggregation" + "." + "switchPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkSwitchLinkAggregation" + "." + "switchPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkSwitchLinkAggregationBody) contextValidateSwitchProfilePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SwitchProfilePorts); i++ {

		if o.SwitchProfilePorts[i] != nil {
			if err := o.SwitchProfilePorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkSwitchLinkAggregation" + "." + "switchProfilePorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkSwitchLinkAggregation" + "." + "switchProfilePorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchLinkAggregationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchLinkAggregationBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchLinkAggregationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0 create network switch link aggregation params body switch ports items0
swagger:model CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0
*/
type CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0 struct {

	// Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	// Required: true
	PortID *string `json:"portId"`

	// Serial number of the switch.
	// Required: true
	Serial *string `json:"serial"`
}

// Validate validates this create network switch link aggregation params body switch ports items0
func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) Validate(formats strfmt.Registry) error {
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

func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) validatePortID(formats strfmt.Registry) error {

	if err := validate.Required("portId", "body", o.PortID); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", o.Serial); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network switch link aggregation params body switch ports items0 based on context it is used
func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchLinkAggregationParamsBodySwitchPortsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0 create network switch link aggregation params body switch profile ports items0
swagger:model CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0
*/
type CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0 struct {

	// Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	// Required: true
	PortID *string `json:"portId"`

	// Profile identifier.
	// Required: true
	Profile *string `json:"profile"`
}

// Validate validates this create network switch link aggregation params body switch profile ports items0
func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) Validate(formats strfmt.Registry) error {
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

func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) validatePortID(formats strfmt.Registry) error {

	if err := validate.Required("portId", "body", o.PortID); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) validateProfile(formats strfmt.Registry) error {

	if err := validate.Required("profile", "body", o.Profile); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network switch link aggregation params body switch profile ports items0 based on context it is used
func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchLinkAggregationParamsBodySwitchProfilePortsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
