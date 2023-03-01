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

// GetNetworkMqttBrokersReader is a Reader for the GetNetworkMqttBrokers structure.
type GetNetworkMqttBrokersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkMqttBrokersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkMqttBrokersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkMqttBrokersOK creates a GetNetworkMqttBrokersOK with default headers values
func NewGetNetworkMqttBrokersOK() *GetNetworkMqttBrokersOK {
	return &GetNetworkMqttBrokersOK{}
}

/*
GetNetworkMqttBrokersOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkMqttBrokersOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network mqtt brokers o k response has a 2xx status code
func (o *GetNetworkMqttBrokersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network mqtt brokers o k response has a 3xx status code
func (o *GetNetworkMqttBrokersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network mqtt brokers o k response has a 4xx status code
func (o *GetNetworkMqttBrokersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network mqtt brokers o k response has a 5xx status code
func (o *GetNetworkMqttBrokersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network mqtt brokers o k response a status code equal to that given
func (o *GetNetworkMqttBrokersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network mqtt brokers o k response
func (o *GetNetworkMqttBrokersOK) Code() int {
	return 200
}

func (o *GetNetworkMqttBrokersOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/mqttBrokers][%d] getNetworkMqttBrokersOK  %+v", 200, o.Payload)
}

func (o *GetNetworkMqttBrokersOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/mqttBrokers][%d] getNetworkMqttBrokersOK  %+v", 200, o.Payload)
}

func (o *GetNetworkMqttBrokersOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkMqttBrokersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
