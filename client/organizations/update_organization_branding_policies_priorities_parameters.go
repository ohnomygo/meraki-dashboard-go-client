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

// NewUpdateOrganizationBrandingPoliciesPrioritiesParams creates a new UpdateOrganizationBrandingPoliciesPrioritiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationBrandingPoliciesPrioritiesParams() *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	return &UpdateOrganizationBrandingPoliciesPrioritiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationBrandingPoliciesPrioritiesParamsWithTimeout creates a new UpdateOrganizationBrandingPoliciesPrioritiesParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationBrandingPoliciesPrioritiesParamsWithTimeout(timeout time.Duration) *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	return &UpdateOrganizationBrandingPoliciesPrioritiesParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationBrandingPoliciesPrioritiesParamsWithContext creates a new UpdateOrganizationBrandingPoliciesPrioritiesParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationBrandingPoliciesPrioritiesParamsWithContext(ctx context.Context) *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	return &UpdateOrganizationBrandingPoliciesPrioritiesParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationBrandingPoliciesPrioritiesParamsWithHTTPClient creates a new UpdateOrganizationBrandingPoliciesPrioritiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationBrandingPoliciesPrioritiesParamsWithHTTPClient(client *http.Client) *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	return &UpdateOrganizationBrandingPoliciesPrioritiesParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationBrandingPoliciesPrioritiesParams contains all the parameters to send to the API endpoint

	for the update organization branding policies priorities operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationBrandingPoliciesPrioritiesParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	// UpdateOrganizationBrandingPoliciesPriorities.
	UpdateOrganizationBrandingPoliciesPriorities UpdateOrganizationBrandingPoliciesPrioritiesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization branding policies priorities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) WithDefaults() *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization branding policies priorities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) WithTimeout(timeout time.Duration) *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) WithContext(ctx context.Context) *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) WithHTTPClient(client *http.Client) *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) WithOrganizationID(organizationID string) *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationBrandingPoliciesPriorities adds the updateOrganizationBrandingPoliciesPriorities to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) WithUpdateOrganizationBrandingPoliciesPriorities(updateOrganizationBrandingPoliciesPriorities UpdateOrganizationBrandingPoliciesPrioritiesBody) *UpdateOrganizationBrandingPoliciesPrioritiesParams {
	o.SetUpdateOrganizationBrandingPoliciesPriorities(updateOrganizationBrandingPoliciesPriorities)
	return o
}

// SetUpdateOrganizationBrandingPoliciesPriorities adds the updateOrganizationBrandingPoliciesPriorities to the update organization branding policies priorities params
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) SetUpdateOrganizationBrandingPoliciesPriorities(updateOrganizationBrandingPoliciesPriorities UpdateOrganizationBrandingPoliciesPrioritiesBody) {
	o.UpdateOrganizationBrandingPoliciesPriorities = updateOrganizationBrandingPoliciesPriorities
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationBrandingPoliciesPrioritiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationBrandingPoliciesPriorities); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
