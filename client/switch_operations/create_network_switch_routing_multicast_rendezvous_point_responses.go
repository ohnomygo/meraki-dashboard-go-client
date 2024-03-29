// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateNetworkSwitchRoutingMulticastRendezvousPointReader is a Reader for the CreateNetworkSwitchRoutingMulticastRendezvousPoint structure.
type CreateNetworkSwitchRoutingMulticastRendezvousPointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchRoutingMulticastRendezvousPointCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkSwitchRoutingMulticastRendezvousPointCreated creates a CreateNetworkSwitchRoutingMulticastRendezvousPointCreated with default headers values
func NewCreateNetworkSwitchRoutingMulticastRendezvousPointCreated() *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated {
	return &CreateNetworkSwitchRoutingMulticastRendezvousPointCreated{}
}

/*
CreateNetworkSwitchRoutingMulticastRendezvousPointCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchRoutingMulticastRendezvousPointCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network switch routing multicast rendezvous point created response has a 2xx status code
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network switch routing multicast rendezvous point created response has a 3xx status code
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network switch routing multicast rendezvous point created response has a 4xx status code
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network switch routing multicast rendezvous point created response has a 5xx status code
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network switch routing multicast rendezvous point created response a status code equal to that given
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network switch routing multicast rendezvous point created response
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) Code() int {
	return 201
}

func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/routing/multicast/rendezvousPoints][%d] createNetworkSwitchRoutingMulticastRendezvousPointCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/routing/multicast/rendezvousPoints][%d] createNetworkSwitchRoutingMulticastRendezvousPointCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkSwitchRoutingMulticastRendezvousPointBody create network switch routing multicast rendezvous point body
// Example: {"interfaceIp":"192.168.1.2","multicastGroup":"192.168.128.0/24"}
swagger:model CreateNetworkSwitchRoutingMulticastRendezvousPointBody
*/
type CreateNetworkSwitchRoutingMulticastRendezvousPointBody struct {

	// The IP address of the interface where the RP needs to be created.
	// Required: true
	InterfaceIP *string `json:"interfaceIp"`

	// 'Any', or the IP address of a multicast group
	// Required: true
	MulticastGroup *string `json:"multicastGroup"`
}

// Validate validates this create network switch routing multicast rendezvous point body
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterfaceIP(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMulticastGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointBody) validateInterfaceIP(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchRoutingMulticastRendezvousPoint"+"."+"interfaceIp", "body", o.InterfaceIP); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointBody) validateMulticastGroup(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchRoutingMulticastRendezvousPoint"+"."+"multicastGroup", "body", o.MulticastGroup); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network switch routing multicast rendezvous point body based on context it is used
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchRoutingMulticastRendezvousPointBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
