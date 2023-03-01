// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// NewGetOrganizationApplianceSecurityIntrusionParams creates a new GetOrganizationApplianceSecurityIntrusionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationApplianceSecurityIntrusionParams() *GetOrganizationApplianceSecurityIntrusionParams {
	return &GetOrganizationApplianceSecurityIntrusionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationApplianceSecurityIntrusionParamsWithTimeout creates a new GetOrganizationApplianceSecurityIntrusionParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationApplianceSecurityIntrusionParamsWithTimeout(timeout time.Duration) *GetOrganizationApplianceSecurityIntrusionParams {
	return &GetOrganizationApplianceSecurityIntrusionParams{
		timeout: timeout,
	}
}

// NewGetOrganizationApplianceSecurityIntrusionParamsWithContext creates a new GetOrganizationApplianceSecurityIntrusionParams object
// with the ability to set a context for a request.
func NewGetOrganizationApplianceSecurityIntrusionParamsWithContext(ctx context.Context) *GetOrganizationApplianceSecurityIntrusionParams {
	return &GetOrganizationApplianceSecurityIntrusionParams{
		Context: ctx,
	}
}

// NewGetOrganizationApplianceSecurityIntrusionParamsWithHTTPClient creates a new GetOrganizationApplianceSecurityIntrusionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationApplianceSecurityIntrusionParamsWithHTTPClient(client *http.Client) *GetOrganizationApplianceSecurityIntrusionParams {
	return &GetOrganizationApplianceSecurityIntrusionParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationApplianceSecurityIntrusionParams contains all the parameters to send to the API endpoint

	for the get organization appliance security intrusion operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationApplianceSecurityIntrusionParams struct {

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization appliance security intrusion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationApplianceSecurityIntrusionParams) WithDefaults() *GetOrganizationApplianceSecurityIntrusionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization appliance security intrusion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationApplianceSecurityIntrusionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization appliance security intrusion params
func (o *GetOrganizationApplianceSecurityIntrusionParams) WithTimeout(timeout time.Duration) *GetOrganizationApplianceSecurityIntrusionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization appliance security intrusion params
func (o *GetOrganizationApplianceSecurityIntrusionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization appliance security intrusion params
func (o *GetOrganizationApplianceSecurityIntrusionParams) WithContext(ctx context.Context) *GetOrganizationApplianceSecurityIntrusionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization appliance security intrusion params
func (o *GetOrganizationApplianceSecurityIntrusionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization appliance security intrusion params
func (o *GetOrganizationApplianceSecurityIntrusionParams) WithHTTPClient(client *http.Client) *GetOrganizationApplianceSecurityIntrusionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization appliance security intrusion params
func (o *GetOrganizationApplianceSecurityIntrusionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization appliance security intrusion params
func (o *GetOrganizationApplianceSecurityIntrusionParams) WithOrganizationID(organizationID string) *GetOrganizationApplianceSecurityIntrusionParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization appliance security intrusion params
func (o *GetOrganizationApplianceSecurityIntrusionParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationApplianceSecurityIntrusionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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