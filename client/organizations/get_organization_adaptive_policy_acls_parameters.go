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

// NewGetOrganizationAdaptivePolicyAclsParams creates a new GetOrganizationAdaptivePolicyAclsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationAdaptivePolicyAclsParams() *GetOrganizationAdaptivePolicyAclsParams {
	return &GetOrganizationAdaptivePolicyAclsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationAdaptivePolicyAclsParamsWithTimeout creates a new GetOrganizationAdaptivePolicyAclsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationAdaptivePolicyAclsParamsWithTimeout(timeout time.Duration) *GetOrganizationAdaptivePolicyAclsParams {
	return &GetOrganizationAdaptivePolicyAclsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationAdaptivePolicyAclsParamsWithContext creates a new GetOrganizationAdaptivePolicyAclsParams object
// with the ability to set a context for a request.
func NewGetOrganizationAdaptivePolicyAclsParamsWithContext(ctx context.Context) *GetOrganizationAdaptivePolicyAclsParams {
	return &GetOrganizationAdaptivePolicyAclsParams{
		Context: ctx,
	}
}

// NewGetOrganizationAdaptivePolicyAclsParamsWithHTTPClient creates a new GetOrganizationAdaptivePolicyAclsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationAdaptivePolicyAclsParamsWithHTTPClient(client *http.Client) *GetOrganizationAdaptivePolicyAclsParams {
	return &GetOrganizationAdaptivePolicyAclsParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationAdaptivePolicyAclsParams contains all the parameters to send to the API endpoint

	for the get organization adaptive policy acls operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationAdaptivePolicyAclsParams struct {

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization adaptive policy acls params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAdaptivePolicyAclsParams) WithDefaults() *GetOrganizationAdaptivePolicyAclsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization adaptive policy acls params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAdaptivePolicyAclsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization adaptive policy acls params
func (o *GetOrganizationAdaptivePolicyAclsParams) WithTimeout(timeout time.Duration) *GetOrganizationAdaptivePolicyAclsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization adaptive policy acls params
func (o *GetOrganizationAdaptivePolicyAclsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization adaptive policy acls params
func (o *GetOrganizationAdaptivePolicyAclsParams) WithContext(ctx context.Context) *GetOrganizationAdaptivePolicyAclsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization adaptive policy acls params
func (o *GetOrganizationAdaptivePolicyAclsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization adaptive policy acls params
func (o *GetOrganizationAdaptivePolicyAclsParams) WithHTTPClient(client *http.Client) *GetOrganizationAdaptivePolicyAclsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization adaptive policy acls params
func (o *GetOrganizationAdaptivePolicyAclsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization adaptive policy acls params
func (o *GetOrganizationAdaptivePolicyAclsParams) WithOrganizationID(organizationID string) *GetOrganizationAdaptivePolicyAclsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization adaptive policy acls params
func (o *GetOrganizationAdaptivePolicyAclsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationAdaptivePolicyAclsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
