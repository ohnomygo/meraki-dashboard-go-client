// Code generated by go-swagger; DO NOT EDIT.

package sensor

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

// NewCreateNetworkSensorAlertsProfileParams creates a new CreateNetworkSensorAlertsProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkSensorAlertsProfileParams() *CreateNetworkSensorAlertsProfileParams {
	return &CreateNetworkSensorAlertsProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkSensorAlertsProfileParamsWithTimeout creates a new CreateNetworkSensorAlertsProfileParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkSensorAlertsProfileParamsWithTimeout(timeout time.Duration) *CreateNetworkSensorAlertsProfileParams {
	return &CreateNetworkSensorAlertsProfileParams{
		timeout: timeout,
	}
}

// NewCreateNetworkSensorAlertsProfileParamsWithContext creates a new CreateNetworkSensorAlertsProfileParams object
// with the ability to set a context for a request.
func NewCreateNetworkSensorAlertsProfileParamsWithContext(ctx context.Context) *CreateNetworkSensorAlertsProfileParams {
	return &CreateNetworkSensorAlertsProfileParams{
		Context: ctx,
	}
}

// NewCreateNetworkSensorAlertsProfileParamsWithHTTPClient creates a new CreateNetworkSensorAlertsProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkSensorAlertsProfileParamsWithHTTPClient(client *http.Client) *CreateNetworkSensorAlertsProfileParams {
	return &CreateNetworkSensorAlertsProfileParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkSensorAlertsProfileParams contains all the parameters to send to the API endpoint

	for the create network sensor alerts profile operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkSensorAlertsProfileParams struct {

	// CreateNetworkSensorAlertsProfile.
	CreateNetworkSensorAlertsProfile CreateNetworkSensorAlertsProfileBody

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network sensor alerts profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSensorAlertsProfileParams) WithDefaults() *CreateNetworkSensorAlertsProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network sensor alerts profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSensorAlertsProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) WithTimeout(timeout time.Duration) *CreateNetworkSensorAlertsProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) WithContext(ctx context.Context) *CreateNetworkSensorAlertsProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) WithHTTPClient(client *http.Client) *CreateNetworkSensorAlertsProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkSensorAlertsProfile adds the createNetworkSensorAlertsProfile to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) WithCreateNetworkSensorAlertsProfile(createNetworkSensorAlertsProfile CreateNetworkSensorAlertsProfileBody) *CreateNetworkSensorAlertsProfileParams {
	o.SetCreateNetworkSensorAlertsProfile(createNetworkSensorAlertsProfile)
	return o
}

// SetCreateNetworkSensorAlertsProfile adds the createNetworkSensorAlertsProfile to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) SetCreateNetworkSensorAlertsProfile(createNetworkSensorAlertsProfile CreateNetworkSensorAlertsProfileBody) {
	o.CreateNetworkSensorAlertsProfile = createNetworkSensorAlertsProfile
}

// WithNetworkID adds the networkID to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) WithNetworkID(networkID string) *CreateNetworkSensorAlertsProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network sensor alerts profile params
func (o *CreateNetworkSensorAlertsProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkSensorAlertsProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkSensorAlertsProfile); err != nil {
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
