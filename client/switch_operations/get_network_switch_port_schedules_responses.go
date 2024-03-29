// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSwitchPortSchedulesReader is a Reader for the GetNetworkSwitchPortSchedules structure.
type GetNetworkSwitchPortSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchPortSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchPortSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSwitchPortSchedulesOK creates a GetNetworkSwitchPortSchedulesOK with default headers values
func NewGetNetworkSwitchPortSchedulesOK() *GetNetworkSwitchPortSchedulesOK {
	return &GetNetworkSwitchPortSchedulesOK{}
}

/*
GetNetworkSwitchPortSchedulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchPortSchedulesOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network switch port schedules o k response has a 2xx status code
func (o *GetNetworkSwitchPortSchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch port schedules o k response has a 3xx status code
func (o *GetNetworkSwitchPortSchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch port schedules o k response has a 4xx status code
func (o *GetNetworkSwitchPortSchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch port schedules o k response has a 5xx status code
func (o *GetNetworkSwitchPortSchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch port schedules o k response a status code equal to that given
func (o *GetNetworkSwitchPortSchedulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network switch port schedules o k response
func (o *GetNetworkSwitchPortSchedulesOK) Code() int {
	return 200
}

func (o *GetNetworkSwitchPortSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/portSchedules][%d] getNetworkSwitchPortSchedulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchPortSchedulesOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/portSchedules][%d] getNetworkSwitchPortSchedulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchPortSchedulesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkSwitchPortSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
