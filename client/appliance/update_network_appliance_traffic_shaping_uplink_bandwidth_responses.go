// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// UpdateNetworkApplianceTrafficShapingUplinkBandwidthReader is a Reader for the UpdateNetworkApplianceTrafficShapingUplinkBandwidth structure.
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthOK creates a UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK with default headers values
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthOK() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK {
	return &UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK{}
}

/*
UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance traffic shaping uplink bandwidth o k response has a 2xx status code
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance traffic shaping uplink bandwidth o k response has a 3xx status code
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance traffic shaping uplink bandwidth o k response has a 4xx status code
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance traffic shaping uplink bandwidth o k response has a 5xx status code
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance traffic shaping uplink bandwidth o k response a status code equal to that given
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network appliance traffic shaping uplink bandwidth o k response
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) Code() int {
	return 200
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth][%d] updateNetworkApplianceTrafficShapingUplinkBandwidthOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth][%d] updateNetworkApplianceTrafficShapingUplinkBandwidthOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody update network appliance traffic shaping uplink bandwidth body
// Example: {"bandwidthLimits":{"cellular":{"limitDown":51200,"limitUp":51200},"wan1":{"limitDown":1000000,"limitUp":1000000},"wan2":{"limitDown":1000000,"limitUp":1000000}}}
swagger:model UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody
*/
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody struct {

	// bandwidth limits
	BandwidthLimits *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// Validate validates this update network appliance traffic shaping uplink bandwidth body
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBandwidthLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody) validateBandwidthLimits(formats strfmt.Registry) error {
	if swag.IsZero(o.BandwidthLimits) { // not required
		return nil
	}

	if o.BandwidthLimits != nil {
		if err := o.BandwidthLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance traffic shaping uplink bandwidth body based on the context it is used
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBandwidthLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody) contextValidateBandwidthLimits(ctx context.Context, formats strfmt.Registry) error {

	if o.BandwidthLimits != nil {
		if err := o.BandwidthLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingUplinkBandwidthBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits A mapping of uplinks to their bandwidth settings (be sure to check which uplinks are supported for your network)
swagger:model UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits
*/
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits struct {

	// cellular
	Cellular *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular `json:"cellular,omitempty"`

	// wan1
	Wan1 *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1 `json:"wan1,omitempty"`

	// wan2
	Wan2 *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2 `json:"wan2,omitempty"`
}

// Validate validates this update network appliance traffic shaping uplink bandwidth params body bandwidth limits
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCellular(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWan1(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWan2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) validateCellular(formats strfmt.Registry) error {
	if swag.IsZero(o.Cellular) { // not required
		return nil
	}

	if o.Cellular != nil {
		if err := o.Cellular.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "cellular")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "cellular")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) validateWan1(formats strfmt.Registry) error {
	if swag.IsZero(o.Wan1) { // not required
		return nil
	}

	if o.Wan1 != nil {
		if err := o.Wan1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "wan1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "wan1")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) validateWan2(formats strfmt.Registry) error {
	if swag.IsZero(o.Wan2) { // not required
		return nil
	}

	if o.Wan2 != nil {
		if err := o.Wan2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "wan2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "wan2")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance traffic shaping uplink bandwidth params body bandwidth limits based on the context it is used
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCellular(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWan1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWan2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) contextValidateCellular(ctx context.Context, formats strfmt.Registry) error {

	if o.Cellular != nil {
		if err := o.Cellular.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "cellular")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "cellular")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) contextValidateWan1(ctx context.Context, formats strfmt.Registry) error {

	if o.Wan1 != nil {
		if err := o.Wan1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "wan1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "wan1")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) contextValidateWan2(ctx context.Context, formats strfmt.Registry) error {

	if o.Wan2 != nil {
		if err := o.Wan2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "wan2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShapingUplinkBandwidth" + "." + "bandwidthLimits" + "." + "wan2")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular The bandwidth settings for the 'cellular' uplink
swagger:model UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular
*/
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular struct {

	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown int64 `json:"limitDown,omitempty"`

	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp int64 `json:"limitUp,omitempty"`
}

// Validate validates this update network appliance traffic shaping uplink bandwidth params body bandwidth limits cellular
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance traffic shaping uplink bandwidth params body bandwidth limits cellular based on context it is used
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsCellular
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1 The bandwidth settings for the 'wan1' uplink
swagger:model UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1
*/
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1 struct {

	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown int64 `json:"limitDown,omitempty"`

	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp int64 `json:"limitUp,omitempty"`
}

// Validate validates this update network appliance traffic shaping uplink bandwidth params body bandwidth limits wan1
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance traffic shaping uplink bandwidth params body bandwidth limits wan1 based on context it is used
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2 The bandwidth settings for the 'wan2' uplink
swagger:model UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2
*/
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2 struct {

	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown int64 `json:"limitDown,omitempty"`

	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp int64 `json:"limitUp,omitempty"`
}

// Validate validates this update network appliance traffic shaping uplink bandwidth params body bandwidth limits wan2
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance traffic shaping uplink bandwidth params body bandwidth limits wan2 based on context it is used
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingUplinkBandwidthParamsBodyBandwidthLimitsWan2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
