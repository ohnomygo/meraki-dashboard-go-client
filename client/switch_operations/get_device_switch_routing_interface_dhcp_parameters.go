// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// NewGetDeviceSwitchRoutingInterfaceDhcpParams creates a new GetDeviceSwitchRoutingInterfaceDhcpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceSwitchRoutingInterfaceDhcpParams() *GetDeviceSwitchRoutingInterfaceDhcpParams {
	return &GetDeviceSwitchRoutingInterfaceDhcpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceSwitchRoutingInterfaceDhcpParamsWithTimeout creates a new GetDeviceSwitchRoutingInterfaceDhcpParams object
// with the ability to set a timeout on a request.
func NewGetDeviceSwitchRoutingInterfaceDhcpParamsWithTimeout(timeout time.Duration) *GetDeviceSwitchRoutingInterfaceDhcpParams {
	return &GetDeviceSwitchRoutingInterfaceDhcpParams{
		timeout: timeout,
	}
}

// NewGetDeviceSwitchRoutingInterfaceDhcpParamsWithContext creates a new GetDeviceSwitchRoutingInterfaceDhcpParams object
// with the ability to set a context for a request.
func NewGetDeviceSwitchRoutingInterfaceDhcpParamsWithContext(ctx context.Context) *GetDeviceSwitchRoutingInterfaceDhcpParams {
	return &GetDeviceSwitchRoutingInterfaceDhcpParams{
		Context: ctx,
	}
}

// NewGetDeviceSwitchRoutingInterfaceDhcpParamsWithHTTPClient creates a new GetDeviceSwitchRoutingInterfaceDhcpParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceSwitchRoutingInterfaceDhcpParamsWithHTTPClient(client *http.Client) *GetDeviceSwitchRoutingInterfaceDhcpParams {
	return &GetDeviceSwitchRoutingInterfaceDhcpParams{
		HTTPClient: client,
	}
}

/*
GetDeviceSwitchRoutingInterfaceDhcpParams contains all the parameters to send to the API endpoint

	for the get device switch routing interface dhcp operation.

	Typically these are written to a http.Request.
*/
type GetDeviceSwitchRoutingInterfaceDhcpParams struct {

	/* InterfaceID.

	   Interface ID
	*/
	InterfaceID string

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device switch routing interface dhcp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) WithDefaults() *GetDeviceSwitchRoutingInterfaceDhcpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device switch routing interface dhcp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) WithTimeout(timeout time.Duration) *GetDeviceSwitchRoutingInterfaceDhcpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) WithContext(ctx context.Context) *GetDeviceSwitchRoutingInterfaceDhcpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) WithHTTPClient(client *http.Client) *GetDeviceSwitchRoutingInterfaceDhcpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInterfaceID adds the interfaceID to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) WithInterfaceID(interfaceID string) *GetDeviceSwitchRoutingInterfaceDhcpParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) SetInterfaceID(interfaceID string) {
	o.InterfaceID = interfaceID
}

// WithSerial adds the serial to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) WithSerial(serial string) *GetDeviceSwitchRoutingInterfaceDhcpParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device switch routing interface dhcp params
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceSwitchRoutingInterfaceDhcpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param interfaceId
	if err := r.SetPathParam("interfaceId", o.InterfaceID); err != nil {
		return err
	}

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}