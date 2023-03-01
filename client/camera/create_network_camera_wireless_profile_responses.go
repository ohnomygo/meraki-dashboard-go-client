// Code generated by go-swagger; DO NOT EDIT.

package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateNetworkCameraWirelessProfileReader is a Reader for the CreateNetworkCameraWirelessProfile structure.
type CreateNetworkCameraWirelessProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkCameraWirelessProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNetworkCameraWirelessProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkCameraWirelessProfileOK creates a CreateNetworkCameraWirelessProfileOK with default headers values
func NewCreateNetworkCameraWirelessProfileOK() *CreateNetworkCameraWirelessProfileOK {
	return &CreateNetworkCameraWirelessProfileOK{}
}

/*
CreateNetworkCameraWirelessProfileOK describes a response with status code 200, with default header values.

Successful operation
*/
type CreateNetworkCameraWirelessProfileOK struct {
	Payload interface{}
}

// IsSuccess returns true when this create network camera wireless profile o k response has a 2xx status code
func (o *CreateNetworkCameraWirelessProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network camera wireless profile o k response has a 3xx status code
func (o *CreateNetworkCameraWirelessProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network camera wireless profile o k response has a 4xx status code
func (o *CreateNetworkCameraWirelessProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network camera wireless profile o k response has a 5xx status code
func (o *CreateNetworkCameraWirelessProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create network camera wireless profile o k response a status code equal to that given
func (o *CreateNetworkCameraWirelessProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create network camera wireless profile o k response
func (o *CreateNetworkCameraWirelessProfileOK) Code() int {
	return 200
}

func (o *CreateNetworkCameraWirelessProfileOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/camera/wirelessProfiles][%d] createNetworkCameraWirelessProfileOK  %+v", 200, o.Payload)
}

func (o *CreateNetworkCameraWirelessProfileOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/camera/wirelessProfiles][%d] createNetworkCameraWirelessProfileOK  %+v", 200, o.Payload)
}

func (o *CreateNetworkCameraWirelessProfileOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkCameraWirelessProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkCameraWirelessProfileBody create network camera wireless profile body
// Example: {"name":"wireless profile A","ssid":{"authMode":"8021x-radius","encryptionMode":"wpa-eap","name":"ssid test"}}
swagger:model CreateNetworkCameraWirelessProfileBody
*/
type CreateNetworkCameraWirelessProfileBody struct {

	// identity
	Identity *CreateNetworkCameraWirelessProfileParamsBodyIdentity `json:"identity,omitempty"`

	// The name of the camera wireless profile. This parameter is required.
	// Required: true
	Name *string `json:"name"`

	// ssid
	// Required: true
	Ssid *CreateNetworkCameraWirelessProfileParamsBodySsid `json:"ssid"`
}

// Validate validates this create network camera wireless profile body
func (o *CreateNetworkCameraWirelessProfileBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSsid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkCameraWirelessProfileBody) validateIdentity(formats strfmt.Registry) error {
	if swag.IsZero(o.Identity) { // not required
		return nil
	}

	if o.Identity != nil {
		if err := o.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkCameraWirelessProfile" + "." + "identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkCameraWirelessProfile" + "." + "identity")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkCameraWirelessProfileBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkCameraWirelessProfile"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkCameraWirelessProfileBody) validateSsid(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkCameraWirelessProfile"+"."+"ssid", "body", o.Ssid); err != nil {
		return err
	}

	if o.Ssid != nil {
		if err := o.Ssid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkCameraWirelessProfile" + "." + "ssid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkCameraWirelessProfile" + "." + "ssid")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network camera wireless profile body based on the context it is used
func (o *CreateNetworkCameraWirelessProfileBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSsid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkCameraWirelessProfileBody) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if o.Identity != nil {
		if err := o.Identity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkCameraWirelessProfile" + "." + "identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkCameraWirelessProfile" + "." + "identity")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkCameraWirelessProfileBody) contextValidateSsid(ctx context.Context, formats strfmt.Registry) error {

	if o.Ssid != nil {
		if err := o.Ssid.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkCameraWirelessProfile" + "." + "ssid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkCameraWirelessProfile" + "." + "ssid")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkCameraWirelessProfileBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkCameraWirelessProfileBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkCameraWirelessProfileBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkCameraWirelessProfileParamsBodyIdentity The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
swagger:model CreateNetworkCameraWirelessProfileParamsBodyIdentity
*/
type CreateNetworkCameraWirelessProfileParamsBodyIdentity struct {

	// The password of the identity.
	Password string `json:"password,omitempty"`

	// The username of the identity.
	Username string `json:"username,omitempty"`
}

// Validate validates this create network camera wireless profile params body identity
func (o *CreateNetworkCameraWirelessProfileParamsBodyIdentity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network camera wireless profile params body identity based on context it is used
func (o *CreateNetworkCameraWirelessProfileParamsBodyIdentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkCameraWirelessProfileParamsBodyIdentity) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkCameraWirelessProfileParamsBodyIdentity) UnmarshalBinary(b []byte) error {
	var res CreateNetworkCameraWirelessProfileParamsBodyIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkCameraWirelessProfileParamsBodySsid The details of the SSID config.
swagger:model CreateNetworkCameraWirelessProfileParamsBodySsid
*/
type CreateNetworkCameraWirelessProfileParamsBodySsid struct {

	// The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	// Enum: [8021x-radius psk]
	AuthMode string `json:"authMode,omitempty"`

	// The encryption mode of the SSID. It can be set to ('wpa', 'wpa-eap'). With 'wpa' mode, the authMode should be 'psk' and with 'wpa-eap' the authMode should be '8021x-radius'
	EncryptionMode string `json:"encryptionMode,omitempty"`

	// The name of the SSID.
	Name string `json:"name,omitempty"`

	// The pre-shared key of the SSID.
	Psk string `json:"psk,omitempty"`
}

// Validate validates this create network camera wireless profile params body ssid
func (o *CreateNetworkCameraWirelessProfileParamsBodySsid) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createNetworkCameraWirelessProfileParamsBodySsidTypeAuthModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8021x-radius","psk"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkCameraWirelessProfileParamsBodySsidTypeAuthModePropEnum = append(createNetworkCameraWirelessProfileParamsBodySsidTypeAuthModePropEnum, v)
	}
}

const (

	// CreateNetworkCameraWirelessProfileParamsBodySsidAuthModeNr8021xDashRadius captures enum value "8021x-radius"
	CreateNetworkCameraWirelessProfileParamsBodySsidAuthModeNr8021xDashRadius string = "8021x-radius"

	// CreateNetworkCameraWirelessProfileParamsBodySsidAuthModePsk captures enum value "psk"
	CreateNetworkCameraWirelessProfileParamsBodySsidAuthModePsk string = "psk"
)

// prop value enum
func (o *CreateNetworkCameraWirelessProfileParamsBodySsid) validateAuthModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkCameraWirelessProfileParamsBodySsidTypeAuthModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkCameraWirelessProfileParamsBodySsid) validateAuthMode(formats strfmt.Registry) error {
	if swag.IsZero(o.AuthMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateAuthModeEnum("createNetworkCameraWirelessProfile"+"."+"ssid"+"."+"authMode", "body", o.AuthMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network camera wireless profile params body ssid based on context it is used
func (o *CreateNetworkCameraWirelessProfileParamsBodySsid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkCameraWirelessProfileParamsBodySsid) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkCameraWirelessProfileParamsBodySsid) UnmarshalBinary(b []byte) error {
	var res CreateNetworkCameraWirelessProfileParamsBodySsid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
