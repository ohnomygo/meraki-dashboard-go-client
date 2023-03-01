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

// NewGetOrganizationPolicyObjectParams creates a new GetOrganizationPolicyObjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationPolicyObjectParams() *GetOrganizationPolicyObjectParams {
	return &GetOrganizationPolicyObjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationPolicyObjectParamsWithTimeout creates a new GetOrganizationPolicyObjectParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationPolicyObjectParamsWithTimeout(timeout time.Duration) *GetOrganizationPolicyObjectParams {
	return &GetOrganizationPolicyObjectParams{
		timeout: timeout,
	}
}

// NewGetOrganizationPolicyObjectParamsWithContext creates a new GetOrganizationPolicyObjectParams object
// with the ability to set a context for a request.
func NewGetOrganizationPolicyObjectParamsWithContext(ctx context.Context) *GetOrganizationPolicyObjectParams {
	return &GetOrganizationPolicyObjectParams{
		Context: ctx,
	}
}

// NewGetOrganizationPolicyObjectParamsWithHTTPClient creates a new GetOrganizationPolicyObjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationPolicyObjectParamsWithHTTPClient(client *http.Client) *GetOrganizationPolicyObjectParams {
	return &GetOrganizationPolicyObjectParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationPolicyObjectParams contains all the parameters to send to the API endpoint

	for the get organization policy object operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationPolicyObjectParams struct {

	// OrganizationID.
	OrganizationID string

	// PolicyObjectID.
	PolicyObjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization policy object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationPolicyObjectParams) WithDefaults() *GetOrganizationPolicyObjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization policy object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationPolicyObjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) WithTimeout(timeout time.Duration) *GetOrganizationPolicyObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) WithContext(ctx context.Context) *GetOrganizationPolicyObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) WithHTTPClient(client *http.Client) *GetOrganizationPolicyObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) WithOrganizationID(organizationID string) *GetOrganizationPolicyObjectParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPolicyObjectID adds the policyObjectID to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) WithPolicyObjectID(policyObjectID string) *GetOrganizationPolicyObjectParams {
	o.SetPolicyObjectID(policyObjectID)
	return o
}

// SetPolicyObjectID adds the policyObjectId to the get organization policy object params
func (o *GetOrganizationPolicyObjectParams) SetPolicyObjectID(policyObjectID string) {
	o.PolicyObjectID = policyObjectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationPolicyObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	// path param policyObjectId
	if err := r.SetPathParam("policyObjectId", o.PolicyObjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
