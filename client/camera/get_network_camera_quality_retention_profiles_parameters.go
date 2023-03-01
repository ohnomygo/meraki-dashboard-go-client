// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// NewGetNetworkCameraQualityRetentionProfilesParams creates a new GetNetworkCameraQualityRetentionProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkCameraQualityRetentionProfilesParams() *GetNetworkCameraQualityRetentionProfilesParams {
	return &GetNetworkCameraQualityRetentionProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkCameraQualityRetentionProfilesParamsWithTimeout creates a new GetNetworkCameraQualityRetentionProfilesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkCameraQualityRetentionProfilesParamsWithTimeout(timeout time.Duration) *GetNetworkCameraQualityRetentionProfilesParams {
	return &GetNetworkCameraQualityRetentionProfilesParams{
		timeout: timeout,
	}
}

// NewGetNetworkCameraQualityRetentionProfilesParamsWithContext creates a new GetNetworkCameraQualityRetentionProfilesParams object
// with the ability to set a context for a request.
func NewGetNetworkCameraQualityRetentionProfilesParamsWithContext(ctx context.Context) *GetNetworkCameraQualityRetentionProfilesParams {
	return &GetNetworkCameraQualityRetentionProfilesParams{
		Context: ctx,
	}
}

// NewGetNetworkCameraQualityRetentionProfilesParamsWithHTTPClient creates a new GetNetworkCameraQualityRetentionProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkCameraQualityRetentionProfilesParamsWithHTTPClient(client *http.Client) *GetNetworkCameraQualityRetentionProfilesParams {
	return &GetNetworkCameraQualityRetentionProfilesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkCameraQualityRetentionProfilesParams contains all the parameters to send to the API endpoint

	for the get network camera quality retention profiles operation.

	Typically these are written to a http.Request.
*/
type GetNetworkCameraQualityRetentionProfilesParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network camera quality retention profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCameraQualityRetentionProfilesParams) WithDefaults() *GetNetworkCameraQualityRetentionProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network camera quality retention profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCameraQualityRetentionProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network camera quality retention profiles params
func (o *GetNetworkCameraQualityRetentionProfilesParams) WithTimeout(timeout time.Duration) *GetNetworkCameraQualityRetentionProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network camera quality retention profiles params
func (o *GetNetworkCameraQualityRetentionProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network camera quality retention profiles params
func (o *GetNetworkCameraQualityRetentionProfilesParams) WithContext(ctx context.Context) *GetNetworkCameraQualityRetentionProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network camera quality retention profiles params
func (o *GetNetworkCameraQualityRetentionProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network camera quality retention profiles params
func (o *GetNetworkCameraQualityRetentionProfilesParams) WithHTTPClient(client *http.Client) *GetNetworkCameraQualityRetentionProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network camera quality retention profiles params
func (o *GetNetworkCameraQualityRetentionProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network camera quality retention profiles params
func (o *GetNetworkCameraQualityRetentionProfilesParams) WithNetworkID(networkID string) *GetNetworkCameraQualityRetentionProfilesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network camera quality retention profiles params
func (o *GetNetworkCameraQualityRetentionProfilesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkCameraQualityRetentionProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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