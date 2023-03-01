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

// NewAssignOrganizationLicensesSeatsParams creates a new AssignOrganizationLicensesSeatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignOrganizationLicensesSeatsParams() *AssignOrganizationLicensesSeatsParams {
	return &AssignOrganizationLicensesSeatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignOrganizationLicensesSeatsParamsWithTimeout creates a new AssignOrganizationLicensesSeatsParams object
// with the ability to set a timeout on a request.
func NewAssignOrganizationLicensesSeatsParamsWithTimeout(timeout time.Duration) *AssignOrganizationLicensesSeatsParams {
	return &AssignOrganizationLicensesSeatsParams{
		timeout: timeout,
	}
}

// NewAssignOrganizationLicensesSeatsParamsWithContext creates a new AssignOrganizationLicensesSeatsParams object
// with the ability to set a context for a request.
func NewAssignOrganizationLicensesSeatsParamsWithContext(ctx context.Context) *AssignOrganizationLicensesSeatsParams {
	return &AssignOrganizationLicensesSeatsParams{
		Context: ctx,
	}
}

// NewAssignOrganizationLicensesSeatsParamsWithHTTPClient creates a new AssignOrganizationLicensesSeatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignOrganizationLicensesSeatsParamsWithHTTPClient(client *http.Client) *AssignOrganizationLicensesSeatsParams {
	return &AssignOrganizationLicensesSeatsParams{
		HTTPClient: client,
	}
}

/*
AssignOrganizationLicensesSeatsParams contains all the parameters to send to the API endpoint

	for the assign organization licenses seats operation.

	Typically these are written to a http.Request.
*/
type AssignOrganizationLicensesSeatsParams struct {

	// AssignOrganizationLicensesSeats.
	AssignOrganizationLicensesSeats AssignOrganizationLicensesSeatsBody

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign organization licenses seats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignOrganizationLicensesSeatsParams) WithDefaults() *AssignOrganizationLicensesSeatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign organization licenses seats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignOrganizationLicensesSeatsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) WithTimeout(timeout time.Duration) *AssignOrganizationLicensesSeatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) WithContext(ctx context.Context) *AssignOrganizationLicensesSeatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) WithHTTPClient(client *http.Client) *AssignOrganizationLicensesSeatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssignOrganizationLicensesSeats adds the assignOrganizationLicensesSeats to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) WithAssignOrganizationLicensesSeats(assignOrganizationLicensesSeats AssignOrganizationLicensesSeatsBody) *AssignOrganizationLicensesSeatsParams {
	o.SetAssignOrganizationLicensesSeats(assignOrganizationLicensesSeats)
	return o
}

// SetAssignOrganizationLicensesSeats adds the assignOrganizationLicensesSeats to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) SetAssignOrganizationLicensesSeats(assignOrganizationLicensesSeats AssignOrganizationLicensesSeatsBody) {
	o.AssignOrganizationLicensesSeats = assignOrganizationLicensesSeats
}

// WithOrganizationID adds the organizationID to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) WithOrganizationID(organizationID string) *AssignOrganizationLicensesSeatsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the assign organization licenses seats params
func (o *AssignOrganizationLicensesSeatsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *AssignOrganizationLicensesSeatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.AssignOrganizationLicensesSeats); err != nil {
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
