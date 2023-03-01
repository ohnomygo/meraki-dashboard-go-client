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

// NewCloneOrganizationParams creates a new CloneOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloneOrganizationParams() *CloneOrganizationParams {
	return &CloneOrganizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloneOrganizationParamsWithTimeout creates a new CloneOrganizationParams object
// with the ability to set a timeout on a request.
func NewCloneOrganizationParamsWithTimeout(timeout time.Duration) *CloneOrganizationParams {
	return &CloneOrganizationParams{
		timeout: timeout,
	}
}

// NewCloneOrganizationParamsWithContext creates a new CloneOrganizationParams object
// with the ability to set a context for a request.
func NewCloneOrganizationParamsWithContext(ctx context.Context) *CloneOrganizationParams {
	return &CloneOrganizationParams{
		Context: ctx,
	}
}

// NewCloneOrganizationParamsWithHTTPClient creates a new CloneOrganizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloneOrganizationParamsWithHTTPClient(client *http.Client) *CloneOrganizationParams {
	return &CloneOrganizationParams{
		HTTPClient: client,
	}
}

/*
CloneOrganizationParams contains all the parameters to send to the API endpoint

	for the clone organization operation.

	Typically these are written to a http.Request.
*/
type CloneOrganizationParams struct {

	// CloneOrganization.
	CloneOrganization CloneOrganizationBody

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clone organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneOrganizationParams) WithDefaults() *CloneOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clone organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the clone organization params
func (o *CloneOrganizationParams) WithTimeout(timeout time.Duration) *CloneOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone organization params
func (o *CloneOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone organization params
func (o *CloneOrganizationParams) WithContext(ctx context.Context) *CloneOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone organization params
func (o *CloneOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone organization params
func (o *CloneOrganizationParams) WithHTTPClient(client *http.Client) *CloneOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone organization params
func (o *CloneOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloneOrganization adds the cloneOrganization to the clone organization params
func (o *CloneOrganizationParams) WithCloneOrganization(cloneOrganization CloneOrganizationBody) *CloneOrganizationParams {
	o.SetCloneOrganization(cloneOrganization)
	return o
}

// SetCloneOrganization adds the cloneOrganization to the clone organization params
func (o *CloneOrganizationParams) SetCloneOrganization(cloneOrganization CloneOrganizationBody) {
	o.CloneOrganization = cloneOrganization
}

// WithOrganizationID adds the organizationID to the clone organization params
func (o *CloneOrganizationParams) WithOrganizationID(organizationID string) *CloneOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the clone organization params
func (o *CloneOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CloneOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CloneOrganization); err != nil {
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