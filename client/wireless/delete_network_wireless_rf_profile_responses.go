// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkWirelessRfProfileReader is a Reader for the DeleteNetworkWirelessRfProfile structure.
type DeleteNetworkWirelessRfProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkWirelessRfProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkWirelessRfProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkWirelessRfProfileNoContent creates a DeleteNetworkWirelessRfProfileNoContent with default headers values
func NewDeleteNetworkWirelessRfProfileNoContent() *DeleteNetworkWirelessRfProfileNoContent {
	return &DeleteNetworkWirelessRfProfileNoContent{}
}

/*
DeleteNetworkWirelessRfProfileNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkWirelessRfProfileNoContent struct {
}

// IsSuccess returns true when this delete network wireless rf profile no content response has a 2xx status code
func (o *DeleteNetworkWirelessRfProfileNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network wireless rf profile no content response has a 3xx status code
func (o *DeleteNetworkWirelessRfProfileNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network wireless rf profile no content response has a 4xx status code
func (o *DeleteNetworkWirelessRfProfileNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network wireless rf profile no content response has a 5xx status code
func (o *DeleteNetworkWirelessRfProfileNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network wireless rf profile no content response a status code equal to that given
func (o *DeleteNetworkWirelessRfProfileNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete network wireless rf profile no content response
func (o *DeleteNetworkWirelessRfProfileNoContent) Code() int {
	return 204
}

func (o *DeleteNetworkWirelessRfProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/wireless/rfProfiles/{rfProfileId}][%d] deleteNetworkWirelessRfProfileNoContent ", 204)
}

func (o *DeleteNetworkWirelessRfProfileNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/wireless/rfProfiles/{rfProfileId}][%d] deleteNetworkWirelessRfProfileNoContent ", 204)
}

func (o *DeleteNetworkWirelessRfProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}