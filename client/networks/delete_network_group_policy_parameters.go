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

// NewDeleteNetworkGroupPolicyParams creates a new DeleteNetworkGroupPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkGroupPolicyParams() *DeleteNetworkGroupPolicyParams {
	return &DeleteNetworkGroupPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkGroupPolicyParamsWithTimeout creates a new DeleteNetworkGroupPolicyParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkGroupPolicyParamsWithTimeout(timeout time.Duration) *DeleteNetworkGroupPolicyParams {
	return &DeleteNetworkGroupPolicyParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkGroupPolicyParamsWithContext creates a new DeleteNetworkGroupPolicyParams object
// with the ability to set a context for a request.
func NewDeleteNetworkGroupPolicyParamsWithContext(ctx context.Context) *DeleteNetworkGroupPolicyParams {
	return &DeleteNetworkGroupPolicyParams{
		Context: ctx,
	}
}

// NewDeleteNetworkGroupPolicyParamsWithHTTPClient creates a new DeleteNetworkGroupPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkGroupPolicyParamsWithHTTPClient(client *http.Client) *DeleteNetworkGroupPolicyParams {
	return &DeleteNetworkGroupPolicyParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkGroupPolicyParams contains all the parameters to send to the API endpoint

	for the delete network group policy operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkGroupPolicyParams struct {

	// GroupPolicyID.
	GroupPolicyID string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network group policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkGroupPolicyParams) WithDefaults() *DeleteNetworkGroupPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network group policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkGroupPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) WithTimeout(timeout time.Duration) *DeleteNetworkGroupPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) WithContext(ctx context.Context) *DeleteNetworkGroupPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) WithHTTPClient(client *http.Client) *DeleteNetworkGroupPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupPolicyID adds the groupPolicyID to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) WithGroupPolicyID(groupPolicyID string) *DeleteNetworkGroupPolicyParams {
	o.SetGroupPolicyID(groupPolicyID)
	return o
}

// SetGroupPolicyID adds the groupPolicyId to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) SetGroupPolicyID(groupPolicyID string) {
	o.GroupPolicyID = groupPolicyID
}

// WithNetworkID adds the networkID to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) WithNetworkID(networkID string) *DeleteNetworkGroupPolicyParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network group policy params
func (o *DeleteNetworkGroupPolicyParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkGroupPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupPolicyId
	if err := r.SetPathParam("groupPolicyId", o.GroupPolicyID); err != nil {
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
