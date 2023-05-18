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

// GetDeviceSwitchPortsStatusesPacketsReader is a Reader for the GetDeviceSwitchPortsStatusesPackets structure.
type GetDeviceSwitchPortsStatusesPacketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceSwitchPortsStatusesPacketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceSwitchPortsStatusesPacketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceSwitchPortsStatusesPacketsOK creates a GetDeviceSwitchPortsStatusesPacketsOK with default headers values
func NewGetDeviceSwitchPortsStatusesPacketsOK() *GetDeviceSwitchPortsStatusesPacketsOK {
	return &GetDeviceSwitchPortsStatusesPacketsOK{}
}

/*
GetDeviceSwitchPortsStatusesPacketsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceSwitchPortsStatusesPacketsOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get device switch ports statuses packets o k response has a 2xx status code
func (o *GetDeviceSwitchPortsStatusesPacketsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device switch ports statuses packets o k response has a 3xx status code
func (o *GetDeviceSwitchPortsStatusesPacketsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device switch ports statuses packets o k response has a 4xx status code
func (o *GetDeviceSwitchPortsStatusesPacketsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device switch ports statuses packets o k response has a 5xx status code
func (o *GetDeviceSwitchPortsStatusesPacketsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device switch ports statuses packets o k response a status code equal to that given
func (o *GetDeviceSwitchPortsStatusesPacketsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device switch ports statuses packets o k response
func (o *GetDeviceSwitchPortsStatusesPacketsOK) Code() int {
	return 200
}

func (o *GetDeviceSwitchPortsStatusesPacketsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/ports/statuses/packets][%d] getDeviceSwitchPortsStatusesPacketsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchPortsStatusesPacketsOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/ports/statuses/packets][%d] getDeviceSwitchPortsStatusesPacketsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchPortsStatusesPacketsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetDeviceSwitchPortsStatusesPacketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
