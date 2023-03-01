// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationReader is a Reader for the DeleteOrganization structure.
type DeleteOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationNoContent creates a DeleteOrganizationNoContent with default headers values
func NewDeleteOrganizationNoContent() *DeleteOrganizationNoContent {
	return &DeleteOrganizationNoContent{}
}

/*
DeleteOrganizationNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteOrganizationNoContent struct {
}

// IsSuccess returns true when this delete organization no content response has a 2xx status code
func (o *DeleteOrganizationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization no content response has a 3xx status code
func (o *DeleteOrganizationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization no content response has a 4xx status code
func (o *DeleteOrganizationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization no content response has a 5xx status code
func (o *DeleteOrganizationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization no content response a status code equal to that given
func (o *DeleteOrganizationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organization no content response
func (o *DeleteOrganizationNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}][%d] deleteOrganizationNoContent ", 204)
}

func (o *DeleteOrganizationNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}][%d] deleteOrganizationNoContent ", 204)
}

func (o *DeleteOrganizationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
