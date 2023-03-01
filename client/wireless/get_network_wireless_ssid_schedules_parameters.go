// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// NewGetNetworkWirelessSsidSchedulesParams creates a new GetNetworkWirelessSsidSchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessSsidSchedulesParams() *GetNetworkWirelessSsidSchedulesParams {
	return &GetNetworkWirelessSsidSchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessSsidSchedulesParamsWithTimeout creates a new GetNetworkWirelessSsidSchedulesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessSsidSchedulesParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessSsidSchedulesParams {
	return &GetNetworkWirelessSsidSchedulesParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessSsidSchedulesParamsWithContext creates a new GetNetworkWirelessSsidSchedulesParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessSsidSchedulesParamsWithContext(ctx context.Context) *GetNetworkWirelessSsidSchedulesParams {
	return &GetNetworkWirelessSsidSchedulesParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessSsidSchedulesParamsWithHTTPClient creates a new GetNetworkWirelessSsidSchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessSsidSchedulesParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessSsidSchedulesParams {
	return &GetNetworkWirelessSsidSchedulesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkWirelessSsidSchedulesParams contains all the parameters to send to the API endpoint

	for the get network wireless ssid schedules operation.

	Typically these are written to a http.Request.
*/
type GetNetworkWirelessSsidSchedulesParams struct {

	// NetworkID.
	NetworkID string

	// Number.
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless ssid schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidSchedulesParams) WithDefaults() *GetNetworkWirelessSsidSchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless ssid schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidSchedulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessSsidSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) WithContext(ctx context.Context) *GetNetworkWirelessSsidSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessSsidSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) WithNetworkID(networkID string) *GetNetworkWirelessSsidSchedulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) WithNumber(number string) *GetNetworkWirelessSsidSchedulesParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get network wireless ssid schedules params
func (o *GetNetworkWirelessSsidSchedulesParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessSsidSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
