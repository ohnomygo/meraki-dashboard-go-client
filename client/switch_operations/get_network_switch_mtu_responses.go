// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetNetworkSwitchMtuReader is a Reader for the GetNetworkSwitchMtu structure.
type GetNetworkSwitchMtuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchMtuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchMtuOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSwitchMtuOK creates a GetNetworkSwitchMtuOK with default headers values
func NewGetNetworkSwitchMtuOK() *GetNetworkSwitchMtuOK {
	return &GetNetworkSwitchMtuOK{}
}

/*
GetNetworkSwitchMtuOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchMtuOK struct {
	Payload *GetNetworkSwitchMtuOKBody
}

// IsSuccess returns true when this get network switch mtu o k response has a 2xx status code
func (o *GetNetworkSwitchMtuOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch mtu o k response has a 3xx status code
func (o *GetNetworkSwitchMtuOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch mtu o k response has a 4xx status code
func (o *GetNetworkSwitchMtuOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch mtu o k response has a 5xx status code
func (o *GetNetworkSwitchMtuOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch mtu o k response a status code equal to that given
func (o *GetNetworkSwitchMtuOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network switch mtu o k response
func (o *GetNetworkSwitchMtuOK) Code() int {
	return 200
}

func (o *GetNetworkSwitchMtuOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/mtu][%d] getNetworkSwitchMtuOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchMtuOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/mtu][%d] getNetworkSwitchMtuOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchMtuOK) GetPayload() *GetNetworkSwitchMtuOKBody {
	return o.Payload
}

func (o *GetNetworkSwitchMtuOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkSwitchMtuOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSwitchMtuOKBody get network switch mtu o k body
swagger:model GetNetworkSwitchMtuOKBody
*/
type GetNetworkSwitchMtuOKBody struct {

	// MTU size for the entire network. Default value is 9578.
	DefaultMtuSize int64 `json:"defaultMtuSize,omitempty"`

	// Override MTU size for individual switches or switch profiles.
	//       An empty array will clear overrides.
	Overrides []*GetNetworkSwitchMtuOKBodyOverridesItems0 `json:"overrides"`
}

// Validate validates this get network switch mtu o k body
func (o *GetNetworkSwitchMtuOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSwitchMtuOKBody) validateOverrides(formats strfmt.Registry) error {
	if swag.IsZero(o.Overrides) { // not required
		return nil
	}

	for i := 0; i < len(o.Overrides); i++ {
		if swag.IsZero(o.Overrides[i]) { // not required
			continue
		}

		if o.Overrides[i] != nil {
			if err := o.Overrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkSwitchMtuOK" + "." + "overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getNetworkSwitchMtuOK" + "." + "overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get network switch mtu o k body based on the context it is used
func (o *GetNetworkSwitchMtuOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSwitchMtuOKBody) contextValidateOverrides(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Overrides); i++ {

		if o.Overrides[i] != nil {
			if err := o.Overrides[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkSwitchMtuOK" + "." + "overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getNetworkSwitchMtuOK" + "." + "overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchMtuOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchMtuOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchMtuOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSwitchMtuOKBodyOverridesItems0 get network switch mtu o k body overrides items0
swagger:model GetNetworkSwitchMtuOKBodyOverridesItems0
*/
type GetNetworkSwitchMtuOKBodyOverridesItems0 struct {

	// MTU size for the switches or switch profiles.
	// Required: true
	MtuSize *int64 `json:"mtuSize"`

	// List of switch profile IDs. Applicable only for template network.
	SwitchProfiles []string `json:"switchProfiles"`

	// List of switch serials. Applicable only for switch network.
	Switches []string `json:"switches"`
}

// Validate validates this get network switch mtu o k body overrides items0
func (o *GetNetworkSwitchMtuOKBodyOverridesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMtuSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSwitchMtuOKBodyOverridesItems0) validateMtuSize(formats strfmt.Registry) error {

	if err := validate.Required("mtuSize", "body", o.MtuSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network switch mtu o k body overrides items0 based on context it is used
func (o *GetNetworkSwitchMtuOKBodyOverridesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchMtuOKBodyOverridesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchMtuOKBodyOverridesItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchMtuOKBodyOverridesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
