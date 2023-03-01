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

// NewGetNetworkTopologyLinkLayerParams creates a new GetNetworkTopologyLinkLayerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkTopologyLinkLayerParams() *GetNetworkTopologyLinkLayerParams {
	return &GetNetworkTopologyLinkLayerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkTopologyLinkLayerParamsWithTimeout creates a new GetNetworkTopologyLinkLayerParams object
// with the ability to set a timeout on a request.
func NewGetNetworkTopologyLinkLayerParamsWithTimeout(timeout time.Duration) *GetNetworkTopologyLinkLayerParams {
	return &GetNetworkTopologyLinkLayerParams{
		timeout: timeout,
	}
}

// NewGetNetworkTopologyLinkLayerParamsWithContext creates a new GetNetworkTopologyLinkLayerParams object
// with the ability to set a context for a request.
func NewGetNetworkTopologyLinkLayerParamsWithContext(ctx context.Context) *GetNetworkTopologyLinkLayerParams {
	return &GetNetworkTopologyLinkLayerParams{
		Context: ctx,
	}
}

// NewGetNetworkTopologyLinkLayerParamsWithHTTPClient creates a new GetNetworkTopologyLinkLayerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkTopologyLinkLayerParamsWithHTTPClient(client *http.Client) *GetNetworkTopologyLinkLayerParams {
	return &GetNetworkTopologyLinkLayerParams{
		HTTPClient: client,
	}
}

/*
GetNetworkTopologyLinkLayerParams contains all the parameters to send to the API endpoint

	for the get network topology link layer operation.

	Typically these are written to a http.Request.
*/
type GetNetworkTopologyLinkLayerParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network topology link layer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkTopologyLinkLayerParams) WithDefaults() *GetNetworkTopologyLinkLayerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network topology link layer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkTopologyLinkLayerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network topology link layer params
func (o *GetNetworkTopologyLinkLayerParams) WithTimeout(timeout time.Duration) *GetNetworkTopologyLinkLayerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network topology link layer params
func (o *GetNetworkTopologyLinkLayerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network topology link layer params
func (o *GetNetworkTopologyLinkLayerParams) WithContext(ctx context.Context) *GetNetworkTopologyLinkLayerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network topology link layer params
func (o *GetNetworkTopologyLinkLayerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network topology link layer params
func (o *GetNetworkTopologyLinkLayerParams) WithHTTPClient(client *http.Client) *GetNetworkTopologyLinkLayerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network topology link layer params
func (o *GetNetworkTopologyLinkLayerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network topology link layer params
func (o *GetNetworkTopologyLinkLayerParams) WithNetworkID(networkID string) *GetNetworkTopologyLinkLayerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network topology link layer params
func (o *GetNetworkTopologyLinkLayerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkTopologyLinkLayerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
