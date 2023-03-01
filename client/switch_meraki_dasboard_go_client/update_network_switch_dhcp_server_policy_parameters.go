// Code generated by go-swagger; DO NOT EDIT.

package switch_meraki_dasboard_go_client

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

// NewUpdateNetworkSwitchDhcpServerPolicyParams creates a new UpdateNetworkSwitchDhcpServerPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkSwitchDhcpServerPolicyParams() *UpdateNetworkSwitchDhcpServerPolicyParams {
	return &UpdateNetworkSwitchDhcpServerPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchDhcpServerPolicyParamsWithTimeout creates a new UpdateNetworkSwitchDhcpServerPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkSwitchDhcpServerPolicyParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchDhcpServerPolicyParams {
	return &UpdateNetworkSwitchDhcpServerPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchDhcpServerPolicyParamsWithContext creates a new UpdateNetworkSwitchDhcpServerPolicyParams object
// with the ability to set a context for a request.
func NewUpdateNetworkSwitchDhcpServerPolicyParamsWithContext(ctx context.Context) *UpdateNetworkSwitchDhcpServerPolicyParams {
	return &UpdateNetworkSwitchDhcpServerPolicyParams{
		Context: ctx,
	}
}

// NewUpdateNetworkSwitchDhcpServerPolicyParamsWithHTTPClient creates a new UpdateNetworkSwitchDhcpServerPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkSwitchDhcpServerPolicyParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchDhcpServerPolicyParams {
	return &UpdateNetworkSwitchDhcpServerPolicyParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkSwitchDhcpServerPolicyParams contains all the parameters to send to the API endpoint

	for the update network switch dhcp server policy operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkSwitchDhcpServerPolicyParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkSwitchDhcpServerPolicy.
	UpdateNetworkSwitchDhcpServerPolicy UpdateNetworkSwitchDhcpServerPolicyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network switch dhcp server policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) WithDefaults() *UpdateNetworkSwitchDhcpServerPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network switch dhcp server policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchDhcpServerPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) WithContext(ctx context.Context) *UpdateNetworkSwitchDhcpServerPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchDhcpServerPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) WithNetworkID(networkID string) *UpdateNetworkSwitchDhcpServerPolicyParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkSwitchDhcpServerPolicy adds the updateNetworkSwitchDhcpServerPolicy to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) WithUpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy UpdateNetworkSwitchDhcpServerPolicyBody) *UpdateNetworkSwitchDhcpServerPolicyParams {
	o.SetUpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy)
	return o
}

// SetUpdateNetworkSwitchDhcpServerPolicy adds the updateNetworkSwitchDhcpServerPolicy to the update network switch dhcp server policy params
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) SetUpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy UpdateNetworkSwitchDhcpServerPolicyBody) {
	o.UpdateNetworkSwitchDhcpServerPolicy = updateNetworkSwitchDhcpServerPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchDhcpServerPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkSwitchDhcpServerPolicy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
