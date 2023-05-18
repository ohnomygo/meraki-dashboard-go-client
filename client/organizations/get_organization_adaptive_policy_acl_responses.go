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

// GetOrganizationAdaptivePolicyACLReader is a Reader for the GetOrganizationAdaptivePolicyACL structure.
type GetOrganizationAdaptivePolicyACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationAdaptivePolicyACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationAdaptivePolicyACLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationAdaptivePolicyACLOK creates a GetOrganizationAdaptivePolicyACLOK with default headers values
func NewGetOrganizationAdaptivePolicyACLOK() *GetOrganizationAdaptivePolicyACLOK {
	return &GetOrganizationAdaptivePolicyACLOK{}
}

/*
GetOrganizationAdaptivePolicyACLOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationAdaptivePolicyACLOK struct {
	Payload *GetOrganizationAdaptivePolicyACLOKBody
}

// IsSuccess returns true when this get organization adaptive policy Acl o k response has a 2xx status code
func (o *GetOrganizationAdaptivePolicyACLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization adaptive policy Acl o k response has a 3xx status code
func (o *GetOrganizationAdaptivePolicyACLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization adaptive policy Acl o k response has a 4xx status code
func (o *GetOrganizationAdaptivePolicyACLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization adaptive policy Acl o k response has a 5xx status code
func (o *GetOrganizationAdaptivePolicyACLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization adaptive policy Acl o k response a status code equal to that given
func (o *GetOrganizationAdaptivePolicyACLOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization adaptive policy Acl o k response
func (o *GetOrganizationAdaptivePolicyACLOK) Code() int {
	return 200
}

func (o *GetOrganizationAdaptivePolicyACLOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/acls/{aclId}][%d] getOrganizationAdaptivePolicyAclOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdaptivePolicyACLOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/acls/{aclId}][%d] getOrganizationAdaptivePolicyAclOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdaptivePolicyACLOK) GetPayload() *GetOrganizationAdaptivePolicyACLOKBody {
	return o.Payload
}

func (o *GetOrganizationAdaptivePolicyACLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationAdaptivePolicyACLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationAdaptivePolicyACLOKBody get organization adaptive policy ACL o k body
swagger:model GetOrganizationAdaptivePolicyACLOKBody
*/
type GetOrganizationAdaptivePolicyACLOKBody struct {

	// ID of the adaptive policy ACL
	ACLID string `json:"aclId,omitempty"`

	// When the adaptive policy ACL was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Description of the adaptive policy ACL
	Description string `json:"description,omitempty"`

	// IP version of adpative policy ACL
	IPVersion string `json:"ipVersion,omitempty"`

	// Name of the adaptive policy ACL
	Name string `json:"name,omitempty"`

	// An ordered array of the adaptive policy ACL rules
	Rules []*GetOrganizationAdaptivePolicyACLOKBodyRulesItems0 `json:"rules"`

	// When the adaptive policy ACL was last updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this get organization adaptive policy ACL o k body
func (o *GetOrganizationAdaptivePolicyACLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationAdaptivePolicyACLOKBody) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("getOrganizationAdaptivePolicyAclOK"+"."+"createdAt", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationAdaptivePolicyACLOKBody) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(o.Rules) { // not required
		return nil
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationAdaptivePolicyAclOK" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOrganizationAdaptivePolicyAclOK" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationAdaptivePolicyACLOKBody) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("getOrganizationAdaptivePolicyAclOK"+"."+"updatedAt", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get organization adaptive policy ACL o k body based on the context it is used
func (o *GetOrganizationAdaptivePolicyACLOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationAdaptivePolicyACLOKBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {
			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationAdaptivePolicyAclOK" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOrganizationAdaptivePolicyAclOK" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyACLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyACLOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationAdaptivePolicyACLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationAdaptivePolicyACLOKBodyRulesItems0 get organization adaptive policy ACL o k body rules items0
swagger:model GetOrganizationAdaptivePolicyACLOKBodyRulesItems0
*/
type GetOrganizationAdaptivePolicyACLOKBodyRulesItems0 struct {

	// Destination port
	DstPort string `json:"dstPort,omitempty"`

	// 'allow' or 'deny' traffic specified by this rule
	Policy string `json:"policy,omitempty"`

	// The type of protocol
	Protocol string `json:"protocol,omitempty"`

	// Source port
	SrcPort string `json:"srcPort,omitempty"`
}

// Validate validates this get organization adaptive policy ACL o k body rules items0
func (o *GetOrganizationAdaptivePolicyACLOKBodyRulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization adaptive policy ACL o k body rules items0 based on context it is used
func (o *GetOrganizationAdaptivePolicyACLOKBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyACLOKBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyACLOKBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationAdaptivePolicyACLOKBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
