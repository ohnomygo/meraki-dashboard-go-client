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

// NewGetDeviceCameraQualityAndRetentionParams creates a new GetDeviceCameraQualityAndRetentionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceCameraQualityAndRetentionParams() *GetDeviceCameraQualityAndRetentionParams {
	return &GetDeviceCameraQualityAndRetentionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceCameraQualityAndRetentionParamsWithTimeout creates a new GetDeviceCameraQualityAndRetentionParams object
// with the ability to set a timeout on a request.
func NewGetDeviceCameraQualityAndRetentionParamsWithTimeout(timeout time.Duration) *GetDeviceCameraQualityAndRetentionParams {
	return &GetDeviceCameraQualityAndRetentionParams{
		timeout: timeout,
	}
}

// NewGetDeviceCameraQualityAndRetentionParamsWithContext creates a new GetDeviceCameraQualityAndRetentionParams object
// with the ability to set a context for a request.
func NewGetDeviceCameraQualityAndRetentionParamsWithContext(ctx context.Context) *GetDeviceCameraQualityAndRetentionParams {
	return &GetDeviceCameraQualityAndRetentionParams{
		Context: ctx,
	}
}

// NewGetDeviceCameraQualityAndRetentionParamsWithHTTPClient creates a new GetDeviceCameraQualityAndRetentionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceCameraQualityAndRetentionParamsWithHTTPClient(client *http.Client) *GetDeviceCameraQualityAndRetentionParams {
	return &GetDeviceCameraQualityAndRetentionParams{
		HTTPClient: client,
	}
}

/*
GetDeviceCameraQualityAndRetentionParams contains all the parameters to send to the API endpoint

	for the get device camera quality and retention operation.

	Typically these are written to a http.Request.
*/
type GetDeviceCameraQualityAndRetentionParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device camera quality and retention params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraQualityAndRetentionParams) WithDefaults() *GetDeviceCameraQualityAndRetentionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device camera quality and retention params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraQualityAndRetentionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device camera quality and retention params
func (o *GetDeviceCameraQualityAndRetentionParams) WithTimeout(timeout time.Duration) *GetDeviceCameraQualityAndRetentionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device camera quality and retention params
func (o *GetDeviceCameraQualityAndRetentionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device camera quality and retention params
func (o *GetDeviceCameraQualityAndRetentionParams) WithContext(ctx context.Context) *GetDeviceCameraQualityAndRetentionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device camera quality and retention params
func (o *GetDeviceCameraQualityAndRetentionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device camera quality and retention params
func (o *GetDeviceCameraQualityAndRetentionParams) WithHTTPClient(client *http.Client) *GetDeviceCameraQualityAndRetentionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device camera quality and retention params
func (o *GetDeviceCameraQualityAndRetentionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device camera quality and retention params
func (o *GetDeviceCameraQualityAndRetentionParams) WithSerial(serial string) *GetDeviceCameraQualityAndRetentionParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device camera quality and retention params
func (o *GetDeviceCameraQualityAndRetentionParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceCameraQualityAndRetentionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
