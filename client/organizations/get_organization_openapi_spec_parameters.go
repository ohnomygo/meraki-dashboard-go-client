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

// NewGetOrganizationOpenapiSpecParams creates a new GetOrganizationOpenapiSpecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationOpenapiSpecParams() *GetOrganizationOpenapiSpecParams {
	return &GetOrganizationOpenapiSpecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationOpenapiSpecParamsWithTimeout creates a new GetOrganizationOpenapiSpecParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationOpenapiSpecParamsWithTimeout(timeout time.Duration) *GetOrganizationOpenapiSpecParams {
	return &GetOrganizationOpenapiSpecParams{
		timeout: timeout,
	}
}

// NewGetOrganizationOpenapiSpecParamsWithContext creates a new GetOrganizationOpenapiSpecParams object
// with the ability to set a context for a request.
func NewGetOrganizationOpenapiSpecParamsWithContext(ctx context.Context) *GetOrganizationOpenapiSpecParams {
	return &GetOrganizationOpenapiSpecParams{
		Context: ctx,
	}
}

// NewGetOrganizationOpenapiSpecParamsWithHTTPClient creates a new GetOrganizationOpenapiSpecParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationOpenapiSpecParamsWithHTTPClient(client *http.Client) *GetOrganizationOpenapiSpecParams {
	return &GetOrganizationOpenapiSpecParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationOpenapiSpecParams contains all the parameters to send to the API endpoint

	for the get organization openapi spec operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationOpenapiSpecParams struct {

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization openapi spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationOpenapiSpecParams) WithDefaults() *GetOrganizationOpenapiSpecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization openapi spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationOpenapiSpecParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization openapi spec params
func (o *GetOrganizationOpenapiSpecParams) WithTimeout(timeout time.Duration) *GetOrganizationOpenapiSpecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization openapi spec params
func (o *GetOrganizationOpenapiSpecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization openapi spec params
func (o *GetOrganizationOpenapiSpecParams) WithContext(ctx context.Context) *GetOrganizationOpenapiSpecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization openapi spec params
func (o *GetOrganizationOpenapiSpecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization openapi spec params
func (o *GetOrganizationOpenapiSpecParams) WithHTTPClient(client *http.Client) *GetOrganizationOpenapiSpecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization openapi spec params
func (o *GetOrganizationOpenapiSpecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization openapi spec params
func (o *GetOrganizationOpenapiSpecParams) WithOrganizationID(organizationID string) *GetOrganizationOpenapiSpecParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization openapi spec params
func (o *GetOrganizationOpenapiSpecParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationOpenapiSpecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
