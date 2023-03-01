// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// NewUpdateNetworkApplianceSsidParams creates a new UpdateNetworkApplianceSsidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceSsidParams() *UpdateNetworkApplianceSsidParams {
	return &UpdateNetworkApplianceSsidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceSsidParamsWithTimeout creates a new UpdateNetworkApplianceSsidParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceSsidParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceSsidParams {
	return &UpdateNetworkApplianceSsidParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceSsidParamsWithContext creates a new UpdateNetworkApplianceSsidParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceSsidParamsWithContext(ctx context.Context) *UpdateNetworkApplianceSsidParams {
	return &UpdateNetworkApplianceSsidParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceSsidParamsWithHTTPClient creates a new UpdateNetworkApplianceSsidParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceSsidParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceSsidParams {
	return &UpdateNetworkApplianceSsidParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceSsidParams contains all the parameters to send to the API endpoint

	for the update network appliance ssid operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceSsidParams struct {

	// NetworkID.
	NetworkID string

	// Number.
	Number string

	// UpdateNetworkApplianceSsid.
	UpdateNetworkApplianceSsid UpdateNetworkApplianceSsidBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance ssid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceSsidParams) WithDefaults() *UpdateNetworkApplianceSsidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance ssid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceSsidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceSsidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) WithContext(ctx context.Context) *UpdateNetworkApplianceSsidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceSsidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) WithNetworkID(networkID string) *UpdateNetworkApplianceSsidParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) WithNumber(number string) *UpdateNetworkApplianceSsidParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) SetNumber(number string) {
	o.Number = number
}

// WithUpdateNetworkApplianceSsid adds the updateNetworkApplianceSsid to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) WithUpdateNetworkApplianceSsid(updateNetworkApplianceSsid UpdateNetworkApplianceSsidBody) *UpdateNetworkApplianceSsidParams {
	o.SetUpdateNetworkApplianceSsid(updateNetworkApplianceSsid)
	return o
}

// SetUpdateNetworkApplianceSsid adds the updateNetworkApplianceSsid to the update network appliance ssid params
func (o *UpdateNetworkApplianceSsidParams) SetUpdateNetworkApplianceSsid(updateNetworkApplianceSsid UpdateNetworkApplianceSsidBody) {
	o.UpdateNetworkApplianceSsid = updateNetworkApplianceSsid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceSsidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.UpdateNetworkApplianceSsid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}