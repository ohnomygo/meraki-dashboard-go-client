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

// UpdateNetworkWirelessSsidEapOverrideReader is a Reader for the UpdateNetworkWirelessSsidEapOverride structure.
type UpdateNetworkWirelessSsidEapOverrideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkWirelessSsidEapOverrideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkWirelessSsidEapOverrideOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkWirelessSsidEapOverrideOK creates a UpdateNetworkWirelessSsidEapOverrideOK with default headers values
func NewUpdateNetworkWirelessSsidEapOverrideOK() *UpdateNetworkWirelessSsidEapOverrideOK {
	return &UpdateNetworkWirelessSsidEapOverrideOK{}
}

/*
UpdateNetworkWirelessSsidEapOverrideOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkWirelessSsidEapOverrideOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network wireless ssid eap override o k response has a 2xx status code
func (o *UpdateNetworkWirelessSsidEapOverrideOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network wireless ssid eap override o k response has a 3xx status code
func (o *UpdateNetworkWirelessSsidEapOverrideOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network wireless ssid eap override o k response has a 4xx status code
func (o *UpdateNetworkWirelessSsidEapOverrideOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network wireless ssid eap override o k response has a 5xx status code
func (o *UpdateNetworkWirelessSsidEapOverrideOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network wireless ssid eap override o k response a status code equal to that given
func (o *UpdateNetworkWirelessSsidEapOverrideOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network wireless ssid eap override o k response
func (o *UpdateNetworkWirelessSsidEapOverrideOK) Code() int {
	return 200
}

func (o *UpdateNetworkWirelessSsidEapOverrideOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/ssids/{number}/eapOverride][%d] updateNetworkWirelessSsidEapOverrideOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWirelessSsidEapOverrideOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/ssids/{number}/eapOverride][%d] updateNetworkWirelessSsidEapOverrideOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWirelessSsidEapOverrideOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkWirelessSsidEapOverrideOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkWirelessSsidEapOverrideBody update network wireless ssid eap override body
// Example: {}
swagger:model UpdateNetworkWirelessSsidEapOverrideBody
*/
type UpdateNetworkWirelessSsidEapOverrideBody struct {

	// eapol key
	EapolKey *UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey `json:"eapolKey,omitempty"`

	// identity
	Identity *UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity `json:"identity,omitempty"`

	// Maximum number of general EAP retries.
	MaxRetries int64 `json:"maxRetries,omitempty"`

	// General EAP timeout in seconds.
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this update network wireless ssid eap override body
func (o *UpdateNetworkWirelessSsidEapOverrideBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateNetworkWirelessSsidEapOverrideBody) validateEapolKey(formats strfmt.Registry) error {
	if swag.IsZero(o.EapolKey) { // not required
		return nil
	}

	if o.EapolKey != nil {
		if err := o.EapolKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidEapOverride" + "." + "eapolKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidEapOverride" + "." + "eapolKey")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkWirelessSsidEapOverrideBody) validateIdentity(formats strfmt.Registry) error {
	if swag.IsZero(o.Identity) { // not required
		return nil
	}

	if o.Identity != nil {
		if err := o.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidEapOverride" + "." + "identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidEapOverride" + "." + "identity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network wireless ssid eap override body based on the context it is used
func (o *UpdateNetworkWirelessSsidEapOverrideBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *UpdateNetworkWirelessSsidEapOverrideBody) contextValidateEapolKey(ctx context.Context, formats strfmt.Registry) error {

	if o.EapolKey != nil {
		if err := o.EapolKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidEapOverride" + "." + "eapolKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidEapOverride" + "." + "eapolKey")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkWirelessSsidEapOverrideBody) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if o.Identity != nil {
		if err := o.Identity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidEapOverride" + "." + "identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidEapOverride" + "." + "identity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidEapOverrideBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidEapOverrideBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidEapOverrideBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey EAPOL Key settings.
swagger:model UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey
*/
type UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey struct {

	// Maximum number of EAPOL key retries.
	Retries int64 `json:"retries,omitempty"`

	// EAPOL Key timeout in milliseconds.
	TimeoutInMs int64 `json:"timeoutInMs,omitempty"`
}

// Validate validates this update network wireless ssid eap override params body eapol key
func (o *UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network wireless ssid eap override params body eapol key based on context it is used
func (o *UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidEapOverrideParamsBodyEapolKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity EAP settings for identity requests.
swagger:model UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity
*/
type UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity struct {

	// Maximum number of EAP retries.
	Retries int64 `json:"retries,omitempty"`

	// EAP timeout in seconds.
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this update network wireless ssid eap override params body identity
func (o *UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network wireless ssid eap override params body identity based on context it is used
func (o *UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidEapOverrideParamsBodyIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
