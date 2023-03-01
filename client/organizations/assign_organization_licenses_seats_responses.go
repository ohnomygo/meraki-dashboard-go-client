// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// AssignOrganizationLicensesSeatsReader is a Reader for the AssignOrganizationLicensesSeats structure.
type AssignOrganizationLicensesSeatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignOrganizationLicensesSeatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignOrganizationLicensesSeatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignOrganizationLicensesSeatsOK creates a AssignOrganizationLicensesSeatsOK with default headers values
func NewAssignOrganizationLicensesSeatsOK() *AssignOrganizationLicensesSeatsOK {
	return &AssignOrganizationLicensesSeatsOK{}
}

/*
AssignOrganizationLicensesSeatsOK describes a response with status code 200, with default header values.

Successful operation
*/
type AssignOrganizationLicensesSeatsOK struct {
	Payload *AssignOrganizationLicensesSeatsOKBody
}

// IsSuccess returns true when this assign organization licenses seats o k response has a 2xx status code
func (o *AssignOrganizationLicensesSeatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign organization licenses seats o k response has a 3xx status code
func (o *AssignOrganizationLicensesSeatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign organization licenses seats o k response has a 4xx status code
func (o *AssignOrganizationLicensesSeatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign organization licenses seats o k response has a 5xx status code
func (o *AssignOrganizationLicensesSeatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign organization licenses seats o k response a status code equal to that given
func (o *AssignOrganizationLicensesSeatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign organization licenses seats o k response
func (o *AssignOrganizationLicensesSeatsOK) Code() int {
	return 200
}

func (o *AssignOrganizationLicensesSeatsOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/licenses/assignSeats][%d] assignOrganizationLicensesSeatsOK  %+v", 200, o.Payload)
}

func (o *AssignOrganizationLicensesSeatsOK) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/licenses/assignSeats][%d] assignOrganizationLicensesSeatsOK  %+v", 200, o.Payload)
}

func (o *AssignOrganizationLicensesSeatsOK) GetPayload() *AssignOrganizationLicensesSeatsOKBody {
	return o.Payload
}

func (o *AssignOrganizationLicensesSeatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AssignOrganizationLicensesSeatsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AssignOrganizationLicensesSeatsBody assign organization licenses seats body
// Example: {"licenseId":"1234","networkId":"N_24329156","seatCount":20}
swagger:model AssignOrganizationLicensesSeatsBody
*/
type AssignOrganizationLicensesSeatsBody struct {

	// The ID of the SM license to assign seats from
	// Required: true
	LicenseID *string `json:"licenseId"`

	// The ID of the SM network to assign the seats to
	// Required: true
	NetworkID *string `json:"networkId"`

	// The number of seats to assign to the SM network. Must be less than or equal to the total number of seats of the license
	// Required: true
	SeatCount *int64 `json:"seatCount"`
}

// Validate validates this assign organization licenses seats body
func (o *AssignOrganizationLicensesSeatsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLicenseID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSeatCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignOrganizationLicensesSeatsBody) validateLicenseID(formats strfmt.Registry) error {

	if err := validate.Required("assignOrganizationLicensesSeats"+"."+"licenseId", "body", o.LicenseID); err != nil {
		return err
	}

	return nil
}

func (o *AssignOrganizationLicensesSeatsBody) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("assignOrganizationLicensesSeats"+"."+"networkId", "body", o.NetworkID); err != nil {
		return err
	}

	return nil
}

func (o *AssignOrganizationLicensesSeatsBody) validateSeatCount(formats strfmt.Registry) error {

	if err := validate.Required("assignOrganizationLicensesSeats"+"."+"seatCount", "body", o.SeatCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this assign organization licenses seats body based on context it is used
func (o *AssignOrganizationLicensesSeatsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsBody) UnmarshalBinary(b []byte) error {
	var res AssignOrganizationLicensesSeatsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AssignOrganizationLicensesSeatsOKBody assign organization licenses seats o k body
swagger:model AssignOrganizationLicensesSeatsOKBody
*/
type AssignOrganizationLicensesSeatsOKBody struct {

	// Resulting licenses from the move
	ResultingLicenses []*AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0 `json:"resultingLicenses"`
}

// Validate validates this assign organization licenses seats o k body
func (o *AssignOrganizationLicensesSeatsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResultingLicenses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignOrganizationLicensesSeatsOKBody) validateResultingLicenses(formats strfmt.Registry) error {
	if swag.IsZero(o.ResultingLicenses) { // not required
		return nil
	}

	for i := 0; i < len(o.ResultingLicenses); i++ {
		if swag.IsZero(o.ResultingLicenses[i]) { // not required
			continue
		}

		if o.ResultingLicenses[i] != nil {
			if err := o.ResultingLicenses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignOrganizationLicensesSeatsOK" + "." + "resultingLicenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignOrganizationLicensesSeatsOK" + "." + "resultingLicenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this assign organization licenses seats o k body based on the context it is used
func (o *AssignOrganizationLicensesSeatsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResultingLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignOrganizationLicensesSeatsOKBody) contextValidateResultingLicenses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ResultingLicenses); i++ {

		if o.ResultingLicenses[i] != nil {
			if err := o.ResultingLicenses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignOrganizationLicensesSeatsOK" + "." + "resultingLicenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignOrganizationLicensesSeatsOK" + "." + "resultingLicenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsOKBody) UnmarshalBinary(b []byte) error {
	var res AssignOrganizationLicensesSeatsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0 assign organization licenses seats o k body resulting licenses items0
swagger:model AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0
*/
type AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0 struct {

	// The date the license started burning
	ActivationDate string `json:"activationDate,omitempty"`

	// The date the license was claimed into the organization
	ClaimDate string `json:"claimDate,omitempty"`

	// Serial number of the device the license is assigned to
	DeviceSerial string `json:"deviceSerial,omitempty"`

	// The duration of the individual license
	DurationInDays int64 `json:"durationInDays,omitempty"`

	// The date the license will expire
	ExpirationDate string `json:"expirationDate,omitempty"`

	// The id of the head license this license is queued behind. If there is no head license, it returns nil.
	HeadLicenseID string `json:"headLicenseId,omitempty"`

	// License ID
	ID string `json:"id,omitempty"`

	// License key
	LicenseKey string `json:"licenseKey,omitempty"`

	// License type
	LicenseType string `json:"licenseType,omitempty"`

	// ID of the network the license is assigned to
	NetworkID string `json:"networkId,omitempty"`

	// Order number
	OrderNumber string `json:"orderNumber,omitempty"`

	// DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.
	PermanentlyQueuedLicenses []*AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0 `json:"permanentlyQueuedLicenses"`

	// The number of seats of the license. Only applicable to SM licenses.
	SeatCount int64 `json:"seatCount,omitempty"`

	// The state of the license. All queued licenses have a status of `recentlyQueued`.
	// Enum: [active expired expiring recentlyQueued unused unusedActive]
	State string `json:"state,omitempty"`

	// The duration of the license plus all permanently queued licenses associated with it
	TotalDurationInDays int64 `json:"totalDurationInDays,omitempty"`
}

// Validate validates this assign organization licenses seats o k body resulting licenses items0
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePermanentlyQueuedLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0) validatePermanentlyQueuedLicenses(formats strfmt.Registry) error {
	if swag.IsZero(o.PermanentlyQueuedLicenses) { // not required
		return nil
	}

	for i := 0; i < len(o.PermanentlyQueuedLicenses); i++ {
		if swag.IsZero(o.PermanentlyQueuedLicenses[i]) { // not required
			continue
		}

		if o.PermanentlyQueuedLicenses[i] != nil {
			if err := o.PermanentlyQueuedLicenses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var assignOrganizationLicensesSeatsOKBodyResultingLicensesItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","expired","expiring","recentlyQueued","unused","unusedActive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assignOrganizationLicensesSeatsOKBodyResultingLicensesItems0TypeStatePropEnum = append(assignOrganizationLicensesSeatsOKBodyResultingLicensesItems0TypeStatePropEnum, v)
	}
}

const (

	// AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateActive captures enum value "active"
	AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateActive string = "active"

	// AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateExpired captures enum value "expired"
	AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateExpired string = "expired"

	// AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateExpiring captures enum value "expiring"
	AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateExpiring string = "expiring"

	// AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateRecentlyQueued captures enum value "recentlyQueued"
	AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateRecentlyQueued string = "recentlyQueued"

	// AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateUnused captures enum value "unused"
	AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateUnused string = "unused"

	// AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateUnusedActive captures enum value "unusedActive"
	AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0StateUnusedActive string = "unusedActive"
)

// prop value enum
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assignOrganizationLicensesSeatsOKBodyResultingLicensesItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("state", "body", o.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this assign organization licenses seats o k body resulting licenses items0 based on the context it is used
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePermanentlyQueuedLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0) contextValidatePermanentlyQueuedLicenses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PermanentlyQueuedLicenses); i++ {

		if o.PermanentlyQueuedLicenses[i] != nil {
			if err := o.PermanentlyQueuedLicenses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0) UnmarshalBinary(b []byte) error {
	var res AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0 assign organization licenses seats o k body resulting licenses items0 permanently queued licenses items0
swagger:model AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0
*/
type AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0 struct {

	// The duration of the individual license
	DurationInDays int64 `json:"durationInDays,omitempty"`

	// Permanently queued license ID
	ID string `json:"id,omitempty"`

	// License key
	LicenseKey string `json:"licenseKey,omitempty"`

	// License type
	LicenseType string `json:"licenseType,omitempty"`

	// Order number
	OrderNumber string `json:"orderNumber,omitempty"`
}

// Validate validates this assign organization licenses seats o k body resulting licenses items0 permanently queued licenses items0
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assign organization licenses seats o k body resulting licenses items0 permanently queued licenses items0 based on context it is used
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0) UnmarshalBinary(b []byte) error {
	var res AssignOrganizationLicensesSeatsOKBodyResultingLicensesItems0PermanentlyQueuedLicensesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
