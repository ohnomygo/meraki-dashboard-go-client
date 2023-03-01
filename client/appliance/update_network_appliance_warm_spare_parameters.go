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
)

// NewUpdateNetworkApplianceWarmSpareParams creates a new UpdateNetworkApplianceWarmSpareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceWarmSpareParams() *UpdateNetworkApplianceWarmSpareParams {
	return &UpdateNetworkApplianceWarmSpareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceWarmSpareParamsWithTimeout creates a new UpdateNetworkApplianceWarmSpareParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceWarmSpareParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceWarmSpareParams {
	return &UpdateNetworkApplianceWarmSpareParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceWarmSpareParamsWithContext creates a new UpdateNetworkApplianceWarmSpareParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceWarmSpareParamsWithContext(ctx context.Context) *UpdateNetworkApplianceWarmSpareParams {
	return &UpdateNetworkApplianceWarmSpareParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceWarmSpareParamsWithHTTPClient creates a new UpdateNetworkApplianceWarmSpareParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceWarmSpareParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceWarmSpareParams {
	return &UpdateNetworkApplianceWarmSpareParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceWarmSpareParams contains all the parameters to send to the API endpoint

	for the update network appliance warm spare operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceWarmSpareParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkApplianceWarmSpare.
	UpdateNetworkApplianceWarmSpare UpdateNetworkApplianceWarmSpareBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance warm spare params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceWarmSpareParams) WithDefaults() *UpdateNetworkApplianceWarmSpareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance warm spare params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceWarmSpareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceWarmSpareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) WithContext(ctx context.Context) *UpdateNetworkApplianceWarmSpareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceWarmSpareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) WithNetworkID(networkID string) *UpdateNetworkApplianceWarmSpareParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceWarmSpare adds the updateNetworkApplianceWarmSpare to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) WithUpdateNetworkApplianceWarmSpare(updateNetworkApplianceWarmSpare UpdateNetworkApplianceWarmSpareBody) *UpdateNetworkApplianceWarmSpareParams {
	o.SetUpdateNetworkApplianceWarmSpare(updateNetworkApplianceWarmSpare)
	return o
}

// SetUpdateNetworkApplianceWarmSpare adds the updateNetworkApplianceWarmSpare to the update network appliance warm spare params
func (o *UpdateNetworkApplianceWarmSpareParams) SetUpdateNetworkApplianceWarmSpare(updateNetworkApplianceWarmSpare UpdateNetworkApplianceWarmSpareBody) {
	o.UpdateNetworkApplianceWarmSpare = updateNetworkApplianceWarmSpare
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceWarmSpareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceWarmSpare); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
