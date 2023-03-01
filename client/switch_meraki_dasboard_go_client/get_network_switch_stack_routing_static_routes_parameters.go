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

// NewGetNetworkSwitchStackRoutingStaticRoutesParams creates a new GetNetworkSwitchStackRoutingStaticRoutesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSwitchStackRoutingStaticRoutesParams() *GetNetworkSwitchStackRoutingStaticRoutesParams {
	return &GetNetworkSwitchStackRoutingStaticRoutesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchStackRoutingStaticRoutesParamsWithTimeout creates a new GetNetworkSwitchStackRoutingStaticRoutesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSwitchStackRoutingStaticRoutesParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchStackRoutingStaticRoutesParams {
	return &GetNetworkSwitchStackRoutingStaticRoutesParams{
		timeout: timeout,
	}
}

// NewGetNetworkSwitchStackRoutingStaticRoutesParamsWithContext creates a new GetNetworkSwitchStackRoutingStaticRoutesParams object
// with the ability to set a context for a request.
func NewGetNetworkSwitchStackRoutingStaticRoutesParamsWithContext(ctx context.Context) *GetNetworkSwitchStackRoutingStaticRoutesParams {
	return &GetNetworkSwitchStackRoutingStaticRoutesParams{
		Context: ctx,
	}
}

// NewGetNetworkSwitchStackRoutingStaticRoutesParamsWithHTTPClient creates a new GetNetworkSwitchStackRoutingStaticRoutesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSwitchStackRoutingStaticRoutesParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchStackRoutingStaticRoutesParams {
	return &GetNetworkSwitchStackRoutingStaticRoutesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSwitchStackRoutingStaticRoutesParams contains all the parameters to send to the API endpoint

	for the get network switch stack routing static routes operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSwitchStackRoutingStaticRoutesParams struct {

	// NetworkID.
	NetworkID string

	// SwitchStackID.
	SwitchStackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network switch stack routing static routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) WithDefaults() *GetNetworkSwitchStackRoutingStaticRoutesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network switch stack routing static routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchStackRoutingStaticRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) WithContext(ctx context.Context) *GetNetworkSwitchStackRoutingStaticRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchStackRoutingStaticRoutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) WithNetworkID(networkID string) *GetNetworkSwitchStackRoutingStaticRoutesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSwitchStackID adds the switchStackID to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) WithSwitchStackID(switchStackID string) *GetNetworkSwitchStackRoutingStaticRoutesParams {
	o.SetSwitchStackID(switchStackID)
	return o
}

// SetSwitchStackID adds the switchStackId to the get network switch stack routing static routes params
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) SetSwitchStackID(switchStackID string) {
	o.SwitchStackID = switchStackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchStackRoutingStaticRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
