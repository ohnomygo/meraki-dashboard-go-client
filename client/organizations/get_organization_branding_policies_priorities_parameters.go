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

// NewGetOrganizationBrandingPoliciesPrioritiesParams creates a new GetOrganizationBrandingPoliciesPrioritiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationBrandingPoliciesPrioritiesParams() *GetOrganizationBrandingPoliciesPrioritiesParams {
	return &GetOrganizationBrandingPoliciesPrioritiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationBrandingPoliciesPrioritiesParamsWithTimeout creates a new GetOrganizationBrandingPoliciesPrioritiesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationBrandingPoliciesPrioritiesParamsWithTimeout(timeout time.Duration) *GetOrganizationBrandingPoliciesPrioritiesParams {
	return &GetOrganizationBrandingPoliciesPrioritiesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationBrandingPoliciesPrioritiesParamsWithContext creates a new GetOrganizationBrandingPoliciesPrioritiesParams object
// with the ability to set a context for a request.
func NewGetOrganizationBrandingPoliciesPrioritiesParamsWithContext(ctx context.Context) *GetOrganizationBrandingPoliciesPrioritiesParams {
	return &GetOrganizationBrandingPoliciesPrioritiesParams{
		Context: ctx,
	}
}

// NewGetOrganizationBrandingPoliciesPrioritiesParamsWithHTTPClient creates a new GetOrganizationBrandingPoliciesPrioritiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationBrandingPoliciesPrioritiesParamsWithHTTPClient(client *http.Client) *GetOrganizationBrandingPoliciesPrioritiesParams {
	return &GetOrganizationBrandingPoliciesPrioritiesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationBrandingPoliciesPrioritiesParams contains all the parameters to send to the API endpoint

	for the get organization branding policies priorities operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationBrandingPoliciesPrioritiesParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization branding policies priorities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) WithDefaults() *GetOrganizationBrandingPoliciesPrioritiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization branding policies priorities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization branding policies priorities params
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) WithTimeout(timeout time.Duration) *GetOrganizationBrandingPoliciesPrioritiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization branding policies priorities params
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization branding policies priorities params
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) WithContext(ctx context.Context) *GetOrganizationBrandingPoliciesPrioritiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization branding policies priorities params
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization branding policies priorities params
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) WithHTTPClient(client *http.Client) *GetOrganizationBrandingPoliciesPrioritiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization branding policies priorities params
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization branding policies priorities params
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) WithOrganizationID(organizationID string) *GetOrganizationBrandingPoliciesPrioritiesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization branding policies priorities params
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationBrandingPoliciesPrioritiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
