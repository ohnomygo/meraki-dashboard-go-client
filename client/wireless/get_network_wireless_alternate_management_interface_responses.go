// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkWirelessAlternateManagementInterfaceReader is a Reader for the GetNetworkWirelessAlternateManagementInterface structure.
type GetNetworkWirelessAlternateManagementInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessAlternateManagementInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessAlternateManagementInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessAlternateManagementInterfaceOK creates a GetNetworkWirelessAlternateManagementInterfaceOK with default headers values
func NewGetNetworkWirelessAlternateManagementInterfaceOK() *GetNetworkWirelessAlternateManagementInterfaceOK {
	return &GetNetworkWirelessAlternateManagementInterfaceOK{}
}

/*
GetNetworkWirelessAlternateManagementInterfaceOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessAlternateManagementInterfaceOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network wireless alternate management interface o k response has a 2xx status code
func (o *GetNetworkWirelessAlternateManagementInterfaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless alternate management interface o k response has a 3xx status code
func (o *GetNetworkWirelessAlternateManagementInterfaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless alternate management interface o k response has a 4xx status code
func (o *GetNetworkWirelessAlternateManagementInterfaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless alternate management interface o k response has a 5xx status code
func (o *GetNetworkWirelessAlternateManagementInterfaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless alternate management interface o k response a status code equal to that given
func (o *GetNetworkWirelessAlternateManagementInterfaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless alternate management interface o k response
func (o *GetNetworkWirelessAlternateManagementInterfaceOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessAlternateManagementInterfaceOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/alternateManagementInterface][%d] getNetworkWirelessAlternateManagementInterfaceOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessAlternateManagementInterfaceOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/alternateManagementInterface][%d] getNetworkWirelessAlternateManagementInterfaceOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessAlternateManagementInterfaceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkWirelessAlternateManagementInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}