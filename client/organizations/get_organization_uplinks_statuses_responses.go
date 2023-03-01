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

// GetOrganizationUplinksStatusesReader is a Reader for the GetOrganizationUplinksStatuses structure.
type GetOrganizationUplinksStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationUplinksStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationUplinksStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationUplinksStatusesOK creates a GetOrganizationUplinksStatusesOK with default headers values
func NewGetOrganizationUplinksStatusesOK() *GetOrganizationUplinksStatusesOK {
	return &GetOrganizationUplinksStatusesOK{}
}

/*
GetOrganizationUplinksStatusesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationUplinksStatusesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetOrganizationUplinksStatusesOKBodyItems0
}

// IsSuccess returns true when this get organization uplinks statuses o k response has a 2xx status code
func (o *GetOrganizationUplinksStatusesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization uplinks statuses o k response has a 3xx status code
func (o *GetOrganizationUplinksStatusesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization uplinks statuses o k response has a 4xx status code
func (o *GetOrganizationUplinksStatusesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization uplinks statuses o k response has a 5xx status code
func (o *GetOrganizationUplinksStatusesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization uplinks statuses o k response a status code equal to that given
func (o *GetOrganizationUplinksStatusesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization uplinks statuses o k response
func (o *GetOrganizationUplinksStatusesOK) Code() int {
	return 200
}

func (o *GetOrganizationUplinksStatusesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/uplinks/statuses][%d] getOrganizationUplinksStatusesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationUplinksStatusesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/uplinks/statuses][%d] getOrganizationUplinksStatusesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationUplinksStatusesOK) GetPayload() []*GetOrganizationUplinksStatusesOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationUplinksStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetOrganizationUplinksStatusesOKBodyItems0 get organization uplinks statuses o k body items0
swagger:model GetOrganizationUplinksStatusesOKBodyItems0
*/
type GetOrganizationUplinksStatusesOKBodyItems0 struct {

	// Last reported time for the device
	// Format: date-time
	LastReportedAt strfmt.DateTime `json:"lastReportedAt,omitempty"`

	// The uplink model
	Model string `json:"model,omitempty"`

	// Network identifier
	NetworkID string `json:"networkId,omitempty"`

	// The uplink serial
	Serial string `json:"serial,omitempty"`

	// Uplinks
	Uplinks []*GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0 `json:"uplinks"`
}

// Validate validates this get organization uplinks statuses o k body items0
func (o *GetOrganizationUplinksStatusesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLastReportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUplinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationUplinksStatusesOKBodyItems0) validateLastReportedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.LastReportedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastReportedAt", "body", "date-time", o.LastReportedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationUplinksStatusesOKBodyItems0) validateUplinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Uplinks) { // not required
		return nil
	}

	for i := 0; i < len(o.Uplinks); i++ {
		if swag.IsZero(o.Uplinks[i]) { // not required
			continue
		}

		if o.Uplinks[i] != nil {
			if err := o.Uplinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("uplinks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("uplinks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get organization uplinks statuses o k body items0 based on the context it is used
func (o *GetOrganizationUplinksStatusesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUplinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationUplinksStatusesOKBodyItems0) contextValidateUplinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Uplinks); i++ {

		if o.Uplinks[i] != nil {
			if err := o.Uplinks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("uplinks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("uplinks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationUplinksStatusesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationUplinksStatusesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationUplinksStatusesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0 get organization uplinks statuses o k body items0 uplinks items0
swagger:model GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0
*/
type GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0 struct {

	// Access Point Name
	Apn string `json:"apn,omitempty"`

	// Connection Type
	ConnectionType string `json:"connectionType,omitempty"`

	// Primary DNS IP
	Dns1 string `json:"dns1,omitempty"`

	// Secondary DNS IP
	Dns2 string `json:"dns2,omitempty"`

	// Gateway IP
	Gateway string `json:"gateway,omitempty"`

	// Integrated Circuit Card Identification Number
	Iccid string `json:"iccid,omitempty"`

	// Uplink interface
	// Enum: [cellular wan1 wan2]
	Interface string `json:"interface,omitempty"`

	// Uplink IP
	IP string `json:"ip,omitempty"`

	// The way in which the IP is assigned
	IPAssignedBy string `json:"ipAssignedBy,omitempty"`

	// Primary DNS IP
	PrimaryDNS string `json:"primaryDns,omitempty"`

	// Network Provider
	Provider string `json:"provider,omitempty"`

	// Public IP
	PublicIP string `json:"publicIp,omitempty"`

	// Secondary DNS IP
	SecondaryDNS string `json:"secondaryDns,omitempty"`

	// signal stat
	SignalStat *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat `json:"signalStat,omitempty"`

	// Signal Type
	SignalType string `json:"signalType,omitempty"`

	// Uplink status
	// Enum: [active connecting failed not connected ready]
	Status string `json:"status,omitempty"`
}

// Validate validates this get organization uplinks statuses o k body items0 uplinks items0
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignalStat(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getOrganizationUplinksStatusesOKBodyItems0UplinksItems0TypeInterfacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cellular","wan1","wan2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationUplinksStatusesOKBodyItems0UplinksItems0TypeInterfacePropEnum = append(getOrganizationUplinksStatusesOKBodyItems0UplinksItems0TypeInterfacePropEnum, v)
	}
}

const (

	// GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0InterfaceCellular captures enum value "cellular"
	GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0InterfaceCellular string = "cellular"

	// GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0InterfaceWan1 captures enum value "wan1"
	GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0InterfaceWan1 string = "wan1"

	// GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0InterfaceWan2 captures enum value "wan2"
	GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0InterfaceWan2 string = "wan2"
)

// prop value enum
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) validateInterfaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationUplinksStatusesOKBodyItems0UplinksItems0TypeInterfacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(o.Interface) { // not required
		return nil
	}

	// value enum
	if err := o.validateInterfaceEnum("interface", "body", o.Interface); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) validateSignalStat(formats strfmt.Registry) error {
	if swag.IsZero(o.SignalStat) { // not required
		return nil
	}

	if o.SignalStat != nil {
		if err := o.SignalStat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signalStat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signalStat")
			}
			return err
		}
	}

	return nil
}

var getOrganizationUplinksStatusesOKBodyItems0UplinksItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","connecting","failed","not connected","ready"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationUplinksStatusesOKBodyItems0UplinksItems0TypeStatusPropEnum = append(getOrganizationUplinksStatusesOKBodyItems0UplinksItems0TypeStatusPropEnum, v)
	}
}

const (

	// GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusActive captures enum value "active"
	GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusActive string = "active"

	// GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusConnecting captures enum value "connecting"
	GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusConnecting string = "connecting"

	// GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusFailed captures enum value "failed"
	GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusFailed string = "failed"

	// GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusNotConnected captures enum value "not connected"
	GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusNotConnected string = "not connected"

	// GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusReady captures enum value "ready"
	GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0StatusReady string = "ready"
)

// prop value enum
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationUplinksStatusesOKBodyItems0UplinksItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get organization uplinks statuses o k body items0 uplinks items0 based on the context it is used
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSignalStat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) contextValidateSignalStat(ctx context.Context, formats strfmt.Registry) error {

	if o.SignalStat != nil {
		if err := o.SignalStat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signalStat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signalStat")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat Tower Signal Status
swagger:model GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat
*/
type GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat struct {

	// Reference Signal Received Power
	Rsrp string `json:"rsrp,omitempty"`

	// Reference Signal Received Quality
	Rsrq string `json:"rsrq,omitempty"`
}

// Validate validates this get organization uplinks statuses o k body items0 uplinks items0 signal stat
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization uplinks statuses o k body items0 uplinks items0 signal stat based on context it is used
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat) UnmarshalBinary(b []byte) error {
	var res GetOrganizationUplinksStatusesOKBodyItems0UplinksItems0SignalStat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}