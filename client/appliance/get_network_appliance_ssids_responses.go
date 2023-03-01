// Code generated by go-swagger; DO NOT EDIT.

package appliance

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
)

// GetNetworkApplianceSsidsReader is a Reader for the GetNetworkApplianceSsids structure.
type GetNetworkApplianceSsidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceSsidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceSsidsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceSsidsOK creates a GetNetworkApplianceSsidsOK with default headers values
func NewGetNetworkApplianceSsidsOK() *GetNetworkApplianceSsidsOK {
	return &GetNetworkApplianceSsidsOK{}
}

/*
GetNetworkApplianceSsidsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceSsidsOK struct {
	Payload []*GetNetworkApplianceSsidsOKBodyItems0
}

// IsSuccess returns true when this get network appliance ssids o k response has a 2xx status code
func (o *GetNetworkApplianceSsidsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance ssids o k response has a 3xx status code
func (o *GetNetworkApplianceSsidsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance ssids o k response has a 4xx status code
func (o *GetNetworkApplianceSsidsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance ssids o k response has a 5xx status code
func (o *GetNetworkApplianceSsidsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance ssids o k response a status code equal to that given
func (o *GetNetworkApplianceSsidsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network appliance ssids o k response
func (o *GetNetworkApplianceSsidsOK) Code() int {
	return 200
}

func (o *GetNetworkApplianceSsidsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/ssids][%d] getNetworkApplianceSsidsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceSsidsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/ssids][%d] getNetworkApplianceSsidsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceSsidsOK) GetPayload() []*GetNetworkApplianceSsidsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkApplianceSsidsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkApplianceSsidsOKBodyItems0 get network appliance ssids o k body items0
swagger:model GetNetworkApplianceSsidsOKBodyItems0
*/
type GetNetworkApplianceSsidsOKBodyItems0 struct {

	// The association control method for the SSID.
	AuthMode string `json:"authMode,omitempty"`

	// The VLAN ID of the VLAN associated to this SSID.
	DefaultVlanID int64 `json:"defaultVlanId,omitempty"`

	// Whether or not the SSID is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// The psk encryption mode for the SSID.
	EncryptionMode string `json:"encryptionMode,omitempty"`

	// The name of the SSID.
	Name string `json:"name,omitempty"`

	// The number of the SSID.
	Number int64 `json:"number,omitempty"`

	// The RADIUS 802.1x servers to be used for authentication.
	RadiusServers []*GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0 `json:"radiusServers"`

	// Boolean indicating whether the MX should advertise or hide this SSID.
	Visible bool `json:"visible,omitempty"`

	// WPA encryption mode for the SSID.
	WpaEncryptionMode string `json:"wpaEncryptionMode,omitempty"`
}

// Validate validates this get network appliance ssids o k body items0
func (o *GetNetworkApplianceSsidsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRadiusServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkApplianceSsidsOKBodyItems0) validateRadiusServers(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusServers) { // not required
		return nil
	}

	for i := 0; i < len(o.RadiusServers); i++ {
		if swag.IsZero(o.RadiusServers[i]) { // not required
			continue
		}

		if o.RadiusServers[i] != nil {
			if err := o.RadiusServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get network appliance ssids o k body items0 based on the context it is used
func (o *GetNetworkApplianceSsidsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRadiusServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkApplianceSsidsOKBodyItems0) contextValidateRadiusServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusServers); i++ {

		if o.RadiusServers[i] != nil {
			if err := o.RadiusServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceSsidsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceSsidsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceSsidsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0 get network appliance ssids o k body items0 radius servers items0
swagger:model GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0
*/
type GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0 struct {

	// The IP address of your RADIUS server.
	Host string `json:"host,omitempty"`

	// The UDP port your RADIUS servers listens on for Access-requests.
	Port int64 `json:"port,omitempty"`
}

// Validate validates this get network appliance ssids o k body items0 radius servers items0
func (o *GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network appliance ssids o k body items0 radius servers items0 based on context it is used
func (o *GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceSsidsOKBodyItems0RadiusServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}