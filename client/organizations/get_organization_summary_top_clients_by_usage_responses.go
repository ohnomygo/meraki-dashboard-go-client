// Code generated by go-swagger; DO NOT EDIT.

package organizations

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
)

// GetOrganizationSummaryTopClientsByUsageReader is a Reader for the GetOrganizationSummaryTopClientsByUsage structure.
type GetOrganizationSummaryTopClientsByUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSummaryTopClientsByUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSummaryTopClientsByUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationSummaryTopClientsByUsageOK creates a GetOrganizationSummaryTopClientsByUsageOK with default headers values
func NewGetOrganizationSummaryTopClientsByUsageOK() *GetOrganizationSummaryTopClientsByUsageOK {
	return &GetOrganizationSummaryTopClientsByUsageOK{}
}

/*
GetOrganizationSummaryTopClientsByUsageOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationSummaryTopClientsByUsageOK struct {
	Payload []*GetOrganizationSummaryTopClientsByUsageOKBodyItems0
}

// IsSuccess returns true when this get organization summary top clients by usage o k response has a 2xx status code
func (o *GetOrganizationSummaryTopClientsByUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization summary top clients by usage o k response has a 3xx status code
func (o *GetOrganizationSummaryTopClientsByUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization summary top clients by usage o k response has a 4xx status code
func (o *GetOrganizationSummaryTopClientsByUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization summary top clients by usage o k response has a 5xx status code
func (o *GetOrganizationSummaryTopClientsByUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization summary top clients by usage o k response a status code equal to that given
func (o *GetOrganizationSummaryTopClientsByUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization summary top clients by usage o k response
func (o *GetOrganizationSummaryTopClientsByUsageOK) Code() int {
	return 200
}

func (o *GetOrganizationSummaryTopClientsByUsageOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/summary/top/clients/byUsage][%d] getOrganizationSummaryTopClientsByUsageOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSummaryTopClientsByUsageOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/summary/top/clients/byUsage][%d] getOrganizationSummaryTopClientsByUsageOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSummaryTopClientsByUsageOK) GetPayload() []*GetOrganizationSummaryTopClientsByUsageOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationSummaryTopClientsByUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationSummaryTopClientsByUsageOKBodyItems0 get organization summary top clients by usage o k body items0
swagger:model GetOrganizationSummaryTopClientsByUsageOKBodyItems0
*/
type GetOrganizationSummaryTopClientsByUsageOKBodyItems0 struct {

	// ID of client
	ID string `json:"id,omitempty"`

	// MAC address of client
	Mac string `json:"mac,omitempty"`

	// Name of client
	Name string `json:"name,omitempty"`

	// network
	Network *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network `json:"network,omitempty"`

	// usage
	Usage *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage `json:"usage,omitempty"`
}

// Validate validates this get organization summary top clients by usage o k body items0
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(o.Network) { // not required
		return nil
	}

	if o.Network != nil {
		if err := o.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(o.Usage) { // not required
		return nil
	}

	if o.Usage != nil {
		if err := o.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization summary top clients by usage o k body items0 based on the context it is used
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if o.Network != nil {
		if err := o.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if o.Usage != nil {
		if err := o.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopClientsByUsageOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network get organization summary top clients by usage o k body items0 network
swagger:model GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network
*/
type GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network struct {

	// ID of network
	ID string `json:"id,omitempty"`

	// Name of network
	Name string `json:"name,omitempty"`
}

// Validate validates this get organization summary top clients by usage o k body items0 network
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization summary top clients by usage o k body items0 network based on context it is used
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopClientsByUsageOKBodyItems0Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage Data usage information
swagger:model GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage
*/
type GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage struct {

	// Downstream data usage by client
	Downstream float32 `json:"downstream,omitempty"`

	// Percentage of total data usage by client
	Percentage float32 `json:"percentage,omitempty"`

	// Total data usage by client
	Total float32 `json:"total,omitempty"`

	// Upstream data usage by client
	Upstream float32 `json:"upstream,omitempty"`
}

// Validate validates this get organization summary top clients by usage o k body items0 usage
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization summary top clients by usage o k body items0 usage based on context it is used
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopClientsByUsageOKBodyItems0Usage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}