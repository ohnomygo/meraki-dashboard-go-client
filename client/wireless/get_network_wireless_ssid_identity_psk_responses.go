// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// GetNetworkWirelessSsidIdentityPskReader is a Reader for the GetNetworkWirelessSsidIdentityPsk structure.
type GetNetworkWirelessSsidIdentityPskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessSsidIdentityPskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessSsidIdentityPskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessSsidIdentityPskOK creates a GetNetworkWirelessSsidIdentityPskOK with default headers values
func NewGetNetworkWirelessSsidIdentityPskOK() *GetNetworkWirelessSsidIdentityPskOK {
	return &GetNetworkWirelessSsidIdentityPskOK{}
}

/*
GetNetworkWirelessSsidIdentityPskOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessSsidIdentityPskOK struct {
	Payload *GetNetworkWirelessSsidIdentityPskOKBody
}

// IsSuccess returns true when this get network wireless ssid identity psk o k response has a 2xx status code
func (o *GetNetworkWirelessSsidIdentityPskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless ssid identity psk o k response has a 3xx status code
func (o *GetNetworkWirelessSsidIdentityPskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless ssid identity psk o k response has a 4xx status code
func (o *GetNetworkWirelessSsidIdentityPskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless ssid identity psk o k response has a 5xx status code
func (o *GetNetworkWirelessSsidIdentityPskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless ssid identity psk o k response a status code equal to that given
func (o *GetNetworkWirelessSsidIdentityPskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless ssid identity psk o k response
func (o *GetNetworkWirelessSsidIdentityPskOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessSsidIdentityPskOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}][%d] getNetworkWirelessSsidIdentityPskOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidIdentityPskOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}][%d] getNetworkWirelessSsidIdentityPskOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidIdentityPskOK) GetPayload() *GetNetworkWirelessSsidIdentityPskOKBody {
	return o.Payload
}

func (o *GetNetworkWirelessSsidIdentityPskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkWirelessSsidIdentityPskOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWirelessSsidIdentityPskOKBody get network wireless ssid identity psk o k body
swagger:model GetNetworkWirelessSsidIdentityPskOKBody
*/
type GetNetworkWirelessSsidIdentityPskOKBody struct {

	// The email associated with the System's Manager User
	Email string `json:"email,omitempty"`

	// Timestamp for when the Identity PSK expires, or 'null' to never expire
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expiresAt,omitempty"`

	// The group policy to be applied to clients
	GroupPolicyID string `json:"groupPolicyId,omitempty"`

	// The unique identifier of the Identity PSK
	ID string `json:"id,omitempty"`

	// The name of the Identity PSK
	Name string `json:"name,omitempty"`

	// The passphrase for client authentication
	Passphrase string `json:"passphrase,omitempty"`

	// The WiFi Personal Network unique identifier
	WifiPersonalNetworkID string `json:"wifiPersonalNetworkId,omitempty"`
}

// Validate validates this get network wireless ssid identity psk o k body
func (o *GetNetworkWirelessSsidIdentityPskOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSsidIdentityPskOKBody) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(o.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("getNetworkWirelessSsidIdentityPskOK"+"."+"expiresAt", "body", "date-time", o.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network wireless ssid identity psk o k body based on context it is used
func (o *GetNetworkWirelessSsidIdentityPskOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSsidIdentityPskOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSsidIdentityPskOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSsidIdentityPskOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
