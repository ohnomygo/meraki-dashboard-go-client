// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// CreateOrganizationConfigTemplateReader is a Reader for the CreateOrganizationConfigTemplate structure.
type CreateOrganizationConfigTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationConfigTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationConfigTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrganizationConfigTemplateCreated creates a CreateOrganizationConfigTemplateCreated with default headers values
func NewCreateOrganizationConfigTemplateCreated() *CreateOrganizationConfigTemplateCreated {
	return &CreateOrganizationConfigTemplateCreated{}
}

/*
CreateOrganizationConfigTemplateCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationConfigTemplateCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create organization config template created response has a 2xx status code
func (o *CreateOrganizationConfigTemplateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization config template created response has a 3xx status code
func (o *CreateOrganizationConfigTemplateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization config template created response has a 4xx status code
func (o *CreateOrganizationConfigTemplateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization config template created response has a 5xx status code
func (o *CreateOrganizationConfigTemplateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization config template created response a status code equal to that given
func (o *CreateOrganizationConfigTemplateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create organization config template created response
func (o *CreateOrganizationConfigTemplateCreated) Code() int {
	return 201
}

func (o *CreateOrganizationConfigTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/configTemplates][%d] createOrganizationConfigTemplateCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationConfigTemplateCreated) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/configTemplates][%d] createOrganizationConfigTemplateCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationConfigTemplateCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationConfigTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateOrganizationConfigTemplateBody create organization config template body
// Example: {"name":"My config template","timeZone":"America/Los_Angeles"}
swagger:model CreateOrganizationConfigTemplateBody
*/
type CreateOrganizationConfigTemplateBody struct {

	// The ID of the network or config template to copy configuration from
	CopyFromNetworkID string `json:"copyFromNetworkId,omitempty"`

	// The name of the configuration template
	// Required: true
	Name *string `json:"name"`

	// The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this create organization config template body
func (o *CreateOrganizationConfigTemplateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationConfigTemplateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationConfigTemplate"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization config template body based on context it is used
func (o *CreateOrganizationConfigTemplateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationConfigTemplateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationConfigTemplateBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationConfigTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
