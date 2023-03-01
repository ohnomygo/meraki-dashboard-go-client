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

// NewGetNetworkApplianceFirewallInboundCellularFirewallRulesParams creates a new GetNetworkApplianceFirewallInboundCellularFirewallRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkApplianceFirewallInboundCellularFirewallRulesParams() *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	return &GetNetworkApplianceFirewallInboundCellularFirewallRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkApplianceFirewallInboundCellularFirewallRulesParamsWithTimeout creates a new GetNetworkApplianceFirewallInboundCellularFirewallRulesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkApplianceFirewallInboundCellularFirewallRulesParamsWithTimeout(timeout time.Duration) *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	return &GetNetworkApplianceFirewallInboundCellularFirewallRulesParams{
		timeout: timeout,
	}
}

// NewGetNetworkApplianceFirewallInboundCellularFirewallRulesParamsWithContext creates a new GetNetworkApplianceFirewallInboundCellularFirewallRulesParams object
// with the ability to set a context for a request.
func NewGetNetworkApplianceFirewallInboundCellularFirewallRulesParamsWithContext(ctx context.Context) *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	return &GetNetworkApplianceFirewallInboundCellularFirewallRulesParams{
		Context: ctx,
	}
}

// NewGetNetworkApplianceFirewallInboundCellularFirewallRulesParamsWithHTTPClient creates a new GetNetworkApplianceFirewallInboundCellularFirewallRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkApplianceFirewallInboundCellularFirewallRulesParamsWithHTTPClient(client *http.Client) *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	return &GetNetworkApplianceFirewallInboundCellularFirewallRulesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkApplianceFirewallInboundCellularFirewallRulesParams contains all the parameters to send to the API endpoint

	for the get network appliance firewall inbound cellular firewall rules operation.

	Typically these are written to a http.Request.
*/
type GetNetworkApplianceFirewallInboundCellularFirewallRulesParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network appliance firewall inbound cellular firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) WithDefaults() *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network appliance firewall inbound cellular firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network appliance firewall inbound cellular firewall rules params
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) WithTimeout(timeout time.Duration) *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network appliance firewall inbound cellular firewall rules params
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network appliance firewall inbound cellular firewall rules params
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) WithContext(ctx context.Context) *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network appliance firewall inbound cellular firewall rules params
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network appliance firewall inbound cellular firewall rules params
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) WithHTTPClient(client *http.Client) *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network appliance firewall inbound cellular firewall rules params
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network appliance firewall inbound cellular firewall rules params
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) WithNetworkID(networkID string) *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network appliance firewall inbound cellular firewall rules params
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkApplianceFirewallInboundCellularFirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
