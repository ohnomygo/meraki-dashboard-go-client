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

// NewUpdateOrganizationApplianceSecurityIntrusionParams creates a new UpdateOrganizationApplianceSecurityIntrusionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationApplianceSecurityIntrusionParams() *UpdateOrganizationApplianceSecurityIntrusionParams {
	return &UpdateOrganizationApplianceSecurityIntrusionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationApplianceSecurityIntrusionParamsWithTimeout creates a new UpdateOrganizationApplianceSecurityIntrusionParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationApplianceSecurityIntrusionParamsWithTimeout(timeout time.Duration) *UpdateOrganizationApplianceSecurityIntrusionParams {
	return &UpdateOrganizationApplianceSecurityIntrusionParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationApplianceSecurityIntrusionParamsWithContext creates a new UpdateOrganizationApplianceSecurityIntrusionParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationApplianceSecurityIntrusionParamsWithContext(ctx context.Context) *UpdateOrganizationApplianceSecurityIntrusionParams {
	return &UpdateOrganizationApplianceSecurityIntrusionParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationApplianceSecurityIntrusionParamsWithHTTPClient creates a new UpdateOrganizationApplianceSecurityIntrusionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationApplianceSecurityIntrusionParamsWithHTTPClient(client *http.Client) *UpdateOrganizationApplianceSecurityIntrusionParams {
	return &UpdateOrganizationApplianceSecurityIntrusionParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationApplianceSecurityIntrusionParams contains all the parameters to send to the API endpoint

	for the update organization appliance security intrusion operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationApplianceSecurityIntrusionParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	// UpdateOrganizationApplianceSecurityIntrusion.
	UpdateOrganizationApplianceSecurityIntrusion UpdateOrganizationApplianceSecurityIntrusionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization appliance security intrusion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) WithDefaults() *UpdateOrganizationApplianceSecurityIntrusionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization appliance security intrusion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) WithTimeout(timeout time.Duration) *UpdateOrganizationApplianceSecurityIntrusionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) WithContext(ctx context.Context) *UpdateOrganizationApplianceSecurityIntrusionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) WithHTTPClient(client *http.Client) *UpdateOrganizationApplianceSecurityIntrusionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) WithOrganizationID(organizationID string) *UpdateOrganizationApplianceSecurityIntrusionParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationApplianceSecurityIntrusion adds the updateOrganizationApplianceSecurityIntrusion to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) WithUpdateOrganizationApplianceSecurityIntrusion(updateOrganizationApplianceSecurityIntrusion UpdateOrganizationApplianceSecurityIntrusionBody) *UpdateOrganizationApplianceSecurityIntrusionParams {
	o.SetUpdateOrganizationApplianceSecurityIntrusion(updateOrganizationApplianceSecurityIntrusion)
	return o
}

// SetUpdateOrganizationApplianceSecurityIntrusion adds the updateOrganizationApplianceSecurityIntrusion to the update organization appliance security intrusion params
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) SetUpdateOrganizationApplianceSecurityIntrusion(updateOrganizationApplianceSecurityIntrusion UpdateOrganizationApplianceSecurityIntrusionBody) {
	o.UpdateOrganizationApplianceSecurityIntrusion = updateOrganizationApplianceSecurityIntrusion
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationApplianceSecurityIntrusionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationApplianceSecurityIntrusion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
