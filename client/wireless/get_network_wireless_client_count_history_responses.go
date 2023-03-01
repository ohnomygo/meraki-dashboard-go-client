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

// GetNetworkWirelessClientCountHistoryReader is a Reader for the GetNetworkWirelessClientCountHistory structure.
type GetNetworkWirelessClientCountHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessClientCountHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessClientCountHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessClientCountHistoryOK creates a GetNetworkWirelessClientCountHistoryOK with default headers values
func NewGetNetworkWirelessClientCountHistoryOK() *GetNetworkWirelessClientCountHistoryOK {
	return &GetNetworkWirelessClientCountHistoryOK{}
}

/*
GetNetworkWirelessClientCountHistoryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessClientCountHistoryOK struct {
	Payload []*GetNetworkWirelessClientCountHistoryOKBodyItems0
}

// IsSuccess returns true when this get network wireless client count history o k response has a 2xx status code
func (o *GetNetworkWirelessClientCountHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless client count history o k response has a 3xx status code
func (o *GetNetworkWirelessClientCountHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless client count history o k response has a 4xx status code
func (o *GetNetworkWirelessClientCountHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless client count history o k response has a 5xx status code
func (o *GetNetworkWirelessClientCountHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless client count history o k response a status code equal to that given
func (o *GetNetworkWirelessClientCountHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless client count history o k response
func (o *GetNetworkWirelessClientCountHistoryOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessClientCountHistoryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/clientCountHistory][%d] getNetworkWirelessClientCountHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessClientCountHistoryOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/clientCountHistory][%d] getNetworkWirelessClientCountHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessClientCountHistoryOK) GetPayload() []*GetNetworkWirelessClientCountHistoryOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkWirelessClientCountHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWirelessClientCountHistoryOKBodyItems0 get network wireless client count history o k body items0
swagger:model GetNetworkWirelessClientCountHistoryOKBodyItems0
*/
type GetNetworkWirelessClientCountHistoryOKBodyItems0 struct {

	// Number of connected clients
	ClientCount int64 `json:"clientCount,omitempty"`

	// The end time of the query range
	// Format: date-time
	EndTs strfmt.DateTime `json:"endTs,omitempty"`

	// The start time of the query range
	// Format: date-time
	StartTs strfmt.DateTime `json:"startTs,omitempty"`
}

// Validate validates this get network wireless client count history o k body items0
func (o *GetNetworkWirelessClientCountHistoryOKBodyItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetNetworkWirelessClientCountHistoryOKBodyItems0) validateEndTs(formats strfmt.Registry) error {
	if swag.IsZero(o.EndTs) { // not required
		return nil
	}

	if err := validate.FormatOf("endTs", "body", "date-time", o.EndTs.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkWirelessClientCountHistoryOKBodyItems0) validateStartTs(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTs) { // not required
		return nil
	}

	if err := validate.FormatOf("startTs", "body", "date-time", o.StartTs.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network wireless client count history o k body items0 based on context it is used
func (o *GetNetworkWirelessClientCountHistoryOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessClientCountHistoryOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessClientCountHistoryOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessClientCountHistoryOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
