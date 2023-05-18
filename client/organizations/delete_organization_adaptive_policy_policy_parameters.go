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

// NewDeleteOrganizationAdaptivePolicyPolicyParams creates a new DeleteOrganizationAdaptivePolicyPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationAdaptivePolicyPolicyParams() *DeleteOrganizationAdaptivePolicyPolicyParams {
	return &DeleteOrganizationAdaptivePolicyPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationAdaptivePolicyPolicyParamsWithTimeout creates a new DeleteOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationAdaptivePolicyPolicyParamsWithTimeout(timeout time.Duration) *DeleteOrganizationAdaptivePolicyPolicyParams {
	return &DeleteOrganizationAdaptivePolicyPolicyParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationAdaptivePolicyPolicyParamsWithContext creates a new DeleteOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationAdaptivePolicyPolicyParamsWithContext(ctx context.Context) *DeleteOrganizationAdaptivePolicyPolicyParams {
	return &DeleteOrganizationAdaptivePolicyPolicyParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationAdaptivePolicyPolicyParamsWithHTTPClient creates a new DeleteOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationAdaptivePolicyPolicyParamsWithHTTPClient(client *http.Client) *DeleteOrganizationAdaptivePolicyPolicyParams {
	return &DeleteOrganizationAdaptivePolicyPolicyParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationAdaptivePolicyPolicyParams contains all the parameters to send to the API endpoint

	for the delete organization adaptive policy policy operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationAdaptivePolicyPolicyParams struct {

	/* ID.

	   Id
	*/
	ID string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization adaptive policy policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) WithDefaults() *DeleteOrganizationAdaptivePolicyPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization adaptive policy policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) WithTimeout(timeout time.Duration) *DeleteOrganizationAdaptivePolicyPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) WithContext(ctx context.Context) *DeleteOrganizationAdaptivePolicyPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) WithHTTPClient(client *http.Client) *DeleteOrganizationAdaptivePolicyPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) WithID(id string) *DeleteOrganizationAdaptivePolicyPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) SetID(id string) {
	o.ID = id
}

// WithOrganizationID adds the organizationID to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) WithOrganizationID(organizationID string) *DeleteOrganizationAdaptivePolicyPolicyParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete organization adaptive policy policy params
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationAdaptivePolicyPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
