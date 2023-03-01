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

// NewGetNetworkApplianceFirewallL3FirewallRulesParams creates a new GetNetworkApplianceFirewallL3FirewallRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkApplianceFirewallL3FirewallRulesParams() *GetNetworkApplianceFirewallL3FirewallRulesParams {
	return &GetNetworkApplianceFirewallL3FirewallRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkApplianceFirewallL3FirewallRulesParamsWithTimeout creates a new GetNetworkApplianceFirewallL3FirewallRulesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkApplianceFirewallL3FirewallRulesParamsWithTimeout(timeout time.Duration) *GetNetworkApplianceFirewallL3FirewallRulesParams {
	return &GetNetworkApplianceFirewallL3FirewallRulesParams{
		timeout: timeout,
	}
}

// NewGetNetworkApplianceFirewallL3FirewallRulesParamsWithContext creates a new GetNetworkApplianceFirewallL3FirewallRulesParams object
// with the ability to set a context for a request.
func NewGetNetworkApplianceFirewallL3FirewallRulesParamsWithContext(ctx context.Context) *GetNetworkApplianceFirewallL3FirewallRulesParams {
	return &GetNetworkApplianceFirewallL3FirewallRulesParams{
		Context: ctx,
	}
}

// NewGetNetworkApplianceFirewallL3FirewallRulesParamsWithHTTPClient creates a new GetNetworkApplianceFirewallL3FirewallRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkApplianceFirewallL3FirewallRulesParamsWithHTTPClient(client *http.Client) *GetNetworkApplianceFirewallL3FirewallRulesParams {
	return &GetNetworkApplianceFirewallL3FirewallRulesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkApplianceFirewallL3FirewallRulesParams contains all the parameters to send to the API endpoint

	for the get network appliance firewall l3 firewall rules operation.

	Typically these are written to a http.Request.
*/
type GetNetworkApplianceFirewallL3FirewallRulesParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network appliance firewall l3 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) WithDefaults() *GetNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network appliance firewall l3 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network appliance firewall l3 firewall rules params
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) WithTimeout(timeout time.Duration) *GetNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network appliance firewall l3 firewall rules params
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network appliance firewall l3 firewall rules params
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) WithContext(ctx context.Context) *GetNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network appliance firewall l3 firewall rules params
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network appliance firewall l3 firewall rules params
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) WithHTTPClient(client *http.Client) *GetNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network appliance firewall l3 firewall rules params
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network appliance firewall l3 firewall rules params
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) WithNetworkID(networkID string) *GetNetworkApplianceFirewallL3FirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network appliance firewall l3 firewall rules params
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkApplianceFirewallL3FirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}