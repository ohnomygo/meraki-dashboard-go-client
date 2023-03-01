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

// NewUpdateNetworkWirelessSsidParams creates a new UpdateNetworkWirelessSsidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkWirelessSsidParams() *UpdateNetworkWirelessSsidParams {
	return &UpdateNetworkWirelessSsidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWirelessSsidParamsWithTimeout creates a new UpdateNetworkWirelessSsidParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkWirelessSsidParamsWithTimeout(timeout time.Duration) *UpdateNetworkWirelessSsidParams {
	return &UpdateNetworkWirelessSsidParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkWirelessSsidParamsWithContext creates a new UpdateNetworkWirelessSsidParams object
// with the ability to set a context for a request.
func NewUpdateNetworkWirelessSsidParamsWithContext(ctx context.Context) *UpdateNetworkWirelessSsidParams {
	return &UpdateNetworkWirelessSsidParams{
		Context: ctx,
	}
}

// NewUpdateNetworkWirelessSsidParamsWithHTTPClient creates a new UpdateNetworkWirelessSsidParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkWirelessSsidParamsWithHTTPClient(client *http.Client) *UpdateNetworkWirelessSsidParams {
	return &UpdateNetworkWirelessSsidParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkWirelessSsidParams contains all the parameters to send to the API endpoint

	for the update network wireless ssid operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkWirelessSsidParams struct {

	// NetworkID.
	NetworkID string

	// Number.
	Number string

	// UpdateNetworkWirelessSsid.
	UpdateNetworkWirelessSsid UpdateNetworkWirelessSsidBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network wireless ssid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessSsidParams) WithDefaults() *UpdateNetworkWirelessSsidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network wireless ssid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessSsidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) WithTimeout(timeout time.Duration) *UpdateNetworkWirelessSsidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) WithContext(ctx context.Context) *UpdateNetworkWirelessSsidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) WithHTTPClient(client *http.Client) *UpdateNetworkWirelessSsidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) WithNetworkID(networkID string) *UpdateNetworkWirelessSsidParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) WithNumber(number string) *UpdateNetworkWirelessSsidParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) SetNumber(number string) {
	o.Number = number
}

// WithUpdateNetworkWirelessSsid adds the updateNetworkWirelessSsid to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) WithUpdateNetworkWirelessSsid(updateNetworkWirelessSsid UpdateNetworkWirelessSsidBody) *UpdateNetworkWirelessSsidParams {
	o.SetUpdateNetworkWirelessSsid(updateNetworkWirelessSsid)
	return o
}

// SetUpdateNetworkWirelessSsid adds the updateNetworkWirelessSsid to the update network wireless ssid params
func (o *UpdateNetworkWirelessSsidParams) SetUpdateNetworkWirelessSsid(updateNetworkWirelessSsid UpdateNetworkWirelessSsidBody) {
	o.UpdateNetworkWirelessSsid = updateNetworkWirelessSsid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWirelessSsidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.UpdateNetworkWirelessSsid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
