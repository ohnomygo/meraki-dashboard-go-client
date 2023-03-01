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

// NewGetOrganizationAdaptivePolicyOverviewParams creates a new GetOrganizationAdaptivePolicyOverviewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationAdaptivePolicyOverviewParams() *GetOrganizationAdaptivePolicyOverviewParams {
	return &GetOrganizationAdaptivePolicyOverviewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationAdaptivePolicyOverviewParamsWithTimeout creates a new GetOrganizationAdaptivePolicyOverviewParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationAdaptivePolicyOverviewParamsWithTimeout(timeout time.Duration) *GetOrganizationAdaptivePolicyOverviewParams {
	return &GetOrganizationAdaptivePolicyOverviewParams{
		timeout: timeout,
	}
}

// NewGetOrganizationAdaptivePolicyOverviewParamsWithContext creates a new GetOrganizationAdaptivePolicyOverviewParams object
// with the ability to set a context for a request.
func NewGetOrganizationAdaptivePolicyOverviewParamsWithContext(ctx context.Context) *GetOrganizationAdaptivePolicyOverviewParams {
	return &GetOrganizationAdaptivePolicyOverviewParams{
		Context: ctx,
	}
}

// NewGetOrganizationAdaptivePolicyOverviewParamsWithHTTPClient creates a new GetOrganizationAdaptivePolicyOverviewParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationAdaptivePolicyOverviewParamsWithHTTPClient(client *http.Client) *GetOrganizationAdaptivePolicyOverviewParams {
	return &GetOrganizationAdaptivePolicyOverviewParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationAdaptivePolicyOverviewParams contains all the parameters to send to the API endpoint

	for the get organization adaptive policy overview operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationAdaptivePolicyOverviewParams struct {

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization adaptive policy overview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAdaptivePolicyOverviewParams) WithDefaults() *GetOrganizationAdaptivePolicyOverviewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization adaptive policy overview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAdaptivePolicyOverviewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization adaptive policy overview params
func (o *GetOrganizationAdaptivePolicyOverviewParams) WithTimeout(timeout time.Duration) *GetOrganizationAdaptivePolicyOverviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization adaptive policy overview params
func (o *GetOrganizationAdaptivePolicyOverviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization adaptive policy overview params
func (o *GetOrganizationAdaptivePolicyOverviewParams) WithContext(ctx context.Context) *GetOrganizationAdaptivePolicyOverviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization adaptive policy overview params
func (o *GetOrganizationAdaptivePolicyOverviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization adaptive policy overview params
func (o *GetOrganizationAdaptivePolicyOverviewParams) WithHTTPClient(client *http.Client) *GetOrganizationAdaptivePolicyOverviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization adaptive policy overview params
func (o *GetOrganizationAdaptivePolicyOverviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization adaptive policy overview params
func (o *GetOrganizationAdaptivePolicyOverviewParams) WithOrganizationID(organizationID string) *GetOrganizationAdaptivePolicyOverviewParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization adaptive policy overview params
func (o *GetOrganizationAdaptivePolicyOverviewParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationAdaptivePolicyOverviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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