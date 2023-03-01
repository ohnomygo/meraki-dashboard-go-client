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

// GetOrganizationConfigurationChangesReader is a Reader for the GetOrganizationConfigurationChanges structure.
type GetOrganizationConfigurationChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationConfigurationChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationConfigurationChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationConfigurationChangesOK creates a GetOrganizationConfigurationChangesOK with default headers values
func NewGetOrganizationConfigurationChangesOK() *GetOrganizationConfigurationChangesOK {
	return &GetOrganizationConfigurationChangesOK{}
}

/*
GetOrganizationConfigurationChangesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationConfigurationChangesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

// IsSuccess returns true when this get organization configuration changes o k response has a 2xx status code
func (o *GetOrganizationConfigurationChangesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization configuration changes o k response has a 3xx status code
func (o *GetOrganizationConfigurationChangesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization configuration changes o k response has a 4xx status code
func (o *GetOrganizationConfigurationChangesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization configuration changes o k response has a 5xx status code
func (o *GetOrganizationConfigurationChangesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization configuration changes o k response a status code equal to that given
func (o *GetOrganizationConfigurationChangesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization configuration changes o k response
func (o *GetOrganizationConfigurationChangesOK) Code() int {
	return 200
}

func (o *GetOrganizationConfigurationChangesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/configurationChanges][%d] getOrganizationConfigurationChangesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationConfigurationChangesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/configurationChanges][%d] getOrganizationConfigurationChangesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationConfigurationChangesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationConfigurationChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
