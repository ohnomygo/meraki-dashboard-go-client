// Code generated by go-swagger; DO NOT EDIT.

package insight

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

// GetOrganizationInsightMonitoredMediaServersReader is a Reader for the GetOrganizationInsightMonitoredMediaServers structure.
type GetOrganizationInsightMonitoredMediaServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationInsightMonitoredMediaServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationInsightMonitoredMediaServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationInsightMonitoredMediaServersOK creates a GetOrganizationInsightMonitoredMediaServersOK with default headers values
func NewGetOrganizationInsightMonitoredMediaServersOK() *GetOrganizationInsightMonitoredMediaServersOK {
	return &GetOrganizationInsightMonitoredMediaServersOK{}
}

/*
GetOrganizationInsightMonitoredMediaServersOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationInsightMonitoredMediaServersOK struct {
	Payload []*GetOrganizationInsightMonitoredMediaServersOKBodyItems0
}

// IsSuccess returns true when this get organization insight monitored media servers o k response has a 2xx status code
func (o *GetOrganizationInsightMonitoredMediaServersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization insight monitored media servers o k response has a 3xx status code
func (o *GetOrganizationInsightMonitoredMediaServersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization insight monitored media servers o k response has a 4xx status code
func (o *GetOrganizationInsightMonitoredMediaServersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization insight monitored media servers o k response has a 5xx status code
func (o *GetOrganizationInsightMonitoredMediaServersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization insight monitored media servers o k response a status code equal to that given
func (o *GetOrganizationInsightMonitoredMediaServersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization insight monitored media servers o k response
func (o *GetOrganizationInsightMonitoredMediaServersOK) Code() int {
	return 200
}

func (o *GetOrganizationInsightMonitoredMediaServersOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/insight/monitoredMediaServers][%d] getOrganizationInsightMonitoredMediaServersOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInsightMonitoredMediaServersOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/insight/monitoredMediaServers][%d] getOrganizationInsightMonitoredMediaServersOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInsightMonitoredMediaServersOK) GetPayload() []*GetOrganizationInsightMonitoredMediaServersOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationInsightMonitoredMediaServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationInsightMonitoredMediaServersOKBodyItems0 get organization insight monitored media servers o k body items0
swagger:model GetOrganizationInsightMonitoredMediaServersOKBodyItems0
*/
type GetOrganizationInsightMonitoredMediaServersOKBodyItems0 struct {

	// The IP address (IPv4 only) or hostname of the media server to monitor
	Address string `json:"address,omitempty"`

	// Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead
	BestEffortMonitoringEnabled bool `json:"bestEffortMonitoringEnabled,omitempty"`

	// Monitored media server id
	ID string `json:"id,omitempty"`

	// The name of the VoIP provider
	Name string `json:"name,omitempty"`
}

// Validate validates this get organization insight monitored media servers o k body items0
func (o *GetOrganizationInsightMonitoredMediaServersOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization insight monitored media servers o k body items0 based on context it is used
func (o *GetOrganizationInsightMonitoredMediaServersOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationInsightMonitoredMediaServersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationInsightMonitoredMediaServersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationInsightMonitoredMediaServersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
