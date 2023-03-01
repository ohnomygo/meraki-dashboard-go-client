// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkPiiRequestReader is a Reader for the DeleteNetworkPiiRequest structure.
type DeleteNetworkPiiRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkPiiRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkPiiRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkPiiRequestNoContent creates a DeleteNetworkPiiRequestNoContent with default headers values
func NewDeleteNetworkPiiRequestNoContent() *DeleteNetworkPiiRequestNoContent {
	return &DeleteNetworkPiiRequestNoContent{}
}

/*
DeleteNetworkPiiRequestNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkPiiRequestNoContent struct {
}

// IsSuccess returns true when this delete network pii request no content response has a 2xx status code
func (o *DeleteNetworkPiiRequestNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network pii request no content response has a 3xx status code
func (o *DeleteNetworkPiiRequestNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network pii request no content response has a 4xx status code
func (o *DeleteNetworkPiiRequestNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network pii request no content response has a 5xx status code
func (o *DeleteNetworkPiiRequestNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network pii request no content response a status code equal to that given
func (o *DeleteNetworkPiiRequestNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete network pii request no content response
func (o *DeleteNetworkPiiRequestNoContent) Code() int {
	return 204
}

func (o *DeleteNetworkPiiRequestNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/pii/requests/{requestId}][%d] deleteNetworkPiiRequestNoContent ", 204)
}

func (o *DeleteNetworkPiiRequestNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/pii/requests/{requestId}][%d] deleteNetworkPiiRequestNoContent ", 204)
}

func (o *DeleteNetworkPiiRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
