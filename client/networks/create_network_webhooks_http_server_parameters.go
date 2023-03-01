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

// NewCreateNetworkWebhooksHTTPServerParams creates a new CreateNetworkWebhooksHTTPServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkWebhooksHTTPServerParams() *CreateNetworkWebhooksHTTPServerParams {
	return &CreateNetworkWebhooksHTTPServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkWebhooksHTTPServerParamsWithTimeout creates a new CreateNetworkWebhooksHTTPServerParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkWebhooksHTTPServerParamsWithTimeout(timeout time.Duration) *CreateNetworkWebhooksHTTPServerParams {
	return &CreateNetworkWebhooksHTTPServerParams{
		timeout: timeout,
	}
}

// NewCreateNetworkWebhooksHTTPServerParamsWithContext creates a new CreateNetworkWebhooksHTTPServerParams object
// with the ability to set a context for a request.
func NewCreateNetworkWebhooksHTTPServerParamsWithContext(ctx context.Context) *CreateNetworkWebhooksHTTPServerParams {
	return &CreateNetworkWebhooksHTTPServerParams{
		Context: ctx,
	}
}

// NewCreateNetworkWebhooksHTTPServerParamsWithHTTPClient creates a new CreateNetworkWebhooksHTTPServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkWebhooksHTTPServerParamsWithHTTPClient(client *http.Client) *CreateNetworkWebhooksHTTPServerParams {
	return &CreateNetworkWebhooksHTTPServerParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkWebhooksHTTPServerParams contains all the parameters to send to the API endpoint

	for the create network webhooks Http server operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkWebhooksHTTPServerParams struct {

	// CreateNetworkWebhooksHTTPServer.
	CreateNetworkWebhooksHTTPServer CreateNetworkWebhooksHTTPServerBody

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network webhooks Http server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkWebhooksHTTPServerParams) WithDefaults() *CreateNetworkWebhooksHTTPServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network webhooks Http server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkWebhooksHTTPServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) WithTimeout(timeout time.Duration) *CreateNetworkWebhooksHTTPServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) WithContext(ctx context.Context) *CreateNetworkWebhooksHTTPServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) WithHTTPClient(client *http.Client) *CreateNetworkWebhooksHTTPServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkWebhooksHTTPServer adds the createNetworkWebhooksHTTPServer to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) WithCreateNetworkWebhooksHTTPServer(createNetworkWebhooksHTTPServer CreateNetworkWebhooksHTTPServerBody) *CreateNetworkWebhooksHTTPServerParams {
	o.SetCreateNetworkWebhooksHTTPServer(createNetworkWebhooksHTTPServer)
	return o
}

// SetCreateNetworkWebhooksHTTPServer adds the createNetworkWebhooksHttpServer to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) SetCreateNetworkWebhooksHTTPServer(createNetworkWebhooksHTTPServer CreateNetworkWebhooksHTTPServerBody) {
	o.CreateNetworkWebhooksHTTPServer = createNetworkWebhooksHTTPServer
}

// WithNetworkID adds the networkID to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) WithNetworkID(networkID string) *CreateNetworkWebhooksHTTPServerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network webhooks Http server params
func (o *CreateNetworkWebhooksHTTPServerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkWebhooksHTTPServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkWebhooksHTTPServer); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
