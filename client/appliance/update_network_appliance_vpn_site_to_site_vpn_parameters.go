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

// NewUpdateNetworkApplianceVpnSiteToSiteVpnParams creates a new UpdateNetworkApplianceVpnSiteToSiteVpnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceVpnSiteToSiteVpnParams() *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	return &UpdateNetworkApplianceVpnSiteToSiteVpnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnParamsWithTimeout creates a new UpdateNetworkApplianceVpnSiteToSiteVpnParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceVpnSiteToSiteVpnParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	return &UpdateNetworkApplianceVpnSiteToSiteVpnParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnParamsWithContext creates a new UpdateNetworkApplianceVpnSiteToSiteVpnParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceVpnSiteToSiteVpnParamsWithContext(ctx context.Context) *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	return &UpdateNetworkApplianceVpnSiteToSiteVpnParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnParamsWithHTTPClient creates a new UpdateNetworkApplianceVpnSiteToSiteVpnParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceVpnSiteToSiteVpnParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	return &UpdateNetworkApplianceVpnSiteToSiteVpnParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceVpnSiteToSiteVpnParams contains all the parameters to send to the API endpoint

	for the update network appliance vpn site to site vpn operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceVpnSiteToSiteVpnParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkApplianceVpnSiteToSiteVpn.
	UpdateNetworkApplianceVpnSiteToSiteVpn UpdateNetworkApplianceVpnSiteToSiteVpnBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance vpn site to site vpn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) WithDefaults() *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance vpn site to site vpn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) WithContext(ctx context.Context) *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) WithNetworkID(networkID string) *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceVpnSiteToSiteVpn adds the updateNetworkApplianceVpnSiteToSiteVpn to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) WithUpdateNetworkApplianceVpnSiteToSiteVpn(updateNetworkApplianceVpnSiteToSiteVpn UpdateNetworkApplianceVpnSiteToSiteVpnBody) *UpdateNetworkApplianceVpnSiteToSiteVpnParams {
	o.SetUpdateNetworkApplianceVpnSiteToSiteVpn(updateNetworkApplianceVpnSiteToSiteVpn)
	return o
}

// SetUpdateNetworkApplianceVpnSiteToSiteVpn adds the updateNetworkApplianceVpnSiteToSiteVpn to the update network appliance vpn site to site vpn params
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) SetUpdateNetworkApplianceVpnSiteToSiteVpn(updateNetworkApplianceVpnSiteToSiteVpn UpdateNetworkApplianceVpnSiteToSiteVpnBody) {
	o.UpdateNetworkApplianceVpnSiteToSiteVpn = updateNetworkApplianceVpnSiteToSiteVpn
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceVpnSiteToSiteVpn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
