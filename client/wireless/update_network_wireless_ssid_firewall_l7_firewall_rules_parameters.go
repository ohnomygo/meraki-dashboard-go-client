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

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesParams creates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesParams() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	return &UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsWithTimeout creates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsWithTimeout(timeout time.Duration) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	return &UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsWithContext creates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams object
// with the ability to set a context for a request.
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsWithContext(ctx context.Context) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	return &UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams{
		Context: ctx,
	}
}

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsWithHTTPClient creates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsWithHTTPClient(client *http.Client) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	return &UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams contains all the parameters to send to the API endpoint

	for the update network wireless ssid firewall l7 firewall rules operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Number.

	   Number
	*/
	Number string

	// UpdateNetworkWirelessSsidFirewallL7FirewallRules.
	UpdateNetworkWirelessSsidFirewallL7FirewallRules UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network wireless ssid firewall l7 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) WithDefaults() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network wireless ssid firewall l7 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) WithTimeout(timeout time.Duration) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) WithContext(ctx context.Context) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) WithHTTPClient(client *http.Client) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) WithNetworkID(networkID string) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) WithNumber(number string) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) SetNumber(number string) {
	o.Number = number
}

// WithUpdateNetworkWirelessSsidFirewallL7FirewallRules adds the updateNetworkWirelessSsidFirewallL7FirewallRules to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) WithUpdateNetworkWirelessSsidFirewallL7FirewallRules(updateNetworkWirelessSsidFirewallL7FirewallRules UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody) *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetUpdateNetworkWirelessSsidFirewallL7FirewallRules(updateNetworkWirelessSsidFirewallL7FirewallRules)
	return o
}

// SetUpdateNetworkWirelessSsidFirewallL7FirewallRules adds the updateNetworkWirelessSsidFirewallL7FirewallRules to the update network wireless ssid firewall l7 firewall rules params
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) SetUpdateNetworkWirelessSsidFirewallL7FirewallRules(updateNetworkWirelessSsidFirewallL7FirewallRules UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody) {
	o.UpdateNetworkWirelessSsidFirewallL7FirewallRules = updateNetworkWirelessSsidFirewallL7FirewallRules
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.UpdateNetworkWirelessSsidFirewallL7FirewallRules); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
