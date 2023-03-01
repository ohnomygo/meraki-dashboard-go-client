// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// GetNetworkWirelessLatencyHistoryReader is a Reader for the GetNetworkWirelessLatencyHistory structure.
type GetNetworkWirelessLatencyHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessLatencyHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessLatencyHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessLatencyHistoryOK creates a GetNetworkWirelessLatencyHistoryOK with default headers values
func NewGetNetworkWirelessLatencyHistoryOK() *GetNetworkWirelessLatencyHistoryOK {
	return &GetNetworkWirelessLatencyHistoryOK{}
}

/*
GetNetworkWirelessLatencyHistoryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessLatencyHistoryOK struct {
	Payload []*GetNetworkWirelessLatencyHistoryOKBodyItems0
}

// IsSuccess returns true when this get network wireless latency history o k response has a 2xx status code
func (o *GetNetworkWirelessLatencyHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless latency history o k response has a 3xx status code
func (o *GetNetworkWirelessLatencyHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless latency history o k response has a 4xx status code
func (o *GetNetworkWirelessLatencyHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless latency history o k response has a 5xx status code
func (o *GetNetworkWirelessLatencyHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless latency history o k response a status code equal to that given
func (o *GetNetworkWirelessLatencyHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless latency history o k response
func (o *GetNetworkWirelessLatencyHistoryOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessLatencyHistoryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/latencyHistory][%d] getNetworkWirelessLatencyHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessLatencyHistoryOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/latencyHistory][%d] getNetworkWirelessLatencyHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessLatencyHistoryOK) GetPayload() []*GetNetworkWirelessLatencyHistoryOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkWirelessLatencyHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWirelessLatencyHistoryOKBodyItems0 get network wireless latency history o k body items0
swagger:model GetNetworkWirelessLatencyHistoryOKBodyItems0
*/
type GetNetworkWirelessLatencyHistoryOKBodyItems0 struct {

	// Average latency in milliseconds
	AvgLatencyMs int64 `json:"avgLatencyMs,omitempty"`

	// The end time of the query range
	// Format: date-time
	EndTs strfmt.DateTime `json:"endTs,omitempty"`

	// The start time of the query range
	// Format: date-time
	StartTs strfmt.DateTime `json:"startTs,omitempty"`
}

// Validate validates this get network wireless latency history o k body items0
func (o *GetNetworkWirelessLatencyHistoryOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEndTs(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartTs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessLatencyHistoryOKBodyItems0) validateEndTs(formats strfmt.Registry) error {
	if swag.IsZero(o.EndTs) { // not required
		return nil
	}

	if err := validate.FormatOf("endTs", "body", "date-time", o.EndTs.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkWirelessLatencyHistoryOKBodyItems0) validateStartTs(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTs) { // not required
		return nil
	}

	if err := validate.FormatOf("startTs", "body", "date-time", o.StartTs.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network wireless latency history o k body items0 based on context it is used
func (o *GetNetworkWirelessLatencyHistoryOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessLatencyHistoryOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessLatencyHistoryOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessLatencyHistoryOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
