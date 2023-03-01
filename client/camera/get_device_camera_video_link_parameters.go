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

// NewGetDeviceCameraVideoLinkParams creates a new GetDeviceCameraVideoLinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceCameraVideoLinkParams() *GetDeviceCameraVideoLinkParams {
	return &GetDeviceCameraVideoLinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceCameraVideoLinkParamsWithTimeout creates a new GetDeviceCameraVideoLinkParams object
// with the ability to set a timeout on a request.
func NewGetDeviceCameraVideoLinkParamsWithTimeout(timeout time.Duration) *GetDeviceCameraVideoLinkParams {
	return &GetDeviceCameraVideoLinkParams{
		timeout: timeout,
	}
}

// NewGetDeviceCameraVideoLinkParamsWithContext creates a new GetDeviceCameraVideoLinkParams object
// with the ability to set a context for a request.
func NewGetDeviceCameraVideoLinkParamsWithContext(ctx context.Context) *GetDeviceCameraVideoLinkParams {
	return &GetDeviceCameraVideoLinkParams{
		Context: ctx,
	}
}

// NewGetDeviceCameraVideoLinkParamsWithHTTPClient creates a new GetDeviceCameraVideoLinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceCameraVideoLinkParamsWithHTTPClient(client *http.Client) *GetDeviceCameraVideoLinkParams {
	return &GetDeviceCameraVideoLinkParams{
		HTTPClient: client,
	}
}

/*
GetDeviceCameraVideoLinkParams contains all the parameters to send to the API endpoint

	for the get device camera video link operation.

	Typically these are written to a http.Request.
*/
type GetDeviceCameraVideoLinkParams struct {

	// Serial.
	Serial string

	/* Timestamp.

	   [optional] The video link will start at this time. The timestamp should be a string in ISO8601 format. If no timestamp is specified, we will assume current time.

	   Format: date-time
	*/
	Timestamp *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device camera video link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraVideoLinkParams) WithDefaults() *GetDeviceCameraVideoLinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device camera video link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraVideoLinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) WithTimeout(timeout time.Duration) *GetDeviceCameraVideoLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) WithContext(ctx context.Context) *GetDeviceCameraVideoLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) WithHTTPClient(client *http.Client) *GetDeviceCameraVideoLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) WithSerial(serial string) *GetDeviceCameraVideoLinkParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithTimestamp adds the timestamp to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) WithTimestamp(timestamp *strfmt.DateTime) *GetDeviceCameraVideoLinkParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the get device camera video link params
func (o *GetDeviceCameraVideoLinkParams) SetTimestamp(timestamp *strfmt.DateTime) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceCameraVideoLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if o.Timestamp != nil {

		// query param timestamp
		var qrTimestamp strfmt.DateTime

		if o.Timestamp != nil {
			qrTimestamp = *o.Timestamp
		}
		qTimestamp := qrTimestamp.String()
		if qTimestamp != "" {

			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
