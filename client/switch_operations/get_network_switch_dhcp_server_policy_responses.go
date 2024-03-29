// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSwitchDhcpServerPolicyReader is a Reader for the GetNetworkSwitchDhcpServerPolicy structure.
type GetNetworkSwitchDhcpServerPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchDhcpServerPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchDhcpServerPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSwitchDhcpServerPolicyOK creates a GetNetworkSwitchDhcpServerPolicyOK with default headers values
func NewGetNetworkSwitchDhcpServerPolicyOK() *GetNetworkSwitchDhcpServerPolicyOK {
	return &GetNetworkSwitchDhcpServerPolicyOK{}
}

/*
GetNetworkSwitchDhcpServerPolicyOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchDhcpServerPolicyOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network switch dhcp server policy o k response has a 2xx status code
func (o *GetNetworkSwitchDhcpServerPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch dhcp server policy o k response has a 3xx status code
func (o *GetNetworkSwitchDhcpServerPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch dhcp server policy o k response has a 4xx status code
func (o *GetNetworkSwitchDhcpServerPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch dhcp server policy o k response has a 5xx status code
func (o *GetNetworkSwitchDhcpServerPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch dhcp server policy o k response a status code equal to that given
func (o *GetNetworkSwitchDhcpServerPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network switch dhcp server policy o k response
func (o *GetNetworkSwitchDhcpServerPolicyOK) Code() int {
	return 200
}

func (o *GetNetworkSwitchDhcpServerPolicyOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/dhcpServerPolicy][%d] getNetworkSwitchDhcpServerPolicyOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchDhcpServerPolicyOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/dhcpServerPolicy][%d] getNetworkSwitchDhcpServerPolicyOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchDhcpServerPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSwitchDhcpServerPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
