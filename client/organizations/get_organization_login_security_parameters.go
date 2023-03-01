// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrganizationLoginSecurityParams creates a new GetOrganizationLoginSecurityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationLoginSecurityParams() *GetOrganizationLoginSecurityParams {
	return &GetOrganizationLoginSecurityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationLoginSecurityParamsWithTimeout creates a new GetOrganizationLoginSecurityParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationLoginSecurityParamsWithTimeout(timeout time.Duration) *GetOrganizationLoginSecurityParams {
	return &GetOrganizationLoginSecurityParams{
		timeout: timeout,
	}
}

// NewGetOrganizationLoginSecurityParamsWithContext creates a new GetOrganizationLoginSecurityParams object
// with the ability to set a context for a request.
func NewGetOrganizationLoginSecurityParamsWithContext(ctx context.Context) *GetOrganizationLoginSecurityParams {
	return &GetOrganizationLoginSecurityParams{
		Context: ctx,
	}
}

// NewGetOrganizationLoginSecurityParamsWithHTTPClient creates a new GetOrganizationLoginSecurityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationLoginSecurityParamsWithHTTPClient(client *http.Client) *GetOrganizationLoginSecurityParams {
	return &GetOrganizationLoginSecurityParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationLoginSecurityParams contains all the parameters to send to the API endpoint

	for the get organization login security operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationLoginSecurityParams struct {

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization login security params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationLoginSecurityParams) WithDefaults() *GetOrganizationLoginSecurityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization login security params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationLoginSecurityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization login security params
func (o *GetOrganizationLoginSecurityParams) WithTimeout(timeout time.Duration) *GetOrganizationLoginSecurityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization login security params
func (o *GetOrganizationLoginSecurityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization login security params
func (o *GetOrganizationLoginSecurityParams) WithContext(ctx context.Context) *GetOrganizationLoginSecurityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization login security params
func (o *GetOrganizationLoginSecurityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization login security params
func (o *GetOrganizationLoginSecurityParams) WithHTTPClient(client *http.Client) *GetOrganizationLoginSecurityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization login security params
func (o *GetOrganizationLoginSecurityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization login security params
func (o *GetOrganizationLoginSecurityParams) WithOrganizationID(organizationID string) *GetOrganizationLoginSecurityParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization login security params
func (o *GetOrganizationLoginSecurityParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationLoginSecurityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}