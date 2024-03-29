// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// NewUpdateNetworkCameraWirelessProfileParams creates a new UpdateNetworkCameraWirelessProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkCameraWirelessProfileParams() *UpdateNetworkCameraWirelessProfileParams {
	return &UpdateNetworkCameraWirelessProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkCameraWirelessProfileParamsWithTimeout creates a new UpdateNetworkCameraWirelessProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkCameraWirelessProfileParamsWithTimeout(timeout time.Duration) *UpdateNetworkCameraWirelessProfileParams {
	return &UpdateNetworkCameraWirelessProfileParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkCameraWirelessProfileParamsWithContext creates a new UpdateNetworkCameraWirelessProfileParams object
// with the ability to set a context for a request.
func NewUpdateNetworkCameraWirelessProfileParamsWithContext(ctx context.Context) *UpdateNetworkCameraWirelessProfileParams {
	return &UpdateNetworkCameraWirelessProfileParams{
		Context: ctx,
	}
}

// NewUpdateNetworkCameraWirelessProfileParamsWithHTTPClient creates a new UpdateNetworkCameraWirelessProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkCameraWirelessProfileParamsWithHTTPClient(client *http.Client) *UpdateNetworkCameraWirelessProfileParams {
	return &UpdateNetworkCameraWirelessProfileParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkCameraWirelessProfileParams contains all the parameters to send to the API endpoint

	for the update network camera wireless profile operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkCameraWirelessProfileParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkCameraWirelessProfile.
	UpdateNetworkCameraWirelessProfile UpdateNetworkCameraWirelessProfileBody

	/* WirelessProfileID.

	   Wireless profile ID
	*/
	WirelessProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network camera wireless profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkCameraWirelessProfileParams) WithDefaults() *UpdateNetworkCameraWirelessProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network camera wireless profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkCameraWirelessProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) WithTimeout(timeout time.Duration) *UpdateNetworkCameraWirelessProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) WithContext(ctx context.Context) *UpdateNetworkCameraWirelessProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) WithHTTPClient(client *http.Client) *UpdateNetworkCameraWirelessProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) WithNetworkID(networkID string) *UpdateNetworkCameraWirelessProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkCameraWirelessProfile adds the updateNetworkCameraWirelessProfile to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) WithUpdateNetworkCameraWirelessProfile(updateNetworkCameraWirelessProfile UpdateNetworkCameraWirelessProfileBody) *UpdateNetworkCameraWirelessProfileParams {
	o.SetUpdateNetworkCameraWirelessProfile(updateNetworkCameraWirelessProfile)
	return o
}

// SetUpdateNetworkCameraWirelessProfile adds the updateNetworkCameraWirelessProfile to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) SetUpdateNetworkCameraWirelessProfile(updateNetworkCameraWirelessProfile UpdateNetworkCameraWirelessProfileBody) {
	o.UpdateNetworkCameraWirelessProfile = updateNetworkCameraWirelessProfile
}

// WithWirelessProfileID adds the wirelessProfileID to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) WithWirelessProfileID(wirelessProfileID string) *UpdateNetworkCameraWirelessProfileParams {
	o.SetWirelessProfileID(wirelessProfileID)
	return o
}

// SetWirelessProfileID adds the wirelessProfileId to the update network camera wireless profile params
func (o *UpdateNetworkCameraWirelessProfileParams) SetWirelessProfileID(wirelessProfileID string) {
	o.WirelessProfileID = wirelessProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkCameraWirelessProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkCameraWirelessProfile); err != nil {
		return err
	}

	// path param wirelessProfileId
	if err := r.SetPathParam("wirelessProfileId", o.WirelessProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
