// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationAdaptivePolicyGroupReader is a Reader for the DeleteOrganizationAdaptivePolicyGroup structure.
type DeleteOrganizationAdaptivePolicyGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationAdaptivePolicyGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationAdaptivePolicyGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationAdaptivePolicyGroupNoContent creates a DeleteOrganizationAdaptivePolicyGroupNoContent with default headers values
func NewDeleteOrganizationAdaptivePolicyGroupNoContent() *DeleteOrganizationAdaptivePolicyGroupNoContent {
	return &DeleteOrganizationAdaptivePolicyGroupNoContent{}
}

/*
DeleteOrganizationAdaptivePolicyGroupNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteOrganizationAdaptivePolicyGroupNoContent struct {
}

// IsSuccess returns true when this delete organization adaptive policy group no content response has a 2xx status code
func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization adaptive policy group no content response has a 3xx status code
func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization adaptive policy group no content response has a 4xx status code
func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization adaptive policy group no content response has a 5xx status code
func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization adaptive policy group no content response a status code equal to that given
func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organization adaptive policy group no content response
func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/adaptivePolicy/groups/{id}][%d] deleteOrganizationAdaptivePolicyGroupNoContent ", 204)
}

func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/adaptivePolicy/groups/{id}][%d] deleteOrganizationAdaptivePolicyGroupNoContent ", 204)
}

func (o *DeleteOrganizationAdaptivePolicyGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}