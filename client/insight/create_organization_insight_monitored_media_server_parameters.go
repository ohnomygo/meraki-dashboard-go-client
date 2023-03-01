// Code generated by go-swagger; DO NOT EDIT.

package insight

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

// NewCreateOrganizationInsightMonitoredMediaServerParams creates a new CreateOrganizationInsightMonitoredMediaServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationInsightMonitoredMediaServerParams() *CreateOrganizationInsightMonitoredMediaServerParams {
	return &CreateOrganizationInsightMonitoredMediaServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationInsightMonitoredMediaServerParamsWithTimeout creates a new CreateOrganizationInsightMonitoredMediaServerParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationInsightMonitoredMediaServerParamsWithTimeout(timeout time.Duration) *CreateOrganizationInsightMonitoredMediaServerParams {
	return &CreateOrganizationInsightMonitoredMediaServerParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationInsightMonitoredMediaServerParamsWithContext creates a new CreateOrganizationInsightMonitoredMediaServerParams object
// with the ability to set a context for a request.
func NewCreateOrganizationInsightMonitoredMediaServerParamsWithContext(ctx context.Context) *CreateOrganizationInsightMonitoredMediaServerParams {
	return &CreateOrganizationInsightMonitoredMediaServerParams{
		Context: ctx,
	}
}

// NewCreateOrganizationInsightMonitoredMediaServerParamsWithHTTPClient creates a new CreateOrganizationInsightMonitoredMediaServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationInsightMonitoredMediaServerParamsWithHTTPClient(client *http.Client) *CreateOrganizationInsightMonitoredMediaServerParams {
	return &CreateOrganizationInsightMonitoredMediaServerParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationInsightMonitoredMediaServerParams contains all the parameters to send to the API endpoint

	for the create organization insight monitored media server operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationInsightMonitoredMediaServerParams struct {

	// CreateOrganizationInsightMonitoredMediaServer.
	CreateOrganizationInsightMonitoredMediaServer CreateOrganizationInsightMonitoredMediaServerBody

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization insight monitored media server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationInsightMonitoredMediaServerParams) WithDefaults() *CreateOrganizationInsightMonitoredMediaServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization insight monitored media server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationInsightMonitoredMediaServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) WithTimeout(timeout time.Duration) *CreateOrganizationInsightMonitoredMediaServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) WithContext(ctx context.Context) *CreateOrganizationInsightMonitoredMediaServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) WithHTTPClient(client *http.Client) *CreateOrganizationInsightMonitoredMediaServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationInsightMonitoredMediaServer adds the createOrganizationInsightMonitoredMediaServer to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) WithCreateOrganizationInsightMonitoredMediaServer(createOrganizationInsightMonitoredMediaServer CreateOrganizationInsightMonitoredMediaServerBody) *CreateOrganizationInsightMonitoredMediaServerParams {
	o.SetCreateOrganizationInsightMonitoredMediaServer(createOrganizationInsightMonitoredMediaServer)
	return o
}

// SetCreateOrganizationInsightMonitoredMediaServer adds the createOrganizationInsightMonitoredMediaServer to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) SetCreateOrganizationInsightMonitoredMediaServer(createOrganizationInsightMonitoredMediaServer CreateOrganizationInsightMonitoredMediaServerBody) {
	o.CreateOrganizationInsightMonitoredMediaServer = createOrganizationInsightMonitoredMediaServer
}

// WithOrganizationID adds the organizationID to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) WithOrganizationID(organizationID string) *CreateOrganizationInsightMonitoredMediaServerParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization insight monitored media server params
func (o *CreateOrganizationInsightMonitoredMediaServerParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationInsightMonitoredMediaServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationInsightMonitoredMediaServer); err != nil {
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
