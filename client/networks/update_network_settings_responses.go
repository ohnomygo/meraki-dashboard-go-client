// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// UpdateNetworkSettingsReader is a Reader for the UpdateNetworkSettings structure.
type UpdateNetworkSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSettingsOK creates a UpdateNetworkSettingsOK with default headers values
func NewUpdateNetworkSettingsOK() *UpdateNetworkSettingsOK {
	return &UpdateNetworkSettingsOK{}
}

/*
UpdateNetworkSettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSettingsOK struct {
	Payload *UpdateNetworkSettingsOKBody
}

// IsSuccess returns true when this update network settings o k response has a 2xx status code
func (o *UpdateNetworkSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network settings o k response has a 3xx status code
func (o *UpdateNetworkSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network settings o k response has a 4xx status code
func (o *UpdateNetworkSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network settings o k response has a 5xx status code
func (o *UpdateNetworkSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network settings o k response a status code equal to that given
func (o *UpdateNetworkSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network settings o k response
func (o *UpdateNetworkSettingsOK) Code() int {
	return 200
}

func (o *UpdateNetworkSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/settings][%d] updateNetworkSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSettingsOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/settings][%d] updateNetworkSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSettingsOK) GetPayload() *UpdateNetworkSettingsOKBody {
	return o.Payload
}

func (o *UpdateNetworkSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateNetworkSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkSettingsBody update network settings body
// Example: {"localStatusPage":{"authentication":{"enabled":false,"password":"miles123"}},"localStatusPageEnabled":true,"remoteStatusPageEnabled":true,"securePort":{"enabled":false}}
swagger:model UpdateNetworkSettingsBody
*/
type UpdateNetworkSettingsBody struct {

	// local status page
	LocalStatusPage *UpdateNetworkSettingsParamsBodyLocalStatusPage `json:"localStatusPage,omitempty"`

	// Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	LocalStatusPageEnabled bool `json:"localStatusPageEnabled,omitempty"`

	// Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	RemoteStatusPageEnabled bool `json:"remoteStatusPageEnabled,omitempty"`

	// secure port
	SecurePort *UpdateNetworkSettingsParamsBodySecurePort `json:"securePort,omitempty"`
}

// Validate validates this update network settings body
func (o *UpdateNetworkSettingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocalStatusPage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecurePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsBody) validateLocalStatusPage(formats strfmt.Registry) error {
	if swag.IsZero(o.LocalStatusPage) { // not required
		return nil
	}

	if o.LocalStatusPage != nil {
		if err := o.LocalStatusPage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettings" + "." + "localStatusPage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettings" + "." + "localStatusPage")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsBody) validateSecurePort(formats strfmt.Registry) error {
	if swag.IsZero(o.SecurePort) { // not required
		return nil
	}

	if o.SecurePort != nil {
		if err := o.SecurePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettings" + "." + "securePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettings" + "." + "securePort")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network settings body based on the context it is used
func (o *UpdateNetworkSettingsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocalStatusPage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSecurePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsBody) contextValidateLocalStatusPage(ctx context.Context, formats strfmt.Registry) error {

	if o.LocalStatusPage != nil {
		if err := o.LocalStatusPage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettings" + "." + "localStatusPage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettings" + "." + "localStatusPage")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsBody) contextValidateSecurePort(ctx context.Context, formats strfmt.Registry) error {

	if o.SecurePort != nil {
		if err := o.SecurePort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettings" + "." + "securePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettings" + "." + "securePort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsOKBody update network settings o k body
swagger:model UpdateNetworkSettingsOKBody
*/
type UpdateNetworkSettingsOKBody struct {

	// client privacy
	ClientPrivacy *UpdateNetworkSettingsOKBodyClientPrivacy `json:"clientPrivacy,omitempty"`

	// fips
	Fips *UpdateNetworkSettingsOKBodyFips `json:"fips,omitempty"`

	// local status page
	LocalStatusPage *UpdateNetworkSettingsOKBodyLocalStatusPage `json:"localStatusPage,omitempty"`

	// Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	LocalStatusPageEnabled bool `json:"localStatusPageEnabled,omitempty"`

	// named vlans
	NamedVlans *UpdateNetworkSettingsOKBodyNamedVlans `json:"namedVlans,omitempty"`

	// Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	RemoteStatusPageEnabled bool `json:"remoteStatusPageEnabled,omitempty"`

	// secure port
	SecurePort *UpdateNetworkSettingsOKBodySecurePort `json:"securePort,omitempty"`
}

// Validate validates this update network settings o k body
func (o *UpdateNetworkSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientPrivacy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFips(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocalStatusPage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNamedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecurePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsOKBody) validateClientPrivacy(formats strfmt.Registry) error {
	if swag.IsZero(o.ClientPrivacy) { // not required
		return nil
	}

	if o.ClientPrivacy != nil {
		if err := o.ClientPrivacy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "clientPrivacy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "clientPrivacy")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsOKBody) validateFips(formats strfmt.Registry) error {
	if swag.IsZero(o.Fips) { // not required
		return nil
	}

	if o.Fips != nil {
		if err := o.Fips.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "fips")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "fips")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsOKBody) validateLocalStatusPage(formats strfmt.Registry) error {
	if swag.IsZero(o.LocalStatusPage) { // not required
		return nil
	}

	if o.LocalStatusPage != nil {
		if err := o.LocalStatusPage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "localStatusPage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "localStatusPage")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsOKBody) validateNamedVlans(formats strfmt.Registry) error {
	if swag.IsZero(o.NamedVlans) { // not required
		return nil
	}

	if o.NamedVlans != nil {
		if err := o.NamedVlans.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "namedVlans")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "namedVlans")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsOKBody) validateSecurePort(formats strfmt.Registry) error {
	if swag.IsZero(o.SecurePort) { // not required
		return nil
	}

	if o.SecurePort != nil {
		if err := o.SecurePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "securePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "securePort")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network settings o k body based on the context it is used
func (o *UpdateNetworkSettingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateClientPrivacy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFips(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocalStatusPage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNamedVlans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSecurePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsOKBody) contextValidateClientPrivacy(ctx context.Context, formats strfmt.Registry) error {

	if o.ClientPrivacy != nil {
		if err := o.ClientPrivacy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "clientPrivacy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "clientPrivacy")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsOKBody) contextValidateFips(ctx context.Context, formats strfmt.Registry) error {

	if o.Fips != nil {
		if err := o.Fips.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "fips")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "fips")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsOKBody) contextValidateLocalStatusPage(ctx context.Context, formats strfmt.Registry) error {

	if o.LocalStatusPage != nil {
		if err := o.LocalStatusPage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "localStatusPage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "localStatusPage")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsOKBody) contextValidateNamedVlans(ctx context.Context, formats strfmt.Registry) error {

	if o.NamedVlans != nil {
		if err := o.NamedVlans.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "namedVlans")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "namedVlans")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSettingsOKBody) contextValidateSecurePort(ctx context.Context, formats strfmt.Registry) error {

	if o.SecurePort != nil {
		if err := o.SecurePort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "securePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "securePort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsOKBodyClientPrivacy Privacy settings
swagger:model UpdateNetworkSettingsOKBodyClientPrivacy
*/
type UpdateNetworkSettingsOKBodyClientPrivacy struct {

	// The date to expire the data before
	// Format: date-time
	ExpireDataBefore strfmt.DateTime `json:"expireDataBefore,omitempty"`

	// The number of days, weeks, or months in Epoch time to expire the data before
	ExpireDataOlderThan int64 `json:"expireDataOlderThan,omitempty"`
}

// Validate validates this update network settings o k body client privacy
func (o *UpdateNetworkSettingsOKBodyClientPrivacy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExpireDataBefore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsOKBodyClientPrivacy) validateExpireDataBefore(formats strfmt.Registry) error {
	if swag.IsZero(o.ExpireDataBefore) { // not required
		return nil
	}

	if err := validate.FormatOf("updateNetworkSettingsOK"+"."+"clientPrivacy"+"."+"expireDataBefore", "body", "date-time", o.ExpireDataBefore.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network settings o k body client privacy based on context it is used
func (o *UpdateNetworkSettingsOKBodyClientPrivacy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyClientPrivacy) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyClientPrivacy) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsOKBodyClientPrivacy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsOKBodyFips A hash of FIPS options applied to the Network
swagger:model UpdateNetworkSettingsOKBodyFips
*/
type UpdateNetworkSettingsOKBodyFips struct {

	// Enables / disables FIPS on the network.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network settings o k body fips
func (o *UpdateNetworkSettingsOKBodyFips) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network settings o k body fips based on context it is used
func (o *UpdateNetworkSettingsOKBodyFips) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyFips) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyFips) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsOKBodyFips
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsOKBodyLocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
swagger:model UpdateNetworkSettingsOKBodyLocalStatusPage
*/
type UpdateNetworkSettingsOKBodyLocalStatusPage struct {

	// authentication
	Authentication *UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// Validate validates this update network settings o k body local status page
func (o *UpdateNetworkSettingsOKBodyLocalStatusPage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsOKBodyLocalStatusPage) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(o.Authentication) { // not required
		return nil
	}

	if o.Authentication != nil {
		if err := o.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "localStatusPage" + "." + "authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "localStatusPage" + "." + "authentication")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network settings o k body local status page based on the context it is used
func (o *UpdateNetworkSettingsOKBodyLocalStatusPage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsOKBodyLocalStatusPage) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if o.Authentication != nil {
		if err := o.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettingsOK" + "." + "localStatusPage" + "." + "authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettingsOK" + "." + "localStatusPage" + "." + "authentication")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyLocalStatusPage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyLocalStatusPage) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsOKBodyLocalStatusPage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication A hash of Local Status page(s)' authentication options applied to the Network.
swagger:model UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication
*/
type UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication struct {

	// Enables / disables the authentication on Local Status page(s).
	Enabled bool `json:"enabled,omitempty"`

	// The username used for Local Status Page(s).
	Username string `json:"username,omitempty"`
}

// Validate validates this update network settings o k body local status page authentication
func (o *UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network settings o k body local status page authentication based on context it is used
func (o *UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsOKBodyLocalStatusPageAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsOKBodyNamedVlans A hash of Named VLANs options applied to the Network.
swagger:model UpdateNetworkSettingsOKBodyNamedVlans
*/
type UpdateNetworkSettingsOKBodyNamedVlans struct {

	// Enables / disables Named VLANs on the Network.
	// Required: true
	Enabled *bool `json:"enabled"`
}

// Validate validates this update network settings o k body named vlans
func (o *UpdateNetworkSettingsOKBodyNamedVlans) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsOKBodyNamedVlans) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("updateNetworkSettingsOK"+"."+"namedVlans"+"."+"enabled", "body", o.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network settings o k body named vlans based on context it is used
func (o *UpdateNetworkSettingsOKBodyNamedVlans) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyNamedVlans) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodyNamedVlans) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsOKBodyNamedVlans
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsOKBodySecurePort A hash of SecureConnect options applied to the Network.
swagger:model UpdateNetworkSettingsOKBodySecurePort
*/
type UpdateNetworkSettingsOKBodySecurePort struct {

	// Enables / disables SecureConnect on the network. Optional.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network settings o k body secure port
func (o *UpdateNetworkSettingsOKBodySecurePort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network settings o k body secure port based on context it is used
func (o *UpdateNetworkSettingsOKBodySecurePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodySecurePort) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsOKBodySecurePort) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsOKBodySecurePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsParamsBodyLocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
swagger:model UpdateNetworkSettingsParamsBodyLocalStatusPage
*/
type UpdateNetworkSettingsParamsBodyLocalStatusPage struct {

	// authentication
	Authentication *UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// Validate validates this update network settings params body local status page
func (o *UpdateNetworkSettingsParamsBodyLocalStatusPage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsParamsBodyLocalStatusPage) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(o.Authentication) { // not required
		return nil
	}

	if o.Authentication != nil {
		if err := o.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettings" + "." + "localStatusPage" + "." + "authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettings" + "." + "localStatusPage" + "." + "authentication")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network settings params body local status page based on the context it is used
func (o *UpdateNetworkSettingsParamsBodyLocalStatusPage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSettingsParamsBodyLocalStatusPage) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if o.Authentication != nil {
		if err := o.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSettings" + "." + "localStatusPage" + "." + "authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSettings" + "." + "localStatusPage" + "." + "authentication")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsParamsBodyLocalStatusPage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsParamsBodyLocalStatusPage) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsParamsBodyLocalStatusPage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication A hash of Local Status page(s)' authentication options applied to the Network.
swagger:model UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication
*/
type UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication struct {

	// Enables / disables the authentication on Local Status page(s).
	Enabled bool `json:"enabled,omitempty"`

	// The password used for Local Status Page(s). Set this to null to clear the password.
	Password string `json:"password,omitempty"`
}

// Validate validates this update network settings params body local status page authentication
func (o *UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network settings params body local status page authentication based on context it is used
func (o *UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsParamsBodyLocalStatusPageAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSettingsParamsBodySecurePort A hash of SecureConnect options applied to the Network.
swagger:model UpdateNetworkSettingsParamsBodySecurePort
*/
type UpdateNetworkSettingsParamsBodySecurePort struct {

	// Enables / disables SecureConnect on the network. Optional.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network settings params body secure port
func (o *UpdateNetworkSettingsParamsBodySecurePort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network settings params body secure port based on context it is used
func (o *UpdateNetworkSettingsParamsBodySecurePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSettingsParamsBodySecurePort) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSettingsParamsBodySecurePort) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSettingsParamsBodySecurePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}