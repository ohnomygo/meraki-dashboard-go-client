// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// NewGetNetworkSwitchStormControlParams creates a new GetNetworkSwitchStormControlParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSwitchStormControlParams() *GetNetworkSwitchStormControlParams {
	return &GetNetworkSwitchStormControlParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchStormControlParamsWithTimeout creates a new GetNetworkSwitchStormControlParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSwitchStormControlParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchStormControlParams {
	return &GetNetworkSwitchStormControlParams{
		timeout: timeout,
	}
}

// NewGetNetworkSwitchStormControlParamsWithContext creates a new GetNetworkSwitchStormControlParams object
// with the ability to set a context for a request.
func NewGetNetworkSwitchStormControlParamsWithContext(ctx context.Context) *GetNetworkSwitchStormControlParams {
	return &GetNetworkSwitchStormControlParams{
		Context: ctx,
	}
}

// NewGetNetworkSwitchStormControlParamsWithHTTPClient creates a new GetNetworkSwitchStormControlParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSwitchStormControlParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchStormControlParams {
	return &GetNetworkSwitchStormControlParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSwitchStormControlParams contains all the parameters to send to the API endpoint

	for the get network switch storm control operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSwitchStormControlParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network switch storm control params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchStormControlParams) WithDefaults() *GetNetworkSwitchStormControlParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network switch storm control params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchStormControlParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network switch storm control params
func (o *GetNetworkSwitchStormControlParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchStormControlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch storm control params
func (o *GetNetworkSwitchStormControlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch storm control params
func (o *GetNetworkSwitchStormControlParams) WithContext(ctx context.Context) *GetNetworkSwitchStormControlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch storm control params
func (o *GetNetworkSwitchStormControlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch storm control params
func (o *GetNetworkSwitchStormControlParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchStormControlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch storm control params
func (o *GetNetworkSwitchStormControlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network switch storm control params
func (o *GetNetworkSwitchStormControlParams) WithNetworkID(networkID string) *GetNetworkSwitchStormControlParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch storm control params
func (o *GetNetworkSwitchStormControlParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchStormControlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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