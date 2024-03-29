// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// UpdateNetworkSwitchStpReader is a Reader for the UpdateNetworkSwitchStp structure.
type UpdateNetworkSwitchStpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchStpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchStpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSwitchStpOK creates a UpdateNetworkSwitchStpOK with default headers values
func NewUpdateNetworkSwitchStpOK() *UpdateNetworkSwitchStpOK {
	return &UpdateNetworkSwitchStpOK{}
}

/*
UpdateNetworkSwitchStpOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchStpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network switch stp o k response has a 2xx status code
func (o *UpdateNetworkSwitchStpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network switch stp o k response has a 3xx status code
func (o *UpdateNetworkSwitchStpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network switch stp o k response has a 4xx status code
func (o *UpdateNetworkSwitchStpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network switch stp o k response has a 5xx status code
func (o *UpdateNetworkSwitchStpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network switch stp o k response a status code equal to that given
func (o *UpdateNetworkSwitchStpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network switch stp o k response
func (o *UpdateNetworkSwitchStpOK) Code() int {
	return 200
}

func (o *UpdateNetworkSwitchStpOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/stp][%d] updateNetworkSwitchStpOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchStpOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/stp][%d] updateNetworkSwitchStpOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchStpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchStpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkSwitchStpBody update network switch stp body
// Example: {"rstpEnabled":true,"stpBridgePriority":[{"stpPriority":4096,"switches":["Q234-ABCD-0001","Q234-ABCD-0002","Q234-ABCD-0003"]},{"stacks":["789102","123456","129102"],"stpPriority":28672}]}
swagger:model UpdateNetworkSwitchStpBody
*/
type UpdateNetworkSwitchStpBody struct {

	// The spanning tree protocol status in network
	RstpEnabled bool `json:"rstpEnabled,omitempty"`

	// STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings.
	StpBridgePriority []*UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0 `json:"stpBridgePriority"`
}

// Validate validates this update network switch stp body
func (o *UpdateNetworkSwitchStpBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStpBridgePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchStpBody) validateStpBridgePriority(formats strfmt.Registry) error {
	if swag.IsZero(o.StpBridgePriority) { // not required
		return nil
	}

	for i := 0; i < len(o.StpBridgePriority); i++ {
		if swag.IsZero(o.StpBridgePriority[i]) { // not required
			continue
		}

		if o.StpBridgePriority[i] != nil {
			if err := o.StpBridgePriority[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchStp" + "." + "stpBridgePriority" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchStp" + "." + "stpBridgePriority" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network switch stp body based on the context it is used
func (o *UpdateNetworkSwitchStpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateStpBridgePriority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchStpBody) contextValidateStpBridgePriority(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.StpBridgePriority); i++ {

		if o.StpBridgePriority[i] != nil {
			if err := o.StpBridgePriority[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchStp" + "." + "stpBridgePriority" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchStp" + "." + "stpBridgePriority" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchStpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchStpBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchStpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0 update network switch stp params body stp bridge priority items0
swagger:model UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0
*/
type UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0 struct {

	// List of stack IDs
	Stacks []string `json:"stacks"`

	// STP priority for switch, stacks, or switch profiles
	// Required: true
	StpPriority *int64 `json:"stpPriority"`

	// List of switch profile IDs
	SwitchProfiles []string `json:"switchProfiles"`

	// List of switch serial numbers
	Switches []string `json:"switches"`
}

// Validate validates this update network switch stp params body stp bridge priority items0
func (o *UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStpPriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0) validateStpPriority(formats strfmt.Registry) error {

	if err := validate.Required("stpPriority", "body", o.StpPriority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch stp params body stp bridge priority items0 based on context it is used
func (o *UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchStpParamsBodyStpBridgePriorityItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
