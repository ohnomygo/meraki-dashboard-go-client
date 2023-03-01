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

// NewUpdateNetworkWebhooksHTTPServerParams creates a new UpdateNetworkWebhooksHTTPServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkWebhooksHTTPServerParams() *UpdateNetworkWebhooksHTTPServerParams {
	return &UpdateNetworkWebhooksHTTPServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWebhooksHTTPServerParamsWithTimeout creates a new UpdateNetworkWebhooksHTTPServerParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkWebhooksHTTPServerParamsWithTimeout(timeout time.Duration) *UpdateNetworkWebhooksHTTPServerParams {
	return &UpdateNetworkWebhooksHTTPServerParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkWebhooksHTTPServerParamsWithContext creates a new UpdateNetworkWebhooksHTTPServerParams object
// with the ability to set a context for a request.
func NewUpdateNetworkWebhooksHTTPServerParamsWithContext(ctx context.Context) *UpdateNetworkWebhooksHTTPServerParams {
	return &UpdateNetworkWebhooksHTTPServerParams{
		Context: ctx,
	}
}

// NewUpdateNetworkWebhooksHTTPServerParamsWithHTTPClient creates a new UpdateNetworkWebhooksHTTPServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkWebhooksHTTPServerParamsWithHTTPClient(client *http.Client) *UpdateNetworkWebhooksHTTPServerParams {
	return &UpdateNetworkWebhooksHTTPServerParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkWebhooksHTTPServerParams contains all the parameters to send to the API endpoint

	for the update network webhooks Http server operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkWebhooksHTTPServerParams struct {

	// HTTPServerID.
	HTTPServerID string

	// NetworkID.
	NetworkID string

	// UpdateNetworkWebhooksHTTPServer.
	UpdateNetworkWebhooksHTTPServer UpdateNetworkWebhooksHTTPServerBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network webhooks Http server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWebhooksHTTPServerParams) WithDefaults() *UpdateNetworkWebhooksHTTPServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network webhooks Http server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWebhooksHTTPServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) WithTimeout(timeout time.Duration) *UpdateNetworkWebhooksHTTPServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) WithContext(ctx context.Context) *UpdateNetworkWebhooksHTTPServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) WithHTTPClient(client *http.Client) *UpdateNetworkWebhooksHTTPServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHTTPServerID adds the hTTPServerID to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) WithHTTPServerID(hTTPServerID string) *UpdateNetworkWebhooksHTTPServerParams {
	o.SetHTTPServerID(hTTPServerID)
	return o
}

// SetHTTPServerID adds the httpServerId to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) SetHTTPServerID(hTTPServerID string) {
	o.HTTPServerID = hTTPServerID
}

// WithNetworkID adds the networkID to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) WithNetworkID(networkID string) *UpdateNetworkWebhooksHTTPServerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkWebhooksHTTPServer adds the updateNetworkWebhooksHTTPServer to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) WithUpdateNetworkWebhooksHTTPServer(updateNetworkWebhooksHTTPServer UpdateNetworkWebhooksHTTPServerBody) *UpdateNetworkWebhooksHTTPServerParams {
	o.SetUpdateNetworkWebhooksHTTPServer(updateNetworkWebhooksHTTPServer)
	return o
}

// SetUpdateNetworkWebhooksHTTPServer adds the updateNetworkWebhooksHttpServer to the update network webhooks Http server params
func (o *UpdateNetworkWebhooksHTTPServerParams) SetUpdateNetworkWebhooksHTTPServer(updateNetworkWebhooksHTTPServer UpdateNetworkWebhooksHTTPServerBody) {
	o.UpdateNetworkWebhooksHTTPServer = updateNetworkWebhooksHTTPServer
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWebhooksHTTPServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param httpServerId
	if err := r.SetPathParam("httpServerId", o.HTTPServerID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkWebhooksHTTPServer); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}