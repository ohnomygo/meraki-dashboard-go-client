// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UnbindNetworkReader is a Reader for the UnbindNetwork structure.
type UnbindNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnbindNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnbindNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnbindNetworkOK creates a UnbindNetworkOK with default headers values
func NewUnbindNetworkOK() *UnbindNetworkOK {
	return &UnbindNetworkOK{}
}

/*
UnbindNetworkOK describes a response with status code 200, with default header values.

Successful operation
*/
type UnbindNetworkOK struct {
	Payload *UnbindNetworkOKBody
}

// IsSuccess returns true when this unbind network o k response has a 2xx status code
func (o *UnbindNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unbind network o k response has a 3xx status code
func (o *UnbindNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind network o k response has a 4xx status code
func (o *UnbindNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unbind network o k response has a 5xx status code
func (o *UnbindNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unbind network o k response a status code equal to that given
func (o *UnbindNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unbind network o k response
func (o *UnbindNetworkOK) Code() int {
	return 200
}

func (o *UnbindNetworkOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/unbind][%d] unbindNetworkOK  %+v", 200, o.Payload)
}

func (o *UnbindNetworkOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/unbind][%d] unbindNetworkOK  %+v", 200, o.Payload)
}

func (o *UnbindNetworkOK) GetPayload() *UnbindNetworkOKBody {
	return o.Payload
}

func (o *UnbindNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UnbindNetworkOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UnbindNetworkBody unbind network body
// Example: {"retainConfigs":true}
swagger:model UnbindNetworkBody
*/
type UnbindNetworkBody struct {

	// Optional boolean to retain all the current configs given by the template.
	RetainConfigs bool `json:"retainConfigs,omitempty"`
}

// Validate validates this unbind network body
func (o *UnbindNetworkBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unbind network body based on context it is used
func (o *UnbindNetworkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnbindNetworkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnbindNetworkBody) UnmarshalBinary(b []byte) error {
	var res UnbindNetworkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UnbindNetworkOKBody unbind network o k body
swagger:model UnbindNetworkOKBody
*/
type UnbindNetworkOKBody struct {

	// Enrollment string for the network
	EnrollmentString string `json:"enrollmentString,omitempty"`

	// Network ID
	ID string `json:"id,omitempty"`

	// If the network is bound to a config template
	IsBoundToConfigTemplate bool `json:"isBoundToConfigTemplate,omitempty"`

	// Network name
	Name string `json:"name,omitempty"`

	// Notes for the network
	Notes string `json:"notes,omitempty"`

	// Organization ID
	OrganizationID string `json:"organizationId,omitempty"`

	// List of the product types that the network supports
	ProductTypes []string `json:"productTypes"`

	// Network tags
	Tags []string `json:"tags"`

	// Timezone of the network
	TimeZone string `json:"timeZone,omitempty"`

	// URL to the network Dashboard UI
	URL string `json:"url,omitempty"`
}

// Validate validates this unbind network o k body
func (o *UnbindNetworkOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unbind network o k body based on context it is used
func (o *UnbindNetworkOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnbindNetworkOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnbindNetworkOKBody) UnmarshalBinary(b []byte) error {
	var res UnbindNetworkOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
