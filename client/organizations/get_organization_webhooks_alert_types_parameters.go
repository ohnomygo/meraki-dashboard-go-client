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

// NewGetOrganizationWebhooksAlertTypesParams creates a new GetOrganizationWebhooksAlertTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationWebhooksAlertTypesParams() *GetOrganizationWebhooksAlertTypesParams {
	return &GetOrganizationWebhooksAlertTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationWebhooksAlertTypesParamsWithTimeout creates a new GetOrganizationWebhooksAlertTypesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationWebhooksAlertTypesParamsWithTimeout(timeout time.Duration) *GetOrganizationWebhooksAlertTypesParams {
	return &GetOrganizationWebhooksAlertTypesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationWebhooksAlertTypesParamsWithContext creates a new GetOrganizationWebhooksAlertTypesParams object
// with the ability to set a context for a request.
func NewGetOrganizationWebhooksAlertTypesParamsWithContext(ctx context.Context) *GetOrganizationWebhooksAlertTypesParams {
	return &GetOrganizationWebhooksAlertTypesParams{
		Context: ctx,
	}
}

// NewGetOrganizationWebhooksAlertTypesParamsWithHTTPClient creates a new GetOrganizationWebhooksAlertTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationWebhooksAlertTypesParamsWithHTTPClient(client *http.Client) *GetOrganizationWebhooksAlertTypesParams {
	return &GetOrganizationWebhooksAlertTypesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationWebhooksAlertTypesParams contains all the parameters to send to the API endpoint

	for the get organization webhooks alert types operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationWebhooksAlertTypesParams struct {

	// OrganizationID.
	OrganizationID string

	/* ProductType.

	   Filter sample alerts to a specific product type
	*/
	ProductType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization webhooks alert types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationWebhooksAlertTypesParams) WithDefaults() *GetOrganizationWebhooksAlertTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization webhooks alert types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationWebhooksAlertTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) WithTimeout(timeout time.Duration) *GetOrganizationWebhooksAlertTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) WithContext(ctx context.Context) *GetOrganizationWebhooksAlertTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) WithHTTPClient(client *http.Client) *GetOrganizationWebhooksAlertTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) WithOrganizationID(organizationID string) *GetOrganizationWebhooksAlertTypesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProductType adds the productType to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) WithProductType(productType *string) *GetOrganizationWebhooksAlertTypesParams {
	o.SetProductType(productType)
	return o
}

// SetProductType adds the productType to the get organization webhooks alert types params
func (o *GetOrganizationWebhooksAlertTypesParams) SetProductType(productType *string) {
	o.ProductType = productType
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationWebhooksAlertTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.ProductType != nil {

		// query param productType
		var qrProductType string

		if o.ProductType != nil {
			qrProductType = *o.ProductType
		}
		qProductType := qrProductType
		if qProductType != "" {

			if err := r.SetQueryParam("productType", qProductType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
