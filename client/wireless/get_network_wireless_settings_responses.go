// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// GetNetworkWirelessSettingsReader is a Reader for the GetNetworkWirelessSettings structure.
type GetNetworkWirelessSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessSettingsOK creates a GetNetworkWirelessSettingsOK with default headers values
func NewGetNetworkWirelessSettingsOK() *GetNetworkWirelessSettingsOK {
	return &GetNetworkWirelessSettingsOK{}
}

/*
GetNetworkWirelessSettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessSettingsOK struct {
	Payload *GetNetworkWirelessSettingsOKBody
}

// IsSuccess returns true when this get network wireless settings o k response has a 2xx status code
func (o *GetNetworkWirelessSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless settings o k response has a 3xx status code
func (o *GetNetworkWirelessSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless settings o k response has a 4xx status code
func (o *GetNetworkWirelessSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless settings o k response has a 5xx status code
func (o *GetNetworkWirelessSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless settings o k response a status code equal to that given
func (o *GetNetworkWirelessSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless settings o k response
func (o *GetNetworkWirelessSettingsOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessSettingsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/settings][%d] getNetworkWirelessSettingsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSettingsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/settings][%d] getNetworkWirelessSettingsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSettingsOK) GetPayload() *GetNetworkWirelessSettingsOKBody {
	return o.Payload
}

func (o *GetNetworkWirelessSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkWirelessSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWirelessSettingsOKBody get network wireless settings o k body
swagger:model GetNetworkWirelessSettingsOKBody
*/
type GetNetworkWirelessSettingsOKBody struct {

	// Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	IPV6BridgeEnabled bool `json:"ipv6BridgeEnabled,omitempty"`

	// Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LedLightsOn bool `json:"ledLightsOn,omitempty"`

	// Toggle for enabling or disabling location analytics for your network
	LocationAnalyticsEnabled bool `json:"locationAnalyticsEnabled,omitempty"`

	// Toggle for enabling or disabling meshing in a network
	MeshingEnabled bool `json:"meshingEnabled,omitempty"`

	// named vlans
	NamedVlans *GetNetworkWirelessSettingsOKBodyNamedVlans `json:"namedVlans,omitempty"`

	// The upgrade strategy to apply to the network. Must be one of 'minimizeUpgradeTime' or 'minimizeClientDowntime'. Requires firmware version MR 26.8 or higher'
	// Enum: [minimizeClientDowntime minimizeUpgradeTime]
	UpgradeStrategy string `json:"upgradeStrategy,omitempty"`
}

// Validate validates this get network wireless settings o k body
func (o *GetNetworkWirelessSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNamedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpgradeStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSettingsOKBody) validateNamedVlans(formats strfmt.Registry) error {
	if swag.IsZero(o.NamedVlans) { // not required
		return nil
	}

	if o.NamedVlans != nil {
		if err := o.NamedVlans.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWirelessSettingsOK" + "." + "namedVlans")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWirelessSettingsOK" + "." + "namedVlans")
			}
			return err
		}
	}

	return nil
}

var getNetworkWirelessSettingsOKBodyTypeUpgradeStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minimizeClientDowntime","minimizeUpgradeTime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSettingsOKBodyTypeUpgradeStrategyPropEnum = append(getNetworkWirelessSettingsOKBodyTypeUpgradeStrategyPropEnum, v)
	}
}

const (

	// GetNetworkWirelessSettingsOKBodyUpgradeStrategyMinimizeClientDowntime captures enum value "minimizeClientDowntime"
	GetNetworkWirelessSettingsOKBodyUpgradeStrategyMinimizeClientDowntime string = "minimizeClientDowntime"

	// GetNetworkWirelessSettingsOKBodyUpgradeStrategyMinimizeUpgradeTime captures enum value "minimizeUpgradeTime"
	GetNetworkWirelessSettingsOKBodyUpgradeStrategyMinimizeUpgradeTime string = "minimizeUpgradeTime"
)

// prop value enum
func (o *GetNetworkWirelessSettingsOKBody) validateUpgradeStrategyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSettingsOKBodyTypeUpgradeStrategyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSettingsOKBody) validateUpgradeStrategy(formats strfmt.Registry) error {
	if swag.IsZero(o.UpgradeStrategy) { // not required
		return nil
	}

	// value enum
	if err := o.validateUpgradeStrategyEnum("getNetworkWirelessSettingsOK"+"."+"upgradeStrategy", "body", o.UpgradeStrategy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get network wireless settings o k body based on the context it is used
func (o *GetNetworkWirelessSettingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNamedVlans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSettingsOKBody) contextValidateNamedVlans(ctx context.Context, formats strfmt.Registry) error {

	if o.NamedVlans != nil {
		if err := o.NamedVlans.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWirelessSettingsOK" + "." + "namedVlans")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWirelessSettingsOK" + "." + "namedVlans")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkWirelessSettingsOKBodyNamedVlans Named VLAN settings for wireless networks.
swagger:model GetNetworkWirelessSettingsOKBodyNamedVlans
*/
type GetNetworkWirelessSettingsOKBodyNamedVlans struct {

	// pool dhcp monitoring
	PoolDhcpMonitoring *GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"`
}

// Validate validates this get network wireless settings o k body named vlans
func (o *GetNetworkWirelessSettingsOKBodyNamedVlans) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePoolDhcpMonitoring(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSettingsOKBodyNamedVlans) validatePoolDhcpMonitoring(formats strfmt.Registry) error {
	if swag.IsZero(o.PoolDhcpMonitoring) { // not required
		return nil
	}

	if o.PoolDhcpMonitoring != nil {
		if err := o.PoolDhcpMonitoring.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWirelessSettingsOK" + "." + "namedVlans" + "." + "poolDhcpMonitoring")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWirelessSettingsOK" + "." + "namedVlans" + "." + "poolDhcpMonitoring")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network wireless settings o k body named vlans based on the context it is used
func (o *GetNetworkWirelessSettingsOKBodyNamedVlans) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePoolDhcpMonitoring(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSettingsOKBodyNamedVlans) contextValidatePoolDhcpMonitoring(ctx context.Context, formats strfmt.Registry) error {

	if o.PoolDhcpMonitoring != nil {
		if err := o.PoolDhcpMonitoring.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWirelessSettingsOK" + "." + "namedVlans" + "." + "poolDhcpMonitoring")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWirelessSettingsOK" + "." + "namedVlans" + "." + "poolDhcpMonitoring")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSettingsOKBodyNamedVlans) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSettingsOKBodyNamedVlans) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSettingsOKBodyNamedVlans
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring Named VLAN Pool DHCP Monitoring settings.
swagger:model GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring
*/
type GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring struct {

	// The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Duration int64 `json:"duration,omitempty"`

	// Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this get network wireless settings o k body named vlans pool dhcp monitoring
func (o *GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network wireless settings o k body named vlans pool dhcp monitoring based on context it is used
func (o *GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSettingsOKBodyNamedVlansPoolDhcpMonitoring
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}