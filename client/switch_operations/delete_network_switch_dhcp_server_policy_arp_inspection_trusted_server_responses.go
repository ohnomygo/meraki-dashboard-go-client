// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerReader is a Reader for the DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer structure.
type DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent creates a DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent with default headers values
func NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent() *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent {
	return &DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent{}
}

/*
DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent struct {
}

// IsSuccess returns true when this delete network switch dhcp server policy arp inspection trusted server no content response has a 2xx status code
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network switch dhcp server policy arp inspection trusted server no content response has a 3xx status code
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network switch dhcp server policy arp inspection trusted server no content response has a 4xx status code
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network switch dhcp server policy arp inspection trusted server no content response has a 5xx status code
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network switch dhcp server policy arp inspection trusted server no content response a status code equal to that given
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete network switch dhcp server policy arp inspection trusted server no content response
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) Code() int {
	return 204
}

func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}][%d] deleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent ", 204)
}

func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}][%d] deleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent ", 204)
}

func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
