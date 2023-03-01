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

// CreateNetworkFirmwareUpgradesStagedGroupReader is a Reader for the CreateNetworkFirmwareUpgradesStagedGroup structure.
type CreateNetworkFirmwareUpgradesStagedGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkFirmwareUpgradesStagedGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNetworkFirmwareUpgradesStagedGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkFirmwareUpgradesStagedGroupOK creates a CreateNetworkFirmwareUpgradesStagedGroupOK with default headers values
func NewCreateNetworkFirmwareUpgradesStagedGroupOK() *CreateNetworkFirmwareUpgradesStagedGroupOK {
	return &CreateNetworkFirmwareUpgradesStagedGroupOK{}
}

/*
CreateNetworkFirmwareUpgradesStagedGroupOK describes a response with status code 200, with default header values.

Successful operation
*/
type CreateNetworkFirmwareUpgradesStagedGroupOK struct {
	Payload interface{}
}

// IsSuccess returns true when this create network firmware upgrades staged group o k response has a 2xx status code
func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network firmware upgrades staged group o k response has a 3xx status code
func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network firmware upgrades staged group o k response has a 4xx status code
func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network firmware upgrades staged group o k response has a 5xx status code
func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create network firmware upgrades staged group o k response a status code equal to that given
func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create network firmware upgrades staged group o k response
func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) Code() int {
	return 200
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/firmwareUpgrades/staged/groups][%d] createNetworkFirmwareUpgradesStagedGroupOK  %+v", 200, o.Payload)
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/firmwareUpgrades/staged/groups][%d] createNetworkFirmwareUpgradesStagedGroupOK  %+v", 200, o.Payload)
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkFirmwareUpgradesStagedGroupBody create network firmware upgrades staged group body
// Example: {"assignedDevices":{"devices":[{"name":"Device Name","serial":"Q234-ABCD-5678"}],"switchStacks":[{"id":"1234","name":"Stack Name"}]},"description":"The description of the group","isDefault":false,"name":"My Staged Upgrade Group"}
swagger:model CreateNetworkFirmwareUpgradesStagedGroupBody
*/
type CreateNetworkFirmwareUpgradesStagedGroupBody struct {

	// assigned devices
	AssignedDevices *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices `json:"assignedDevices,omitempty"`

	// Description of the Staged Upgrade Group. Length must be 1 to 255 characters
	Description string `json:"description,omitempty"`

	// Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	// Required: true
	IsDefault *bool `json:"isDefault"`

	// Name of the Staged Upgrade Group. Length must be 1 to 255 characters
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this create network firmware upgrades staged group body
func (o *CreateNetworkFirmwareUpgradesStagedGroupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAssignedDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupBody) validateAssignedDevices(formats strfmt.Registry) error {
	if swag.IsZero(o.AssignedDevices) { // not required
		return nil
	}

	if o.AssignedDevices != nil {
		if err := o.AssignedDevices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupBody) validateIsDefault(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkFirmwareUpgradesStagedGroup"+"."+"isDefault", "body", o.IsDefault); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkFirmwareUpgradesStagedGroup"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network firmware upgrades staged group body based on the context it is used
func (o *CreateNetworkFirmwareUpgradesStagedGroupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAssignedDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupBody) contextValidateAssignedDevices(ctx context.Context, formats strfmt.Registry) error {

	if o.AssignedDevices != nil {
		if err := o.AssignedDevices.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFirmwareUpgradesStagedGroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFirmwareUpgradesStagedGroupBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFirmwareUpgradesStagedGroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices The devices and Switch Stacks assigned to the Group
swagger:model CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices
*/
type CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices struct {

	// Data Array of Devices containing the name and serial
	Devices []*CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0 `json:"devices"`

	// Data Array of Switch Stacks containing the name and id
	SwitchStacks []*CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0 `json:"switchStacks"`
}

// Validate validates this create network firmware upgrades staged group params body assigned devices
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSwitchStacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices) validateDevices(formats strfmt.Registry) error {
	if swag.IsZero(o.Devices) { // not required
		return nil
	}

	for i := 0; i < len(o.Devices); i++ {
		if swag.IsZero(o.Devices[i]) { // not required
			continue
		}

		if o.Devices[i] != nil {
			if err := o.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices" + "." + "devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices" + "." + "devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices) validateSwitchStacks(formats strfmt.Registry) error {
	if swag.IsZero(o.SwitchStacks) { // not required
		return nil
	}

	for i := 0; i < len(o.SwitchStacks); i++ {
		if swag.IsZero(o.SwitchStacks[i]) { // not required
			continue
		}

		if o.SwitchStacks[i] != nil {
			if err := o.SwitchStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices" + "." + "switchStacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices" + "." + "switchStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create network firmware upgrades staged group params body assigned devices based on the context it is used
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSwitchStacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Devices); i++ {

		if o.Devices[i] != nil {
			if err := o.Devices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices" + "." + "devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices" + "." + "devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices) contextValidateSwitchStacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SwitchStacks); i++ {

		if o.SwitchStacks[i] != nil {
			if err := o.SwitchStacks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices" + "." + "switchStacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkFirmwareUpgradesStagedGroup" + "." + "assignedDevices" + "." + "switchStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0 create network firmware upgrades staged group params body assigned devices devices items0
swagger:model CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0
*/
type CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0 struct {

	// Name of the device
	Name string `json:"name,omitempty"`

	// Serial of the device
	// Required: true
	Serial *string `json:"serial"`
}

// Validate validates this create network firmware upgrades staged group params body assigned devices devices items0
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", o.Serial); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network firmware upgrades staged group params body assigned devices devices items0 based on context it is used
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesDevicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0 create network firmware upgrades staged group params body assigned devices switch stacks items0
swagger:model CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0
*/
type CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0 struct {

	// ID of the Switch Stack
	// Required: true
	ID *string `json:"id"`

	// Name of the Switch Stack
	Name string `json:"name,omitempty"`
}

// Validate validates this create network firmware upgrades staged group params body assigned devices switch stacks items0
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network firmware upgrades staged group params body assigned devices switch stacks items0 based on context it is used
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFirmwareUpgradesStagedGroupParamsBodyAssignedDevicesSwitchStacksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
