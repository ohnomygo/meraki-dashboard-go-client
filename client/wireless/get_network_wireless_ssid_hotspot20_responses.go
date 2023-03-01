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

// GetNetworkWirelessSsidHotspot20Reader is a Reader for the GetNetworkWirelessSsidHotspot20 structure.
type GetNetworkWirelessSsidHotspot20Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessSsidHotspot20Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessSsidHotspot20OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessSsidHotspot20OK creates a GetNetworkWirelessSsidHotspot20OK with default headers values
func NewGetNetworkWirelessSsidHotspot20OK() *GetNetworkWirelessSsidHotspot20OK {
	return &GetNetworkWirelessSsidHotspot20OK{}
}

/*
GetNetworkWirelessSsidHotspot20OK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessSsidHotspot20OK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network wireless ssid hotspot20 o k response has a 2xx status code
func (o *GetNetworkWirelessSsidHotspot20OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless ssid hotspot20 o k response has a 3xx status code
func (o *GetNetworkWirelessSsidHotspot20OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless ssid hotspot20 o k response has a 4xx status code
func (o *GetNetworkWirelessSsidHotspot20OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless ssid hotspot20 o k response has a 5xx status code
func (o *GetNetworkWirelessSsidHotspot20OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless ssid hotspot20 o k response a status code equal to that given
func (o *GetNetworkWirelessSsidHotspot20OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless ssid hotspot20 o k response
func (o *GetNetworkWirelessSsidHotspot20OK) Code() int {
	return 200
}

func (o *GetNetworkWirelessSsidHotspot20OK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/hotspot20][%d] getNetworkWirelessSsidHotspot20OK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidHotspot20OK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/hotspot20][%d] getNetworkWirelessSsidHotspot20OK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidHotspot20OK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkWirelessSsidHotspot20OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
