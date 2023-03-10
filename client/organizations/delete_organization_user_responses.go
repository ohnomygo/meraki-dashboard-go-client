// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationUserReader is a Reader for the DeleteOrganizationUser structure.
type DeleteOrganizationUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationUserNoContent creates a DeleteOrganizationUserNoContent with default headers values
func NewDeleteOrganizationUserNoContent() *DeleteOrganizationUserNoContent {
	return &DeleteOrganizationUserNoContent{}
}

/*
DeleteOrganizationUserNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteOrganizationUserNoContent struct {
}

// IsSuccess returns true when this delete organization user no content response has a 2xx status code
func (o *DeleteOrganizationUserNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization user no content response has a 3xx status code
func (o *DeleteOrganizationUserNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization user no content response has a 4xx status code
func (o *DeleteOrganizationUserNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization user no content response has a 5xx status code
func (o *DeleteOrganizationUserNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization user no content response a status code equal to that given
func (o *DeleteOrganizationUserNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organization user no content response
func (o *DeleteOrganizationUserNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/users/{userId}][%d] deleteOrganizationUserNoContent ", 204)
}

func (o *DeleteOrganizationUserNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/users/{userId}][%d] deleteOrganizationUserNoContent ", 204)
}

func (o *DeleteOrganizationUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
