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

// NewUpdateDeviceCameraWirelessProfilesParams creates a new UpdateDeviceCameraWirelessProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceCameraWirelessProfilesParams() *UpdateDeviceCameraWirelessProfilesParams {
	return &UpdateDeviceCameraWirelessProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceCameraWirelessProfilesParamsWithTimeout creates a new UpdateDeviceCameraWirelessProfilesParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceCameraWirelessProfilesParamsWithTimeout(timeout time.Duration) *UpdateDeviceCameraWirelessProfilesParams {
	return &UpdateDeviceCameraWirelessProfilesParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceCameraWirelessProfilesParamsWithContext creates a new UpdateDeviceCameraWirelessProfilesParams object
// with the ability to set a context for a request.
func NewUpdateDeviceCameraWirelessProfilesParamsWithContext(ctx context.Context) *UpdateDeviceCameraWirelessProfilesParams {
	return &UpdateDeviceCameraWirelessProfilesParams{
		Context: ctx,
	}
}

// NewUpdateDeviceCameraWirelessProfilesParamsWithHTTPClient creates a new UpdateDeviceCameraWirelessProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceCameraWirelessProfilesParamsWithHTTPClient(client *http.Client) *UpdateDeviceCameraWirelessProfilesParams {
	return &UpdateDeviceCameraWirelessProfilesParams{
		HTTPClient: client,
	}
}

/*
UpdateDeviceCameraWirelessProfilesParams contains all the parameters to send to the API endpoint

	for the update device camera wireless profiles operation.

	Typically these are written to a http.Request.
*/
type UpdateDeviceCameraWirelessProfilesParams struct {

	// Serial.
	Serial string

	// UpdateDeviceCameraWirelessProfiles.
	UpdateDeviceCameraWirelessProfiles UpdateDeviceCameraWirelessProfilesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device camera wireless profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceCameraWirelessProfilesParams) WithDefaults() *UpdateDeviceCameraWirelessProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device camera wireless profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceCameraWirelessProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) WithTimeout(timeout time.Duration) *UpdateDeviceCameraWirelessProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) WithContext(ctx context.Context) *UpdateDeviceCameraWirelessProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) WithHTTPClient(client *http.Client) *UpdateDeviceCameraWirelessProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) WithSerial(serial string) *UpdateDeviceCameraWirelessProfilesParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceCameraWirelessProfiles adds the updateDeviceCameraWirelessProfiles to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) WithUpdateDeviceCameraWirelessProfiles(updateDeviceCameraWirelessProfiles UpdateDeviceCameraWirelessProfilesBody) *UpdateDeviceCameraWirelessProfilesParams {
	o.SetUpdateDeviceCameraWirelessProfiles(updateDeviceCameraWirelessProfiles)
	return o
}

// SetUpdateDeviceCameraWirelessProfiles adds the updateDeviceCameraWirelessProfiles to the update device camera wireless profiles params
func (o *UpdateDeviceCameraWirelessProfilesParams) SetUpdateDeviceCameraWirelessProfiles(updateDeviceCameraWirelessProfiles UpdateDeviceCameraWirelessProfilesBody) {
	o.UpdateDeviceCameraWirelessProfiles = updateDeviceCameraWirelessProfiles
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceCameraWirelessProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateDeviceCameraWirelessProfiles); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
