// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// BindNetworkReader is a Reader for the BindNetwork structure.
type BindNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBindNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBindNetworkOK creates a BindNetworkOK with default headers values
func NewBindNetworkOK() *BindNetworkOK {
	return &BindNetworkOK{}
}

/*
BindNetworkOK describes a response with status code 200, with default header values.

Successful operation
*/
type BindNetworkOK struct {
	Payload interface{}
}

// IsSuccess returns true when this bind network o k response has a 2xx status code
func (o *BindNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bind network o k response has a 3xx status code
func (o *BindNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bind network o k response has a 4xx status code
func (o *BindNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bind network o k response has a 5xx status code
func (o *BindNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bind network o k response a status code equal to that given
func (o *BindNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bind network o k response
func (o *BindNetworkOK) Code() int {
	return 200
}

func (o *BindNetworkOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/bind][%d] bindNetworkOK  %+v", 200, o.Payload)
}

func (o *BindNetworkOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/bind][%d] bindNetworkOK  %+v", 200, o.Payload)
}

func (o *BindNetworkOK) GetPayload() interface{} {
	return o.Payload
}

func (o *BindNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
BindNetworkBody bind network body
// Example: {"autoBind":false,"configTemplateId":"N_23952905"}
swagger:model BindNetworkBody
*/
type BindNetworkBody struct {

	// Optional boolean indicating whether the network's switches should automatically bind to profiles of the same model. Defaults to false if left unspecified. This option only affects switch networks and switch templates. Auto-bind is not valid unless the switch template has at least one profile and has at most one profile per switch model.
	AutoBind bool `json:"autoBind,omitempty"`

	// The ID of the template to which the network should be bound.
	// Required: true
	ConfigTemplateID *string `json:"configTemplateId"`
}

// Validate validates this bind network body
func (o *BindNetworkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfigTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BindNetworkBody) validateConfigTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("bindNetwork"+"."+"configTemplateId", "body", o.ConfigTemplateID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bind network body based on context it is used
func (o *BindNetworkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BindNetworkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BindNetworkBody) UnmarshalBinary(b []byte) error {
	var res BindNetworkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
