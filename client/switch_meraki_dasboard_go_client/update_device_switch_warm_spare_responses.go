// Code generated by go-swagger; DO NOT EDIT.

package switch_meraki_dasboard_go_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateDeviceSwitchWarmSpareReader is a Reader for the UpdateDeviceSwitchWarmSpare structure.
type UpdateDeviceSwitchWarmSpareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceSwitchWarmSpareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceSwitchWarmSpareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDeviceSwitchWarmSpareOK creates a UpdateDeviceSwitchWarmSpareOK with default headers values
func NewUpdateDeviceSwitchWarmSpareOK() *UpdateDeviceSwitchWarmSpareOK {
	return &UpdateDeviceSwitchWarmSpareOK{}
}

/*
UpdateDeviceSwitchWarmSpareOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceSwitchWarmSpareOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update device switch warm spare o k response has a 2xx status code
func (o *UpdateDeviceSwitchWarmSpareOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device switch warm spare o k response has a 3xx status code
func (o *UpdateDeviceSwitchWarmSpareOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device switch warm spare o k response has a 4xx status code
func (o *UpdateDeviceSwitchWarmSpareOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device switch warm spare o k response has a 5xx status code
func (o *UpdateDeviceSwitchWarmSpareOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device switch warm spare o k response a status code equal to that given
func (o *UpdateDeviceSwitchWarmSpareOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device switch warm spare o k response
func (o *UpdateDeviceSwitchWarmSpareOK) Code() int {
	return 200
}

func (o *UpdateDeviceSwitchWarmSpareOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/switch/warmSpare][%d] updateDeviceSwitchWarmSpareOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceSwitchWarmSpareOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/switch/warmSpare][%d] updateDeviceSwitchWarmSpareOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceSwitchWarmSpareOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceSwitchWarmSpareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceSwitchWarmSpareBody update device switch warm spare body
// Example: {"enabled":true,"spareSerial":"Q234-ABCD-0002"}
swagger:model UpdateDeviceSwitchWarmSpareBody
*/
type UpdateDeviceSwitchWarmSpareBody struct {

	// Enable or disable warm spare for a switch
	// Required: true
	Enabled *bool `json:"enabled"`

	// Serial number of the warm spare switch
	SpareSerial string `json:"spareSerial,omitempty"`
}

// Validate validates this update device switch warm spare body
func (o *UpdateDeviceSwitchWarmSpareBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceSwitchWarmSpareBody) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("updateDeviceSwitchWarmSpare"+"."+"enabled", "body", o.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device switch warm spare body based on context it is used
func (o *UpdateDeviceSwitchWarmSpareBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchWarmSpareBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchWarmSpareBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchWarmSpareBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}