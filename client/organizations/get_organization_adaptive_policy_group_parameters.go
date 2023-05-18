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

// NewGetOrganizationAdaptivePolicyGroupParams creates a new GetOrganizationAdaptivePolicyGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationAdaptivePolicyGroupParams() *GetOrganizationAdaptivePolicyGroupParams {
	return &GetOrganizationAdaptivePolicyGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationAdaptivePolicyGroupParamsWithTimeout creates a new GetOrganizationAdaptivePolicyGroupParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationAdaptivePolicyGroupParamsWithTimeout(timeout time.Duration) *GetOrganizationAdaptivePolicyGroupParams {
	return &GetOrganizationAdaptivePolicyGroupParams{
		timeout: timeout,
	}
}

// NewGetOrganizationAdaptivePolicyGroupParamsWithContext creates a new GetOrganizationAdaptivePolicyGroupParams object
// with the ability to set a context for a request.
func NewGetOrganizationAdaptivePolicyGroupParamsWithContext(ctx context.Context) *GetOrganizationAdaptivePolicyGroupParams {
	return &GetOrganizationAdaptivePolicyGroupParams{
		Context: ctx,
	}
}

// NewGetOrganizationAdaptivePolicyGroupParamsWithHTTPClient creates a new GetOrganizationAdaptivePolicyGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationAdaptivePolicyGroupParamsWithHTTPClient(client *http.Client) *GetOrganizationAdaptivePolicyGroupParams {
	return &GetOrganizationAdaptivePolicyGroupParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationAdaptivePolicyGroupParams contains all the parameters to send to the API endpoint

	for the get organization adaptive policy group operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationAdaptivePolicyGroupParams struct {

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

// WithDefaults hydrates default values in the get organization adaptive policy group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAdaptivePolicyGroupParams) WithDefaults() *GetOrganizationAdaptivePolicyGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization adaptive policy group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAdaptivePolicyGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) WithTimeout(timeout time.Duration) *GetOrganizationAdaptivePolicyGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) WithContext(ctx context.Context) *GetOrganizationAdaptivePolicyGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) WithHTTPClient(client *http.Client) *GetOrganizationAdaptivePolicyGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) WithID(id string) *GetOrganizationAdaptivePolicyGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) SetID(id string) {
	o.ID = id
}

// WithOrganizationID adds the organizationID to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) WithOrganizationID(organizationID string) *GetOrganizationAdaptivePolicyGroupParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization adaptive policy group params
func (o *GetOrganizationAdaptivePolicyGroupParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationAdaptivePolicyGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
