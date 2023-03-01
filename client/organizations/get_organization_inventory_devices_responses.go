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
	"github.com/go-openapi/validate"
)

// GetOrganizationInventoryDevicesReader is a Reader for the GetOrganizationInventoryDevices structure.
type GetOrganizationInventoryDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationInventoryDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationInventoryDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationInventoryDevicesOK creates a GetOrganizationInventoryDevicesOK with default headers values
func NewGetOrganizationInventoryDevicesOK() *GetOrganizationInventoryDevicesOK {
	return &GetOrganizationInventoryDevicesOK{}
}

/*
GetOrganizationInventoryDevicesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationInventoryDevicesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetOrganizationInventoryDevicesOKBodyItems0
}

// IsSuccess returns true when this get organization inventory devices o k response has a 2xx status code
func (o *GetOrganizationInventoryDevicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization inventory devices o k response has a 3xx status code
func (o *GetOrganizationInventoryDevicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization inventory devices o k response has a 4xx status code
func (o *GetOrganizationInventoryDevicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization inventory devices o k response has a 5xx status code
func (o *GetOrganizationInventoryDevicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization inventory devices o k response a status code equal to that given
func (o *GetOrganizationInventoryDevicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization inventory devices o k response
func (o *GetOrganizationInventoryDevicesOK) Code() int {
	return 200
}

func (o *GetOrganizationInventoryDevicesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/inventory/devices][%d] getOrganizationInventoryDevicesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInventoryDevicesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/inventory/devices][%d] getOrganizationInventoryDevicesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInventoryDevicesOK) GetPayload() []*GetOrganizationInventoryDevicesOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationInventoryDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationInventoryDevicesOKBodyItems0 get organization inventory devices o k body items0
swagger:model GetOrganizationInventoryDevicesOKBodyItems0
*/
type GetOrganizationInventoryDevicesOKBodyItems0 struct {

	// Claimed time of the device
	// Format: date-time
	ClaimedAt strfmt.DateTime `json:"claimedAt,omitempty"`

	// License expiration date of the device
	// Format: date-time
	LicenseExpirationDate strfmt.DateTime `json:"licenseExpirationDate,omitempty"`

	// MAC address of the device
	Mac string `json:"mac,omitempty"`

	// Model type of the device
	Model string `json:"model,omitempty"`

	// Name of the device
	Name string `json:"name,omitempty"`

	// Network Id of the device
	NetworkID string `json:"networkId,omitempty"`

	// Order number of the device
	OrderNumber string `json:"orderNumber,omitempty"`

	// Product type of the device
	ProductType string `json:"productType,omitempty"`

	// Serial number of the device
	Serial string `json:"serial,omitempty"`

	// Device tags
	Tags []string `json:"tags"`
}

// Validate validates this get organization inventory devices o k body items0
func (o *GetOrganizationInventoryDevicesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClaimedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLicenseExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationInventoryDevicesOKBodyItems0) validateClaimedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.ClaimedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("claimedAt", "body", "date-time", o.ClaimedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationInventoryDevicesOKBodyItems0) validateLicenseExpirationDate(formats strfmt.Registry) error {
	if swag.IsZero(o.LicenseExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("licenseExpirationDate", "body", "date-time", o.LicenseExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get organization inventory devices o k body items0 based on context it is used
func (o *GetOrganizationInventoryDevicesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationInventoryDevicesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationInventoryDevicesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationInventoryDevicesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
