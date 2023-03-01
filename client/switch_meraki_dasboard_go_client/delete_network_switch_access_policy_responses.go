// Code generated by go-swagger; DO NOT EDIT.

package switch_meraki_dasboard_go_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkSwitchAccessPolicyReader is a Reader for the DeleteNetworkSwitchAccessPolicy structure.
type DeleteNetworkSwitchAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkSwitchAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkSwitchAccessPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkSwitchAccessPolicyNoContent creates a DeleteNetworkSwitchAccessPolicyNoContent with default headers values
func NewDeleteNetworkSwitchAccessPolicyNoContent() *DeleteNetworkSwitchAccessPolicyNoContent {
	return &DeleteNetworkSwitchAccessPolicyNoContent{}
}

/*
DeleteNetworkSwitchAccessPolicyNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkSwitchAccessPolicyNoContent struct {
}

// IsSuccess returns true when this delete network switch access policy no content response has a 2xx status code
func (o *DeleteNetworkSwitchAccessPolicyNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network switch access policy no content response has a 3xx status code
func (o *DeleteNetworkSwitchAccessPolicyNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network switch access policy no content response has a 4xx status code
func (o *DeleteNetworkSwitchAccessPolicyNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network switch access policy no content response has a 5xx status code
func (o *DeleteNetworkSwitchAccessPolicyNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network switch access policy no content response a status code equal to that given
func (o *DeleteNetworkSwitchAccessPolicyNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete network switch access policy no content response
func (o *DeleteNetworkSwitchAccessPolicyNoContent) Code() int {
	return 204
}

func (o *DeleteNetworkSwitchAccessPolicyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}][%d] deleteNetworkSwitchAccessPolicyNoContent ", 204)
}

func (o *DeleteNetworkSwitchAccessPolicyNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}][%d] deleteNetworkSwitchAccessPolicyNoContent ", 204)
}

func (o *DeleteNetworkSwitchAccessPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}