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

// NewCreateOrganizationSamlRoleParams creates a new CreateOrganizationSamlRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationSamlRoleParams() *CreateOrganizationSamlRoleParams {
	return &CreateOrganizationSamlRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationSamlRoleParamsWithTimeout creates a new CreateOrganizationSamlRoleParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationSamlRoleParamsWithTimeout(timeout time.Duration) *CreateOrganizationSamlRoleParams {
	return &CreateOrganizationSamlRoleParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationSamlRoleParamsWithContext creates a new CreateOrganizationSamlRoleParams object
// with the ability to set a context for a request.
func NewCreateOrganizationSamlRoleParamsWithContext(ctx context.Context) *CreateOrganizationSamlRoleParams {
	return &CreateOrganizationSamlRoleParams{
		Context: ctx,
	}
}

// NewCreateOrganizationSamlRoleParamsWithHTTPClient creates a new CreateOrganizationSamlRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationSamlRoleParamsWithHTTPClient(client *http.Client) *CreateOrganizationSamlRoleParams {
	return &CreateOrganizationSamlRoleParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationSamlRoleParams contains all the parameters to send to the API endpoint

	for the create organization saml role operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationSamlRoleParams struct {

	// CreateOrganizationSamlRole.
	CreateOrganizationSamlRole CreateOrganizationSamlRoleBody

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization saml role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationSamlRoleParams) WithDefaults() *CreateOrganizationSamlRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization saml role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationSamlRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) WithTimeout(timeout time.Duration) *CreateOrganizationSamlRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) WithContext(ctx context.Context) *CreateOrganizationSamlRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) WithHTTPClient(client *http.Client) *CreateOrganizationSamlRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationSamlRole adds the createOrganizationSamlRole to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) WithCreateOrganizationSamlRole(createOrganizationSamlRole CreateOrganizationSamlRoleBody) *CreateOrganizationSamlRoleParams {
	o.SetCreateOrganizationSamlRole(createOrganizationSamlRole)
	return o
}

// SetCreateOrganizationSamlRole adds the createOrganizationSamlRole to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) SetCreateOrganizationSamlRole(createOrganizationSamlRole CreateOrganizationSamlRoleBody) {
	o.CreateOrganizationSamlRole = createOrganizationSamlRole
}

// WithOrganizationID adds the organizationID to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) WithOrganizationID(organizationID string) *CreateOrganizationSamlRoleParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization saml role params
func (o *CreateOrganizationSamlRoleParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationSamlRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationSamlRole); err != nil {
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
