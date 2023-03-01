// Code generated by go-swagger; DO NOT EDIT.

package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkCameraQualityRetentionProfileReader is a Reader for the DeleteNetworkCameraQualityRetentionProfile structure.
type DeleteNetworkCameraQualityRetentionProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkCameraQualityRetentionProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkCameraQualityRetentionProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkCameraQualityRetentionProfileNoContent creates a DeleteNetworkCameraQualityRetentionProfileNoContent with default headers values
func NewDeleteNetworkCameraQualityRetentionProfileNoContent() *DeleteNetworkCameraQualityRetentionProfileNoContent {
	return &DeleteNetworkCameraQualityRetentionProfileNoContent{}
}

/*
DeleteNetworkCameraQualityRetentionProfileNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkCameraQualityRetentionProfileNoContent struct {
}

// IsSuccess returns true when this delete network camera quality retention profile no content response has a 2xx status code
func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network camera quality retention profile no content response has a 3xx status code
func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network camera quality retention profile no content response has a 4xx status code
func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network camera quality retention profile no content response has a 5xx status code
func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network camera quality retention profile no content response a status code equal to that given
func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete network camera quality retention profile no content response
func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) Code() int {
	return 204
}

func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}][%d] deleteNetworkCameraQualityRetentionProfileNoContent ", 204)
}

func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}][%d] deleteNetworkCameraQualityRetentionProfileNoContent ", 204)
}

func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}