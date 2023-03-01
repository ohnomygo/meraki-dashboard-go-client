// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// GetOrganizationDevicesUplinksLossAndLatencyReader is a Reader for the GetOrganizationDevicesUplinksLossAndLatency structure.
type GetOrganizationDevicesUplinksLossAndLatencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationDevicesUplinksLossAndLatencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationDevicesUplinksLossAndLatencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationDevicesUplinksLossAndLatencyOK creates a GetOrganizationDevicesUplinksLossAndLatencyOK with default headers values
func NewGetOrganizationDevicesUplinksLossAndLatencyOK() *GetOrganizationDevicesUplinksLossAndLatencyOK {
	return &GetOrganizationDevicesUplinksLossAndLatencyOK{}
}

/*
GetOrganizationDevicesUplinksLossAndLatencyOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationDevicesUplinksLossAndLatencyOK struct {
	Payload []*GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0
}

// IsSuccess returns true when this get organization devices uplinks loss and latency o k response has a 2xx status code
func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization devices uplinks loss and latency o k response has a 3xx status code
func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization devices uplinks loss and latency o k response has a 4xx status code
func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization devices uplinks loss and latency o k response has a 5xx status code
func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization devices uplinks loss and latency o k response a status code equal to that given
func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization devices uplinks loss and latency o k response
func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) Code() int {
	return 200
}

func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/devices/uplinksLossAndLatency][%d] getOrganizationDevicesUplinksLossAndLatencyOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/devices/uplinksLossAndLatency][%d] getOrganizationDevicesUplinksLossAndLatencyOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) GetPayload() []*GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationDevicesUplinksLossAndLatencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0 get organization devices uplinks loss and latency o k body items0
swagger:model GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0
*/
type GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0 struct {

	// IP address of uplink
	IP string `json:"ip,omitempty"`

	// Network ID
	NetworkID string `json:"networkId,omitempty"`

	// Serial of MX device
	Serial string `json:"serial,omitempty"`

	// Loss and latency timeseries data
	TimeSeries []*GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0 `json:"timeSeries"`

	// Uplink interface (wan1, wan2, or cellular)
	Uplink string `json:"uplink,omitempty"`
}

// Validate validates this get organization devices uplinks loss and latency o k body items0
func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimeSeries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0) validateTimeSeries(formats strfmt.Registry) error {
	if swag.IsZero(o.TimeSeries) { // not required
		return nil
	}

	for i := 0; i < len(o.TimeSeries); i++ {
		if swag.IsZero(o.TimeSeries[i]) { // not required
			continue
		}

		if o.TimeSeries[i] != nil {
			if err := o.TimeSeries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("timeSeries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("timeSeries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get organization devices uplinks loss and latency o k body items0 based on the context it is used
func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTimeSeries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0) contextValidateTimeSeries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.TimeSeries); i++ {

		if o.TimeSeries[i] != nil {
			if err := o.TimeSeries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("timeSeries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("timeSeries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0 get organization devices uplinks loss and latency o k body items0 time series items0
swagger:model GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0
*/
type GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0 struct {

	// Latency in milliseconds
	LatencyMs float32 `json:"latencyMs,omitempty"`

	// Loss percentage
	LossPercent float32 `json:"lossPercent,omitempty"`

	// Timestamp for this data point
	// Format: date-time
	Ts strfmt.DateTime `json:"ts,omitempty"`
}

// Validate validates this get organization devices uplinks loss and latency o k body items0 time series items0
func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0) validateTs(formats strfmt.Registry) error {
	if swag.IsZero(o.Ts) { // not required
		return nil
	}

	if err := validate.FormatOf("ts", "body", "date-time", o.Ts.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get organization devices uplinks loss and latency o k body items0 time series items0 based on context it is used
func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationDevicesUplinksLossAndLatencyOKBodyItems0TimeSeriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
