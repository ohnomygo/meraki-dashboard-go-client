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
	"github.com/go-openapi/swag"
)

// NewGetOrganizationApplianceUplinkStatusesParams creates a new GetOrganizationApplianceUplinkStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationApplianceUplinkStatusesParams() *GetOrganizationApplianceUplinkStatusesParams {
	return &GetOrganizationApplianceUplinkStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationApplianceUplinkStatusesParamsWithTimeout creates a new GetOrganizationApplianceUplinkStatusesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationApplianceUplinkStatusesParamsWithTimeout(timeout time.Duration) *GetOrganizationApplianceUplinkStatusesParams {
	return &GetOrganizationApplianceUplinkStatusesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationApplianceUplinkStatusesParamsWithContext creates a new GetOrganizationApplianceUplinkStatusesParams object
// with the ability to set a context for a request.
func NewGetOrganizationApplianceUplinkStatusesParamsWithContext(ctx context.Context) *GetOrganizationApplianceUplinkStatusesParams {
	return &GetOrganizationApplianceUplinkStatusesParams{
		Context: ctx,
	}
}

// NewGetOrganizationApplianceUplinkStatusesParamsWithHTTPClient creates a new GetOrganizationApplianceUplinkStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationApplianceUplinkStatusesParamsWithHTTPClient(client *http.Client) *GetOrganizationApplianceUplinkStatusesParams {
	return &GetOrganizationApplianceUplinkStatusesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationApplianceUplinkStatusesParams contains all the parameters to send to the API endpoint

	for the get organization appliance uplink statuses operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationApplianceUplinkStatusesParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* Iccids.

	   A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
	*/
	Iccids []string

	/* NetworkIds.

	   A list of network IDs. The returned devices will be filtered to only include these networks.
	*/
	NetworkIds []string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	*/
	PerPage *int64

	/* Serials.

	   A list of serial numbers. The returned devices will be filtered to only include these serials.
	*/
	Serials []string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization appliance uplink statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationApplianceUplinkStatusesParams) WithDefaults() *GetOrganizationApplianceUplinkStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization appliance uplink statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationApplianceUplinkStatusesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithTimeout(timeout time.Duration) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithContext(ctx context.Context) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithHTTPClient(client *http.Client) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithEndingBefore(endingBefore *string) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithIccids adds the iccids to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithIccids(iccids []string) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetIccids(iccids)
	return o
}

// SetIccids adds the iccids to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetIccids(iccids []string) {
	o.Iccids = iccids
}

// WithNetworkIds adds the networkIds to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithNetworkIds(networkIds []string) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetNetworkIds(networkIds)
	return o
}

// SetNetworkIds adds the networkIds to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetNetworkIds(networkIds []string) {
	o.NetworkIds = networkIds
}

// WithOrganizationID adds the organizationID to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithOrganizationID(organizationID string) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithPerPage(perPage *int64) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSerials adds the serials to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithSerials(serials []string) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetSerials(serials)
	return o
}

// SetSerials adds the serials to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetSerials(serials []string) {
	o.Serials = serials
}

// WithStartingAfter adds the startingAfter to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) WithStartingAfter(startingAfter *string) *GetOrganizationApplianceUplinkStatusesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization appliance uplink statuses params
func (o *GetOrganizationApplianceUplinkStatusesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationApplianceUplinkStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Iccids != nil {

		// binding items for iccids
		joinedIccids := o.bindParamIccids(reg)

		// query array param iccids
		if err := r.SetQueryParam("iccids", joinedIccids...); err != nil {
			return err
		}
	}

	if o.NetworkIds != nil {

		// binding items for networkIds
		joinedNetworkIds := o.bindParamNetworkIds(reg)

		// query array param networkIds
		if err := r.SetQueryParam("networkIds", joinedNetworkIds...); err != nil {
			return err
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
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

	if o.Serials != nil {

		// binding items for serials
		joinedSerials := o.bindParamSerials(reg)

		// query array param serials
		if err := r.SetQueryParam("serials", joinedSerials...); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrganizationApplianceUplinkStatuses binds the parameter iccids
func (o *GetOrganizationApplianceUplinkStatusesParams) bindParamIccids(formats strfmt.Registry) []string {
	iccidsIR := o.Iccids

	var iccidsIC []string
	for _, iccidsIIR := range iccidsIR { // explode []string

		iccidsIIV := iccidsIIR // string as string
		iccidsIC = append(iccidsIC, iccidsIIV)
	}

	// items.CollectionFormat: ""
	iccidsIS := swag.JoinByFormat(iccidsIC, "")

	return iccidsIS
}

// bindParamGetOrganizationApplianceUplinkStatuses binds the parameter networkIds
func (o *GetOrganizationApplianceUplinkStatusesParams) bindParamNetworkIds(formats strfmt.Registry) []string {
	networkIdsIR := o.NetworkIds

	var networkIdsIC []string
	for _, networkIdsIIR := range networkIdsIR { // explode []string

		networkIdsIIV := networkIdsIIR // string as string
		networkIdsIC = append(networkIdsIC, networkIdsIIV)
	}

	// items.CollectionFormat: ""
	networkIdsIS := swag.JoinByFormat(networkIdsIC, "")

	return networkIdsIS
}

// bindParamGetOrganizationApplianceUplinkStatuses binds the parameter serials
func (o *GetOrganizationApplianceUplinkStatusesParams) bindParamSerials(formats strfmt.Registry) []string {
	serialsIR := o.Serials

	var serialsIC []string
	for _, serialsIIR := range serialsIR { // explode []string

		serialsIIV := serialsIIR // string as string
		serialsIC = append(serialsIC, serialsIIV)
	}

	// items.CollectionFormat: ""
	serialsIS := swag.JoinByFormat(serialsIC, "")

	return serialsIS
}
