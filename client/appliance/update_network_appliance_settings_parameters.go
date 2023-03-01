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

// NewUpdateNetworkApplianceSettingsParams creates a new UpdateNetworkApplianceSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceSettingsParams() *UpdateNetworkApplianceSettingsParams {
	return &UpdateNetworkApplianceSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceSettingsParamsWithTimeout creates a new UpdateNetworkApplianceSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceSettingsParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceSettingsParams {
	return &UpdateNetworkApplianceSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceSettingsParamsWithContext creates a new UpdateNetworkApplianceSettingsParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceSettingsParamsWithContext(ctx context.Context) *UpdateNetworkApplianceSettingsParams {
	return &UpdateNetworkApplianceSettingsParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceSettingsParamsWithHTTPClient creates a new UpdateNetworkApplianceSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceSettingsParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceSettingsParams {
	return &UpdateNetworkApplianceSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceSettingsParams contains all the parameters to send to the API endpoint

	for the update network appliance settings operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceSettingsParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkApplianceSettings.
	UpdateNetworkApplianceSettings UpdateNetworkApplianceSettingsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceSettingsParams) WithDefaults() *UpdateNetworkApplianceSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) WithContext(ctx context.Context) *UpdateNetworkApplianceSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) WithNetworkID(networkID string) *UpdateNetworkApplianceSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceSettings adds the updateNetworkApplianceSettings to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) WithUpdateNetworkApplianceSettings(updateNetworkApplianceSettings UpdateNetworkApplianceSettingsBody) *UpdateNetworkApplianceSettingsParams {
	o.SetUpdateNetworkApplianceSettings(updateNetworkApplianceSettings)
	return o
}

// SetUpdateNetworkApplianceSettings adds the updateNetworkApplianceSettings to the update network appliance settings params
func (o *UpdateNetworkApplianceSettingsParams) SetUpdateNetworkApplianceSettings(updateNetworkApplianceSettings UpdateNetworkApplianceSettingsBody) {
	o.UpdateNetworkApplianceSettings = updateNetworkApplianceSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceSettings); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
