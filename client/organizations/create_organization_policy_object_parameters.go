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

// NewCreateOrganizationPolicyObjectParams creates a new CreateOrganizationPolicyObjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationPolicyObjectParams() *CreateOrganizationPolicyObjectParams {
	return &CreateOrganizationPolicyObjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationPolicyObjectParamsWithTimeout creates a new CreateOrganizationPolicyObjectParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationPolicyObjectParamsWithTimeout(timeout time.Duration) *CreateOrganizationPolicyObjectParams {
	return &CreateOrganizationPolicyObjectParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationPolicyObjectParamsWithContext creates a new CreateOrganizationPolicyObjectParams object
// with the ability to set a context for a request.
func NewCreateOrganizationPolicyObjectParamsWithContext(ctx context.Context) *CreateOrganizationPolicyObjectParams {
	return &CreateOrganizationPolicyObjectParams{
		Context: ctx,
	}
}

// NewCreateOrganizationPolicyObjectParamsWithHTTPClient creates a new CreateOrganizationPolicyObjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationPolicyObjectParamsWithHTTPClient(client *http.Client) *CreateOrganizationPolicyObjectParams {
	return &CreateOrganizationPolicyObjectParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationPolicyObjectParams contains all the parameters to send to the API endpoint

	for the create organization policy object operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationPolicyObjectParams struct {

	// CreateOrganizationPolicyObject.
	CreateOrganizationPolicyObject CreateOrganizationPolicyObjectBody

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization policy object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationPolicyObjectParams) WithDefaults() *CreateOrganizationPolicyObjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization policy object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationPolicyObjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) WithTimeout(timeout time.Duration) *CreateOrganizationPolicyObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) WithContext(ctx context.Context) *CreateOrganizationPolicyObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) WithHTTPClient(client *http.Client) *CreateOrganizationPolicyObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationPolicyObject adds the createOrganizationPolicyObject to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) WithCreateOrganizationPolicyObject(createOrganizationPolicyObject CreateOrganizationPolicyObjectBody) *CreateOrganizationPolicyObjectParams {
	o.SetCreateOrganizationPolicyObject(createOrganizationPolicyObject)
	return o
}

// SetCreateOrganizationPolicyObject adds the createOrganizationPolicyObject to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) SetCreateOrganizationPolicyObject(createOrganizationPolicyObject CreateOrganizationPolicyObjectBody) {
	o.CreateOrganizationPolicyObject = createOrganizationPolicyObject
}

// WithOrganizationID adds the organizationID to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) WithOrganizationID(organizationID string) *CreateOrganizationPolicyObjectParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization policy object params
func (o *CreateOrganizationPolicyObjectParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationPolicyObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationPolicyObject); err != nil {
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
