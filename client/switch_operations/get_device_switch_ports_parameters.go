// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// NewGetDeviceSwitchPortsParams creates a new GetDeviceSwitchPortsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceSwitchPortsParams() *GetDeviceSwitchPortsParams {
	return &GetDeviceSwitchPortsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceSwitchPortsParamsWithTimeout creates a new GetDeviceSwitchPortsParams object
// with the ability to set a timeout on a request.
func NewGetDeviceSwitchPortsParamsWithTimeout(timeout time.Duration) *GetDeviceSwitchPortsParams {
	return &GetDeviceSwitchPortsParams{
		timeout: timeout,
	}
}

// NewGetDeviceSwitchPortsParamsWithContext creates a new GetDeviceSwitchPortsParams object
// with the ability to set a context for a request.
func NewGetDeviceSwitchPortsParamsWithContext(ctx context.Context) *GetDeviceSwitchPortsParams {
	return &GetDeviceSwitchPortsParams{
		Context: ctx,
	}
}

// NewGetDeviceSwitchPortsParamsWithHTTPClient creates a new GetDeviceSwitchPortsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceSwitchPortsParamsWithHTTPClient(client *http.Client) *GetDeviceSwitchPortsParams {
	return &GetDeviceSwitchPortsParams{
		HTTPClient: client,
	}
}

/*
GetDeviceSwitchPortsParams contains all the parameters to send to the API endpoint

	for the get device switch ports operation.

	Typically these are written to a http.Request.
*/
type GetDeviceSwitchPortsParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device switch ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceSwitchPortsParams) WithDefaults() *GetDeviceSwitchPortsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device switch ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceSwitchPortsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device switch ports params
func (o *GetDeviceSwitchPortsParams) WithTimeout(timeout time.Duration) *GetDeviceSwitchPortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device switch ports params
func (o *GetDeviceSwitchPortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device switch ports params
func (o *GetDeviceSwitchPortsParams) WithContext(ctx context.Context) *GetDeviceSwitchPortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device switch ports params
func (o *GetDeviceSwitchPortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device switch ports params
func (o *GetDeviceSwitchPortsParams) WithHTTPClient(client *http.Client) *GetDeviceSwitchPortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device switch ports params
func (o *GetDeviceSwitchPortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device switch ports params
func (o *GetDeviceSwitchPortsParams) WithSerial(serial string) *GetDeviceSwitchPortsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device switch ports params
func (o *GetDeviceSwitchPortsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceSwitchPortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
