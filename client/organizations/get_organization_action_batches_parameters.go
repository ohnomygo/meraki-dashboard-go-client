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

// NewGetOrganizationActionBatchesParams creates a new GetOrganizationActionBatchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationActionBatchesParams() *GetOrganizationActionBatchesParams {
	return &GetOrganizationActionBatchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationActionBatchesParamsWithTimeout creates a new GetOrganizationActionBatchesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationActionBatchesParamsWithTimeout(timeout time.Duration) *GetOrganizationActionBatchesParams {
	return &GetOrganizationActionBatchesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationActionBatchesParamsWithContext creates a new GetOrganizationActionBatchesParams object
// with the ability to set a context for a request.
func NewGetOrganizationActionBatchesParamsWithContext(ctx context.Context) *GetOrganizationActionBatchesParams {
	return &GetOrganizationActionBatchesParams{
		Context: ctx,
	}
}

// NewGetOrganizationActionBatchesParamsWithHTTPClient creates a new GetOrganizationActionBatchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationActionBatchesParamsWithHTTPClient(client *http.Client) *GetOrganizationActionBatchesParams {
	return &GetOrganizationActionBatchesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationActionBatchesParams contains all the parameters to send to the API endpoint

	for the get organization action batches operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationActionBatchesParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* Status.

	   Filter batches by status. Valid types are pending, completed, and failed.
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization action batches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationActionBatchesParams) WithDefaults() *GetOrganizationActionBatchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization action batches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationActionBatchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) WithTimeout(timeout time.Duration) *GetOrganizationActionBatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) WithContext(ctx context.Context) *GetOrganizationActionBatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) WithHTTPClient(client *http.Client) *GetOrganizationActionBatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) WithOrganizationID(organizationID string) *GetOrganizationActionBatchesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithStatus adds the status to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) WithStatus(status *string) *GetOrganizationActionBatchesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get organization action batches params
func (o *GetOrganizationActionBatchesParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationActionBatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
