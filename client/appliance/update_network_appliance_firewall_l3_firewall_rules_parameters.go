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

// NewUpdateNetworkApplianceFirewallL3FirewallRulesParams creates a new UpdateNetworkApplianceFirewallL3FirewallRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceFirewallL3FirewallRulesParams() *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	return &UpdateNetworkApplianceFirewallL3FirewallRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceFirewallL3FirewallRulesParamsWithTimeout creates a new UpdateNetworkApplianceFirewallL3FirewallRulesParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceFirewallL3FirewallRulesParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	return &UpdateNetworkApplianceFirewallL3FirewallRulesParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceFirewallL3FirewallRulesParamsWithContext creates a new UpdateNetworkApplianceFirewallL3FirewallRulesParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceFirewallL3FirewallRulesParamsWithContext(ctx context.Context) *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	return &UpdateNetworkApplianceFirewallL3FirewallRulesParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceFirewallL3FirewallRulesParamsWithHTTPClient creates a new UpdateNetworkApplianceFirewallL3FirewallRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceFirewallL3FirewallRulesParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	return &UpdateNetworkApplianceFirewallL3FirewallRulesParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceFirewallL3FirewallRulesParams contains all the parameters to send to the API endpoint

	for the update network appliance firewall l3 firewall rules operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceFirewallL3FirewallRulesParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkApplianceFirewallL3FirewallRules.
	UpdateNetworkApplianceFirewallL3FirewallRules UpdateNetworkApplianceFirewallL3FirewallRulesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance firewall l3 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) WithDefaults() *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance firewall l3 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) WithContext(ctx context.Context) *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) WithNetworkID(networkID string) *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceFirewallL3FirewallRules adds the updateNetworkApplianceFirewallL3FirewallRules to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) WithUpdateNetworkApplianceFirewallL3FirewallRules(updateNetworkApplianceFirewallL3FirewallRules UpdateNetworkApplianceFirewallL3FirewallRulesBody) *UpdateNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetUpdateNetworkApplianceFirewallL3FirewallRules(updateNetworkApplianceFirewallL3FirewallRules)
	return o
}

// SetUpdateNetworkApplianceFirewallL3FirewallRules adds the updateNetworkApplianceFirewallL3FirewallRules to the update network appliance firewall l3 firewall rules params
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) SetUpdateNetworkApplianceFirewallL3FirewallRules(updateNetworkApplianceFirewallL3FirewallRules UpdateNetworkApplianceFirewallL3FirewallRulesBody) {
	o.UpdateNetworkApplianceFirewallL3FirewallRules = updateNetworkApplianceFirewallL3FirewallRules
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceFirewallL3FirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceFirewallL3FirewallRules); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
