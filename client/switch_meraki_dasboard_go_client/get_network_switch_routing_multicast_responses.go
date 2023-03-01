// Code generated by go-swagger; DO NOT EDIT.

package switch_meraki_dasboard_go_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSwitchRoutingMulticastReader is a Reader for the GetNetworkSwitchRoutingMulticast structure.
type GetNetworkSwitchRoutingMulticastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchRoutingMulticastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchRoutingMulticastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSwitchRoutingMulticastOK creates a GetNetworkSwitchRoutingMulticastOK with default headers values
func NewGetNetworkSwitchRoutingMulticastOK() *GetNetworkSwitchRoutingMulticastOK {
	return &GetNetworkSwitchRoutingMulticastOK{}
}

/*
GetNetworkSwitchRoutingMulticastOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchRoutingMulticastOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network switch routing multicast o k response has a 2xx status code
func (o *GetNetworkSwitchRoutingMulticastOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch routing multicast o k response has a 3xx status code
func (o *GetNetworkSwitchRoutingMulticastOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch routing multicast o k response has a 4xx status code
func (o *GetNetworkSwitchRoutingMulticastOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch routing multicast o k response has a 5xx status code
func (o *GetNetworkSwitchRoutingMulticastOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch routing multicast o k response a status code equal to that given
func (o *GetNetworkSwitchRoutingMulticastOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network switch routing multicast o k response
func (o *GetNetworkSwitchRoutingMulticastOK) Code() int {
	return 200
}

func (o *GetNetworkSwitchRoutingMulticastOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/routing/multicast][%d] getNetworkSwitchRoutingMulticastOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchRoutingMulticastOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/routing/multicast][%d] getNetworkSwitchRoutingMulticastOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchRoutingMulticastOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSwitchRoutingMulticastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
