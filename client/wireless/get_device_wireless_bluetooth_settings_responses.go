// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDeviceWirelessBluetoothSettingsReader is a Reader for the GetDeviceWirelessBluetoothSettings structure.
type GetDeviceWirelessBluetoothSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceWirelessBluetoothSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceWirelessBluetoothSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceWirelessBluetoothSettingsOK creates a GetDeviceWirelessBluetoothSettingsOK with default headers values
func NewGetDeviceWirelessBluetoothSettingsOK() *GetDeviceWirelessBluetoothSettingsOK {
	return &GetDeviceWirelessBluetoothSettingsOK{}
}

/*
GetDeviceWirelessBluetoothSettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceWirelessBluetoothSettingsOK struct {
	Payload *GetDeviceWirelessBluetoothSettingsOKBody
}

// IsSuccess returns true when this get device wireless bluetooth settings o k response has a 2xx status code
func (o *GetDeviceWirelessBluetoothSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device wireless bluetooth settings o k response has a 3xx status code
func (o *GetDeviceWirelessBluetoothSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device wireless bluetooth settings o k response has a 4xx status code
func (o *GetDeviceWirelessBluetoothSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device wireless bluetooth settings o k response has a 5xx status code
func (o *GetDeviceWirelessBluetoothSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device wireless bluetooth settings o k response a status code equal to that given
func (o *GetDeviceWirelessBluetoothSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device wireless bluetooth settings o k response
func (o *GetDeviceWirelessBluetoothSettingsOK) Code() int {
	return 200
}

func (o *GetDeviceWirelessBluetoothSettingsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/wireless/bluetooth/settings][%d] getDeviceWirelessBluetoothSettingsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceWirelessBluetoothSettingsOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/wireless/bluetooth/settings][%d] getDeviceWirelessBluetoothSettingsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceWirelessBluetoothSettingsOK) GetPayload() *GetDeviceWirelessBluetoothSettingsOKBody {
	return o.Payload
}

func (o *GetDeviceWirelessBluetoothSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDeviceWirelessBluetoothSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetDeviceWirelessBluetoothSettingsOKBody get device wireless bluetooth settings o k body
swagger:model GetDeviceWirelessBluetoothSettingsOKBody
*/
type GetDeviceWirelessBluetoothSettingsOKBody struct {

	// Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Major int64 `json:"major,omitempty"`

	// Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor int64 `json:"minor,omitempty"`

	// Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this get device wireless bluetooth settings o k body
func (o *GetDeviceWirelessBluetoothSettingsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get device wireless bluetooth settings o k body based on context it is used
func (o *GetDeviceWirelessBluetoothSettingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceWirelessBluetoothSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceWirelessBluetoothSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res GetDeviceWirelessBluetoothSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}