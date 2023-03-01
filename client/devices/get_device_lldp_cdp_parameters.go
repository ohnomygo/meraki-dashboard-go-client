// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewGetDeviceLldpCdpParams creates a new GetDeviceLldpCdpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceLldpCdpParams() *GetDeviceLldpCdpParams {
	return &GetDeviceLldpCdpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceLldpCdpParamsWithTimeout creates a new GetDeviceLldpCdpParams object
// with the ability to set a timeout on a request.
func NewGetDeviceLldpCdpParamsWithTimeout(timeout time.Duration) *GetDeviceLldpCdpParams {
	return &GetDeviceLldpCdpParams{
		timeout: timeout,
	}
}

// NewGetDeviceLldpCdpParamsWithContext creates a new GetDeviceLldpCdpParams object
// with the ability to set a context for a request.
func NewGetDeviceLldpCdpParamsWithContext(ctx context.Context) *GetDeviceLldpCdpParams {
	return &GetDeviceLldpCdpParams{
		Context: ctx,
	}
}

// NewGetDeviceLldpCdpParamsWithHTTPClient creates a new GetDeviceLldpCdpParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceLldpCdpParamsWithHTTPClient(client *http.Client) *GetDeviceLldpCdpParams {
	return &GetDeviceLldpCdpParams{
		HTTPClient: client,
	}
}

/*
GetDeviceLldpCdpParams contains all the parameters to send to the API endpoint

	for the get device lldp cdp operation.

	Typically these are written to a http.Request.
*/
type GetDeviceLldpCdpParams struct {

	// Serial.
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device lldp cdp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceLldpCdpParams) WithDefaults() *GetDeviceLldpCdpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device lldp cdp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceLldpCdpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device lldp cdp params
func (o *GetDeviceLldpCdpParams) WithTimeout(timeout time.Duration) *GetDeviceLldpCdpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device lldp cdp params
func (o *GetDeviceLldpCdpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device lldp cdp params
func (o *GetDeviceLldpCdpParams) WithContext(ctx context.Context) *GetDeviceLldpCdpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device lldp cdp params
func (o *GetDeviceLldpCdpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device lldp cdp params
func (o *GetDeviceLldpCdpParams) WithHTTPClient(client *http.Client) *GetDeviceLldpCdpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device lldp cdp params
func (o *GetDeviceLldpCdpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device lldp cdp params
func (o *GetDeviceLldpCdpParams) WithSerial(serial string) *GetDeviceLldpCdpParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device lldp cdp params
func (o *GetDeviceLldpCdpParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceLldpCdpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
