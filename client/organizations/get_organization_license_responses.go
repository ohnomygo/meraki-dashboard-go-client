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

// GetOrganizationLicenseReader is a Reader for the GetOrganizationLicense structure.
type GetOrganizationLicenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationLicenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationLicenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationLicenseOK creates a GetOrganizationLicenseOK with default headers values
func NewGetOrganizationLicenseOK() *GetOrganizationLicenseOK {
	return &GetOrganizationLicenseOK{}
}

/*
GetOrganizationLicenseOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationLicenseOK struct {
	Payload *GetOrganizationLicenseOKBody
}

// IsSuccess returns true when this get organization license o k response has a 2xx status code
func (o *GetOrganizationLicenseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization license o k response has a 3xx status code
func (o *GetOrganizationLicenseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization license o k response has a 4xx status code
func (o *GetOrganizationLicenseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization license o k response has a 5xx status code
func (o *GetOrganizationLicenseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization license o k response a status code equal to that given
func (o *GetOrganizationLicenseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization license o k response
func (o *GetOrganizationLicenseOK) Code() int {
	return 200
}

func (o *GetOrganizationLicenseOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/licenses/{licenseId}][%d] getOrganizationLicenseOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLicenseOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/licenses/{licenseId}][%d] getOrganizationLicenseOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLicenseOK) GetPayload() *GetOrganizationLicenseOKBody {
	return o.Payload
}

func (o *GetOrganizationLicenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationLicenseOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationLicenseOKBody get organization license o k body
swagger:model GetOrganizationLicenseOKBody
*/
type GetOrganizationLicenseOKBody struct {

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
	PermanentlyQueuedLicenses []*GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0 `json:"permanentlyQueuedLicenses"`

	// The number of seats of the license. Only applicable to SM licenses.
	SeatCount int64 `json:"seatCount,omitempty"`

	// The state of the license. All queued licenses have a status of `recentlyQueued`.
	// Enum: [active expired expiring recentlyQueued unused unusedActive]
	State string `json:"state,omitempty"`

	// The duration of the license plus all permanently queued licenses associated with it
	TotalDurationInDays int64 `json:"totalDurationInDays,omitempty"`
}

// Validate validates this get organization license o k body
func (o *GetOrganizationLicenseOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetOrganizationLicenseOKBody) validatePermanentlyQueuedLicenses(formats strfmt.Registry) error {
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
					return ve.ValidateName("getOrganizationLicenseOK" + "." + "permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOrganizationLicenseOK" + "." + "permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getOrganizationLicenseOKBodyTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","expired","expiring","recentlyQueued","unused","unusedActive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationLicenseOKBodyTypeStatePropEnum = append(getOrganizationLicenseOKBodyTypeStatePropEnum, v)
	}
}

const (

	// GetOrganizationLicenseOKBodyStateActive captures enum value "active"
	GetOrganizationLicenseOKBodyStateActive string = "active"

	// GetOrganizationLicenseOKBodyStateExpired captures enum value "expired"
	GetOrganizationLicenseOKBodyStateExpired string = "expired"

	// GetOrganizationLicenseOKBodyStateExpiring captures enum value "expiring"
	GetOrganizationLicenseOKBodyStateExpiring string = "expiring"

	// GetOrganizationLicenseOKBodyStateRecentlyQueued captures enum value "recentlyQueued"
	GetOrganizationLicenseOKBodyStateRecentlyQueued string = "recentlyQueued"

	// GetOrganizationLicenseOKBodyStateUnused captures enum value "unused"
	GetOrganizationLicenseOKBodyStateUnused string = "unused"

	// GetOrganizationLicenseOKBodyStateUnusedActive captures enum value "unusedActive"
	GetOrganizationLicenseOKBodyStateUnusedActive string = "unusedActive"
)

// prop value enum
func (o *GetOrganizationLicenseOKBody) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationLicenseOKBodyTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationLicenseOKBody) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("getOrganizationLicenseOK"+"."+"state", "body", o.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get organization license o k body based on the context it is used
func (o *GetOrganizationLicenseOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePermanentlyQueuedLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLicenseOKBody) contextValidatePermanentlyQueuedLicenses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PermanentlyQueuedLicenses); i++ {

		if o.PermanentlyQueuedLicenses[i] != nil {
			if err := o.PermanentlyQueuedLicenses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationLicenseOK" + "." + "permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOrganizationLicenseOK" + "." + "permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLicenseOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLicenseOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLicenseOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0 get organization license o k body permanently queued licenses items0
swagger:model GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0
*/
type GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0 struct {

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

// Validate validates this get organization license o k body permanently queued licenses items0
func (o *GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization license o k body permanently queued licenses items0 based on context it is used
func (o *GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLicenseOKBodyPermanentlyQueuedLicensesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}