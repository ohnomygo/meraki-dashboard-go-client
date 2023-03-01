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

// NewGetOrganizationSnmpParams creates a new GetOrganizationSnmpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationSnmpParams() *GetOrganizationSnmpParams {
	return &GetOrganizationSnmpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationSnmpParamsWithTimeout creates a new GetOrganizationSnmpParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationSnmpParamsWithTimeout(timeout time.Duration) *GetOrganizationSnmpParams {
	return &GetOrganizationSnmpParams{
		timeout: timeout,
	}
}

// NewGetOrganizationSnmpParamsWithContext creates a new GetOrganizationSnmpParams object
// with the ability to set a context for a request.
func NewGetOrganizationSnmpParamsWithContext(ctx context.Context) *GetOrganizationSnmpParams {
	return &GetOrganizationSnmpParams{
		Context: ctx,
	}
}

// NewGetOrganizationSnmpParamsWithHTTPClient creates a new GetOrganizationSnmpParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationSnmpParamsWithHTTPClient(client *http.Client) *GetOrganizationSnmpParams {
	return &GetOrganizationSnmpParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationSnmpParams contains all the parameters to send to the API endpoint

	for the get organization snmp operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationSnmpParams struct {

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization snmp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSnmpParams) WithDefaults() *GetOrganizationSnmpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization snmp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSnmpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization snmp params
func (o *GetOrganizationSnmpParams) WithTimeout(timeout time.Duration) *GetOrganizationSnmpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization snmp params
func (o *GetOrganizationSnmpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization snmp params
func (o *GetOrganizationSnmpParams) WithContext(ctx context.Context) *GetOrganizationSnmpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization snmp params
func (o *GetOrganizationSnmpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization snmp params
func (o *GetOrganizationSnmpParams) WithHTTPClient(client *http.Client) *GetOrganizationSnmpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization snmp params
func (o *GetOrganizationSnmpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization snmp params
func (o *GetOrganizationSnmpParams) WithOrganizationID(organizationID string) *GetOrganizationSnmpParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization snmp params
func (o *GetOrganizationSnmpParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationSnmpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
