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

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportParams creates a new CreateOrganizationInventoryOnboardingCloudMonitoringImportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportParams() *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	return &CreateOrganizationInventoryOnboardingCloudMonitoringImportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportParamsWithTimeout creates a new CreateOrganizationInventoryOnboardingCloudMonitoringImportParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportParamsWithTimeout(timeout time.Duration) *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	return &CreateOrganizationInventoryOnboardingCloudMonitoringImportParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportParamsWithContext creates a new CreateOrganizationInventoryOnboardingCloudMonitoringImportParams object
// with the ability to set a context for a request.
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportParamsWithContext(ctx context.Context) *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	return &CreateOrganizationInventoryOnboardingCloudMonitoringImportParams{
		Context: ctx,
	}
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportParamsWithHTTPClient creates a new CreateOrganizationInventoryOnboardingCloudMonitoringImportParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportParamsWithHTTPClient(client *http.Client) *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	return &CreateOrganizationInventoryOnboardingCloudMonitoringImportParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImportParams contains all the parameters to send to the API endpoint

	for the create organization inventory onboarding cloud monitoring import operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationInventoryOnboardingCloudMonitoringImportParams struct {

	// CreateOrganizationInventoryOnboardingCloudMonitoringImport.
	CreateOrganizationInventoryOnboardingCloudMonitoringImport CreateOrganizationInventoryOnboardingCloudMonitoringImportBody

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization inventory onboarding cloud monitoring import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) WithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization inventory onboarding cloud monitoring import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) WithTimeout(timeout time.Duration) *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) WithContext(ctx context.Context) *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) WithHTTPClient(client *http.Client) *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationInventoryOnboardingCloudMonitoringImport adds the createOrganizationInventoryOnboardingCloudMonitoringImport to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) WithCreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport CreateOrganizationInventoryOnboardingCloudMonitoringImportBody) *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	o.SetCreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport)
	return o
}

// SetCreateOrganizationInventoryOnboardingCloudMonitoringImport adds the createOrganizationInventoryOnboardingCloudMonitoringImport to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) SetCreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport CreateOrganizationInventoryOnboardingCloudMonitoringImportBody) {
	o.CreateOrganizationInventoryOnboardingCloudMonitoringImport = createOrganizationInventoryOnboardingCloudMonitoringImport
}

// WithOrganizationID adds the organizationID to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) WithOrganizationID(organizationID string) *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization inventory onboarding cloud monitoring import params
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationInventoryOnboardingCloudMonitoringImport); err != nil {
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
