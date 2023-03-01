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

// NewDeleteNetworkSwitchStackRoutingStaticRouteParams creates a new DeleteNetworkSwitchStackRoutingStaticRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkSwitchStackRoutingStaticRouteParams() *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	return &DeleteNetworkSwitchStackRoutingStaticRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkSwitchStackRoutingStaticRouteParamsWithTimeout creates a new DeleteNetworkSwitchStackRoutingStaticRouteParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkSwitchStackRoutingStaticRouteParamsWithTimeout(timeout time.Duration) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	return &DeleteNetworkSwitchStackRoutingStaticRouteParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkSwitchStackRoutingStaticRouteParamsWithContext creates a new DeleteNetworkSwitchStackRoutingStaticRouteParams object
// with the ability to set a context for a request.
func NewDeleteNetworkSwitchStackRoutingStaticRouteParamsWithContext(ctx context.Context) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	return &DeleteNetworkSwitchStackRoutingStaticRouteParams{
		Context: ctx,
	}
}

// NewDeleteNetworkSwitchStackRoutingStaticRouteParamsWithHTTPClient creates a new DeleteNetworkSwitchStackRoutingStaticRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkSwitchStackRoutingStaticRouteParamsWithHTTPClient(client *http.Client) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	return &DeleteNetworkSwitchStackRoutingStaticRouteParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkSwitchStackRoutingStaticRouteParams contains all the parameters to send to the API endpoint

	for the delete network switch stack routing static route operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkSwitchStackRoutingStaticRouteParams struct {

	// NetworkID.
	NetworkID string

	// StaticRouteID.
	StaticRouteID string

	// SwitchStackID.
	SwitchStackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network switch stack routing static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) WithDefaults() *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network switch stack routing static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) WithTimeout(timeout time.Duration) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) WithContext(ctx context.Context) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) WithHTTPClient(client *http.Client) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) WithNetworkID(networkID string) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithStaticRouteID adds the staticRouteID to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) WithStaticRouteID(staticRouteID string) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	o.SetStaticRouteID(staticRouteID)
	return o
}

// SetStaticRouteID adds the staticRouteId to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) SetStaticRouteID(staticRouteID string) {
	o.StaticRouteID = staticRouteID
}

// WithSwitchStackID adds the switchStackID to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) WithSwitchStackID(switchStackID string) *DeleteNetworkSwitchStackRoutingStaticRouteParams {
	o.SetSwitchStackID(switchStackID)
	return o
}

// SetSwitchStackID adds the switchStackId to the delete network switch stack routing static route params
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) SetSwitchStackID(switchStackID string) {
	o.SwitchStackID = switchStackID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkSwitchStackRoutingStaticRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param staticRouteId
	if err := r.SetPathParam("staticRouteId", o.StaticRouteID); err != nil {
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
