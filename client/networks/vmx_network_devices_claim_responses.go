// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VmxNetworkDevicesClaimReader is a Reader for the VmxNetworkDevicesClaim structure.
type VmxNetworkDevicesClaimReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VmxNetworkDevicesClaimReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVmxNetworkDevicesClaimOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVmxNetworkDevicesClaimOK creates a VmxNetworkDevicesClaimOK with default headers values
func NewVmxNetworkDevicesClaimOK() *VmxNetworkDevicesClaimOK {
	return &VmxNetworkDevicesClaimOK{}
}

/*
VmxNetworkDevicesClaimOK describes a response with status code 200, with default header values.

Successful operation
*/
type VmxNetworkDevicesClaimOK struct {
	Payload interface{}
}

// IsSuccess returns true when this vmx network devices claim o k response has a 2xx status code
func (o *VmxNetworkDevicesClaimOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vmx network devices claim o k response has a 3xx status code
func (o *VmxNetworkDevicesClaimOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vmx network devices claim o k response has a 4xx status code
func (o *VmxNetworkDevicesClaimOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vmx network devices claim o k response has a 5xx status code
func (o *VmxNetworkDevicesClaimOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vmx network devices claim o k response a status code equal to that given
func (o *VmxNetworkDevicesClaimOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vmx network devices claim o k response
func (o *VmxNetworkDevicesClaimOK) Code() int {
	return 200
}

func (o *VmxNetworkDevicesClaimOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/devices/claim/vmx][%d] vmxNetworkDevicesClaimOK  %+v", 200, o.Payload)
}

func (o *VmxNetworkDevicesClaimOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/devices/claim/vmx][%d] vmxNetworkDevicesClaimOK  %+v", 200, o.Payload)
}

func (o *VmxNetworkDevicesClaimOK) GetPayload() interface{} {
	return o.Payload
}

func (o *VmxNetworkDevicesClaimOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VmxNetworkDevicesClaimBody vmx network devices claim body
// Example: {"size":"small"}
swagger:model VmxNetworkDevicesClaimBody
*/
type VmxNetworkDevicesClaimBody struct {

	// The size of the vMX you claim. It can be one of: small, medium, large, 100
	// Required: true
	// Enum: [100 large medium small]
	Size *string `json:"size"`
}

// Validate validates this vmx network devices claim body
func (o *VmxNetworkDevicesClaimBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vmxNetworkDevicesClaimBodyTypeSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["100","large","medium","small"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmxNetworkDevicesClaimBodyTypeSizePropEnum = append(vmxNetworkDevicesClaimBodyTypeSizePropEnum, v)
	}
}

const (

	// VmxNetworkDevicesClaimBodySizeNr100 captures enum value "100"
	VmxNetworkDevicesClaimBodySizeNr100 string = "100"

	// VmxNetworkDevicesClaimBodySizeLarge captures enum value "large"
	VmxNetworkDevicesClaimBodySizeLarge string = "large"

	// VmxNetworkDevicesClaimBodySizeMedium captures enum value "medium"
	VmxNetworkDevicesClaimBodySizeMedium string = "medium"

	// VmxNetworkDevicesClaimBodySizeSmall captures enum value "small"
	VmxNetworkDevicesClaimBodySizeSmall string = "small"
)

// prop value enum
func (o *VmxNetworkDevicesClaimBody) validateSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vmxNetworkDevicesClaimBodyTypeSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *VmxNetworkDevicesClaimBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("vmxNetworkDevicesClaim"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	// value enum
	if err := o.validateSizeEnum("vmxNetworkDevicesClaim"+"."+"size", "body", *o.Size); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vmx network devices claim body based on context it is used
func (o *VmxNetworkDevicesClaimBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VmxNetworkDevicesClaimBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VmxNetworkDevicesClaimBody) UnmarshalBinary(b []byte) error {
	var res VmxNetworkDevicesClaimBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}