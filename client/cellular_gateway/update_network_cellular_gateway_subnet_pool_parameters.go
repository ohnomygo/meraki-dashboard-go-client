// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

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

// NewUpdateNetworkCellularGatewaySubnetPoolParams creates a new UpdateNetworkCellularGatewaySubnetPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkCellularGatewaySubnetPoolParams() *UpdateNetworkCellularGatewaySubnetPoolParams {
	return &UpdateNetworkCellularGatewaySubnetPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkCellularGatewaySubnetPoolParamsWithTimeout creates a new UpdateNetworkCellularGatewaySubnetPoolParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkCellularGatewaySubnetPoolParamsWithTimeout(timeout time.Duration) *UpdateNetworkCellularGatewaySubnetPoolParams {
	return &UpdateNetworkCellularGatewaySubnetPoolParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkCellularGatewaySubnetPoolParamsWithContext creates a new UpdateNetworkCellularGatewaySubnetPoolParams object
// with the ability to set a context for a request.
func NewUpdateNetworkCellularGatewaySubnetPoolParamsWithContext(ctx context.Context) *UpdateNetworkCellularGatewaySubnetPoolParams {
	return &UpdateNetworkCellularGatewaySubnetPoolParams{
		Context: ctx,
	}
}

// NewUpdateNetworkCellularGatewaySubnetPoolParamsWithHTTPClient creates a new UpdateNetworkCellularGatewaySubnetPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkCellularGatewaySubnetPoolParamsWithHTTPClient(client *http.Client) *UpdateNetworkCellularGatewaySubnetPoolParams {
	return &UpdateNetworkCellularGatewaySubnetPoolParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkCellularGatewaySubnetPoolParams contains all the parameters to send to the API endpoint

	for the update network cellular gateway subnet pool operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkCellularGatewaySubnetPoolParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetworkCellularGatewaySubnetPool.
	UpdateNetworkCellularGatewaySubnetPool UpdateNetworkCellularGatewaySubnetPoolBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network cellular gateway subnet pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) WithDefaults() *UpdateNetworkCellularGatewaySubnetPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network cellular gateway subnet pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) WithTimeout(timeout time.Duration) *UpdateNetworkCellularGatewaySubnetPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) WithContext(ctx context.Context) *UpdateNetworkCellularGatewaySubnetPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) WithHTTPClient(client *http.Client) *UpdateNetworkCellularGatewaySubnetPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) WithNetworkID(networkID string) *UpdateNetworkCellularGatewaySubnetPoolParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkCellularGatewaySubnetPool adds the updateNetworkCellularGatewaySubnetPool to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) WithUpdateNetworkCellularGatewaySubnetPool(updateNetworkCellularGatewaySubnetPool UpdateNetworkCellularGatewaySubnetPoolBody) *UpdateNetworkCellularGatewaySubnetPoolParams {
	o.SetUpdateNetworkCellularGatewaySubnetPool(updateNetworkCellularGatewaySubnetPool)
	return o
}

// SetUpdateNetworkCellularGatewaySubnetPool adds the updateNetworkCellularGatewaySubnetPool to the update network cellular gateway subnet pool params
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) SetUpdateNetworkCellularGatewaySubnetPool(updateNetworkCellularGatewaySubnetPool UpdateNetworkCellularGatewaySubnetPoolBody) {
	o.UpdateNetworkCellularGatewaySubnetPool = updateNetworkCellularGatewaySubnetPool
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkCellularGatewaySubnetPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkCellularGatewaySubnetPool); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
