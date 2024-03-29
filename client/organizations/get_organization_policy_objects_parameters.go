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
	"github.com/go-openapi/swag"
)

// NewGetOrganizationPolicyObjectsParams creates a new GetOrganizationPolicyObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationPolicyObjectsParams() *GetOrganizationPolicyObjectsParams {
	return &GetOrganizationPolicyObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationPolicyObjectsParamsWithTimeout creates a new GetOrganizationPolicyObjectsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationPolicyObjectsParamsWithTimeout(timeout time.Duration) *GetOrganizationPolicyObjectsParams {
	return &GetOrganizationPolicyObjectsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationPolicyObjectsParamsWithContext creates a new GetOrganizationPolicyObjectsParams object
// with the ability to set a context for a request.
func NewGetOrganizationPolicyObjectsParamsWithContext(ctx context.Context) *GetOrganizationPolicyObjectsParams {
	return &GetOrganizationPolicyObjectsParams{
		Context: ctx,
	}
}

// NewGetOrganizationPolicyObjectsParamsWithHTTPClient creates a new GetOrganizationPolicyObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationPolicyObjectsParamsWithHTTPClient(client *http.Client) *GetOrganizationPolicyObjectsParams {
	return &GetOrganizationPolicyObjectsParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationPolicyObjectsParams contains all the parameters to send to the API endpoint

	for the get organization policy objects operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationPolicyObjectsParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 10 - 5000. Default is 5000.
	*/
	PerPage *int64

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization policy objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationPolicyObjectsParams) WithDefaults() *GetOrganizationPolicyObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization policy objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationPolicyObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) WithTimeout(timeout time.Duration) *GetOrganizationPolicyObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) WithContext(ctx context.Context) *GetOrganizationPolicyObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) WithHTTPClient(client *http.Client) *GetOrganizationPolicyObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) WithEndingBefore(endingBefore *string) *GetOrganizationPolicyObjectsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithOrganizationID adds the organizationID to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) WithOrganizationID(organizationID string) *GetOrganizationPolicyObjectsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) WithPerPage(perPage *int64) *GetOrganizationPolicyObjectsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) WithStartingAfter(startingAfter *string) *GetOrganizationPolicyObjectsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization policy objects params
func (o *GetOrganizationPolicyObjectsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationPolicyObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingBefore != nil {

		// query param endingBefore
		var qrEndingBefore string

		if o.EndingBefore != nil {
			qrEndingBefore = *o.EndingBefore
		}
		qEndingBefore := qrEndingBefore
		if qEndingBefore != "" {

			if err := r.SetQueryParam("endingBefore", qEndingBefore); err != nil {
				return err
			}
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.StartingAfter != nil {

		// query param startingAfter
		var qrStartingAfter string

		if o.StartingAfter != nil {
			qrStartingAfter = *o.StartingAfter
		}
		qStartingAfter := qrStartingAfter
		if qStartingAfter != "" {

			if err := r.SetQueryParam("startingAfter", qStartingAfter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
