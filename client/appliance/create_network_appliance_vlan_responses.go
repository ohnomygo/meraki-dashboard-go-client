// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// CreateNetworkApplianceVlanReader is a Reader for the CreateNetworkApplianceVlan structure.
type CreateNetworkApplianceVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkApplianceVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkApplianceVlanCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkApplianceVlanCreated creates a CreateNetworkApplianceVlanCreated with default headers values
func NewCreateNetworkApplianceVlanCreated() *CreateNetworkApplianceVlanCreated {
	return &CreateNetworkApplianceVlanCreated{}
}

/*
CreateNetworkApplianceVlanCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkApplianceVlanCreated struct {
	Payload *CreateNetworkApplianceVlanCreatedBody
}

// IsSuccess returns true when this create network appliance vlan created response has a 2xx status code
func (o *CreateNetworkApplianceVlanCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network appliance vlan created response has a 3xx status code
func (o *CreateNetworkApplianceVlanCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network appliance vlan created response has a 4xx status code
func (o *CreateNetworkApplianceVlanCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network appliance vlan created response has a 5xx status code
func (o *CreateNetworkApplianceVlanCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network appliance vlan created response a status code equal to that given
func (o *CreateNetworkApplianceVlanCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network appliance vlan created response
func (o *CreateNetworkApplianceVlanCreated) Code() int {
	return 201
}

func (o *CreateNetworkApplianceVlanCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/appliance/vlans][%d] createNetworkApplianceVlanCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkApplianceVlanCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/appliance/vlans][%d] createNetworkApplianceVlanCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkApplianceVlanCreated) GetPayload() *CreateNetworkApplianceVlanCreatedBody {
	return o.Payload
}

func (o *CreateNetworkApplianceVlanCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateNetworkApplianceVlanCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkApplianceVlanBody create network appliance vlan body
// Example: {"applianceIp":"192.168.1.2","groupPolicyId":"101","id":"1234","name":"My VLAN","subnet":"192.168.1.0/24"}
swagger:model CreateNetworkApplianceVlanBody
*/
type CreateNetworkApplianceVlanBody struct {

	// The local IP of the appliance on the VLAN
	ApplianceIP string `json:"applianceIp,omitempty"`

	// CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	Cidr string `json:"cidr,omitempty"`

	// The id of the desired group policy to apply to the VLAN
	GroupPolicyID string `json:"groupPolicyId,omitempty"`

	// The VLAN ID of the new VLAN (must be between 1 and 4094)
	// Required: true
	ID *string `json:"id"`

	// ipv6
	IPV6 *CreateNetworkApplianceVlanParamsBodyIPV6 `json:"ipv6,omitempty"`

	// mandatory dhcp
	MandatoryDhcp *CreateNetworkApplianceVlanParamsBodyMandatoryDhcp `json:"mandatoryDhcp,omitempty"`

	// Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Mask int64 `json:"mask,omitempty"`

	// The name of the new VLAN
	// Required: true
	Name *string `json:"name"`

	// The subnet of the VLAN
	Subnet string `json:"subnet,omitempty"`

	// Type of subnetting of the VLAN. Applicable only for template network.
	// Enum: [same unique]
	TemplateVlanType *string `json:"templateVlanType,omitempty"`
}

// Validate validates this create network appliance vlan body
func (o *CreateNetworkApplianceVlanBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMandatoryDhcp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTemplateVlanType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkApplianceVlan"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkApplianceVlanBody) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV6) { // not required
		return nil
	}

	if o.IPV6 != nil {
		if err := o.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkApplianceVlan" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkApplianceVlan" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkApplianceVlanBody) validateMandatoryDhcp(formats strfmt.Registry) error {
	if swag.IsZero(o.MandatoryDhcp) { // not required
		return nil
	}

	if o.MandatoryDhcp != nil {
		if err := o.MandatoryDhcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkApplianceVlan" + "." + "mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkApplianceVlan" + "." + "mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkApplianceVlanBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkApplianceVlan"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

var createNetworkApplianceVlanBodyTypeTemplateVlanTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["same","unique"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkApplianceVlanBodyTypeTemplateVlanTypePropEnum = append(createNetworkApplianceVlanBodyTypeTemplateVlanTypePropEnum, v)
	}
}

const (

	// CreateNetworkApplianceVlanBodyTemplateVlanTypeSame captures enum value "same"
	CreateNetworkApplianceVlanBodyTemplateVlanTypeSame string = "same"

	// CreateNetworkApplianceVlanBodyTemplateVlanTypeUnique captures enum value "unique"
	CreateNetworkApplianceVlanBodyTemplateVlanTypeUnique string = "unique"
)

// prop value enum
func (o *CreateNetworkApplianceVlanBody) validateTemplateVlanTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkApplianceVlanBodyTypeTemplateVlanTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkApplianceVlanBody) validateTemplateVlanType(formats strfmt.Registry) error {
	if swag.IsZero(o.TemplateVlanType) { // not required
		return nil
	}

	// value enum
	if err := o.validateTemplateVlanTypeEnum("createNetworkApplianceVlan"+"."+"templateVlanType", "body", *o.TemplateVlanType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network appliance vlan body based on the context it is used
func (o *CreateNetworkApplianceVlanBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMandatoryDhcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanBody) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV6 != nil {
		if err := o.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkApplianceVlan" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkApplianceVlan" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkApplianceVlanBody) contextValidateMandatoryDhcp(ctx context.Context, formats strfmt.Registry) error {

	if o.MandatoryDhcp != nil {
		if err := o.MandatoryDhcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkApplianceVlan" + "." + "mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkApplianceVlan" + "." + "mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanCreatedBody create network appliance vlan created body
swagger:model CreateNetworkApplianceVlanCreatedBody
*/
type CreateNetworkApplianceVlanCreatedBody struct {

	// The local IP of the appliance on the VLAN
	ApplianceIP string `json:"applianceIp,omitempty"`

	// CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	Cidr string `json:"cidr,omitempty"`

	// The id of the desired group policy to apply to the VLAN
	GroupPolicyID string `json:"groupPolicyId,omitempty"`

	// The VLAN ID of the VLAN
	ID string `json:"id,omitempty"`

	// The interface ID of the VLAN
	InterfaceID string `json:"interfaceId,omitempty"`

	// ipv6
	IPV6 *CreateNetworkApplianceVlanCreatedBodyIPV6 `json:"ipv6,omitempty"`

	// mandatory dhcp
	MandatoryDhcp *CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp `json:"mandatoryDhcp,omitempty"`

	// Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Mask int64 `json:"mask,omitempty"`

	// The name of the VLAN
	Name string `json:"name,omitempty"`

	// The subnet of the VLAN
	Subnet string `json:"subnet,omitempty"`

	// Type of subnetting of the VLAN. Applicable only for template network.
	// Enum: [same unique]
	TemplateVlanType *string `json:"templateVlanType,omitempty"`
}

// Validate validates this create network appliance vlan created body
func (o *CreateNetworkApplianceVlanCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMandatoryDhcp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTemplateVlanType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBody) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV6) { // not required
		return nil
	}

	if o.IPV6 != nil {
		if err := o.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkApplianceVlanCreated" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkApplianceVlanCreated" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBody) validateMandatoryDhcp(formats strfmt.Registry) error {
	if swag.IsZero(o.MandatoryDhcp) { // not required
		return nil
	}

	if o.MandatoryDhcp != nil {
		if err := o.MandatoryDhcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkApplianceVlanCreated" + "." + "mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkApplianceVlanCreated" + "." + "mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

var createNetworkApplianceVlanCreatedBodyTypeTemplateVlanTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["same","unique"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkApplianceVlanCreatedBodyTypeTemplateVlanTypePropEnum = append(createNetworkApplianceVlanCreatedBodyTypeTemplateVlanTypePropEnum, v)
	}
}

const (

	// CreateNetworkApplianceVlanCreatedBodyTemplateVlanTypeSame captures enum value "same"
	CreateNetworkApplianceVlanCreatedBodyTemplateVlanTypeSame string = "same"

	// CreateNetworkApplianceVlanCreatedBodyTemplateVlanTypeUnique captures enum value "unique"
	CreateNetworkApplianceVlanCreatedBodyTemplateVlanTypeUnique string = "unique"
)

// prop value enum
func (o *CreateNetworkApplianceVlanCreatedBody) validateTemplateVlanTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkApplianceVlanCreatedBodyTypeTemplateVlanTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBody) validateTemplateVlanType(formats strfmt.Registry) error {
	if swag.IsZero(o.TemplateVlanType) { // not required
		return nil
	}

	// value enum
	if err := o.validateTemplateVlanTypeEnum("createNetworkApplianceVlanCreated"+"."+"templateVlanType", "body", *o.TemplateVlanType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network appliance vlan created body based on the context it is used
func (o *CreateNetworkApplianceVlanCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMandatoryDhcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBody) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV6 != nil {
		if err := o.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkApplianceVlanCreated" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkApplianceVlanCreated" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBody) contextValidateMandatoryDhcp(ctx context.Context, formats strfmt.Registry) error {

	if o.MandatoryDhcp != nil {
		if err := o.MandatoryDhcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkApplianceVlanCreated" + "." + "mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkApplianceVlanCreated" + "." + "mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanCreatedBodyIPV6 IPv6 configuration on the VLAN
swagger:model CreateNetworkApplianceVlanCreatedBodyIPV6
*/
type CreateNetworkApplianceVlanCreatedBodyIPV6 struct {

	// Enable IPv6 on VLAN
	Enabled bool `json:"enabled,omitempty"`

	// Prefix assignments on the VLAN
	PrefixAssignments []*CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0 `json:"prefixAssignments"`
}

// Validate validates this create network appliance vlan created body IP v6
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrefixAssignments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBodyIPV6) validatePrefixAssignments(formats strfmt.Registry) error {
	if swag.IsZero(o.PrefixAssignments) { // not required
		return nil
	}

	for i := 0; i < len(o.PrefixAssignments); i++ {
		if swag.IsZero(o.PrefixAssignments[i]) { // not required
			continue
		}

		if o.PrefixAssignments[i] != nil {
			if err := o.PrefixAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkApplianceVlanCreated" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkApplianceVlanCreated" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create network appliance vlan created body IP v6 based on the context it is used
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePrefixAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBodyIPV6) contextValidatePrefixAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PrefixAssignments); i++ {

		if o.PrefixAssignments[i] != nil {
			if err := o.PrefixAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkApplianceVlanCreated" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkApplianceVlanCreated" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanCreatedBodyIPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0 create network appliance vlan created body IP v6 prefix assignments items0
swagger:model CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0
*/
type CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0 struct {

	// Auto assign a /64 prefix from the origin to the VLAN
	Autonomous bool `json:"autonomous,omitempty"`

	// origin
	Origin *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin `json:"origin,omitempty"`

	// Manual configuration of the IPv6 Appliance IP
	StaticApplianceIp6 string `json:"staticApplianceIp6,omitempty"`

	// Manual configuration of a /64 prefix on the VLAN
	StaticPrefix string `json:"staticPrefix,omitempty"`
}

// Validate validates this create network appliance vlan created body IP v6 prefix assignments items0
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(o.Origin) { // not required
		return nil
	}

	if o.Origin != nil {
		if err := o.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network appliance vlan created body IP v6 prefix assignments items0 based on the context it is used
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if o.Origin != nil {
		if err := o.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin The origin of the prefix
swagger:model CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin
*/
type CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin struct {

	// Interfaces associated with the prefix
	Interfaces []string `json:"interfaces"`

	// Type of the origin
	// Enum: [independent internet]
	Type string `json:"type,omitempty"`
}

// Validate validates this create network appliance vlan created body IP v6 prefix assignments items0 origin
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createNetworkApplianceVlanCreatedBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["independent","internet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkApplianceVlanCreatedBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum = append(createNetworkApplianceVlanCreatedBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, v)
	}
}

const (

	// CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0OriginTypeIndependent captures enum value "independent"
	CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0OriginTypeIndependent string = "independent"

	// CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0OriginTypeInternet captures enum value "internet"
	CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0OriginTypeInternet string = "internet"
)

// prop value enum
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkApplianceVlanCreatedBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("origin"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network appliance vlan created body IP v6 prefix assignments items0 origin based on context it is used
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanCreatedBodyIPV6PrefixAssignmentsItems0Origin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
swagger:model CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp
*/
type CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp struct {

	// Enable Mandatory DHCP on VLAN.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this create network appliance vlan created body mandatory dhcp
func (o *CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network appliance vlan created body mandatory dhcp based on context it is used
func (o *CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanCreatedBodyMandatoryDhcp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanParamsBodyIPV6 IPv6 configuration on the VLAN
swagger:model CreateNetworkApplianceVlanParamsBodyIPV6
*/
type CreateNetworkApplianceVlanParamsBodyIPV6 struct {

	// Enable IPv6 on VLAN.
	Enabled bool `json:"enabled,omitempty"`

	// Prefix assignments on the VLAN
	PrefixAssignments []*CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0 `json:"prefixAssignments"`
}

// Validate validates this create network appliance vlan params body IP v6
func (o *CreateNetworkApplianceVlanParamsBodyIPV6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrefixAssignments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanParamsBodyIPV6) validatePrefixAssignments(formats strfmt.Registry) error {
	if swag.IsZero(o.PrefixAssignments) { // not required
		return nil
	}

	for i := 0; i < len(o.PrefixAssignments); i++ {
		if swag.IsZero(o.PrefixAssignments[i]) { // not required
			continue
		}

		if o.PrefixAssignments[i] != nil {
			if err := o.PrefixAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkApplianceVlan" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkApplianceVlan" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create network appliance vlan params body IP v6 based on the context it is used
func (o *CreateNetworkApplianceVlanParamsBodyIPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePrefixAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanParamsBodyIPV6) contextValidatePrefixAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PrefixAssignments); i++ {

		if o.PrefixAssignments[i] != nil {
			if err := o.PrefixAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkApplianceVlan" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkApplianceVlan" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanParamsBodyIPV6) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanParamsBodyIPV6) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanParamsBodyIPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0 create network appliance vlan params body IP v6 prefix assignments items0
swagger:model CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0
*/
type CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0 struct {

	// Auto assign a /64 prefix from the origin to the VLAN
	Autonomous bool `json:"autonomous,omitempty"`

	// origin
	Origin *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin `json:"origin,omitempty"`

	// Manual configuration of the IPv6 Appliance IP
	StaticApplianceIp6 string `json:"staticApplianceIp6,omitempty"`

	// Manual configuration of a /64 prefix on the VLAN
	StaticPrefix string `json:"staticPrefix,omitempty"`
}

// Validate validates this create network appliance vlan params body IP v6 prefix assignments items0
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(o.Origin) { // not required
		return nil
	}

	if o.Origin != nil {
		if err := o.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network appliance vlan params body IP v6 prefix assignments items0 based on the context it is used
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if o.Origin != nil {
		if err := o.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin The origin of the prefix
swagger:model CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin
*/
type CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin struct {

	// Interfaces associated with the prefix
	Interfaces []string `json:"interfaces"`

	// Type of the origin
	// Required: true
	// Enum: [independent internet]
	Type *string `json:"type"`
}

// Validate validates this create network appliance vlan params body IP v6 prefix assignments items0 origin
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createNetworkApplianceVlanParamsBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["independent","internet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkApplianceVlanParamsBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum = append(createNetworkApplianceVlanParamsBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, v)
	}
}

const (

	// CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0OriginTypeIndependent captures enum value "independent"
	CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0OriginTypeIndependent string = "independent"

	// CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0OriginTypeInternet captures enum value "internet"
	CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0OriginTypeInternet string = "internet"
)

// prop value enum
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkApplianceVlanParamsBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin) validateType(formats strfmt.Registry) error {

	if err := validate.Required("origin"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("origin"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network appliance vlan params body IP v6 prefix assignments items0 origin based on context it is used
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanParamsBodyIPV6PrefixAssignmentsItems0Origin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkApplianceVlanParamsBodyMandatoryDhcp Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
swagger:model CreateNetworkApplianceVlanParamsBodyMandatoryDhcp
*/
type CreateNetworkApplianceVlanParamsBodyMandatoryDhcp struct {

	// Enable Mandatory DHCP on VLAN.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this create network appliance vlan params body mandatory dhcp
func (o *CreateNetworkApplianceVlanParamsBodyMandatoryDhcp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network appliance vlan params body mandatory dhcp based on context it is used
func (o *CreateNetworkApplianceVlanParamsBodyMandatoryDhcp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanParamsBodyMandatoryDhcp) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanParamsBodyMandatoryDhcp) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanParamsBodyMandatoryDhcp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
