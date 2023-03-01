// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// GetNetworkSmDeviceConnectivityReader is a Reader for the GetNetworkSmDeviceConnectivity structure.
type GetNetworkSmDeviceConnectivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceConnectivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceConnectivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmDeviceConnectivityOK creates a GetNetworkSmDeviceConnectivityOK with default headers values
func NewGetNetworkSmDeviceConnectivityOK() *GetNetworkSmDeviceConnectivityOK {
	return &GetNetworkSmDeviceConnectivityOK{}
}

/*
GetNetworkSmDeviceConnectivityOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmDeviceConnectivityOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetNetworkSmDeviceConnectivityOKBodyItems0
}

// IsSuccess returns true when this get network sm device connectivity o k response has a 2xx status code
func (o *GetNetworkSmDeviceConnectivityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm device connectivity o k response has a 3xx status code
func (o *GetNetworkSmDeviceConnectivityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm device connectivity o k response has a 4xx status code
func (o *GetNetworkSmDeviceConnectivityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm device connectivity o k response has a 5xx status code
func (o *GetNetworkSmDeviceConnectivityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm device connectivity o k response a status code equal to that given
func (o *GetNetworkSmDeviceConnectivityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sm device connectivity o k response
func (o *GetNetworkSmDeviceConnectivityOK) Code() int {
	return 200
}

func (o *GetNetworkSmDeviceConnectivityOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/connectivity][%d] getNetworkSmDeviceConnectivityOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceConnectivityOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/connectivity][%d] getNetworkSmDeviceConnectivityOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceConnectivityOK) GetPayload() []*GetNetworkSmDeviceConnectivityOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSmDeviceConnectivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetNetworkSmDeviceConnectivityOKBodyItems0 get network sm device connectivity o k body items0
swagger:model GetNetworkSmDeviceConnectivityOKBodyItems0
*/
type GetNetworkSmDeviceConnectivityOKBodyItems0 struct {

	// When the device was first seen as connected to the internet in each connection.
	FirstSeenAt string `json:"firstSeenAt,omitempty"`

	// When the device was last seen as connected to the internet in each connection.
	LastSeenAt string `json:"lastSeenAt,omitempty"`
}

// Validate validates this get network sm device connectivity o k body items0
func (o *GetNetworkSmDeviceConnectivityOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network sm device connectivity o k body items0 based on context it is used
func (o *GetNetworkSmDeviceConnectivityOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSmDeviceConnectivityOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSmDeviceConnectivityOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSmDeviceConnectivityOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
