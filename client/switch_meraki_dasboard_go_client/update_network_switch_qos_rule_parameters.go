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

// NewUpdateNetworkSwitchQosRuleParams creates a new UpdateNetworkSwitchQosRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkSwitchQosRuleParams() *UpdateNetworkSwitchQosRuleParams {
	return &UpdateNetworkSwitchQosRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchQosRuleParamsWithTimeout creates a new UpdateNetworkSwitchQosRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkSwitchQosRuleParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchQosRuleParams {
	return &UpdateNetworkSwitchQosRuleParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchQosRuleParamsWithContext creates a new UpdateNetworkSwitchQosRuleParams object
// with the ability to set a context for a request.
func NewUpdateNetworkSwitchQosRuleParamsWithContext(ctx context.Context) *UpdateNetworkSwitchQosRuleParams {
	return &UpdateNetworkSwitchQosRuleParams{
		Context: ctx,
	}
}

// NewUpdateNetworkSwitchQosRuleParamsWithHTTPClient creates a new UpdateNetworkSwitchQosRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkSwitchQosRuleParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchQosRuleParams {
	return &UpdateNetworkSwitchQosRuleParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkSwitchQosRuleParams contains all the parameters to send to the API endpoint

	for the update network switch qos rule operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkSwitchQosRuleParams struct {

	// NetworkID.
	NetworkID string

	// QosRuleID.
	QosRuleID string

	// UpdateNetworkSwitchQosRule.
	UpdateNetworkSwitchQosRule UpdateNetworkSwitchQosRuleBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network switch qos rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchQosRuleParams) WithDefaults() *UpdateNetworkSwitchQosRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network switch qos rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchQosRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchQosRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) WithContext(ctx context.Context) *UpdateNetworkSwitchQosRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchQosRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) WithNetworkID(networkID string) *UpdateNetworkSwitchQosRuleParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithQosRuleID adds the qosRuleID to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) WithQosRuleID(qosRuleID string) *UpdateNetworkSwitchQosRuleParams {
	o.SetQosRuleID(qosRuleID)
	return o
}

// SetQosRuleID adds the qosRuleId to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) SetQosRuleID(qosRuleID string) {
	o.QosRuleID = qosRuleID
}

// WithUpdateNetworkSwitchQosRule adds the updateNetworkSwitchQosRule to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) WithUpdateNetworkSwitchQosRule(updateNetworkSwitchQosRule UpdateNetworkSwitchQosRuleBody) *UpdateNetworkSwitchQosRuleParams {
	o.SetUpdateNetworkSwitchQosRule(updateNetworkSwitchQosRule)
	return o
}

// SetUpdateNetworkSwitchQosRule adds the updateNetworkSwitchQosRule to the update network switch qos rule params
func (o *UpdateNetworkSwitchQosRuleParams) SetUpdateNetworkSwitchQosRule(updateNetworkSwitchQosRule UpdateNetworkSwitchQosRuleBody) {
	o.UpdateNetworkSwitchQosRule = updateNetworkSwitchQosRule
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchQosRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param qosRuleId
	if err := r.SetPathParam("qosRuleId", o.QosRuleID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkSwitchQosRule); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}