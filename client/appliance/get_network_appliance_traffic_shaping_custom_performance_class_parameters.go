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

// NewGetNetworkApplianceTrafficShapingCustomPerformanceClassParams creates a new GetNetworkApplianceTrafficShapingCustomPerformanceClassParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkApplianceTrafficShapingCustomPerformanceClassParams() *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	return &GetNetworkApplianceTrafficShapingCustomPerformanceClassParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithTimeout creates a new GetNetworkApplianceTrafficShapingCustomPerformanceClassParams object
// with the ability to set a timeout on a request.
func NewGetNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithTimeout(timeout time.Duration) *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	return &GetNetworkApplianceTrafficShapingCustomPerformanceClassParams{
		timeout: timeout,
	}
}

// NewGetNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithContext creates a new GetNetworkApplianceTrafficShapingCustomPerformanceClassParams object
// with the ability to set a context for a request.
func NewGetNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithContext(ctx context.Context) *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	return &GetNetworkApplianceTrafficShapingCustomPerformanceClassParams{
		Context: ctx,
	}
}

// NewGetNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithHTTPClient creates a new GetNetworkApplianceTrafficShapingCustomPerformanceClassParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithHTTPClient(client *http.Client) *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	return &GetNetworkApplianceTrafficShapingCustomPerformanceClassParams{
		HTTPClient: client,
	}
}

/*
GetNetworkApplianceTrafficShapingCustomPerformanceClassParams contains all the parameters to send to the API endpoint

	for the get network appliance traffic shaping custom performance class operation.

	Typically these are written to a http.Request.
*/
type GetNetworkApplianceTrafficShapingCustomPerformanceClassParams struct {

	// CustomPerformanceClassID.
	CustomPerformanceClassID string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network appliance traffic shaping custom performance class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithDefaults() *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network appliance traffic shaping custom performance class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithTimeout(timeout time.Duration) *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithContext(ctx context.Context) *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithHTTPClient(client *http.Client) *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomPerformanceClassID adds the customPerformanceClassID to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithCustomPerformanceClassID(customPerformanceClassID string) *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetCustomPerformanceClassID(customPerformanceClassID)
	return o
}

// SetCustomPerformanceClassID adds the customPerformanceClassId to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetCustomPerformanceClassID(customPerformanceClassID string) {
	o.CustomPerformanceClassID = customPerformanceClassID
}

// WithNetworkID adds the networkID to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithNetworkID(networkID string) *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network appliance traffic shaping custom performance class params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customPerformanceClassId
	if err := r.SetPathParam("customPerformanceClassId", o.CustomPerformanceClassID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
