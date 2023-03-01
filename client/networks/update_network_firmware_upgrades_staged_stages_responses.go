// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// UpdateNetworkFirmwareUpgradesStagedStagesReader is a Reader for the UpdateNetworkFirmwareUpgradesStagedStages structure.
type UpdateNetworkFirmwareUpgradesStagedStagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkFirmwareUpgradesStagedStagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkFirmwareUpgradesStagedStagesOK creates a UpdateNetworkFirmwareUpgradesStagedStagesOK with default headers values
func NewUpdateNetworkFirmwareUpgradesStagedStagesOK() *UpdateNetworkFirmwareUpgradesStagedStagesOK {
	return &UpdateNetworkFirmwareUpgradesStagedStagesOK{}
}

/*
UpdateNetworkFirmwareUpgradesStagedStagesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkFirmwareUpgradesStagedStagesOK struct {
	Payload []*UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0
}

// IsSuccess returns true when this update network firmware upgrades staged stages o k response has a 2xx status code
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network firmware upgrades staged stages o k response has a 3xx status code
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network firmware upgrades staged stages o k response has a 4xx status code
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network firmware upgrades staged stages o k response has a 5xx status code
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network firmware upgrades staged stages o k response a status code equal to that given
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network firmware upgrades staged stages o k response
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) Code() int {
	return 200
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/firmwareUpgrades/staged/stages][%d] updateNetworkFirmwareUpgradesStagedStagesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/firmwareUpgrades/staged/stages][%d] updateNetworkFirmwareUpgradesStagedStagesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) GetPayload() []*UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0 {
	return o.Payload
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkFirmwareUpgradesStagedStagesBody update network firmware upgrades staged stages body
// Example: {"_json":[{"group":{"id":"1234"}}]}
swagger:model UpdateNetworkFirmwareUpgradesStagedStagesBody
*/
type UpdateNetworkFirmwareUpgradesStagedStagesBody struct {

	// Array of Staged Upgrade Groups
	JSON []*UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0 `json:"_json"`
}

// Validate validates this update network firmware upgrades staged stages body
func (o *UpdateNetworkFirmwareUpgradesStagedStagesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateJSON(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesBody) validateJSON(formats strfmt.Registry) error {
	if swag.IsZero(o.JSON) { // not required
		return nil
	}

	for i := 0; i < len(o.JSON); i++ {
		if swag.IsZero(o.JSON[i]) { // not required
			continue
		}

		if o.JSON[i] != nil {
			if err := o.JSON[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkFirmwareUpgradesStagedStages" + "." + "_json" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkFirmwareUpgradesStagedStages" + "." + "_json" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network firmware upgrades staged stages body based on the context it is used
func (o *UpdateNetworkFirmwareUpgradesStagedStagesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateJSON(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesBody) contextValidateJSON(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.JSON); i++ {

		if o.JSON[i] != nil {
			if err := o.JSON[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkFirmwareUpgradesStagedStages" + "." + "_json" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkFirmwareUpgradesStagedStages" + "." + "_json" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFirmwareUpgradesStagedStagesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0 update network firmware upgrades staged stages o k body items0
swagger:model UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0
*/
type UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0 struct {

	// group
	Group *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group `json:"group,omitempty"`
}

// Validate validates this update network firmware upgrades staged stages o k body items0
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(o.Group) { // not required
		return nil
	}

	if o.Group != nil {
		if err := o.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network firmware upgrades staged stages o k body items0 based on the context it is used
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if o.Group != nil {
		if err := o.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group The Staged Upgrade Group
swagger:model UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group
*/
type UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group struct {

	// Description of the Staged Upgrade Group
	Description string `json:"description,omitempty"`

	// Id of the Staged Upgrade Group
	ID string `json:"id,omitempty"`

	// Name of the Staged Upgrade Group
	Name string `json:"name,omitempty"`
}

// Validate validates this update network firmware upgrades staged stages o k body items0 group
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network firmware upgrades staged stages o k body items0 group based on context it is used
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFirmwareUpgradesStagedStagesOKBodyItems0Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0 update network firmware upgrades staged stages params body JSON items0
swagger:model UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0
*/
type UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0 struct {

	// group
	Group *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group `json:"group,omitempty"`
}

// Validate validates this update network firmware upgrades staged stages params body JSON items0
func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(o.Group) { // not required
		return nil
	}

	if o.Group != nil {
		if err := o.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network firmware upgrades staged stages params body JSON items0 based on the context it is used
func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if o.Group != nil {
		if err := o.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group The Staged Upgrade Group
swagger:model UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group
*/
type UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group struct {

	// ID of the Staged Upgrade Group
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this update network firmware upgrades staged stages params body JSON items0 group
func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group) validateID(formats strfmt.Registry) error {

	if err := validate.Required("group"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network firmware upgrades staged stages params body JSON items0 group based on context it is used
func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFirmwareUpgradesStagedStagesParamsBodyJSONItems0Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}