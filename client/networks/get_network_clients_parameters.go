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
	"github.com/go-openapi/swag"
)

// NewGetNetworkClientsParams creates a new GetNetworkClientsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkClientsParams() *GetNetworkClientsParams {
	return &GetNetworkClientsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkClientsParamsWithTimeout creates a new GetNetworkClientsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkClientsParamsWithTimeout(timeout time.Duration) *GetNetworkClientsParams {
	return &GetNetworkClientsParams{
		timeout: timeout,
	}
}

// NewGetNetworkClientsParamsWithContext creates a new GetNetworkClientsParams object
// with the ability to set a context for a request.
func NewGetNetworkClientsParamsWithContext(ctx context.Context) *GetNetworkClientsParams {
	return &GetNetworkClientsParams{
		Context: ctx,
	}
}

// NewGetNetworkClientsParamsWithHTTPClient creates a new GetNetworkClientsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkClientsParamsWithHTTPClient(client *http.Client) *GetNetworkClientsParams {
	return &GetNetworkClientsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkClientsParams contains all the parameters to send to the API endpoint

	for the get network clients operation.

	Typically these are written to a http.Request.
*/
type GetNetworkClientsParams struct {

	/* Description.

	   Filters clients based on a partial or full match for the description field.
	*/
	Description *string

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* IP.

	   Filters clients based on a partial or full match for the ip address field.
	*/
	IP *string

	/* Ip6.

	   Filters clients based on a partial or full match for the ip6 address field.
	*/
	Ip6 *string

	/* Ip6Local.

	   Filters clients based on a partial or full match for the ip6Local address field.
	*/
	Ip6Local *string

	/* Mac.

	   Filters clients based on a partial or full match for the mac address field.
	*/
	Mac *string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Os.

	   Filters clients based on a partial or full match for the os (operating system) field.
	*/
	Os *string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10.
	*/
	PerPage *int64

	/* RecentDeviceConnections.

	   Filters clients based on recent connection type. Can be one of 'Wired' or 'Wireless'.
	*/
	RecentDeviceConnections []string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* Statuses.

	   Filters clients based on status. Can be one of 'Online' or 'Offline'.
	*/
	Statuses []string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	*/
	T0 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.

	   Format: float
	*/
	Timespan *float32

	/* Vlan.

	   Filters clients based on the full match for the VLAN field.
	*/
	Vlan *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkClientsParams) WithDefaults() *GetNetworkClientsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkClientsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network clients params
func (o *GetNetworkClientsParams) WithTimeout(timeout time.Duration) *GetNetworkClientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network clients params
func (o *GetNetworkClientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network clients params
func (o *GetNetworkClientsParams) WithContext(ctx context.Context) *GetNetworkClientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network clients params
func (o *GetNetworkClientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network clients params
func (o *GetNetworkClientsParams) WithHTTPClient(client *http.Client) *GetNetworkClientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network clients params
func (o *GetNetworkClientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the get network clients params
func (o *GetNetworkClientsParams) WithDescription(description *string) *GetNetworkClientsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get network clients params
func (o *GetNetworkClientsParams) SetDescription(description *string) {
	o.Description = description
}

// WithEndingBefore adds the endingBefore to the get network clients params
func (o *GetNetworkClientsParams) WithEndingBefore(endingBefore *string) *GetNetworkClientsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network clients params
func (o *GetNetworkClientsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithIP adds the ip to the get network clients params
func (o *GetNetworkClientsParams) WithIP(ip *string) *GetNetworkClientsParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get network clients params
func (o *GetNetworkClientsParams) SetIP(ip *string) {
	o.IP = ip
}

// WithIp6 adds the ip6 to the get network clients params
func (o *GetNetworkClientsParams) WithIp6(ip6 *string) *GetNetworkClientsParams {
	o.SetIp6(ip6)
	return o
}

// SetIp6 adds the ip6 to the get network clients params
func (o *GetNetworkClientsParams) SetIp6(ip6 *string) {
	o.Ip6 = ip6
}

// WithIp6Local adds the ip6Local to the get network clients params
func (o *GetNetworkClientsParams) WithIp6Local(ip6Local *string) *GetNetworkClientsParams {
	o.SetIp6Local(ip6Local)
	return o
}

// SetIp6Local adds the ip6Local to the get network clients params
func (o *GetNetworkClientsParams) SetIp6Local(ip6Local *string) {
	o.Ip6Local = ip6Local
}

// WithMac adds the mac to the get network clients params
func (o *GetNetworkClientsParams) WithMac(mac *string) *GetNetworkClientsParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the get network clients params
func (o *GetNetworkClientsParams) SetMac(mac *string) {
	o.Mac = mac
}

// WithNetworkID adds the networkID to the get network clients params
func (o *GetNetworkClientsParams) WithNetworkID(networkID string) *GetNetworkClientsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network clients params
func (o *GetNetworkClientsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithOs adds the os to the get network clients params
func (o *GetNetworkClientsParams) WithOs(os *string) *GetNetworkClientsParams {
	o.SetOs(os)
	return o
}

// SetOs adds the os to the get network clients params
func (o *GetNetworkClientsParams) SetOs(os *string) {
	o.Os = os
}

// WithPerPage adds the perPage to the get network clients params
func (o *GetNetworkClientsParams) WithPerPage(perPage *int64) *GetNetworkClientsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network clients params
func (o *GetNetworkClientsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithRecentDeviceConnections adds the recentDeviceConnections to the get network clients params
func (o *GetNetworkClientsParams) WithRecentDeviceConnections(recentDeviceConnections []string) *GetNetworkClientsParams {
	o.SetRecentDeviceConnections(recentDeviceConnections)
	return o
}

// SetRecentDeviceConnections adds the recentDeviceConnections to the get network clients params
func (o *GetNetworkClientsParams) SetRecentDeviceConnections(recentDeviceConnections []string) {
	o.RecentDeviceConnections = recentDeviceConnections
}

// WithStartingAfter adds the startingAfter to the get network clients params
func (o *GetNetworkClientsParams) WithStartingAfter(startingAfter *string) *GetNetworkClientsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network clients params
func (o *GetNetworkClientsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithStatuses adds the statuses to the get network clients params
func (o *GetNetworkClientsParams) WithStatuses(statuses []string) *GetNetworkClientsParams {
	o.SetStatuses(statuses)
	return o
}

// SetStatuses adds the statuses to the get network clients params
func (o *GetNetworkClientsParams) SetStatuses(statuses []string) {
	o.Statuses = statuses
}

// WithT0 adds the t0 to the get network clients params
func (o *GetNetworkClientsParams) WithT0(t0 *string) *GetNetworkClientsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get network clients params
func (o *GetNetworkClientsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithTimespan adds the timespan to the get network clients params
func (o *GetNetworkClientsParams) WithTimespan(timespan *float32) *GetNetworkClientsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get network clients params
func (o *GetNetworkClientsParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WithVlan adds the vlan to the get network clients params
func (o *GetNetworkClientsParams) WithVlan(vlan *string) *GetNetworkClientsParams {
	o.SetVlan(vlan)
	return o
}

// SetVlan adds the vlan to the get network clients params
func (o *GetNetworkClientsParams) SetVlan(vlan *string) {
	o.Vlan = vlan
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkClientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

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

	if o.IP != nil {

		// query param ip
		var qrIP string

		if o.IP != nil {
			qrIP = *o.IP
		}
		qIP := qrIP
		if qIP != "" {

			if err := r.SetQueryParam("ip", qIP); err != nil {
				return err
			}
		}
	}

	if o.Ip6 != nil {

		// query param ip6
		var qrIp6 string

		if o.Ip6 != nil {
			qrIp6 = *o.Ip6
		}
		qIp6 := qrIp6
		if qIp6 != "" {

			if err := r.SetQueryParam("ip6", qIp6); err != nil {
				return err
			}
		}
	}

	if o.Ip6Local != nil {

		// query param ip6Local
		var qrIp6Local string

		if o.Ip6Local != nil {
			qrIp6Local = *o.Ip6Local
		}
		qIp6Local := qrIp6Local
		if qIp6Local != "" {

			if err := r.SetQueryParam("ip6Local", qIp6Local); err != nil {
				return err
			}
		}
	}

	if o.Mac != nil {

		// query param mac
		var qrMac string

		if o.Mac != nil {
			qrMac = *o.Mac
		}
		qMac := qrMac
		if qMac != "" {

			if err := r.SetQueryParam("mac", qMac); err != nil {
				return err
			}
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.Os != nil {

		// query param os
		var qrOs string

		if o.Os != nil {
			qrOs = *o.Os
		}
		qOs := qrOs
		if qOs != "" {

			if err := r.SetQueryParam("os", qOs); err != nil {
				return err
			}
		}
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

	if o.RecentDeviceConnections != nil {

		// binding items for recentDeviceConnections
		joinedRecentDeviceConnections := o.bindParamRecentDeviceConnections(reg)

		// query array param recentDeviceConnections
		if err := r.SetQueryParam("recentDeviceConnections", joinedRecentDeviceConnections...); err != nil {
			return err
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

	if o.Statuses != nil {

		// binding items for statuses
		joinedStatuses := o.bindParamStatuses(reg)

		// query array param statuses
		if err := r.SetQueryParam("statuses", joinedStatuses...); err != nil {
			return err
		}
	}

	if o.T0 != nil {

		// query param t0
		var qrT0 string

		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {

			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}
	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float32

		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat32(qrTimespan)
		if qTimespan != "" {

			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}
	}

	if o.Vlan != nil {

		// query param vlan
		var qrVlan string

		if o.Vlan != nil {
			qrVlan = *o.Vlan
		}
		qVlan := qrVlan
		if qVlan != "" {

			if err := r.SetQueryParam("vlan", qVlan); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetNetworkClients binds the parameter recentDeviceConnections
func (o *GetNetworkClientsParams) bindParamRecentDeviceConnections(formats strfmt.Registry) []string {
	recentDeviceConnectionsIR := o.RecentDeviceConnections

	var recentDeviceConnectionsIC []string
	for _, recentDeviceConnectionsIIR := range recentDeviceConnectionsIR { // explode []string

		recentDeviceConnectionsIIV := recentDeviceConnectionsIIR // string as string
		recentDeviceConnectionsIC = append(recentDeviceConnectionsIC, recentDeviceConnectionsIIV)
	}

	// items.CollectionFormat: ""
	recentDeviceConnectionsIS := swag.JoinByFormat(recentDeviceConnectionsIC, "")

	return recentDeviceConnectionsIS
}

// bindParamGetNetworkClients binds the parameter statuses
func (o *GetNetworkClientsParams) bindParamStatuses(formats strfmt.Registry) []string {
	statusesIR := o.Statuses

	var statusesIC []string
	for _, statusesIIR := range statusesIR { // explode []string

		statusesIIV := statusesIIR // string as string
		statusesIC = append(statusesIC, statusesIIV)
	}

	// items.CollectionFormat: ""
	statusesIS := swag.JoinByFormat(statusesIC, "")

	return statusesIS
}
