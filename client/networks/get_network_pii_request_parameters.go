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

// NewGetNetworkPiiRequestParams creates a new GetNetworkPiiRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkPiiRequestParams() *GetNetworkPiiRequestParams {
	return &GetNetworkPiiRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkPiiRequestParamsWithTimeout creates a new GetNetworkPiiRequestParams object
// with the ability to set a timeout on a request.
func NewGetNetworkPiiRequestParamsWithTimeout(timeout time.Duration) *GetNetworkPiiRequestParams {
	return &GetNetworkPiiRequestParams{
		timeout: timeout,
	}
}

// NewGetNetworkPiiRequestParamsWithContext creates a new GetNetworkPiiRequestParams object
// with the ability to set a context for a request.
func NewGetNetworkPiiRequestParamsWithContext(ctx context.Context) *GetNetworkPiiRequestParams {
	return &GetNetworkPiiRequestParams{
		Context: ctx,
	}
}

// NewGetNetworkPiiRequestParamsWithHTTPClient creates a new GetNetworkPiiRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkPiiRequestParamsWithHTTPClient(client *http.Client) *GetNetworkPiiRequestParams {
	return &GetNetworkPiiRequestParams{
		HTTPClient: client,
	}
}

/*
GetNetworkPiiRequestParams contains all the parameters to send to the API endpoint

	for the get network pii request operation.

	Typically these are written to a http.Request.
*/
type GetNetworkPiiRequestParams struct {

	// NetworkID.
	NetworkID string

	// RequestID.
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network pii request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkPiiRequestParams) WithDefaults() *GetNetworkPiiRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network pii request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkPiiRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network pii request params
func (o *GetNetworkPiiRequestParams) WithTimeout(timeout time.Duration) *GetNetworkPiiRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network pii request params
func (o *GetNetworkPiiRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network pii request params
func (o *GetNetworkPiiRequestParams) WithContext(ctx context.Context) *GetNetworkPiiRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network pii request params
func (o *GetNetworkPiiRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network pii request params
func (o *GetNetworkPiiRequestParams) WithHTTPClient(client *http.Client) *GetNetworkPiiRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network pii request params
func (o *GetNetworkPiiRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network pii request params
func (o *GetNetworkPiiRequestParams) WithNetworkID(networkID string) *GetNetworkPiiRequestParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network pii request params
func (o *GetNetworkPiiRequestParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRequestID adds the requestID to the get network pii request params
func (o *GetNetworkPiiRequestParams) WithRequestID(requestID string) *GetNetworkPiiRequestParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get network pii request params
func (o *GetNetworkPiiRequestParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkPiiRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
