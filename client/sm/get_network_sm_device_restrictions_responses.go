// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSmDeviceRestrictionsReader is a Reader for the GetNetworkSmDeviceRestrictions structure.
type GetNetworkSmDeviceRestrictionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceRestrictionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceRestrictionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmDeviceRestrictionsOK creates a GetNetworkSmDeviceRestrictionsOK with default headers values
func NewGetNetworkSmDeviceRestrictionsOK() *GetNetworkSmDeviceRestrictionsOK {
	return &GetNetworkSmDeviceRestrictionsOK{}
}

/*
GetNetworkSmDeviceRestrictionsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmDeviceRestrictionsOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network sm device restrictions o k response has a 2xx status code
func (o *GetNetworkSmDeviceRestrictionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm device restrictions o k response has a 3xx status code
func (o *GetNetworkSmDeviceRestrictionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm device restrictions o k response has a 4xx status code
func (o *GetNetworkSmDeviceRestrictionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm device restrictions o k response has a 5xx status code
func (o *GetNetworkSmDeviceRestrictionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm device restrictions o k response a status code equal to that given
func (o *GetNetworkSmDeviceRestrictionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sm device restrictions o k response
func (o *GetNetworkSmDeviceRestrictionsOK) Code() int {
	return 200
}

func (o *GetNetworkSmDeviceRestrictionsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/restrictions][%d] getNetworkSmDeviceRestrictionsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceRestrictionsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/restrictions][%d] getNetworkSmDeviceRestrictionsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceRestrictionsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkSmDeviceRestrictionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
