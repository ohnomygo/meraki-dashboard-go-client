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

// NewUpdateNetworkWirelessBillingParams creates a new UpdateNetworkWirelessBillingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkWirelessBillingParams() *UpdateNetworkWirelessBillingParams {
	return &UpdateNetworkWirelessBillingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWirelessBillingParamsWithTimeout creates a new UpdateNetworkWirelessBillingParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkWirelessBillingParamsWithTimeout(timeout time.Duration) *UpdateNetworkWirelessBillingParams {
	return &UpdateNetworkWirelessBillingParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkWirelessBillingParamsWithContext creates a new UpdateNetworkWirelessBillingParams object
// with the ability to set a context for a request.
func NewUpdateNetworkWirelessBillingParamsWithContext(ctx context.Context) *UpdateNetworkWirelessBillingParams {
	return &UpdateNetworkWirelessBillingParams{
		Context: ctx,
	}
}

// NewUpdateNetworkWirelessBillingParamsWithHTTPClient creates a new UpdateNetworkWirelessBillingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkWirelessBillingParamsWithHTTPClient(client *http.Client) *UpdateNetworkWirelessBillingParams {
	return &UpdateNetworkWirelessBillingParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkWirelessBillingParams contains all the parameters to send to the API endpoint

	for the update network wireless billing operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkWirelessBillingParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkWirelessBilling.
	UpdateNetworkWirelessBilling UpdateNetworkWirelessBillingBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network wireless billing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessBillingParams) WithDefaults() *UpdateNetworkWirelessBillingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network wireless billing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessBillingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) WithTimeout(timeout time.Duration) *UpdateNetworkWirelessBillingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) WithContext(ctx context.Context) *UpdateNetworkWirelessBillingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) WithHTTPClient(client *http.Client) *UpdateNetworkWirelessBillingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) WithNetworkID(networkID string) *UpdateNetworkWirelessBillingParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkWirelessBilling adds the updateNetworkWirelessBilling to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) WithUpdateNetworkWirelessBilling(updateNetworkWirelessBilling UpdateNetworkWirelessBillingBody) *UpdateNetworkWirelessBillingParams {
	o.SetUpdateNetworkWirelessBilling(updateNetworkWirelessBilling)
	return o
}

// SetUpdateNetworkWirelessBilling adds the updateNetworkWirelessBilling to the update network wireless billing params
func (o *UpdateNetworkWirelessBillingParams) SetUpdateNetworkWirelessBilling(updateNetworkWirelessBilling UpdateNetworkWirelessBillingBody) {
	o.UpdateNetworkWirelessBilling = updateNetworkWirelessBilling
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWirelessBillingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkWirelessBilling); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
