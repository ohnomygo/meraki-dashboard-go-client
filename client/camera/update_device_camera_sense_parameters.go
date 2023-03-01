// Code generated by go-swagger; DO NOT EDIT.

package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewUpdateDeviceCameraSenseParams creates a new UpdateDeviceCameraSenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceCameraSenseParams() *UpdateDeviceCameraSenseParams {
	return &UpdateDeviceCameraSenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceCameraSenseParamsWithTimeout creates a new UpdateDeviceCameraSenseParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceCameraSenseParamsWithTimeout(timeout time.Duration) *UpdateDeviceCameraSenseParams {
	return &UpdateDeviceCameraSenseParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceCameraSenseParamsWithContext creates a new UpdateDeviceCameraSenseParams object
// with the ability to set a context for a request.
func NewUpdateDeviceCameraSenseParamsWithContext(ctx context.Context) *UpdateDeviceCameraSenseParams {
	return &UpdateDeviceCameraSenseParams{
		Context: ctx,
	}
}

// NewUpdateDeviceCameraSenseParamsWithHTTPClient creates a new UpdateDeviceCameraSenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceCameraSenseParamsWithHTTPClient(client *http.Client) *UpdateDeviceCameraSenseParams {
	return &UpdateDeviceCameraSenseParams{
		HTTPClient: client,
	}
}

/*
UpdateDeviceCameraSenseParams contains all the parameters to send to the API endpoint

	for the update device camera sense operation.

	Typically these are written to a http.Request.
*/
type UpdateDeviceCameraSenseParams struct {

	// Serial.
	Serial string

	// UpdateDeviceCameraSense.
	UpdateDeviceCameraSense UpdateDeviceCameraSenseBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device camera sense params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceCameraSenseParams) WithDefaults() *UpdateDeviceCameraSenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device camera sense params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceCameraSenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) WithTimeout(timeout time.Duration) *UpdateDeviceCameraSenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) WithContext(ctx context.Context) *UpdateDeviceCameraSenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) WithHTTPClient(client *http.Client) *UpdateDeviceCameraSenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) WithSerial(serial string) *UpdateDeviceCameraSenseParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceCameraSense adds the updateDeviceCameraSense to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) WithUpdateDeviceCameraSense(updateDeviceCameraSense UpdateDeviceCameraSenseBody) *UpdateDeviceCameraSenseParams {
	o.SetUpdateDeviceCameraSense(updateDeviceCameraSense)
	return o
}

// SetUpdateDeviceCameraSense adds the updateDeviceCameraSense to the update device camera sense params
func (o *UpdateDeviceCameraSenseParams) SetUpdateDeviceCameraSense(updateDeviceCameraSense UpdateDeviceCameraSenseBody) {
	o.UpdateDeviceCameraSense = updateDeviceCameraSense
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceCameraSenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateDeviceCameraSense); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
