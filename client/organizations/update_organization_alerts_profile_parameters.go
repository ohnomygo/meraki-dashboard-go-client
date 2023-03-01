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

// NewUpdateOrganizationAlertsProfileParams creates a new UpdateOrganizationAlertsProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationAlertsProfileParams() *UpdateOrganizationAlertsProfileParams {
	return &UpdateOrganizationAlertsProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationAlertsProfileParamsWithTimeout creates a new UpdateOrganizationAlertsProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationAlertsProfileParamsWithTimeout(timeout time.Duration) *UpdateOrganizationAlertsProfileParams {
	return &UpdateOrganizationAlertsProfileParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationAlertsProfileParamsWithContext creates a new UpdateOrganizationAlertsProfileParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationAlertsProfileParamsWithContext(ctx context.Context) *UpdateOrganizationAlertsProfileParams {
	return &UpdateOrganizationAlertsProfileParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationAlertsProfileParamsWithHTTPClient creates a new UpdateOrganizationAlertsProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationAlertsProfileParamsWithHTTPClient(client *http.Client) *UpdateOrganizationAlertsProfileParams {
	return &UpdateOrganizationAlertsProfileParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationAlertsProfileParams contains all the parameters to send to the API endpoint

	for the update organization alerts profile operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationAlertsProfileParams struct {

	// AlertConfigID.
	AlertConfigID string

	// OrganizationID.
	OrganizationID string

	// UpdateOrganizationAlertsProfile.
	UpdateOrganizationAlertsProfile UpdateOrganizationAlertsProfileBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization alerts profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationAlertsProfileParams) WithDefaults() *UpdateOrganizationAlertsProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization alerts profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationAlertsProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) WithTimeout(timeout time.Duration) *UpdateOrganizationAlertsProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) WithContext(ctx context.Context) *UpdateOrganizationAlertsProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) WithHTTPClient(client *http.Client) *UpdateOrganizationAlertsProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertConfigID adds the alertConfigID to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) WithAlertConfigID(alertConfigID string) *UpdateOrganizationAlertsProfileParams {
	o.SetAlertConfigID(alertConfigID)
	return o
}

// SetAlertConfigID adds the alertConfigId to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) SetAlertConfigID(alertConfigID string) {
	o.AlertConfigID = alertConfigID
}

// WithOrganizationID adds the organizationID to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) WithOrganizationID(organizationID string) *UpdateOrganizationAlertsProfileParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationAlertsProfile adds the updateOrganizationAlertsProfile to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) WithUpdateOrganizationAlertsProfile(updateOrganizationAlertsProfile UpdateOrganizationAlertsProfileBody) *UpdateOrganizationAlertsProfileParams {
	o.SetUpdateOrganizationAlertsProfile(updateOrganizationAlertsProfile)
	return o
}

// SetUpdateOrganizationAlertsProfile adds the updateOrganizationAlertsProfile to the update organization alerts profile params
func (o *UpdateOrganizationAlertsProfileParams) SetUpdateOrganizationAlertsProfile(updateOrganizationAlertsProfile UpdateOrganizationAlertsProfileBody) {
	o.UpdateOrganizationAlertsProfile = updateOrganizationAlertsProfile
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationAlertsProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertConfigId
	if err := r.SetPathParam("alertConfigId", o.AlertConfigID); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationAlertsProfile); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
