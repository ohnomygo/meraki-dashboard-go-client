// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewUpdateNetworkFirmwareUpgradesParams creates a new UpdateNetworkFirmwareUpgradesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkFirmwareUpgradesParams() *UpdateNetworkFirmwareUpgradesParams {
	return &UpdateNetworkFirmwareUpgradesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkFirmwareUpgradesParamsWithTimeout creates a new UpdateNetworkFirmwareUpgradesParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkFirmwareUpgradesParamsWithTimeout(timeout time.Duration) *UpdateNetworkFirmwareUpgradesParams {
	return &UpdateNetworkFirmwareUpgradesParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkFirmwareUpgradesParamsWithContext creates a new UpdateNetworkFirmwareUpgradesParams object
// with the ability to set a context for a request.
func NewUpdateNetworkFirmwareUpgradesParamsWithContext(ctx context.Context) *UpdateNetworkFirmwareUpgradesParams {
	return &UpdateNetworkFirmwareUpgradesParams{
		Context: ctx,
	}
}

// NewUpdateNetworkFirmwareUpgradesParamsWithHTTPClient creates a new UpdateNetworkFirmwareUpgradesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkFirmwareUpgradesParamsWithHTTPClient(client *http.Client) *UpdateNetworkFirmwareUpgradesParams {
	return &UpdateNetworkFirmwareUpgradesParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkFirmwareUpgradesParams contains all the parameters to send to the API endpoint

	for the update network firmware upgrades operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkFirmwareUpgradesParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkFirmwareUpgrades.
	UpdateNetworkFirmwareUpgrades UpdateNetworkFirmwareUpgradesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network firmware upgrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkFirmwareUpgradesParams) WithDefaults() *UpdateNetworkFirmwareUpgradesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network firmware upgrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkFirmwareUpgradesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) WithTimeout(timeout time.Duration) *UpdateNetworkFirmwareUpgradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) WithContext(ctx context.Context) *UpdateNetworkFirmwareUpgradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) WithHTTPClient(client *http.Client) *UpdateNetworkFirmwareUpgradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) WithNetworkID(networkID string) *UpdateNetworkFirmwareUpgradesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkFirmwareUpgrades adds the updateNetworkFirmwareUpgrades to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) WithUpdateNetworkFirmwareUpgrades(updateNetworkFirmwareUpgrades UpdateNetworkFirmwareUpgradesBody) *UpdateNetworkFirmwareUpgradesParams {
	o.SetUpdateNetworkFirmwareUpgrades(updateNetworkFirmwareUpgrades)
	return o
}

// SetUpdateNetworkFirmwareUpgrades adds the updateNetworkFirmwareUpgrades to the update network firmware upgrades params
func (o *UpdateNetworkFirmwareUpgradesParams) SetUpdateNetworkFirmwareUpgrades(updateNetworkFirmwareUpgrades UpdateNetworkFirmwareUpgradesBody) {
	o.UpdateNetworkFirmwareUpgrades = updateNetworkFirmwareUpgrades
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkFirmwareUpgradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkFirmwareUpgrades); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
