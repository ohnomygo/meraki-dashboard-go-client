// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceCellularGatewayLanReader is a Reader for the GetDeviceCellularGatewayLan structure.
type GetDeviceCellularGatewayLanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCellularGatewayLanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCellularGatewayLanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceCellularGatewayLanOK creates a GetDeviceCellularGatewayLanOK with default headers values
func NewGetDeviceCellularGatewayLanOK() *GetDeviceCellularGatewayLanOK {
	return &GetDeviceCellularGatewayLanOK{}
}

/*
GetDeviceCellularGatewayLanOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceCellularGatewayLanOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device cellular gateway lan o k response has a 2xx status code
func (o *GetDeviceCellularGatewayLanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device cellular gateway lan o k response has a 3xx status code
func (o *GetDeviceCellularGatewayLanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device cellular gateway lan o k response has a 4xx status code
func (o *GetDeviceCellularGatewayLanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device cellular gateway lan o k response has a 5xx status code
func (o *GetDeviceCellularGatewayLanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device cellular gateway lan o k response a status code equal to that given
func (o *GetDeviceCellularGatewayLanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device cellular gateway lan o k response
func (o *GetDeviceCellularGatewayLanOK) Code() int {
	return 200
}

func (o *GetDeviceCellularGatewayLanOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/cellularGateway/lan][%d] getDeviceCellularGatewayLanOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCellularGatewayLanOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/cellularGateway/lan][%d] getDeviceCellularGatewayLanOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCellularGatewayLanOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceCellularGatewayLanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
