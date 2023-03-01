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

// NewGetDeviceManagementInterfaceParams creates a new GetDeviceManagementInterfaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceManagementInterfaceParams() *GetDeviceManagementInterfaceParams {
	return &GetDeviceManagementInterfaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceManagementInterfaceParamsWithTimeout creates a new GetDeviceManagementInterfaceParams object
// with the ability to set a timeout on a request.
func NewGetDeviceManagementInterfaceParamsWithTimeout(timeout time.Duration) *GetDeviceManagementInterfaceParams {
	return &GetDeviceManagementInterfaceParams{
		timeout: timeout,
	}
}

// NewGetDeviceManagementInterfaceParamsWithContext creates a new GetDeviceManagementInterfaceParams object
// with the ability to set a context for a request.
func NewGetDeviceManagementInterfaceParamsWithContext(ctx context.Context) *GetDeviceManagementInterfaceParams {
	return &GetDeviceManagementInterfaceParams{
		Context: ctx,
	}
}

// NewGetDeviceManagementInterfaceParamsWithHTTPClient creates a new GetDeviceManagementInterfaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceManagementInterfaceParamsWithHTTPClient(client *http.Client) *GetDeviceManagementInterfaceParams {
	return &GetDeviceManagementInterfaceParams{
		HTTPClient: client,
	}
}

/*
GetDeviceManagementInterfaceParams contains all the parameters to send to the API endpoint

	for the get device management interface operation.

	Typically these are written to a http.Request.
*/
type GetDeviceManagementInterfaceParams struct {

	// Serial.
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device management interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceManagementInterfaceParams) WithDefaults() *GetDeviceManagementInterfaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device management interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceManagementInterfaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device management interface params
func (o *GetDeviceManagementInterfaceParams) WithTimeout(timeout time.Duration) *GetDeviceManagementInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device management interface params
func (o *GetDeviceManagementInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device management interface params
func (o *GetDeviceManagementInterfaceParams) WithContext(ctx context.Context) *GetDeviceManagementInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device management interface params
func (o *GetDeviceManagementInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device management interface params
func (o *GetDeviceManagementInterfaceParams) WithHTTPClient(client *http.Client) *GetDeviceManagementInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device management interface params
func (o *GetDeviceManagementInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device management interface params
func (o *GetDeviceManagementInterfaceParams) WithSerial(serial string) *GetDeviceManagementInterfaceParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device management interface params
func (o *GetDeviceManagementInterfaceParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceManagementInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
