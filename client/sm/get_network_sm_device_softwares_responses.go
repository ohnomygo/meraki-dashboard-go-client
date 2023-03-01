// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// GetNetworkSmDeviceSoftwaresReader is a Reader for the GetNetworkSmDeviceSoftwares structure.
type GetNetworkSmDeviceSoftwaresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceSoftwaresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceSoftwaresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmDeviceSoftwaresOK creates a GetNetworkSmDeviceSoftwaresOK with default headers values
func NewGetNetworkSmDeviceSoftwaresOK() *GetNetworkSmDeviceSoftwaresOK {
	return &GetNetworkSmDeviceSoftwaresOK{}
}

/*
GetNetworkSmDeviceSoftwaresOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmDeviceSoftwaresOK struct {
	Payload []*GetNetworkSmDeviceSoftwaresOKBodyItems0
}

// IsSuccess returns true when this get network sm device softwares o k response has a 2xx status code
func (o *GetNetworkSmDeviceSoftwaresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm device softwares o k response has a 3xx status code
func (o *GetNetworkSmDeviceSoftwaresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm device softwares o k response has a 4xx status code
func (o *GetNetworkSmDeviceSoftwaresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm device softwares o k response has a 5xx status code
func (o *GetNetworkSmDeviceSoftwaresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm device softwares o k response a status code equal to that given
func (o *GetNetworkSmDeviceSoftwaresOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sm device softwares o k response
func (o *GetNetworkSmDeviceSoftwaresOK) Code() int {
	return 200
}

func (o *GetNetworkSmDeviceSoftwaresOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/softwares][%d] getNetworkSmDeviceSoftwaresOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceSoftwaresOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/softwares][%d] getNetworkSmDeviceSoftwaresOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceSoftwaresOK) GetPayload() []*GetNetworkSmDeviceSoftwaresOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSmDeviceSoftwaresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSmDeviceSoftwaresOKBodyItems0 get network sm device softwares o k body items0
swagger:model GetNetworkSmDeviceSoftwaresOKBodyItems0
*/
type GetNetworkSmDeviceSoftwaresOKBodyItems0 struct {

	// The Meraki managed application Id for this record on a particular device.
	AppID string `json:"appId,omitempty"`

	// The size of the software bundle.
	BundleSize int64 `json:"bundleSize,omitempty"`

	// When the Meraki record for the software was created.
	CreatedAt string `json:"createdAt,omitempty"`

	// The Meraki managed device Id.
	DeviceID string `json:"deviceId,omitempty"`

	// The size of the data stored in the application.
	DynamicSize int64 `json:"dynamicSize,omitempty"`

	// The Meraki software Id.
	ID string `json:"id,omitempty"`

	// Software bundle identifier.
	Identifier string `json:"identifier,omitempty"`

	// When the Software was installed on the device.
	InstalledAt string `json:"installedAt,omitempty"`

	// A boolean indicating whether or not an iOS redemption code was used.
	IosRedemptionCode bool `json:"iosRedemptionCode,omitempty"`

	// A boolean indicating whether or not the software is managed by Meraki.
	IsManaged bool `json:"isManaged,omitempty"`

	// The itunes numerical identifier.
	ItunesID string `json:"itunesId,omitempty"`

	// The license key associated with this software installation.
	LicenseKey string `json:"licenseKey,omitempty"`

	// The name of the software.
	Name string `json:"name,omitempty"`

	// The path on the device where the software record is located.
	Path string `json:"path,omitempty"`

	// The redemption code used for this software.
	RedemptionCode int64 `json:"redemptionCode,omitempty"`

	// Short version notation for the software.
	ShortVersion string `json:"shortVersion,omitempty"`

	// The management status of the software.
	Status string `json:"status,omitempty"`

	// A boolean indicating this software record should be installed on the associated device.
	ToInstall bool `json:"toInstall,omitempty"`

	// A boolean indicating this software record should be uninstalled on the associated device.
	ToUninstall bool `json:"toUninstall,omitempty"`

	// When the record was uninstalled from the device.
	UninstalledAt string `json:"uninstalledAt,omitempty"`

	// When the record was last updated by Meraki.
	UpdatedAt string `json:"updatedAt,omitempty"`

	// The vendor of the software.
	Vendor string `json:"vendor,omitempty"`

	// Full version notation for the software.
	Version string `json:"version,omitempty"`
}

// Validate validates this get network sm device softwares o k body items0
func (o *GetNetworkSmDeviceSoftwaresOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network sm device softwares o k body items0 based on context it is used
func (o *GetNetworkSmDeviceSoftwaresOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSmDeviceSoftwaresOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSmDeviceSoftwaresOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSmDeviceSoftwaresOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}