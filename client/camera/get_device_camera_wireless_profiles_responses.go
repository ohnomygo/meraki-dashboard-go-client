// Code generated by go-swagger; DO NOT EDIT.

package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceCameraWirelessProfilesReader is a Reader for the GetDeviceCameraWirelessProfiles structure.
type GetDeviceCameraWirelessProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCameraWirelessProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCameraWirelessProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceCameraWirelessProfilesOK creates a GetDeviceCameraWirelessProfilesOK with default headers values
func NewGetDeviceCameraWirelessProfilesOK() *GetDeviceCameraWirelessProfilesOK {
	return &GetDeviceCameraWirelessProfilesOK{}
}

/*
GetDeviceCameraWirelessProfilesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceCameraWirelessProfilesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device camera wireless profiles o k response has a 2xx status code
func (o *GetDeviceCameraWirelessProfilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device camera wireless profiles o k response has a 3xx status code
func (o *GetDeviceCameraWirelessProfilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device camera wireless profiles o k response has a 4xx status code
func (o *GetDeviceCameraWirelessProfilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device camera wireless profiles o k response has a 5xx status code
func (o *GetDeviceCameraWirelessProfilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device camera wireless profiles o k response a status code equal to that given
func (o *GetDeviceCameraWirelessProfilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device camera wireless profiles o k response
func (o *GetDeviceCameraWirelessProfilesOK) Code() int {
	return 200
}

func (o *GetDeviceCameraWirelessProfilesOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/wirelessProfiles][%d] getDeviceCameraWirelessProfilesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCameraWirelessProfilesOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/wirelessProfiles][%d] getDeviceCameraWirelessProfilesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCameraWirelessProfilesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceCameraWirelessProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
