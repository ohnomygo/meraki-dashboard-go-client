// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// GetOrganizationInventoryOnboardingCloudMonitoringImportsReader is a Reader for the GetOrganizationInventoryOnboardingCloudMonitoringImports structure.
type GetOrganizationInventoryOnboardingCloudMonitoringImportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationInventoryOnboardingCloudMonitoringImportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationInventoryOnboardingCloudMonitoringImportsOK creates a GetOrganizationInventoryOnboardingCloudMonitoringImportsOK with default headers values
func NewGetOrganizationInventoryOnboardingCloudMonitoringImportsOK() *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK {
	return &GetOrganizationInventoryOnboardingCloudMonitoringImportsOK{}
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringImportsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationInventoryOnboardingCloudMonitoringImportsOK struct {
	Payload []*GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0
}

// IsSuccess returns true when this get organization inventory onboarding cloud monitoring imports o k response has a 2xx status code
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization inventory onboarding cloud monitoring imports o k response has a 3xx status code
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization inventory onboarding cloud monitoring imports o k response has a 4xx status code
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization inventory onboarding cloud monitoring imports o k response has a 5xx status code
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization inventory onboarding cloud monitoring imports o k response a status code equal to that given
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization inventory onboarding cloud monitoring imports o k response
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) Code() int {
	return 200
}

func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports][%d] getOrganizationInventoryOnboardingCloudMonitoringImportsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports][%d] getOrganizationInventoryOnboardingCloudMonitoringImportsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) GetPayload() []*GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0 get organization inventory onboarding cloud monitoring imports o k body items0
swagger:model GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0
*/
type GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0 struct {

	// device
	Device *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device `json:"device,omitempty"`

	// Database ID for the new entity entry.
	ImportID string `json:"importId,omitempty"`
}

// Validate validates this get organization inventory onboarding cloud monitoring imports o k body items0
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(o.Device) { // not required
		return nil
	}

	if o.Device != nil {
		if err := o.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization inventory onboarding cloud monitoring imports o k body items0 based on the context it is used
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if o.Device != nil {
		if err := o.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device Represents the details of an imported device.
swagger:model GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device
*/
type GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device struct {

	// Whether or not the device was successfully created in dashboard.
	Created bool `json:"created,omitempty"`

	// Represents the current state of importing the device.
	Status string `json:"status,omitempty"`

	// The url to the device details page within dashboard.
	URL string `json:"url,omitempty"`
}

// Validate validates this get organization inventory onboarding cloud monitoring imports o k body items0 device
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization inventory onboarding cloud monitoring imports o k body items0 device based on context it is used
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device) UnmarshalBinary(b []byte) error {
	var res GetOrganizationInventoryOnboardingCloudMonitoringImportsOKBodyItems0Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
