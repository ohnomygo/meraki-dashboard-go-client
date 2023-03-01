// Code generated by go-swagger; DO NOT EDIT.

package switch_meraki_dasboard_go_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetNetworkSwitchStormControlReader is a Reader for the GetNetworkSwitchStormControl structure.
type GetNetworkSwitchStormControlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchStormControlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchStormControlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSwitchStormControlOK creates a GetNetworkSwitchStormControlOK with default headers values
func NewGetNetworkSwitchStormControlOK() *GetNetworkSwitchStormControlOK {
	return &GetNetworkSwitchStormControlOK{}
}

/*
GetNetworkSwitchStormControlOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchStormControlOK struct {
	Payload *GetNetworkSwitchStormControlOKBody
}

// IsSuccess returns true when this get network switch storm control o k response has a 2xx status code
func (o *GetNetworkSwitchStormControlOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch storm control o k response has a 3xx status code
func (o *GetNetworkSwitchStormControlOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch storm control o k response has a 4xx status code
func (o *GetNetworkSwitchStormControlOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch storm control o k response has a 5xx status code
func (o *GetNetworkSwitchStormControlOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch storm control o k response a status code equal to that given
func (o *GetNetworkSwitchStormControlOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network switch storm control o k response
func (o *GetNetworkSwitchStormControlOK) Code() int {
	return 200
}

func (o *GetNetworkSwitchStormControlOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/stormControl][%d] getNetworkSwitchStormControlOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchStormControlOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/stormControl][%d] getNetworkSwitchStormControlOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchStormControlOK) GetPayload() *GetNetworkSwitchStormControlOKBody {
	return o.Payload
}

func (o *GetNetworkSwitchStormControlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkSwitchStormControlOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSwitchStormControlOKBody get network switch storm control o k body
swagger:model GetNetworkSwitchStormControlOKBody
*/
type GetNetworkSwitchStormControlOKBody struct {

	// Broadcast threshold.
	BroadcastThreshold int64 `json:"broadcastThreshold,omitempty"`

	// Multicast threshold.
	MulticastThreshold int64 `json:"multicastThreshold,omitempty"`

	// Unknown Unicast threshold.
	UnknownUnicastThreshold int64 `json:"unknownUnicastThreshold,omitempty"`
}

// Validate validates this get network switch storm control o k body
func (o *GetNetworkSwitchStormControlOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network switch storm control o k body based on context it is used
func (o *GetNetworkSwitchStormControlOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchStormControlOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchStormControlOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchStormControlOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
