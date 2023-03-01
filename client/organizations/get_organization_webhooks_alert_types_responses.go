// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationWebhooksAlertTypesReader is a Reader for the GetOrganizationWebhooksAlertTypes structure.
type GetOrganizationWebhooksAlertTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationWebhooksAlertTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationWebhooksAlertTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationWebhooksAlertTypesOK creates a GetOrganizationWebhooksAlertTypesOK with default headers values
func NewGetOrganizationWebhooksAlertTypesOK() *GetOrganizationWebhooksAlertTypesOK {
	return &GetOrganizationWebhooksAlertTypesOK{}
}

/*
GetOrganizationWebhooksAlertTypesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationWebhooksAlertTypesOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get organization webhooks alert types o k response has a 2xx status code
func (o *GetOrganizationWebhooksAlertTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization webhooks alert types o k response has a 3xx status code
func (o *GetOrganizationWebhooksAlertTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization webhooks alert types o k response has a 4xx status code
func (o *GetOrganizationWebhooksAlertTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization webhooks alert types o k response has a 5xx status code
func (o *GetOrganizationWebhooksAlertTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization webhooks alert types o k response a status code equal to that given
func (o *GetOrganizationWebhooksAlertTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization webhooks alert types o k response
func (o *GetOrganizationWebhooksAlertTypesOK) Code() int {
	return 200
}

func (o *GetOrganizationWebhooksAlertTypesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/webhooks/alertTypes][%d] getOrganizationWebhooksAlertTypesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationWebhooksAlertTypesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/webhooks/alertTypes][%d] getOrganizationWebhooksAlertTypesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationWebhooksAlertTypesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationWebhooksAlertTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}