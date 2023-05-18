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

// NewUpdateOrganizationEarlyAccessFeaturesOptInParams creates a new UpdateOrganizationEarlyAccessFeaturesOptInParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationEarlyAccessFeaturesOptInParams() *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	return &UpdateOrganizationEarlyAccessFeaturesOptInParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationEarlyAccessFeaturesOptInParamsWithTimeout creates a new UpdateOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationEarlyAccessFeaturesOptInParamsWithTimeout(timeout time.Duration) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	return &UpdateOrganizationEarlyAccessFeaturesOptInParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationEarlyAccessFeaturesOptInParamsWithContext creates a new UpdateOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationEarlyAccessFeaturesOptInParamsWithContext(ctx context.Context) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	return &UpdateOrganizationEarlyAccessFeaturesOptInParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationEarlyAccessFeaturesOptInParamsWithHTTPClient creates a new UpdateOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationEarlyAccessFeaturesOptInParamsWithHTTPClient(client *http.Client) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	return &UpdateOrganizationEarlyAccessFeaturesOptInParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationEarlyAccessFeaturesOptInParams contains all the parameters to send to the API endpoint

	for the update organization early access features opt in operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationEarlyAccessFeaturesOptInParams struct {

	/* OptInID.

	   Opt in ID
	*/
	OptInID string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	// UpdateOrganizationEarlyAccessFeaturesOptIn.
	UpdateOrganizationEarlyAccessFeaturesOptIn UpdateOrganizationEarlyAccessFeaturesOptInBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization early access features opt in params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) WithDefaults() *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization early access features opt in params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) WithTimeout(timeout time.Duration) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) WithContext(ctx context.Context) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) WithHTTPClient(client *http.Client) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptInID adds the optInID to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) WithOptInID(optInID string) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetOptInID(optInID)
	return o
}

// SetOptInID adds the optInId to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) SetOptInID(optInID string) {
	o.OptInID = optInID
}

// WithOrganizationID adds the organizationID to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) WithOrganizationID(organizationID string) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationEarlyAccessFeaturesOptIn adds the updateOrganizationEarlyAccessFeaturesOptIn to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) WithUpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn UpdateOrganizationEarlyAccessFeaturesOptInBody) *UpdateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetUpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn)
	return o
}

// SetUpdateOrganizationEarlyAccessFeaturesOptIn adds the updateOrganizationEarlyAccessFeaturesOptIn to the update organization early access features opt in params
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) SetUpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn UpdateOrganizationEarlyAccessFeaturesOptInBody) {
	o.UpdateOrganizationEarlyAccessFeaturesOptIn = updateOrganizationEarlyAccessFeaturesOptIn
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationEarlyAccessFeaturesOptInParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param optInId
	if err := r.SetPathParam("optInId", o.OptInID); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationEarlyAccessFeaturesOptIn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
