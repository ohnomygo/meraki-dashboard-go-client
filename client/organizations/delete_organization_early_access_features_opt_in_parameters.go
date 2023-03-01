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

// NewDeleteOrganizationEarlyAccessFeaturesOptInParams creates a new DeleteOrganizationEarlyAccessFeaturesOptInParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationEarlyAccessFeaturesOptInParams() *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	return &DeleteOrganizationEarlyAccessFeaturesOptInParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationEarlyAccessFeaturesOptInParamsWithTimeout creates a new DeleteOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationEarlyAccessFeaturesOptInParamsWithTimeout(timeout time.Duration) *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	return &DeleteOrganizationEarlyAccessFeaturesOptInParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationEarlyAccessFeaturesOptInParamsWithContext creates a new DeleteOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationEarlyAccessFeaturesOptInParamsWithContext(ctx context.Context) *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	return &DeleteOrganizationEarlyAccessFeaturesOptInParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationEarlyAccessFeaturesOptInParamsWithHTTPClient creates a new DeleteOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationEarlyAccessFeaturesOptInParamsWithHTTPClient(client *http.Client) *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	return &DeleteOrganizationEarlyAccessFeaturesOptInParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationEarlyAccessFeaturesOptInParams contains all the parameters to send to the API endpoint

	for the delete organization early access features opt in operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationEarlyAccessFeaturesOptInParams struct {

	// OptInID.
	OptInID string

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization early access features opt in params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) WithDefaults() *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization early access features opt in params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) WithTimeout(timeout time.Duration) *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) WithContext(ctx context.Context) *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) WithHTTPClient(client *http.Client) *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptInID adds the optInID to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) WithOptInID(optInID string) *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	o.SetOptInID(optInID)
	return o
}

// SetOptInID adds the optInId to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) SetOptInID(optInID string) {
	o.OptInID = optInID
}

// WithOrganizationID adds the organizationID to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) WithOrganizationID(organizationID string) *DeleteOrganizationEarlyAccessFeaturesOptInParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete organization early access features opt in params
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationEarlyAccessFeaturesOptInParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
