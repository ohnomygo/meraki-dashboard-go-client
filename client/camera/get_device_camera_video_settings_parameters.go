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

// NewGetDeviceCameraVideoSettingsParams creates a new GetDeviceCameraVideoSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceCameraVideoSettingsParams() *GetDeviceCameraVideoSettingsParams {
	return &GetDeviceCameraVideoSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceCameraVideoSettingsParamsWithTimeout creates a new GetDeviceCameraVideoSettingsParams object
// with the ability to set a timeout on a request.
func NewGetDeviceCameraVideoSettingsParamsWithTimeout(timeout time.Duration) *GetDeviceCameraVideoSettingsParams {
	return &GetDeviceCameraVideoSettingsParams{
		timeout: timeout,
	}
}

// NewGetDeviceCameraVideoSettingsParamsWithContext creates a new GetDeviceCameraVideoSettingsParams object
// with the ability to set a context for a request.
func NewGetDeviceCameraVideoSettingsParamsWithContext(ctx context.Context) *GetDeviceCameraVideoSettingsParams {
	return &GetDeviceCameraVideoSettingsParams{
		Context: ctx,
	}
}

// NewGetDeviceCameraVideoSettingsParamsWithHTTPClient creates a new GetDeviceCameraVideoSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceCameraVideoSettingsParamsWithHTTPClient(client *http.Client) *GetDeviceCameraVideoSettingsParams {
	return &GetDeviceCameraVideoSettingsParams{
		HTTPClient: client,
	}
}

/*
GetDeviceCameraVideoSettingsParams contains all the parameters to send to the API endpoint

	for the get device camera video settings operation.

	Typically these are written to a http.Request.
*/
type GetDeviceCameraVideoSettingsParams struct {

	// Serial.
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device camera video settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraVideoSettingsParams) WithDefaults() *GetDeviceCameraVideoSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device camera video settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraVideoSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device camera video settings params
func (o *GetDeviceCameraVideoSettingsParams) WithTimeout(timeout time.Duration) *GetDeviceCameraVideoSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device camera video settings params
func (o *GetDeviceCameraVideoSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device camera video settings params
func (o *GetDeviceCameraVideoSettingsParams) WithContext(ctx context.Context) *GetDeviceCameraVideoSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device camera video settings params
func (o *GetDeviceCameraVideoSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device camera video settings params
func (o *GetDeviceCameraVideoSettingsParams) WithHTTPClient(client *http.Client) *GetDeviceCameraVideoSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device camera video settings params
func (o *GetDeviceCameraVideoSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device camera video settings params
func (o *GetDeviceCameraVideoSettingsParams) WithSerial(serial string) *GetDeviceCameraVideoSettingsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device camera video settings params
func (o *GetDeviceCameraVideoSettingsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceCameraVideoSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
