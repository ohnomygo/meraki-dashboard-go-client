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

// NewDeleteOrganizationPolicyObjectParams creates a new DeleteOrganizationPolicyObjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationPolicyObjectParams() *DeleteOrganizationPolicyObjectParams {
	return &DeleteOrganizationPolicyObjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationPolicyObjectParamsWithTimeout creates a new DeleteOrganizationPolicyObjectParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationPolicyObjectParamsWithTimeout(timeout time.Duration) *DeleteOrganizationPolicyObjectParams {
	return &DeleteOrganizationPolicyObjectParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationPolicyObjectParamsWithContext creates a new DeleteOrganizationPolicyObjectParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationPolicyObjectParamsWithContext(ctx context.Context) *DeleteOrganizationPolicyObjectParams {
	return &DeleteOrganizationPolicyObjectParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationPolicyObjectParamsWithHTTPClient creates a new DeleteOrganizationPolicyObjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationPolicyObjectParamsWithHTTPClient(client *http.Client) *DeleteOrganizationPolicyObjectParams {
	return &DeleteOrganizationPolicyObjectParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationPolicyObjectParams contains all the parameters to send to the API endpoint

	for the delete organization policy object operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationPolicyObjectParams struct {

	// OrganizationID.
	OrganizationID string

	// PolicyObjectID.
	PolicyObjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization policy object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationPolicyObjectParams) WithDefaults() *DeleteOrganizationPolicyObjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization policy object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationPolicyObjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) WithTimeout(timeout time.Duration) *DeleteOrganizationPolicyObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) WithContext(ctx context.Context) *DeleteOrganizationPolicyObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) WithHTTPClient(client *http.Client) *DeleteOrganizationPolicyObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) WithOrganizationID(organizationID string) *DeleteOrganizationPolicyObjectParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPolicyObjectID adds the policyObjectID to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) WithPolicyObjectID(policyObjectID string) *DeleteOrganizationPolicyObjectParams {
	o.SetPolicyObjectID(policyObjectID)
	return o
}

// SetPolicyObjectID adds the policyObjectId to the delete organization policy object params
func (o *DeleteOrganizationPolicyObjectParams) SetPolicyObjectID(policyObjectID string) {
	o.PolicyObjectID = policyObjectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationPolicyObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
