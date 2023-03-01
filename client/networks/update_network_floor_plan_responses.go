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
)

// UpdateNetworkFloorPlanReader is a Reader for the UpdateNetworkFloorPlan structure.
type UpdateNetworkFloorPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkFloorPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkFloorPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkFloorPlanOK creates a UpdateNetworkFloorPlanOK with default headers values
func NewUpdateNetworkFloorPlanOK() *UpdateNetworkFloorPlanOK {
	return &UpdateNetworkFloorPlanOK{}
}

/*
UpdateNetworkFloorPlanOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkFloorPlanOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network floor plan o k response has a 2xx status code
func (o *UpdateNetworkFloorPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network floor plan o k response has a 3xx status code
func (o *UpdateNetworkFloorPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network floor plan o k response has a 4xx status code
func (o *UpdateNetworkFloorPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network floor plan o k response has a 5xx status code
func (o *UpdateNetworkFloorPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network floor plan o k response a status code equal to that given
func (o *UpdateNetworkFloorPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network floor plan o k response
func (o *UpdateNetworkFloorPlanOK) Code() int {
	return 200
}

func (o *UpdateNetworkFloorPlanOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/floorPlans/{floorPlanId}][%d] updateNetworkFloorPlanOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkFloorPlanOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/floorPlans/{floorPlanId}][%d] updateNetworkFloorPlanOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkFloorPlanOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkFloorPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkFloorPlanBody update network floor plan body
// Example: {}
swagger:model UpdateNetworkFloorPlanBody
*/
type UpdateNetworkFloorPlanBody struct {

	// bottom left corner
	BottomLeftCorner *UpdateNetworkFloorPlanParamsBodyBottomLeftCorner `json:"bottomLeftCorner,omitempty"`

	// bottom right corner
	BottomRightCorner *UpdateNetworkFloorPlanParamsBodyBottomRightCorner `json:"bottomRightCorner,omitempty"`

	// center
	Center *UpdateNetworkFloorPlanParamsBodyCenter `json:"center,omitempty"`

	// The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.
	// Format: byte
	ImageContents strfmt.Base64 `json:"imageContents,omitempty"`

	// The name of your floor plan.
	Name string `json:"name,omitempty"`

	// top left corner
	TopLeftCorner *UpdateNetworkFloorPlanParamsBodyTopLeftCorner `json:"topLeftCorner,omitempty"`

	// top right corner
	TopRightCorner *UpdateNetworkFloorPlanParamsBodyTopRightCorner `json:"topRightCorner,omitempty"`
}

// Validate validates this update network floor plan body
func (o *UpdateNetworkFloorPlanBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateNetworkFloorPlanBody) validateBottomLeftCorner(formats strfmt.Registry) error {
	if swag.IsZero(o.BottomLeftCorner) { // not required
		return nil
	}

	if o.BottomLeftCorner != nil {
		if err := o.BottomLeftCorner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "bottomLeftCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "bottomLeftCorner")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkFloorPlanBody) validateBottomRightCorner(formats strfmt.Registry) error {
	if swag.IsZero(o.BottomRightCorner) { // not required
		return nil
	}

	if o.BottomRightCorner != nil {
		if err := o.BottomRightCorner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "bottomRightCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "bottomRightCorner")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkFloorPlanBody) validateCenter(formats strfmt.Registry) error {
	if swag.IsZero(o.Center) { // not required
		return nil
	}

	if o.Center != nil {
		if err := o.Center.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "center")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "center")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkFloorPlanBody) validateTopLeftCorner(formats strfmt.Registry) error {
	if swag.IsZero(o.TopLeftCorner) { // not required
		return nil
	}

	if o.TopLeftCorner != nil {
		if err := o.TopLeftCorner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "topLeftCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "topLeftCorner")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkFloorPlanBody) validateTopRightCorner(formats strfmt.Registry) error {
	if swag.IsZero(o.TopRightCorner) { // not required
		return nil
	}

	if o.TopRightCorner != nil {
		if err := o.TopRightCorner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "topRightCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "topRightCorner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network floor plan body based on the context it is used
func (o *UpdateNetworkFloorPlanBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *UpdateNetworkFloorPlanBody) contextValidateBottomLeftCorner(ctx context.Context, formats strfmt.Registry) error {

	if o.BottomLeftCorner != nil {
		if err := o.BottomLeftCorner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "bottomLeftCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "bottomLeftCorner")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkFloorPlanBody) contextValidateBottomRightCorner(ctx context.Context, formats strfmt.Registry) error {

	if o.BottomRightCorner != nil {
		if err := o.BottomRightCorner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "bottomRightCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "bottomRightCorner")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkFloorPlanBody) contextValidateCenter(ctx context.Context, formats strfmt.Registry) error {

	if o.Center != nil {
		if err := o.Center.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "center")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "center")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkFloorPlanBody) contextValidateTopLeftCorner(ctx context.Context, formats strfmt.Registry) error {

	if o.TopLeftCorner != nil {
		if err := o.TopLeftCorner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "topLeftCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "topLeftCorner")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkFloorPlanBody) contextValidateTopRightCorner(ctx context.Context, formats strfmt.Registry) error {

	if o.TopRightCorner != nil {
		if err := o.TopRightCorner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkFloorPlan" + "." + "topRightCorner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkFloorPlan" + "." + "topRightCorner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFloorPlanBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFloorPlanParamsBodyBottomLeftCorner The longitude and latitude of the bottom left corner of your floor plan.
swagger:model UpdateNetworkFloorPlanParamsBodyBottomLeftCorner
*/
type UpdateNetworkFloorPlanParamsBodyBottomLeftCorner struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this update network floor plan params body bottom left corner
func (o *UpdateNetworkFloorPlanParamsBodyBottomLeftCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network floor plan params body bottom left corner based on context it is used
func (o *UpdateNetworkFloorPlanParamsBodyBottomLeftCorner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyBottomLeftCorner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyBottomLeftCorner) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFloorPlanParamsBodyBottomLeftCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFloorPlanParamsBodyBottomRightCorner The longitude and latitude of the bottom right corner of your floor plan.
swagger:model UpdateNetworkFloorPlanParamsBodyBottomRightCorner
*/
type UpdateNetworkFloorPlanParamsBodyBottomRightCorner struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this update network floor plan params body bottom right corner
func (o *UpdateNetworkFloorPlanParamsBodyBottomRightCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network floor plan params body bottom right corner based on context it is used
func (o *UpdateNetworkFloorPlanParamsBodyBottomRightCorner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyBottomRightCorner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyBottomRightCorner) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFloorPlanParamsBodyBottomRightCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFloorPlanParamsBodyCenter The longitude and latitude of the center of your floor plan. If you want to change the geolocation data of your floor plan, either the 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.
swagger:model UpdateNetworkFloorPlanParamsBodyCenter
*/
type UpdateNetworkFloorPlanParamsBodyCenter struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this update network floor plan params body center
func (o *UpdateNetworkFloorPlanParamsBodyCenter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network floor plan params body center based on context it is used
func (o *UpdateNetworkFloorPlanParamsBodyCenter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyCenter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyCenter) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFloorPlanParamsBodyCenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFloorPlanParamsBodyTopLeftCorner The longitude and latitude of the top left corner of your floor plan.
swagger:model UpdateNetworkFloorPlanParamsBodyTopLeftCorner
*/
type UpdateNetworkFloorPlanParamsBodyTopLeftCorner struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this update network floor plan params body top left corner
func (o *UpdateNetworkFloorPlanParamsBodyTopLeftCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network floor plan params body top left corner based on context it is used
func (o *UpdateNetworkFloorPlanParamsBodyTopLeftCorner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyTopLeftCorner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyTopLeftCorner) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFloorPlanParamsBodyTopLeftCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkFloorPlanParamsBodyTopRightCorner The longitude and latitude of the top right corner of your floor plan.
swagger:model UpdateNetworkFloorPlanParamsBodyTopRightCorner
*/
type UpdateNetworkFloorPlanParamsBodyTopRightCorner struct {

	// Latitude
	Lat float32 `json:"lat,omitempty"`

	// Longitude
	Lng float32 `json:"lng,omitempty"`
}

// Validate validates this update network floor plan params body top right corner
func (o *UpdateNetworkFloorPlanParamsBodyTopRightCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network floor plan params body top right corner based on context it is used
func (o *UpdateNetworkFloorPlanParamsBodyTopRightCorner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyTopRightCorner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkFloorPlanParamsBodyTopRightCorner) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkFloorPlanParamsBodyTopRightCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
