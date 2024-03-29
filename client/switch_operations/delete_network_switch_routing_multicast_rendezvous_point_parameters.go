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

// NewDeleteNetworkSwitchRoutingMulticastRendezvousPointParams creates a new DeleteNetworkSwitchRoutingMulticastRendezvousPointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkSwitchRoutingMulticastRendezvousPointParams() *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &DeleteNetworkSwitchRoutingMulticastRendezvousPointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkSwitchRoutingMulticastRendezvousPointParamsWithTimeout creates a new DeleteNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkSwitchRoutingMulticastRendezvousPointParamsWithTimeout(timeout time.Duration) *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &DeleteNetworkSwitchRoutingMulticastRendezvousPointParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkSwitchRoutingMulticastRendezvousPointParamsWithContext creates a new DeleteNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a context for a request.
func NewDeleteNetworkSwitchRoutingMulticastRendezvousPointParamsWithContext(ctx context.Context) *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &DeleteNetworkSwitchRoutingMulticastRendezvousPointParams{
		Context: ctx,
	}
}

// NewDeleteNetworkSwitchRoutingMulticastRendezvousPointParamsWithHTTPClient creates a new DeleteNetworkSwitchRoutingMulticastRendezvousPointParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkSwitchRoutingMulticastRendezvousPointParamsWithHTTPClient(client *http.Client) *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	return &DeleteNetworkSwitchRoutingMulticastRendezvousPointParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkSwitchRoutingMulticastRendezvousPointParams contains all the parameters to send to the API endpoint

	for the delete network switch routing multicast rendezvous point operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkSwitchRoutingMulticastRendezvousPointParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* RendezvousPointID.

	   Rendezvous point ID
	*/
	RendezvousPointID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network switch routing multicast rendezvous point params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) WithDefaults() *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network switch routing multicast rendezvous point params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) WithTimeout(timeout time.Duration) *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) WithContext(ctx context.Context) *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) WithHTTPClient(client *http.Client) *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) WithNetworkID(networkID string) *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRendezvousPointID adds the rendezvousPointID to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) WithRendezvousPointID(rendezvousPointID string) *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams {
	o.SetRendezvousPointID(rendezvousPointID)
	return o
}

// SetRendezvousPointID adds the rendezvousPointId to the delete network switch routing multicast rendezvous point params
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) SetRendezvousPointID(rendezvousPointID string) {
	o.RendezvousPointID = rendezvousPointID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param rendezvousPointId
	if err := r.SetPathParam("rendezvousPointId", o.RendezvousPointID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
