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

// NewUpdateOrganizationAdaptivePolicyPolicyParams creates a new UpdateOrganizationAdaptivePolicyPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationAdaptivePolicyPolicyParams() *UpdateOrganizationAdaptivePolicyPolicyParams {
	return &UpdateOrganizationAdaptivePolicyPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationAdaptivePolicyPolicyParamsWithTimeout creates a new UpdateOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationAdaptivePolicyPolicyParamsWithTimeout(timeout time.Duration) *UpdateOrganizationAdaptivePolicyPolicyParams {
	return &UpdateOrganizationAdaptivePolicyPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationAdaptivePolicyPolicyParamsWithContext creates a new UpdateOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationAdaptivePolicyPolicyParamsWithContext(ctx context.Context) *UpdateOrganizationAdaptivePolicyPolicyParams {
	return &UpdateOrganizationAdaptivePolicyPolicyParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationAdaptivePolicyPolicyParamsWithHTTPClient creates a new UpdateOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationAdaptivePolicyPolicyParamsWithHTTPClient(client *http.Client) *UpdateOrganizationAdaptivePolicyPolicyParams {
	return &UpdateOrganizationAdaptivePolicyPolicyParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationAdaptivePolicyPolicyParams contains all the parameters to send to the API endpoint

	for the update organization adaptive policy policy operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationAdaptivePolicyPolicyParams struct {

	/* ID.

	   Id
	*/
	ID string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	// UpdateOrganizationAdaptivePolicyPolicy.
	UpdateOrganizationAdaptivePolicyPolicy UpdateOrganizationAdaptivePolicyPolicyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization adaptive policy policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) WithDefaults() *UpdateOrganizationAdaptivePolicyPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization adaptive policy policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) WithTimeout(timeout time.Duration) *UpdateOrganizationAdaptivePolicyPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) WithContext(ctx context.Context) *UpdateOrganizationAdaptivePolicyPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) WithHTTPClient(client *http.Client) *UpdateOrganizationAdaptivePolicyPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) WithID(id string) *UpdateOrganizationAdaptivePolicyPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) SetID(id string) {
	o.ID = id
}

// WithOrganizationID adds the organizationID to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) WithOrganizationID(organizationID string) *UpdateOrganizationAdaptivePolicyPolicyParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationAdaptivePolicyPolicy adds the updateOrganizationAdaptivePolicyPolicy to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) WithUpdateOrganizationAdaptivePolicyPolicy(updateOrganizationAdaptivePolicyPolicy UpdateOrganizationAdaptivePolicyPolicyBody) *UpdateOrganizationAdaptivePolicyPolicyParams {
	o.SetUpdateOrganizationAdaptivePolicyPolicy(updateOrganizationAdaptivePolicyPolicy)
	return o
}

// SetUpdateOrganizationAdaptivePolicyPolicy adds the updateOrganizationAdaptivePolicyPolicy to the update organization adaptive policy policy params
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) SetUpdateOrganizationAdaptivePolicyPolicy(updateOrganizationAdaptivePolicyPolicy UpdateOrganizationAdaptivePolicyPolicyBody) {
	o.UpdateOrganizationAdaptivePolicyPolicy = updateOrganizationAdaptivePolicyPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationAdaptivePolicyPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.UpdateOrganizationAdaptivePolicyPolicy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
