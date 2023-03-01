// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationApplianceVpnStatusesReader is a Reader for the GetOrganizationApplianceVpnStatuses structure.
type GetOrganizationApplianceVpnStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationApplianceVpnStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationApplianceVpnStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationApplianceVpnStatusesOK creates a GetOrganizationApplianceVpnStatusesOK with default headers values
func NewGetOrganizationApplianceVpnStatusesOK() *GetOrganizationApplianceVpnStatusesOK {
	return &GetOrganizationApplianceVpnStatusesOK{}
}

/*
GetOrganizationApplianceVpnStatusesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationApplianceVpnStatusesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

// IsSuccess returns true when this get organization appliance vpn statuses o k response has a 2xx status code
func (o *GetOrganizationApplianceVpnStatusesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization appliance vpn statuses o k response has a 3xx status code
func (o *GetOrganizationApplianceVpnStatusesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization appliance vpn statuses o k response has a 4xx status code
func (o *GetOrganizationApplianceVpnStatusesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization appliance vpn statuses o k response has a 5xx status code
func (o *GetOrganizationApplianceVpnStatusesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization appliance vpn statuses o k response a status code equal to that given
func (o *GetOrganizationApplianceVpnStatusesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization appliance vpn statuses o k response
func (o *GetOrganizationApplianceVpnStatusesOK) Code() int {
	return 200
}

func (o *GetOrganizationApplianceVpnStatusesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/appliance/vpn/statuses][%d] getOrganizationApplianceVpnStatusesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationApplianceVpnStatusesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/appliance/vpn/statuses][%d] getOrganizationApplianceVpnStatusesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationApplianceVpnStatusesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationApplianceVpnStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
