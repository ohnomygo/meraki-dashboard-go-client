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
)

// GetNetworkWirelessSsidEapOverrideReader is a Reader for the GetNetworkWirelessSsidEapOverride structure.
type GetNetworkWirelessSsidEapOverrideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessSsidEapOverrideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessSsidEapOverrideOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessSsidEapOverrideOK creates a GetNetworkWirelessSsidEapOverrideOK with default headers values
func NewGetNetworkWirelessSsidEapOverrideOK() *GetNetworkWirelessSsidEapOverrideOK {
	return &GetNetworkWirelessSsidEapOverrideOK{}
}

/*
GetNetworkWirelessSsidEapOverrideOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessSsidEapOverrideOK struct {
	Payload *GetNetworkWirelessSsidEapOverrideOKBody
}

// IsSuccess returns true when this get network wireless ssid eap override o k response has a 2xx status code
func (o *GetNetworkWirelessSsidEapOverrideOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless ssid eap override o k response has a 3xx status code
func (o *GetNetworkWirelessSsidEapOverrideOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless ssid eap override o k response has a 4xx status code
func (o *GetNetworkWirelessSsidEapOverrideOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless ssid eap override o k response has a 5xx status code
func (o *GetNetworkWirelessSsidEapOverrideOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless ssid eap override o k response a status code equal to that given
func (o *GetNetworkWirelessSsidEapOverrideOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless ssid eap override o k response
func (o *GetNetworkWirelessSsidEapOverrideOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessSsidEapOverrideOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/eapOverride][%d] getNetworkWirelessSsidEapOverrideOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidEapOverrideOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/eapOverride][%d] getNetworkWirelessSsidEapOverrideOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidEapOverrideOK) GetPayload() *GetNetworkWirelessSsidEapOverrideOKBody {
	return o.Payload
}

func (o *GetNetworkWirelessSsidEapOverrideOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkWirelessSsidEapOverrideOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWirelessSsidEapOverrideOKBody get network wireless ssid eap override o k body
swagger:model GetNetworkWirelessSsidEapOverrideOKBody
*/
type GetNetworkWirelessSsidEapOverrideOKBody struct {

	// eapol key
	EapolKey *GetNetworkWirelessSsidEapOverrideOKBodyEapolKey `json:"eapolKey,omitempty"`

	// identity
	Identity *GetNetworkWirelessSsidEapOverrideOKBodyIdentity `json:"identity,omitempty"`

	// Maximum number of general EAP retries.
	MaxRetries int64 `json:"maxRetries,omitempty"`

	// General EAP timeout in seconds.
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this get network wireless ssid eap override o k body
func (o *GetNetworkWirelessSsidEapOverrideOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEapolKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSsidEapOverrideOKBody) validateEapolKey(formats strfmt.Registry) error {
	if swag.IsZero(o.EapolKey) { // not required
		return nil
	}

	if o.EapolKey != nil {
		if err := o.EapolKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWirelessSsidEapOverrideOK" + "." + "eapolKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWirelessSsidEapOverrideOK" + "." + "eapolKey")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkWirelessSsidEapOverrideOKBody) validateIdentity(formats strfmt.Registry) error {
	if swag.IsZero(o.Identity) { // not required
		return nil
	}

	if o.Identity != nil {
		if err := o.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWirelessSsidEapOverrideOK" + "." + "identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWirelessSsidEapOverrideOK" + "." + "identity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network wireless ssid eap override o k body based on the context it is used
func (o *GetNetworkWirelessSsidEapOverrideOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEapolKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSsidEapOverrideOKBody) contextValidateEapolKey(ctx context.Context, formats strfmt.Registry) error {

	if o.EapolKey != nil {
		if err := o.EapolKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWirelessSsidEapOverrideOK" + "." + "eapolKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWirelessSsidEapOverrideOK" + "." + "eapolKey")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkWirelessSsidEapOverrideOKBody) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if o.Identity != nil {
		if err := o.Identity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWirelessSsidEapOverrideOK" + "." + "identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWirelessSsidEapOverrideOK" + "." + "identity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSsidEapOverrideOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSsidEapOverrideOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSsidEapOverrideOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkWirelessSsidEapOverrideOKBodyEapolKey EAPOL Key settings.
swagger:model GetNetworkWirelessSsidEapOverrideOKBodyEapolKey
*/
type GetNetworkWirelessSsidEapOverrideOKBodyEapolKey struct {

	// Maximum number of EAPOL key retries.
	Retries int64 `json:"retries,omitempty"`

	// EAPOL Key timeout in milliseconds.
	TimeoutInMs int64 `json:"timeoutInMs,omitempty"`
}

// Validate validates this get network wireless ssid eap override o k body eapol key
func (o *GetNetworkWirelessSsidEapOverrideOKBodyEapolKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network wireless ssid eap override o k body eapol key based on context it is used
func (o *GetNetworkWirelessSsidEapOverrideOKBodyEapolKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSsidEapOverrideOKBodyEapolKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSsidEapOverrideOKBodyEapolKey) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSsidEapOverrideOKBodyEapolKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkWirelessSsidEapOverrideOKBodyIdentity EAP settings for identity requests.
swagger:model GetNetworkWirelessSsidEapOverrideOKBodyIdentity
*/
type GetNetworkWirelessSsidEapOverrideOKBodyIdentity struct {

	// Maximum number of EAP retries.
	Retries int64 `json:"retries,omitempty"`

	// EAP timeout in seconds.
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this get network wireless ssid eap override o k body identity
func (o *GetNetworkWirelessSsidEapOverrideOKBodyIdentity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network wireless ssid eap override o k body identity based on context it is used
func (o *GetNetworkWirelessSsidEapOverrideOKBodyIdentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSsidEapOverrideOKBodyIdentity) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSsidEapOverrideOKBodyIdentity) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSsidEapOverrideOKBodyIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
