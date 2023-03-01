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

// GetOrganizationAdaptivePolicySettingsReader is a Reader for the GetOrganizationAdaptivePolicySettings structure.
type GetOrganizationAdaptivePolicySettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationAdaptivePolicySettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationAdaptivePolicySettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationAdaptivePolicySettingsOK creates a GetOrganizationAdaptivePolicySettingsOK with default headers values
func NewGetOrganizationAdaptivePolicySettingsOK() *GetOrganizationAdaptivePolicySettingsOK {
	return &GetOrganizationAdaptivePolicySettingsOK{}
}

/*
GetOrganizationAdaptivePolicySettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationAdaptivePolicySettingsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization adaptive policy settings o k response has a 2xx status code
func (o *GetOrganizationAdaptivePolicySettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization adaptive policy settings o k response has a 3xx status code
func (o *GetOrganizationAdaptivePolicySettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization adaptive policy settings o k response has a 4xx status code
func (o *GetOrganizationAdaptivePolicySettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization adaptive policy settings o k response has a 5xx status code
func (o *GetOrganizationAdaptivePolicySettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization adaptive policy settings o k response a status code equal to that given
func (o *GetOrganizationAdaptivePolicySettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization adaptive policy settings o k response
func (o *GetOrganizationAdaptivePolicySettingsOK) Code() int {
	return 200
}

func (o *GetOrganizationAdaptivePolicySettingsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/settings][%d] getOrganizationAdaptivePolicySettingsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdaptivePolicySettingsOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/settings][%d] getOrganizationAdaptivePolicySettingsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdaptivePolicySettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationAdaptivePolicySettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
