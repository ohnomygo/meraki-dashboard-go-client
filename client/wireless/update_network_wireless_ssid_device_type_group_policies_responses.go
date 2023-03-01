// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesReader is a Reader for the UpdateNetworkWirelessSsidDeviceTypeGroupPolicies structure.
type UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK creates a UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK with default headers values
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK {
	return &UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK{}
}

/*
UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network wireless ssid device type group policies o k response has a 2xx status code
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network wireless ssid device type group policies o k response has a 3xx status code
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network wireless ssid device type group policies o k response has a 4xx status code
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network wireless ssid device type group policies o k response has a 5xx status code
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network wireless ssid device type group policies o k response a status code equal to that given
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network wireless ssid device type group policies o k response
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) Code() int {
	return 200
}

func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies][%d] updateNetworkWirelessSsidDeviceTypeGroupPoliciesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies][%d] updateNetworkWirelessSsidDeviceTypeGroupPoliciesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody update network wireless ssid device type group policies body
// Example: {"deviceTypePolicies":[{"devicePolicy":"Allowed","deviceType":"Android"},{"devicePolicy":"Group policy","deviceType":"iPhone","groupPolicyId":101}],"enabled":true}
swagger:model UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody
*/
type UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody struct {

	// List of device type policies.
	DeviceTypePolicies []*UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0 `json:"deviceTypePolicies"`

	// If true, the SSID device type group policies are enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network wireless ssid device type group policies body
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDeviceTypePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody) validateDeviceTypePolicies(formats strfmt.Registry) error {
	if swag.IsZero(o.DeviceTypePolicies) { // not required
		return nil
	}

	for i := 0; i < len(o.DeviceTypePolicies); i++ {
		if swag.IsZero(o.DeviceTypePolicies[i]) { // not required
			continue
		}

		if o.DeviceTypePolicies[i] != nil {
			if err := o.DeviceTypePolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkWirelessSsidDeviceTypeGroupPolicies" + "." + "deviceTypePolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkWirelessSsidDeviceTypeGroupPolicies" + "." + "deviceTypePolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network wireless ssid device type group policies body based on the context it is used
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDeviceTypePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody) contextValidateDeviceTypePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DeviceTypePolicies); i++ {

		if o.DeviceTypePolicies[i] != nil {
			if err := o.DeviceTypePolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkWirelessSsidDeviceTypeGroupPolicies" + "." + "deviceTypePolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkWirelessSsidDeviceTypeGroupPolicies" + "." + "deviceTypePolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0 update network wireless ssid device type group policies params body device type policies items0
swagger:model UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0
*/
type UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0 struct {

	// The device policy. Can be one of 'Allowed', 'Blocked' or 'Group policy'
	// Required: true
	// Enum: [Allowed Blocked Group policy]
	DevicePolicy *string `json:"devicePolicy"`

	// The device type. Can be one of 'Android', 'BlackBerry', 'Chrome OS', 'iPad', 'iPhone', 'iPod', 'Mac OS X', 'Windows', 'Windows Phone', 'B&N Nook' or 'Other OS'
	// Required: true
	// Enum: [Android B&N Nook BlackBerry Chrome OS Mac OS X Other OS Windows Windows Phone iPad iPhone iPod]
	DeviceType *string `json:"deviceType"`

	// ID of the group policy object.
	GroupPolicyID int64 `json:"groupPolicyId,omitempty"`
}

// Validate validates this update network wireless ssid device type group policies params body device type policies items0
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevicePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0TypeDevicePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Allowed","Blocked","Group policy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0TypeDevicePolicyPropEnum = append(updateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0TypeDevicePolicyPropEnum, v)
	}
}

const (

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DevicePolicyAllowed captures enum value "Allowed"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DevicePolicyAllowed string = "Allowed"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DevicePolicyBlocked captures enum value "Blocked"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DevicePolicyBlocked string = "Blocked"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DevicePolicyGroupPolicy captures enum value "Group policy"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DevicePolicyGroupPolicy string = "Group policy"
)

// prop value enum
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0) validateDevicePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0TypeDevicePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0) validateDevicePolicy(formats strfmt.Registry) error {

	if err := validate.Required("devicePolicy", "body", o.DevicePolicy); err != nil {
		return err
	}

	// value enum
	if err := o.validateDevicePolicyEnum("devicePolicy", "body", *o.DevicePolicy); err != nil {
		return err
	}

	return nil
}

var updateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0TypeDeviceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Android","B\u0026N Nook","BlackBerry","Chrome OS","Mac OS X","Other OS","Windows","Windows Phone","iPad","iPhone","iPod"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0TypeDeviceTypePropEnum = append(updateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0TypeDeviceTypePropEnum, v)
	}
}

const (

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeAndroid captures enum value "Android"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeAndroid string = "Android"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeBAndNNook captures enum value "B&N Nook"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeBAndNNook string = "B&N Nook"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeBlackBerry captures enum value "BlackBerry"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeBlackBerry string = "BlackBerry"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeChromeOS captures enum value "Chrome OS"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeChromeOS string = "Chrome OS"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeMacOSX captures enum value "Mac OS X"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeMacOSX string = "Mac OS X"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeOtherOS captures enum value "Other OS"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeOtherOS string = "Other OS"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeWindows captures enum value "Windows"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeWindows string = "Windows"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeWindowsPhone captures enum value "Windows Phone"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeWindowsPhone string = "Windows Phone"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeIPad captures enum value "iPad"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeIPad string = "iPad"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeIPhone captures enum value "iPhone"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeIPhone string = "iPhone"

	// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeIPod captures enum value "iPod"
	UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0DeviceTypeIPod string = "iPod"
)

// prop value enum
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0) validateDeviceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0TypeDeviceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("deviceType", "body", o.DeviceType); err != nil {
		return err
	}

	// value enum
	if err := o.validateDeviceTypeEnum("deviceType", "body", *o.DeviceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network wireless ssid device type group policies params body device type policies items0 based on context it is used
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsBodyDeviceTypePoliciesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}