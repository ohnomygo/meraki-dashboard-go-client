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

// NewGetDeviceSwitchWarmSpareParams creates a new GetDeviceSwitchWarmSpareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceSwitchWarmSpareParams() *GetDeviceSwitchWarmSpareParams {
	return &GetDeviceSwitchWarmSpareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceSwitchWarmSpareParamsWithTimeout creates a new GetDeviceSwitchWarmSpareParams object
// with the ability to set a timeout on a request.
func NewGetDeviceSwitchWarmSpareParamsWithTimeout(timeout time.Duration) *GetDeviceSwitchWarmSpareParams {
	return &GetDeviceSwitchWarmSpareParams{
		timeout: timeout,
	}
}

// NewGetDeviceSwitchWarmSpareParamsWithContext creates a new GetDeviceSwitchWarmSpareParams object
// with the ability to set a context for a request.
func NewGetDeviceSwitchWarmSpareParamsWithContext(ctx context.Context) *GetDeviceSwitchWarmSpareParams {
	return &GetDeviceSwitchWarmSpareParams{
		Context: ctx,
	}
}

// NewGetDeviceSwitchWarmSpareParamsWithHTTPClient creates a new GetDeviceSwitchWarmSpareParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceSwitchWarmSpareParamsWithHTTPClient(client *http.Client) *GetDeviceSwitchWarmSpareParams {
	return &GetDeviceSwitchWarmSpareParams{
		HTTPClient: client,
	}
}

/*
GetDeviceSwitchWarmSpareParams contains all the parameters to send to the API endpoint

	for the get device switch warm spare operation.

	Typically these are written to a http.Request.
*/
type GetDeviceSwitchWarmSpareParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device switch warm spare params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceSwitchWarmSpareParams) WithDefaults() *GetDeviceSwitchWarmSpareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device switch warm spare params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceSwitchWarmSpareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device switch warm spare params
func (o *GetDeviceSwitchWarmSpareParams) WithTimeout(timeout time.Duration) *GetDeviceSwitchWarmSpareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device switch warm spare params
func (o *GetDeviceSwitchWarmSpareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device switch warm spare params
func (o *GetDeviceSwitchWarmSpareParams) WithContext(ctx context.Context) *GetDeviceSwitchWarmSpareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device switch warm spare params
func (o *GetDeviceSwitchWarmSpareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device switch warm spare params
func (o *GetDeviceSwitchWarmSpareParams) WithHTTPClient(client *http.Client) *GetDeviceSwitchWarmSpareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device switch warm spare params
func (o *GetDeviceSwitchWarmSpareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device switch warm spare params
func (o *GetDeviceSwitchWarmSpareParams) WithSerial(serial string) *GetDeviceSwitchWarmSpareParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device switch warm spare params
func (o *GetDeviceSwitchWarmSpareParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceSwitchWarmSpareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
