// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewDeleteNetworkFirmwareUpgradesStagedGroupParams creates a new DeleteNetworkFirmwareUpgradesStagedGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkFirmwareUpgradesStagedGroupParams() *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	return &DeleteNetworkFirmwareUpgradesStagedGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkFirmwareUpgradesStagedGroupParamsWithTimeout creates a new DeleteNetworkFirmwareUpgradesStagedGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkFirmwareUpgradesStagedGroupParamsWithTimeout(timeout time.Duration) *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	return &DeleteNetworkFirmwareUpgradesStagedGroupParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkFirmwareUpgradesStagedGroupParamsWithContext creates a new DeleteNetworkFirmwareUpgradesStagedGroupParams object
// with the ability to set a context for a request.
func NewDeleteNetworkFirmwareUpgradesStagedGroupParamsWithContext(ctx context.Context) *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	return &DeleteNetworkFirmwareUpgradesStagedGroupParams{
		Context: ctx,
	}
}

// NewDeleteNetworkFirmwareUpgradesStagedGroupParamsWithHTTPClient creates a new DeleteNetworkFirmwareUpgradesStagedGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkFirmwareUpgradesStagedGroupParamsWithHTTPClient(client *http.Client) *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	return &DeleteNetworkFirmwareUpgradesStagedGroupParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkFirmwareUpgradesStagedGroupParams contains all the parameters to send to the API endpoint

	for the delete network firmware upgrades staged group operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkFirmwareUpgradesStagedGroupParams struct {

	// GroupID.
	GroupID string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network firmware upgrades staged group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) WithDefaults() *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network firmware upgrades staged group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) WithTimeout(timeout time.Duration) *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) WithContext(ctx context.Context) *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) WithHTTPClient(client *http.Client) *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) WithGroupID(groupID string) *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithNetworkID adds the networkID to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) WithNetworkID(networkID string) *DeleteNetworkFirmwareUpgradesStagedGroupParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network firmware upgrades staged group params
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkFirmwareUpgradesStagedGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
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