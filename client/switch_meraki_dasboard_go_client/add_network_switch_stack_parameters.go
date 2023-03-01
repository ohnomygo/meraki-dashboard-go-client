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

// NewAddNetworkSwitchStackParams creates a new AddNetworkSwitchStackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddNetworkSwitchStackParams() *AddNetworkSwitchStackParams {
	return &AddNetworkSwitchStackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddNetworkSwitchStackParamsWithTimeout creates a new AddNetworkSwitchStackParams object
// with the ability to set a timeout on a request.
func NewAddNetworkSwitchStackParamsWithTimeout(timeout time.Duration) *AddNetworkSwitchStackParams {
	return &AddNetworkSwitchStackParams{
		timeout: timeout,
	}
}

// NewAddNetworkSwitchStackParamsWithContext creates a new AddNetworkSwitchStackParams object
// with the ability to set a context for a request.
func NewAddNetworkSwitchStackParamsWithContext(ctx context.Context) *AddNetworkSwitchStackParams {
	return &AddNetworkSwitchStackParams{
		Context: ctx,
	}
}

// NewAddNetworkSwitchStackParamsWithHTTPClient creates a new AddNetworkSwitchStackParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddNetworkSwitchStackParamsWithHTTPClient(client *http.Client) *AddNetworkSwitchStackParams {
	return &AddNetworkSwitchStackParams{
		HTTPClient: client,
	}
}

/*
AddNetworkSwitchStackParams contains all the parameters to send to the API endpoint

	for the add network switch stack operation.

	Typically these are written to a http.Request.
*/
type AddNetworkSwitchStackParams struct {

	// AddNetworkSwitchStack.
	AddNetworkSwitchStack AddNetworkSwitchStackBody

	// NetworkID.
	NetworkID string

	// SwitchStackID.
	SwitchStackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add network switch stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddNetworkSwitchStackParams) WithDefaults() *AddNetworkSwitchStackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add network switch stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddNetworkSwitchStackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add network switch stack params
func (o *AddNetworkSwitchStackParams) WithTimeout(timeout time.Duration) *AddNetworkSwitchStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add network switch stack params
func (o *AddNetworkSwitchStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add network switch stack params
func (o *AddNetworkSwitchStackParams) WithContext(ctx context.Context) *AddNetworkSwitchStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add network switch stack params
func (o *AddNetworkSwitchStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add network switch stack params
func (o *AddNetworkSwitchStackParams) WithHTTPClient(client *http.Client) *AddNetworkSwitchStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add network switch stack params
func (o *AddNetworkSwitchStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddNetworkSwitchStack adds the addNetworkSwitchStack to the add network switch stack params
func (o *AddNetworkSwitchStackParams) WithAddNetworkSwitchStack(addNetworkSwitchStack AddNetworkSwitchStackBody) *AddNetworkSwitchStackParams {
	o.SetAddNetworkSwitchStack(addNetworkSwitchStack)
	return o
}

// SetAddNetworkSwitchStack adds the addNetworkSwitchStack to the add network switch stack params
func (o *AddNetworkSwitchStackParams) SetAddNetworkSwitchStack(addNetworkSwitchStack AddNetworkSwitchStackBody) {
	o.AddNetworkSwitchStack = addNetworkSwitchStack
}

// WithNetworkID adds the networkID to the add network switch stack params
func (o *AddNetworkSwitchStackParams) WithNetworkID(networkID string) *AddNetworkSwitchStackParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the add network switch stack params
func (o *AddNetworkSwitchStackParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSwitchStackID adds the switchStackID to the add network switch stack params
func (o *AddNetworkSwitchStackParams) WithSwitchStackID(switchStackID string) *AddNetworkSwitchStackParams {
	o.SetSwitchStackID(switchStackID)
	return o
}

// SetSwitchStackID adds the switchStackId to the add network switch stack params
func (o *AddNetworkSwitchStackParams) SetSwitchStackID(switchStackID string) {
	o.SwitchStackID = switchStackID
}

// WriteToRequest writes these params to a swagger request
func (o *AddNetworkSwitchStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.AddNetworkSwitchStack); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param switchStackId
	if err := r.SetPathParam("switchStackId", o.SwitchStackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
