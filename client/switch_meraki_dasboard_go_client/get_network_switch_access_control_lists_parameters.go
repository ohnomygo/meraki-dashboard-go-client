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

// NewGetNetworkSwitchAccessControlListsParams creates a new GetNetworkSwitchAccessControlListsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSwitchAccessControlListsParams() *GetNetworkSwitchAccessControlListsParams {
	return &GetNetworkSwitchAccessControlListsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchAccessControlListsParamsWithTimeout creates a new GetNetworkSwitchAccessControlListsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSwitchAccessControlListsParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchAccessControlListsParams {
	return &GetNetworkSwitchAccessControlListsParams{
		timeout: timeout,
	}
}

// NewGetNetworkSwitchAccessControlListsParamsWithContext creates a new GetNetworkSwitchAccessControlListsParams object
// with the ability to set a context for a request.
func NewGetNetworkSwitchAccessControlListsParamsWithContext(ctx context.Context) *GetNetworkSwitchAccessControlListsParams {
	return &GetNetworkSwitchAccessControlListsParams{
		Context: ctx,
	}
}

// NewGetNetworkSwitchAccessControlListsParamsWithHTTPClient creates a new GetNetworkSwitchAccessControlListsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSwitchAccessControlListsParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchAccessControlListsParams {
	return &GetNetworkSwitchAccessControlListsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSwitchAccessControlListsParams contains all the parameters to send to the API endpoint

	for the get network switch access control lists operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSwitchAccessControlListsParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network switch access control lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchAccessControlListsParams) WithDefaults() *GetNetworkSwitchAccessControlListsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network switch access control lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchAccessControlListsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network switch access control lists params
func (o *GetNetworkSwitchAccessControlListsParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchAccessControlListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch access control lists params
func (o *GetNetworkSwitchAccessControlListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch access control lists params
func (o *GetNetworkSwitchAccessControlListsParams) WithContext(ctx context.Context) *GetNetworkSwitchAccessControlListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch access control lists params
func (o *GetNetworkSwitchAccessControlListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch access control lists params
func (o *GetNetworkSwitchAccessControlListsParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchAccessControlListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch access control lists params
func (o *GetNetworkSwitchAccessControlListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network switch access control lists params
func (o *GetNetworkSwitchAccessControlListsParams) WithNetworkID(networkID string) *GetNetworkSwitchAccessControlListsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch access control lists params
func (o *GetNetworkSwitchAccessControlListsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchAccessControlListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
