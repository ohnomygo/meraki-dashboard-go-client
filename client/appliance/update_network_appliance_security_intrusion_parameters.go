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

// NewUpdateNetworkApplianceSecurityIntrusionParams creates a new UpdateNetworkApplianceSecurityIntrusionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceSecurityIntrusionParams() *UpdateNetworkApplianceSecurityIntrusionParams {
	return &UpdateNetworkApplianceSecurityIntrusionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceSecurityIntrusionParamsWithTimeout creates a new UpdateNetworkApplianceSecurityIntrusionParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceSecurityIntrusionParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceSecurityIntrusionParams {
	return &UpdateNetworkApplianceSecurityIntrusionParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceSecurityIntrusionParamsWithContext creates a new UpdateNetworkApplianceSecurityIntrusionParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceSecurityIntrusionParamsWithContext(ctx context.Context) *UpdateNetworkApplianceSecurityIntrusionParams {
	return &UpdateNetworkApplianceSecurityIntrusionParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceSecurityIntrusionParamsWithHTTPClient creates a new UpdateNetworkApplianceSecurityIntrusionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceSecurityIntrusionParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceSecurityIntrusionParams {
	return &UpdateNetworkApplianceSecurityIntrusionParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceSecurityIntrusionParams contains all the parameters to send to the API endpoint

	for the update network appliance security intrusion operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceSecurityIntrusionParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkApplianceSecurityIntrusion.
	UpdateNetworkApplianceSecurityIntrusion UpdateNetworkApplianceSecurityIntrusionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance security intrusion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceSecurityIntrusionParams) WithDefaults() *UpdateNetworkApplianceSecurityIntrusionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance security intrusion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceSecurityIntrusionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceSecurityIntrusionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) WithContext(ctx context.Context) *UpdateNetworkApplianceSecurityIntrusionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceSecurityIntrusionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) WithNetworkID(networkID string) *UpdateNetworkApplianceSecurityIntrusionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceSecurityIntrusion adds the updateNetworkApplianceSecurityIntrusion to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) WithUpdateNetworkApplianceSecurityIntrusion(updateNetworkApplianceSecurityIntrusion UpdateNetworkApplianceSecurityIntrusionBody) *UpdateNetworkApplianceSecurityIntrusionParams {
	o.SetUpdateNetworkApplianceSecurityIntrusion(updateNetworkApplianceSecurityIntrusion)
	return o
}

// SetUpdateNetworkApplianceSecurityIntrusion adds the updateNetworkApplianceSecurityIntrusion to the update network appliance security intrusion params
func (o *UpdateNetworkApplianceSecurityIntrusionParams) SetUpdateNetworkApplianceSecurityIntrusion(updateNetworkApplianceSecurityIntrusion UpdateNetworkApplianceSecurityIntrusionBody) {
	o.UpdateNetworkApplianceSecurityIntrusion = updateNetworkApplianceSecurityIntrusion
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceSecurityIntrusionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceSecurityIntrusion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
