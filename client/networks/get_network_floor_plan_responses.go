// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkFloorPlanReader is a Reader for the GetNetworkFloorPlan structure.
type GetNetworkFloorPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkFloorPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkFloorPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkFloorPlanOK creates a GetNetworkFloorPlanOK with default headers values
func NewGetNetworkFloorPlanOK() *GetNetworkFloorPlanOK {
	return &GetNetworkFloorPlanOK{}
}

/*
GetNetworkFloorPlanOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkFloorPlanOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network floor plan o k response has a 2xx status code
func (o *GetNetworkFloorPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network floor plan o k response has a 3xx status code
func (o *GetNetworkFloorPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network floor plan o k response has a 4xx status code
func (o *GetNetworkFloorPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network floor plan o k response has a 5xx status code
func (o *GetNetworkFloorPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network floor plan o k response a status code equal to that given
func (o *GetNetworkFloorPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network floor plan o k response
func (o *GetNetworkFloorPlanOK) Code() int {
	return 200
}

func (o *GetNetworkFloorPlanOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/floorPlans/{floorPlanId}][%d] getNetworkFloorPlanOK  %+v", 200, o.Payload)
}

func (o *GetNetworkFloorPlanOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/floorPlans/{floorPlanId}][%d] getNetworkFloorPlanOK  %+v", 200, o.Payload)
}

func (o *GetNetworkFloorPlanOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkFloorPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
