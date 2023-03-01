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

// NewUpdateOrganizationLoginSecurityParams creates a new UpdateOrganizationLoginSecurityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationLoginSecurityParams() *UpdateOrganizationLoginSecurityParams {
	return &UpdateOrganizationLoginSecurityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationLoginSecurityParamsWithTimeout creates a new UpdateOrganizationLoginSecurityParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationLoginSecurityParamsWithTimeout(timeout time.Duration) *UpdateOrganizationLoginSecurityParams {
	return &UpdateOrganizationLoginSecurityParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationLoginSecurityParamsWithContext creates a new UpdateOrganizationLoginSecurityParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationLoginSecurityParamsWithContext(ctx context.Context) *UpdateOrganizationLoginSecurityParams {
	return &UpdateOrganizationLoginSecurityParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationLoginSecurityParamsWithHTTPClient creates a new UpdateOrganizationLoginSecurityParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationLoginSecurityParamsWithHTTPClient(client *http.Client) *UpdateOrganizationLoginSecurityParams {
	return &UpdateOrganizationLoginSecurityParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationLoginSecurityParams contains all the parameters to send to the API endpoint

	for the update organization login security operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationLoginSecurityParams struct {

	// OrganizationID.
	OrganizationID string

	// UpdateOrganizationLoginSecurity.
	UpdateOrganizationLoginSecurity UpdateOrganizationLoginSecurityBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization login security params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationLoginSecurityParams) WithDefaults() *UpdateOrganizationLoginSecurityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization login security params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationLoginSecurityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) WithTimeout(timeout time.Duration) *UpdateOrganizationLoginSecurityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) WithContext(ctx context.Context) *UpdateOrganizationLoginSecurityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) WithHTTPClient(client *http.Client) *UpdateOrganizationLoginSecurityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) WithOrganizationID(organizationID string) *UpdateOrganizationLoginSecurityParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationLoginSecurity adds the updateOrganizationLoginSecurity to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) WithUpdateOrganizationLoginSecurity(updateOrganizationLoginSecurity UpdateOrganizationLoginSecurityBody) *UpdateOrganizationLoginSecurityParams {
	o.SetUpdateOrganizationLoginSecurity(updateOrganizationLoginSecurity)
	return o
}

// SetUpdateOrganizationLoginSecurity adds the updateOrganizationLoginSecurity to the update organization login security params
func (o *UpdateOrganizationLoginSecurityParams) SetUpdateOrganizationLoginSecurity(updateOrganizationLoginSecurity UpdateOrganizationLoginSecurityBody) {
	o.UpdateOrganizationLoginSecurity = updateOrganizationLoginSecurity
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationLoginSecurityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationLoginSecurity); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}