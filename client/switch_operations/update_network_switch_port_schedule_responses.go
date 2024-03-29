// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// UpdateNetworkSwitchPortScheduleReader is a Reader for the UpdateNetworkSwitchPortSchedule structure.
type UpdateNetworkSwitchPortScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchPortScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchPortScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkSwitchPortScheduleOK creates a UpdateNetworkSwitchPortScheduleOK with default headers values
func NewUpdateNetworkSwitchPortScheduleOK() *UpdateNetworkSwitchPortScheduleOK {
	return &UpdateNetworkSwitchPortScheduleOK{}
}

/*
UpdateNetworkSwitchPortScheduleOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchPortScheduleOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network switch port schedule o k response has a 2xx status code
func (o *UpdateNetworkSwitchPortScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network switch port schedule o k response has a 3xx status code
func (o *UpdateNetworkSwitchPortScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network switch port schedule o k response has a 4xx status code
func (o *UpdateNetworkSwitchPortScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network switch port schedule o k response has a 5xx status code
func (o *UpdateNetworkSwitchPortScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network switch port schedule o k response a status code equal to that given
func (o *UpdateNetworkSwitchPortScheduleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network switch port schedule o k response
func (o *UpdateNetworkSwitchPortScheduleOK) Code() int {
	return 200
}

func (o *UpdateNetworkSwitchPortScheduleOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/portSchedules/{portScheduleId}][%d] updateNetworkSwitchPortScheduleOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchPortScheduleOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/portSchedules/{portScheduleId}][%d] updateNetworkSwitchPortScheduleOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchPortScheduleOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchPortScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkSwitchPortScheduleBody update network switch port schedule body
// Example: {"name":"Weekdays schedule","portSchedule":{"friday":{"active":true,"from":"9:00","to":"17:00"},"monday":{"active":true,"from":"9:00","to":"17:00"},"saturday":{"active":false,"from":"0:00","to":"24:00"},"sunday":{"active":false,"from":"0:00","to":"24:00"},"thursday":{"active":true,"from":"9:00","to":"17:00"},"tuesday":{"active":true,"from":"9:00","to":"17:00"},"wednesday":{"active":true,"from":"9:00","to":"17:00"}}}
swagger:model UpdateNetworkSwitchPortScheduleBody
*/
type UpdateNetworkSwitchPortScheduleBody struct {

	// The name for your port schedule.
	Name string `json:"name,omitempty"`

	// port schedule
	PortSchedule *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule `json:"portSchedule,omitempty"`
}

// Validate validates this update network switch port schedule body
func (o *UpdateNetworkSwitchPortScheduleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePortSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchPortScheduleBody) validatePortSchedule(formats strfmt.Registry) error {
	if swag.IsZero(o.PortSchedule) { // not required
		return nil
	}

	if o.PortSchedule != nil {
		if err := o.PortSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network switch port schedule body based on the context it is used
func (o *UpdateNetworkSwitchPortScheduleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePortSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchPortScheduleBody) contextValidatePortSchedule(ctx context.Context, formats strfmt.Registry) error {

	if o.PortSchedule != nil {
		if err := o.PortSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule     The schedule for switch port scheduling. Schedules are applied to days of the week.
//     When it's empty, default schedule with all days of a week are configured.
//     Any unspecified day in the schedule is added as a default schedule configuration of the day.
//
swagger:model UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule
*/
type UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule struct {

	// friday
	Friday *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday `json:"friday,omitempty"`

	// monday
	Monday *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday `json:"monday,omitempty"`

	// saturday
	Saturday *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday `json:"saturday,omitempty"`

	// sunday
	Sunday *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday `json:"sunday,omitempty"`

	// thursday
	Thursday *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday `json:"thursday,omitempty"`

	// tuesday
	Tuesday *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday `json:"tuesday,omitempty"`

	// wednesday
	Wednesday *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday `json:"wednesday,omitempty"`
}

// Validate validates this update network switch port schedule params body port schedule
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFriday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMonday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSaturday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSunday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateThursday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTuesday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWednesday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateFriday(formats strfmt.Registry) error {
	if swag.IsZero(o.Friday) { // not required
		return nil
	}

	if o.Friday != nil {
		if err := o.Friday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "friday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "friday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateMonday(formats strfmt.Registry) error {
	if swag.IsZero(o.Monday) { // not required
		return nil
	}

	if o.Monday != nil {
		if err := o.Monday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "monday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "monday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateSaturday(formats strfmt.Registry) error {
	if swag.IsZero(o.Saturday) { // not required
		return nil
	}

	if o.Saturday != nil {
		if err := o.Saturday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "saturday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "saturday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateSunday(formats strfmt.Registry) error {
	if swag.IsZero(o.Sunday) { // not required
		return nil
	}

	if o.Sunday != nil {
		if err := o.Sunday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "sunday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "sunday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateThursday(formats strfmt.Registry) error {
	if swag.IsZero(o.Thursday) { // not required
		return nil
	}

	if o.Thursday != nil {
		if err := o.Thursday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "thursday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "thursday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateTuesday(formats strfmt.Registry) error {
	if swag.IsZero(o.Tuesday) { // not required
		return nil
	}

	if o.Tuesday != nil {
		if err := o.Tuesday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "tuesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "tuesday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateWednesday(formats strfmt.Registry) error {
	if swag.IsZero(o.Wednesday) { // not required
		return nil
	}

	if o.Wednesday != nil {
		if err := o.Wednesday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "wednesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "wednesday")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network switch port schedule params body port schedule based on the context it is used
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFriday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMonday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSaturday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSunday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateThursday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTuesday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWednesday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateFriday(ctx context.Context, formats strfmt.Registry) error {

	if o.Friday != nil {
		if err := o.Friday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "friday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "friday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateMonday(ctx context.Context, formats strfmt.Registry) error {

	if o.Monday != nil {
		if err := o.Monday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "monday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "monday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateSaturday(ctx context.Context, formats strfmt.Registry) error {

	if o.Saturday != nil {
		if err := o.Saturday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "saturday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "saturday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateSunday(ctx context.Context, formats strfmt.Registry) error {

	if o.Sunday != nil {
		if err := o.Sunday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "sunday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "sunday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateThursday(ctx context.Context, formats strfmt.Registry) error {

	if o.Thursday != nil {
		if err := o.Thursday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "thursday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "thursday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateTuesday(ctx context.Context, formats strfmt.Registry) error {

	if o.Tuesday != nil {
		if err := o.Tuesday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "tuesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "tuesday")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateWednesday(ctx context.Context, formats strfmt.Registry) error {

	if o.Wednesday != nil {
		if err := o.Wednesday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "wednesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "wednesday")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleParamsBodyPortSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday The schedule object for Friday.
swagger:model UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday
*/
type UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this update network switch port schedule params body port schedule friday
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch port schedule params body port schedule friday based on context it is used
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday The schedule object for Monday.
swagger:model UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday
*/
type UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this update network switch port schedule params body port schedule monday
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch port schedule params body port schedule monday based on context it is used
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday The schedule object for Saturday.
swagger:model UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday
*/
type UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this update network switch port schedule params body port schedule saturday
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch port schedule params body port schedule saturday based on context it is used
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday The schedule object for Sunday.
swagger:model UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday
*/
type UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this update network switch port schedule params body port schedule sunday
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch port schedule params body port schedule sunday based on context it is used
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday The schedule object for Thursday.
swagger:model UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday
*/
type UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this update network switch port schedule params body port schedule thursday
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch port schedule params body port schedule thursday based on context it is used
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday The schedule object for Tuesday.
swagger:model UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday
*/
type UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this update network switch port schedule params body port schedule tuesday
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch port schedule params body port schedule tuesday based on context it is used
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday The schedule object for Wednesday.
swagger:model UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday
*/
type UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this update network switch port schedule params body port schedule wednesday
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch port schedule params body port schedule wednesday based on context it is used
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
