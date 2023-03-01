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

// NewUpdateNetworkApplianceStaticRouteParams creates a new UpdateNetworkApplianceStaticRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceStaticRouteParams() *UpdateNetworkApplianceStaticRouteParams {
	return &UpdateNetworkApplianceStaticRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceStaticRouteParamsWithTimeout creates a new UpdateNetworkApplianceStaticRouteParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceStaticRouteParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceStaticRouteParams {
	return &UpdateNetworkApplianceStaticRouteParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceStaticRouteParamsWithContext creates a new UpdateNetworkApplianceStaticRouteParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceStaticRouteParamsWithContext(ctx context.Context) *UpdateNetworkApplianceStaticRouteParams {
	return &UpdateNetworkApplianceStaticRouteParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceStaticRouteParamsWithHTTPClient creates a new UpdateNetworkApplianceStaticRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceStaticRouteParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceStaticRouteParams {
	return &UpdateNetworkApplianceStaticRouteParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceStaticRouteParams contains all the parameters to send to the API endpoint

	for the update network appliance static route operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceStaticRouteParams struct {

	// NetworkID.
	NetworkID string

	// StaticRouteID.
	StaticRouteID string

	// UpdateNetworkApplianceStaticRoute.
	UpdateNetworkApplianceStaticRoute UpdateNetworkApplianceStaticRouteBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceStaticRouteParams) WithDefaults() *UpdateNetworkApplianceStaticRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceStaticRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceStaticRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) WithContext(ctx context.Context) *UpdateNetworkApplianceStaticRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceStaticRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) WithNetworkID(networkID string) *UpdateNetworkApplianceStaticRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithStaticRouteID adds the staticRouteID to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) WithStaticRouteID(staticRouteID string) *UpdateNetworkApplianceStaticRouteParams {
	o.SetStaticRouteID(staticRouteID)
	return o
}

// SetStaticRouteID adds the staticRouteId to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) SetStaticRouteID(staticRouteID string) {
	o.StaticRouteID = staticRouteID
}

// WithUpdateNetworkApplianceStaticRoute adds the updateNetworkApplianceStaticRoute to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) WithUpdateNetworkApplianceStaticRoute(updateNetworkApplianceStaticRoute UpdateNetworkApplianceStaticRouteBody) *UpdateNetworkApplianceStaticRouteParams {
	o.SetUpdateNetworkApplianceStaticRoute(updateNetworkApplianceStaticRoute)
	return o
}

// SetUpdateNetworkApplianceStaticRoute adds the updateNetworkApplianceStaticRoute to the update network appliance static route params
func (o *UpdateNetworkApplianceStaticRouteParams) SetUpdateNetworkApplianceStaticRoute(updateNetworkApplianceStaticRoute UpdateNetworkApplianceStaticRouteBody) {
	o.UpdateNetworkApplianceStaticRoute = updateNetworkApplianceStaticRoute
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceStaticRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param staticRouteId
	if err := r.SetPathParam("staticRouteId", o.StaticRouteID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceStaticRoute); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
