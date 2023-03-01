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

// NewUpdateNetworkApplianceVpnBgpParams creates a new UpdateNetworkApplianceVpnBgpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceVpnBgpParams() *UpdateNetworkApplianceVpnBgpParams {
	return &UpdateNetworkApplianceVpnBgpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceVpnBgpParamsWithTimeout creates a new UpdateNetworkApplianceVpnBgpParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceVpnBgpParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceVpnBgpParams {
	return &UpdateNetworkApplianceVpnBgpParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceVpnBgpParamsWithContext creates a new UpdateNetworkApplianceVpnBgpParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceVpnBgpParamsWithContext(ctx context.Context) *UpdateNetworkApplianceVpnBgpParams {
	return &UpdateNetworkApplianceVpnBgpParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceVpnBgpParamsWithHTTPClient creates a new UpdateNetworkApplianceVpnBgpParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceVpnBgpParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceVpnBgpParams {
	return &UpdateNetworkApplianceVpnBgpParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceVpnBgpParams contains all the parameters to send to the API endpoint

	for the update network appliance vpn bgp operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceVpnBgpParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkApplianceVpnBgp.
	UpdateNetworkApplianceVpnBgp UpdateNetworkApplianceVpnBgpBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance vpn bgp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceVpnBgpParams) WithDefaults() *UpdateNetworkApplianceVpnBgpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance vpn bgp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceVpnBgpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceVpnBgpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) WithContext(ctx context.Context) *UpdateNetworkApplianceVpnBgpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceVpnBgpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) WithNetworkID(networkID string) *UpdateNetworkApplianceVpnBgpParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceVpnBgp adds the updateNetworkApplianceVpnBgp to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) WithUpdateNetworkApplianceVpnBgp(updateNetworkApplianceVpnBgp UpdateNetworkApplianceVpnBgpBody) *UpdateNetworkApplianceVpnBgpParams {
	o.SetUpdateNetworkApplianceVpnBgp(updateNetworkApplianceVpnBgp)
	return o
}

// SetUpdateNetworkApplianceVpnBgp adds the updateNetworkApplianceVpnBgp to the update network appliance vpn bgp params
func (o *UpdateNetworkApplianceVpnBgpParams) SetUpdateNetworkApplianceVpnBgp(updateNetworkApplianceVpnBgp UpdateNetworkApplianceVpnBgpBody) {
	o.UpdateNetworkApplianceVpnBgp = updateNetworkApplianceVpnBgp
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceVpnBgpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceVpnBgp); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
