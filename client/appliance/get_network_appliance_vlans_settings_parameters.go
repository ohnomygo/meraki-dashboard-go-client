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

// NewGetNetworkApplianceVlansSettingsParams creates a new GetNetworkApplianceVlansSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkApplianceVlansSettingsParams() *GetNetworkApplianceVlansSettingsParams {
	return &GetNetworkApplianceVlansSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkApplianceVlansSettingsParamsWithTimeout creates a new GetNetworkApplianceVlansSettingsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkApplianceVlansSettingsParamsWithTimeout(timeout time.Duration) *GetNetworkApplianceVlansSettingsParams {
	return &GetNetworkApplianceVlansSettingsParams{
		timeout: timeout,
	}
}

// NewGetNetworkApplianceVlansSettingsParamsWithContext creates a new GetNetworkApplianceVlansSettingsParams object
// with the ability to set a context for a request.
func NewGetNetworkApplianceVlansSettingsParamsWithContext(ctx context.Context) *GetNetworkApplianceVlansSettingsParams {
	return &GetNetworkApplianceVlansSettingsParams{
		Context: ctx,
	}
}

// NewGetNetworkApplianceVlansSettingsParamsWithHTTPClient creates a new GetNetworkApplianceVlansSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkApplianceVlansSettingsParamsWithHTTPClient(client *http.Client) *GetNetworkApplianceVlansSettingsParams {
	return &GetNetworkApplianceVlansSettingsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkApplianceVlansSettingsParams contains all the parameters to send to the API endpoint

	for the get network appliance vlans settings operation.

	Typically these are written to a http.Request.
*/
type GetNetworkApplianceVlansSettingsParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network appliance vlans settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceVlansSettingsParams) WithDefaults() *GetNetworkApplianceVlansSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network appliance vlans settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceVlansSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network appliance vlans settings params
func (o *GetNetworkApplianceVlansSettingsParams) WithTimeout(timeout time.Duration) *GetNetworkApplianceVlansSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network appliance vlans settings params
func (o *GetNetworkApplianceVlansSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network appliance vlans settings params
func (o *GetNetworkApplianceVlansSettingsParams) WithContext(ctx context.Context) *GetNetworkApplianceVlansSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network appliance vlans settings params
func (o *GetNetworkApplianceVlansSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network appliance vlans settings params
func (o *GetNetworkApplianceVlansSettingsParams) WithHTTPClient(client *http.Client) *GetNetworkApplianceVlansSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network appliance vlans settings params
func (o *GetNetworkApplianceVlansSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network appliance vlans settings params
func (o *GetNetworkApplianceVlansSettingsParams) WithNetworkID(networkID string) *GetNetworkApplianceVlansSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network appliance vlans settings params
func (o *GetNetworkApplianceVlansSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkApplianceVlansSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
