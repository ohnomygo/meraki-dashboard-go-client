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

// NewRemoveNetworkSwitchStackParams creates a new RemoveNetworkSwitchStackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveNetworkSwitchStackParams() *RemoveNetworkSwitchStackParams {
	return &RemoveNetworkSwitchStackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveNetworkSwitchStackParamsWithTimeout creates a new RemoveNetworkSwitchStackParams object
// with the ability to set a timeout on a request.
func NewRemoveNetworkSwitchStackParamsWithTimeout(timeout time.Duration) *RemoveNetworkSwitchStackParams {
	return &RemoveNetworkSwitchStackParams{
		timeout: timeout,
	}
}

// NewRemoveNetworkSwitchStackParamsWithContext creates a new RemoveNetworkSwitchStackParams object
// with the ability to set a context for a request.
func NewRemoveNetworkSwitchStackParamsWithContext(ctx context.Context) *RemoveNetworkSwitchStackParams {
	return &RemoveNetworkSwitchStackParams{
		Context: ctx,
	}
}

// NewRemoveNetworkSwitchStackParamsWithHTTPClient creates a new RemoveNetworkSwitchStackParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveNetworkSwitchStackParamsWithHTTPClient(client *http.Client) *RemoveNetworkSwitchStackParams {
	return &RemoveNetworkSwitchStackParams{
		HTTPClient: client,
	}
}

/*
RemoveNetworkSwitchStackParams contains all the parameters to send to the API endpoint

	for the remove network switch stack operation.

	Typically these are written to a http.Request.
*/
type RemoveNetworkSwitchStackParams struct {

	// NetworkID.
	NetworkID string

	// RemoveNetworkSwitchStack.
	RemoveNetworkSwitchStack RemoveNetworkSwitchStackBody

	// SwitchStackID.
	SwitchStackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove network switch stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveNetworkSwitchStackParams) WithDefaults() *RemoveNetworkSwitchStackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove network switch stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveNetworkSwitchStackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) WithTimeout(timeout time.Duration) *RemoveNetworkSwitchStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) WithContext(ctx context.Context) *RemoveNetworkSwitchStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) WithHTTPClient(client *http.Client) *RemoveNetworkSwitchStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) WithNetworkID(networkID string) *RemoveNetworkSwitchStackParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRemoveNetworkSwitchStack adds the removeNetworkSwitchStack to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) WithRemoveNetworkSwitchStack(removeNetworkSwitchStack RemoveNetworkSwitchStackBody) *RemoveNetworkSwitchStackParams {
	o.SetRemoveNetworkSwitchStack(removeNetworkSwitchStack)
	return o
}

// SetRemoveNetworkSwitchStack adds the removeNetworkSwitchStack to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) SetRemoveNetworkSwitchStack(removeNetworkSwitchStack RemoveNetworkSwitchStackBody) {
	o.RemoveNetworkSwitchStack = removeNetworkSwitchStack
}

// WithSwitchStackID adds the switchStackID to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) WithSwitchStackID(switchStackID string) *RemoveNetworkSwitchStackParams {
	o.SetSwitchStackID(switchStackID)
	return o
}

// SetSwitchStackID adds the switchStackId to the remove network switch stack params
func (o *RemoveNetworkSwitchStackParams) SetSwitchStackID(switchStackID string) {
	o.SwitchStackID = switchStackID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveNetworkSwitchStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.RemoveNetworkSwitchStack); err != nil {
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