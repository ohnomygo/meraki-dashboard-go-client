// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

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

// NewGetDeviceCellularGatewayPortForwardingRulesParams creates a new GetDeviceCellularGatewayPortForwardingRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceCellularGatewayPortForwardingRulesParams() *GetDeviceCellularGatewayPortForwardingRulesParams {
	return &GetDeviceCellularGatewayPortForwardingRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceCellularGatewayPortForwardingRulesParamsWithTimeout creates a new GetDeviceCellularGatewayPortForwardingRulesParams object
// with the ability to set a timeout on a request.
func NewGetDeviceCellularGatewayPortForwardingRulesParamsWithTimeout(timeout time.Duration) *GetDeviceCellularGatewayPortForwardingRulesParams {
	return &GetDeviceCellularGatewayPortForwardingRulesParams{
		timeout: timeout,
	}
}

// NewGetDeviceCellularGatewayPortForwardingRulesParamsWithContext creates a new GetDeviceCellularGatewayPortForwardingRulesParams object
// with the ability to set a context for a request.
func NewGetDeviceCellularGatewayPortForwardingRulesParamsWithContext(ctx context.Context) *GetDeviceCellularGatewayPortForwardingRulesParams {
	return &GetDeviceCellularGatewayPortForwardingRulesParams{
		Context: ctx,
	}
}

// NewGetDeviceCellularGatewayPortForwardingRulesParamsWithHTTPClient creates a new GetDeviceCellularGatewayPortForwardingRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceCellularGatewayPortForwardingRulesParamsWithHTTPClient(client *http.Client) *GetDeviceCellularGatewayPortForwardingRulesParams {
	return &GetDeviceCellularGatewayPortForwardingRulesParams{
		HTTPClient: client,
	}
}

/*
GetDeviceCellularGatewayPortForwardingRulesParams contains all the parameters to send to the API endpoint

	for the get device cellular gateway port forwarding rules operation.

	Typically these are written to a http.Request.
*/
type GetDeviceCellularGatewayPortForwardingRulesParams struct {

	// Serial.
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device cellular gateway port forwarding rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) WithDefaults() *GetDeviceCellularGatewayPortForwardingRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device cellular gateway port forwarding rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device cellular gateway port forwarding rules params
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) WithTimeout(timeout time.Duration) *GetDeviceCellularGatewayPortForwardingRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device cellular gateway port forwarding rules params
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device cellular gateway port forwarding rules params
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) WithContext(ctx context.Context) *GetDeviceCellularGatewayPortForwardingRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device cellular gateway port forwarding rules params
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device cellular gateway port forwarding rules params
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) WithHTTPClient(client *http.Client) *GetDeviceCellularGatewayPortForwardingRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device cellular gateway port forwarding rules params
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device cellular gateway port forwarding rules params
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) WithSerial(serial string) *GetDeviceCellularGatewayPortForwardingRulesParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device cellular gateway port forwarding rules params
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceCellularGatewayPortForwardingRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
