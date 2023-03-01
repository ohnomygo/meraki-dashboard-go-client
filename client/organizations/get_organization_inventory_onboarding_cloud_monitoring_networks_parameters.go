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

// NewGetOrganizationInventoryOnboardingCloudMonitoringNetworksParams creates a new GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationInventoryOnboardingCloudMonitoringNetworksParams() *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	return &GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationInventoryOnboardingCloudMonitoringNetworksParamsWithTimeout creates a new GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationInventoryOnboardingCloudMonitoringNetworksParamsWithTimeout(timeout time.Duration) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	return &GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams{
		timeout: timeout,
	}
}

// NewGetOrganizationInventoryOnboardingCloudMonitoringNetworksParamsWithContext creates a new GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams object
// with the ability to set a context for a request.
func NewGetOrganizationInventoryOnboardingCloudMonitoringNetworksParamsWithContext(ctx context.Context) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	return &GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams{
		Context: ctx,
	}
}

// NewGetOrganizationInventoryOnboardingCloudMonitoringNetworksParamsWithHTTPClient creates a new GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationInventoryOnboardingCloudMonitoringNetworksParamsWithHTTPClient(client *http.Client) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	return &GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams contains all the parameters to send to the API endpoint

	for the get organization inventory onboarding cloud monitoring networks operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams struct {

	/* DeviceType.

	   Device Type switch or wireless controller
	*/
	DeviceType string

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	// OrganizationID.
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000.
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

// WithDefaults hydrates default values in the get organization inventory onboarding cloud monitoring networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithDefaults() *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization inventory onboarding cloud monitoring networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithTimeout(timeout time.Duration) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithContext(ctx context.Context) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithHTTPClient(client *http.Client) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceType adds the deviceType to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithDeviceType(deviceType string) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetDeviceType(deviceType)
	return o
}

// SetDeviceType adds the deviceType to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetDeviceType(deviceType string) {
	o.DeviceType = deviceType
}

// WithEndingBefore adds the endingBefore to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithEndingBefore(endingBefore *string) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithOrganizationID adds the organizationID to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithOrganizationID(organizationID string) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithPerPage(perPage *int64) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WithStartingAfter(startingAfter *string) *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization inventory onboarding cloud monitoring networks params
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationInventoryOnboardingCloudMonitoringNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param deviceType
	qrDeviceType := o.DeviceType
	qDeviceType := qrDeviceType
	if qDeviceType != "" {

		if err := r.SetQueryParam("deviceType", qDeviceType); err != nil {
			return err
		}
	}

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
