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
	"github.com/go-openapi/swag"
)

// NewGetOrganizationApplianceSecurityEventsParams creates a new GetOrganizationApplianceSecurityEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationApplianceSecurityEventsParams() *GetOrganizationApplianceSecurityEventsParams {
	return &GetOrganizationApplianceSecurityEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationApplianceSecurityEventsParamsWithTimeout creates a new GetOrganizationApplianceSecurityEventsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationApplianceSecurityEventsParamsWithTimeout(timeout time.Duration) *GetOrganizationApplianceSecurityEventsParams {
	return &GetOrganizationApplianceSecurityEventsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationApplianceSecurityEventsParamsWithContext creates a new GetOrganizationApplianceSecurityEventsParams object
// with the ability to set a context for a request.
func NewGetOrganizationApplianceSecurityEventsParamsWithContext(ctx context.Context) *GetOrganizationApplianceSecurityEventsParams {
	return &GetOrganizationApplianceSecurityEventsParams{
		Context: ctx,
	}
}

// NewGetOrganizationApplianceSecurityEventsParamsWithHTTPClient creates a new GetOrganizationApplianceSecurityEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationApplianceSecurityEventsParamsWithHTTPClient(client *http.Client) *GetOrganizationApplianceSecurityEventsParams {
	return &GetOrganizationApplianceSecurityEventsParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationApplianceSecurityEventsParams contains all the parameters to send to the API endpoint

	for the get organization appliance security events operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationApplianceSecurityEventsParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	// OrganizationID.
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	*/
	PerPage *int64

	/* SortOrder.

	   Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order.
	*/
	SortOrder *string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* T0.

	   The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization appliance security events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationApplianceSecurityEventsParams) WithDefaults() *GetOrganizationApplianceSecurityEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization appliance security events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationApplianceSecurityEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithTimeout(timeout time.Duration) *GetOrganizationApplianceSecurityEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithContext(ctx context.Context) *GetOrganizationApplianceSecurityEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithHTTPClient(client *http.Client) *GetOrganizationApplianceSecurityEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithEndingBefore(endingBefore *string) *GetOrganizationApplianceSecurityEventsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithOrganizationID adds the organizationID to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithOrganizationID(organizationID string) *GetOrganizationApplianceSecurityEventsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithPerPage(perPage *int64) *GetOrganizationApplianceSecurityEventsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSortOrder adds the sortOrder to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithSortOrder(sortOrder *string) *GetOrganizationApplianceSecurityEventsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStartingAfter adds the startingAfter to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithStartingAfter(startingAfter *string) *GetOrganizationApplianceSecurityEventsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithT0 adds the t0 to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithT0(t0 *string) *GetOrganizationApplianceSecurityEventsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithT1(t1 *string) *GetOrganizationApplianceSecurityEventsParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) WithTimespan(timespan *float32) *GetOrganizationApplianceSecurityEventsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get organization appliance security events params
func (o *GetOrganizationApplianceSecurityEventsParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationApplianceSecurityEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
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

	if o.T0 != nil {

		// query param t0
		var qrT0 string

		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {

			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}
	}

	if o.T1 != nil {

		// query param t1
		var qrT1 string

		if o.T1 != nil {
			qrT1 = *o.T1
		}
		qT1 := qrT1
		if qT1 != "" {

			if err := r.SetQueryParam("t1", qT1); err != nil {
				return err
			}
		}
	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float32

		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat32(qrTimespan)
		if qTimespan != "" {

			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
