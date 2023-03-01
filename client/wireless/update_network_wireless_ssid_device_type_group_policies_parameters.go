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

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams creates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	return &UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithTimeout creates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithTimeout(timeout time.Duration) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	return &UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithContext creates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams object
// with the ability to set a context for a request.
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithContext(ctx context.Context) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	return &UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams{
		Context: ctx,
	}
}

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithHTTPClient creates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithHTTPClient(client *http.Client) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	return &UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams contains all the parameters to send to the API endpoint

	for the update network wireless ssid device type group policies operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams struct {

	// NetworkID.
	NetworkID string

	// Number.
	Number string

	// UpdateNetworkWirelessSsidDeviceTypeGroupPolicies.
	UpdateNetworkWirelessSsidDeviceTypeGroupPolicies UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network wireless ssid device type group policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithDefaults() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network wireless ssid device type group policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithTimeout(timeout time.Duration) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithContext(ctx context.Context) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithHTTPClient(client *http.Client) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithNetworkID(networkID string) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithNumber(number string) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetNumber(number string) {
	o.Number = number
}

// WithUpdateNetworkWirelessSsidDeviceTypeGroupPolicies adds the updateNetworkWirelessSsidDeviceTypeGroupPolicies to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithUpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetUpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies)
	return o
}

// SetUpdateNetworkWirelessSsidDeviceTypeGroupPolicies adds the updateNetworkWirelessSsidDeviceTypeGroupPolicies to the update network wireless ssid device type group policies params
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetUpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesBody) {
	o.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies = updateNetworkWirelessSsidDeviceTypeGroupPolicies
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
