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

// GetNetworkApplianceFirewallL3FirewallRulesReader is a Reader for the GetNetworkApplianceFirewallL3FirewallRules structure.
type GetNetworkApplianceFirewallL3FirewallRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceFirewallL3FirewallRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceFirewallL3FirewallRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceFirewallL3FirewallRulesOK creates a GetNetworkApplianceFirewallL3FirewallRulesOK with default headers values
func NewGetNetworkApplianceFirewallL3FirewallRulesOK() *GetNetworkApplianceFirewallL3FirewallRulesOK {
	return &GetNetworkApplianceFirewallL3FirewallRulesOK{}
}

/*
GetNetworkApplianceFirewallL3FirewallRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceFirewallL3FirewallRulesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network appliance firewall l3 firewall rules o k response has a 2xx status code
func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance firewall l3 firewall rules o k response has a 3xx status code
func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance firewall l3 firewall rules o k response has a 4xx status code
func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance firewall l3 firewall rules o k response has a 5xx status code
func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance firewall l3 firewall rules o k response a status code equal to that given
func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network appliance firewall l3 firewall rules o k response
func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) Code() int {
	return 200
}

func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/firewall/l3FirewallRules][%d] getNetworkApplianceFirewallL3FirewallRulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/firewall/l3FirewallRules][%d] getNetworkApplianceFirewallL3FirewallRulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceFirewallL3FirewallRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
