// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// CreateNetworkFloorPlanReader is a Reader for the CreateNetworkFloorPlan structure.
type CreateNetworkFloorPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkFloorPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkFloorPlanCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkFloorPlanCreated creates a CreateNetworkFloorPlanCreated with default headers values
func NewCreateNetworkFloorPlanCreated() *CreateNetworkFloorPlanCreated {
	return &CreateNetworkFloorPlanCreated{}
}

/*
CreateNetworkFloorPlanCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkFloorPlanCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network floor plan created response has a 2xx status code
func (o *CreateNetworkFloorPlanCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network floor plan created response has a 3xx status code
func (o *CreateNetworkFloorPlanCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network floor plan created response has a 4xx status code
func (o *CreateNetworkFloorPlanCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network floor plan created response has a 5xx status code
func (o *CreateNetworkFloorPlanCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network floor plan created response a status code equal to that given
func (o *CreateNetworkFloorPlanCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network floor plan created response
func (o *CreateNetworkFloorPlanCreated) Code() int {
	return 201
}

func (o *CreateNetworkFloorPlanCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/floorPlans][%d] createNetworkFloorPlanCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkFloorPlanCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/floorPlans][%d] createNetworkFloorPlanCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkFloorPlanCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkFloorPlanCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkFloorPlanBody create network floor plan body
// Example: {"imageContents":"Q2lzY28gTWVyYWtp","name":"HQ Floor Plan"}
swagger:model CreateNetworkFloorPlanBody
*/
type CreateNetworkFloorPlanBody struct {

	// bottom left corner
	BottomLeftCorner *CreateNetworkFloorPlanParamsBodyBottomLeftCorner `json:"bottomLeftCorner,omitempty"`

	// bottom right corner
	BottomRightCorner *CreateNetworkFloorPlanParamsBodyBottomRightCorner `json:"bottomRightCorner,omitempty"`

	// center
	Center *CreateNetworkFloorPlanParamsBodyCenter `json:"center,omitempty"`

	// The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.
	// Required: true
	// Format: byte
	ImageContents *strfmt.Base64 `json:"imageContents"`

	// The name of your floor plan.
	// Required: true
	Name *string `json:"name"`

	// top left corner
	TopLeftCorner *CreateNetworkFloorPlanParamsBodyTopLeftCorner `json:"topLeftCorner,omitempty"`

	// top right corner
	TopRightCorner *CreateNetworkFloorPlanParamsBodyTopRightCorner `json:"topRightCorner,omitempty"`
}

// Validate validates this create network floor plan body
func (o *CreateNetworkFloorPlanBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBottomLeftCorner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBottomRightCorner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCenter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageContents(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTopLeftCorner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTopRightCorner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkFloorPlanBody) validateBottomLeftCorner(formats strfmt.Registry) error {
	if swag.IsZero(o.BottomLeftCorner) { // not required
		return nil
	}

	if o.BottomLeftCorner != nil {
		if err := o.BottomLeftCorner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "bottomLeftCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "bottomLeftCorner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) validateBottomRightCorner(formats strfmt.Registry) error {
	if swag.IsZero(o.BottomRightCorner) { // not required
		return nil
	}

	if o.BottomRightCorner != nil {
		if err := o.BottomRightCorner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "bottomRightCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "bottomRightCorner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) validateCenter(formats strfmt.Registry) error {
	if swag.IsZero(o.Center) { // not required
		return nil
	}

	if o.Center != nil {
		if err := o.Center.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "center")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "center")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) validateImageContents(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkFloorPlan"+"."+"imageContents", "body", o.ImageContents); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkFloorPlan"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) validateTopLeftCorner(formats strfmt.Registry) error {
	if swag.IsZero(o.TopLeftCorner) { // not required
		return nil
	}

	if o.TopLeftCorner != nil {
		if err := o.TopLeftCorner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "topLeftCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "topLeftCorner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) validateTopRightCorner(formats strfmt.Registry) error {
	if swag.IsZero(o.TopRightCorner) { // not required
		return nil
	}

	if o.TopRightCorner != nil {
		if err := o.TopRightCorner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "topRightCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "topRightCorner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network floor plan body based on the context it is used
func (o *CreateNetworkFloorPlanBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBottomLeftCorner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateBottomRightCorner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateCenter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTopLeftCorner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTopRightCorner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkFloorPlanBody) contextValidateBottomLeftCorner(ctx context.Context, formats strfmt.Registry) error {

	if o.BottomLeftCorner != nil {
		if err := o.BottomLeftCorner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "bottomLeftCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "bottomLeftCorner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) contextValidateBottomRightCorner(ctx context.Context, formats strfmt.Registry) error {

	if o.BottomRightCorner != nil {
		if err := o.BottomRightCorner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "bottomRightCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "bottomRightCorner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) contextValidateCenter(ctx context.Context, formats strfmt.Registry) error {

	if o.Center != nil {
		if err := o.Center.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "center")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "center")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) contextValidateTopLeftCorner(ctx context.Context, formats strfmt.Registry) error {

	if o.TopLeftCorner != nil {
		if err := o.TopLeftCorner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "topLeftCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "topLeftCorner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkFloorPlanBody) contextValidateTopRightCorner(ctx context.Context, formats strfmt.Registry) error {

	if o.TopRightCorner != nil {
		if err := o.TopRightCorner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkFloorPlan" + "." + "topRightCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkFloorPlan" + "." + "topRightCorner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFloorPlanBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFloorPlanBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFloorPlanBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkFloorPlanParamsBodyBottomLeftCorner The longitude and latitude of the bottom left corner of your floor plan.
swagger:model CreateNetworkFloorPlanParamsBodyBottomLeftCorner
*/
type CreateNetworkFloorPlanParamsBodyBottomLeftCorner struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this create network floor plan params body bottom left corner
func (o *CreateNetworkFloorPlanParamsBodyBottomLeftCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network floor plan params body bottom left corner based on context it is used
func (o *CreateNetworkFloorPlanParamsBodyBottomLeftCorner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyBottomLeftCorner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyBottomLeftCorner) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFloorPlanParamsBodyBottomLeftCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkFloorPlanParamsBodyBottomRightCorner The longitude and latitude of the bottom right corner of your floor plan.
swagger:model CreateNetworkFloorPlanParamsBodyBottomRightCorner
*/
type CreateNetworkFloorPlanParamsBodyBottomRightCorner struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this create network floor plan params body bottom right corner
func (o *CreateNetworkFloorPlanParamsBodyBottomRightCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network floor plan params body bottom right corner based on context it is used
func (o *CreateNetworkFloorPlanParamsBodyBottomRightCorner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyBottomRightCorner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyBottomRightCorner) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFloorPlanParamsBodyBottomRightCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkFloorPlanParamsBodyCenter The longitude and latitude of the center of your floor plan. The 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.
swagger:model CreateNetworkFloorPlanParamsBodyCenter
*/
type CreateNetworkFloorPlanParamsBodyCenter struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this create network floor plan params body center
func (o *CreateNetworkFloorPlanParamsBodyCenter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network floor plan params body center based on context it is used
func (o *CreateNetworkFloorPlanParamsBodyCenter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyCenter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyCenter) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFloorPlanParamsBodyCenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkFloorPlanParamsBodyTopLeftCorner The longitude and latitude of the top left corner of your floor plan.
swagger:model CreateNetworkFloorPlanParamsBodyTopLeftCorner
*/
type CreateNetworkFloorPlanParamsBodyTopLeftCorner struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this create network floor plan params body top left corner
func (o *CreateNetworkFloorPlanParamsBodyTopLeftCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network floor plan params body top left corner based on context it is used
func (o *CreateNetworkFloorPlanParamsBodyTopLeftCorner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyTopLeftCorner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyTopLeftCorner) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFloorPlanParamsBodyTopLeftCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkFloorPlanParamsBodyTopRightCorner The longitude and latitude of the top right corner of your floor plan.
swagger:model CreateNetworkFloorPlanParamsBodyTopRightCorner
*/
type CreateNetworkFloorPlanParamsBodyTopRightCorner struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this create network floor plan params body top right corner
func (o *CreateNetworkFloorPlanParamsBodyTopRightCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network floor plan params body top right corner based on context it is used
func (o *CreateNetworkFloorPlanParamsBodyTopRightCorner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyTopRightCorner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkFloorPlanParamsBodyTopRightCorner) UnmarshalBinary(b []byte) error {
	var res CreateNetworkFloorPlanParamsBodyTopRightCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
