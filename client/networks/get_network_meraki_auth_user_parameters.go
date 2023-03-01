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

// NewGetNetworkMerakiAuthUserParams creates a new GetNetworkMerakiAuthUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkMerakiAuthUserParams() *GetNetworkMerakiAuthUserParams {
	return &GetNetworkMerakiAuthUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkMerakiAuthUserParamsWithTimeout creates a new GetNetworkMerakiAuthUserParams object
// with the ability to set a timeout on a request.
func NewGetNetworkMerakiAuthUserParamsWithTimeout(timeout time.Duration) *GetNetworkMerakiAuthUserParams {
	return &GetNetworkMerakiAuthUserParams{
		timeout: timeout,
	}
}

// NewGetNetworkMerakiAuthUserParamsWithContext creates a new GetNetworkMerakiAuthUserParams object
// with the ability to set a context for a request.
func NewGetNetworkMerakiAuthUserParamsWithContext(ctx context.Context) *GetNetworkMerakiAuthUserParams {
	return &GetNetworkMerakiAuthUserParams{
		Context: ctx,
	}
}

// NewGetNetworkMerakiAuthUserParamsWithHTTPClient creates a new GetNetworkMerakiAuthUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkMerakiAuthUserParamsWithHTTPClient(client *http.Client) *GetNetworkMerakiAuthUserParams {
	return &GetNetworkMerakiAuthUserParams{
		HTTPClient: client,
	}
}

/*
GetNetworkMerakiAuthUserParams contains all the parameters to send to the API endpoint

	for the get network meraki auth user operation.

	Typically these are written to a http.Request.
*/
type GetNetworkMerakiAuthUserParams struct {

	// MerakiAuthUserID.
	MerakiAuthUserID string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network meraki auth user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkMerakiAuthUserParams) WithDefaults() *GetNetworkMerakiAuthUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network meraki auth user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkMerakiAuthUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) WithTimeout(timeout time.Duration) *GetNetworkMerakiAuthUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) WithContext(ctx context.Context) *GetNetworkMerakiAuthUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) WithHTTPClient(client *http.Client) *GetNetworkMerakiAuthUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMerakiAuthUserID adds the merakiAuthUserID to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) WithMerakiAuthUserID(merakiAuthUserID string) *GetNetworkMerakiAuthUserParams {
	o.SetMerakiAuthUserID(merakiAuthUserID)
	return o
}

// SetMerakiAuthUserID adds the merakiAuthUserId to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) SetMerakiAuthUserID(merakiAuthUserID string) {
	o.MerakiAuthUserID = merakiAuthUserID
}

// WithNetworkID adds the networkID to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) WithNetworkID(networkID string) *GetNetworkMerakiAuthUserParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network meraki auth user params
func (o *GetNetworkMerakiAuthUserParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkMerakiAuthUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param merakiAuthUserId
	if err := r.SetPathParam("merakiAuthUserId", o.MerakiAuthUserID); err != nil {
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
