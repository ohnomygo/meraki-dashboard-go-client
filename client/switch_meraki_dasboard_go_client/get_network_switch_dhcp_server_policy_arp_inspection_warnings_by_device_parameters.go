// Code generated by go-swagger; DO NOT EDIT.

package switch_meraki_dasboard_go_client

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
	"github.com/go-openapi/swag"
)

// NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams creates a new GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams() *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	return &GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParamsWithTimeout creates a new GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	return &GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams{
		timeout: timeout,
	}
}

// NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParamsWithContext creates a new GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams object
// with the ability to set a context for a request.
func NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParamsWithContext(ctx context.Context) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	return &GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams{
		Context: ctx,
	}
}

// NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParamsWithHTTPClient creates a new GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	return &GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams contains all the parameters to send to the API endpoint

	for the get network switch dhcp server policy arp inspection warnings by device operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	// NetworkID.
	NetworkID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	*/
	PerPage *int64

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network switch dhcp server policy arp inspection warnings by device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WithDefaults() *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network switch dhcp server policy arp inspection warnings by device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WithContext(ctx context.Context) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WithEndingBefore(endingBefore *string) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithNetworkID adds the networkID to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WithNetworkID(networkID string) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPerPage adds the perPage to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WithPerPage(perPage *int64) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WithStartingAfter(startingAfter *string) *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network switch dhcp server policy arp inspection warnings by device params
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingBefore != nil {

		// query param endingBefore
		var qrEndingBefore string

		if o.EndingBefore != nil {
			qrEndingBefore = *o.EndingBefore
		}
		qEndingBefore := qrEndingBefore
		if qEndingBefore != "" {

			if err := r.SetQueryParam("endingBefore", qEndingBefore); err != nil {
				return err
			}
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.StartingAfter != nil {

		// query param startingAfter
		var qrStartingAfter string

		if o.StartingAfter != nil {
			qrStartingAfter = *o.StartingAfter
		}
		qStartingAfter := qrStartingAfter
		if qStartingAfter != "" {

			if err := r.SetQueryParam("startingAfter", qStartingAfter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
