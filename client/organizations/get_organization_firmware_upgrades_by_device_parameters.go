// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrganizationFirmwareUpgradesByDeviceParams creates a new GetOrganizationFirmwareUpgradesByDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationFirmwareUpgradesByDeviceParams() *GetOrganizationFirmwareUpgradesByDeviceParams {
	return &GetOrganizationFirmwareUpgradesByDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationFirmwareUpgradesByDeviceParamsWithTimeout creates a new GetOrganizationFirmwareUpgradesByDeviceParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationFirmwareUpgradesByDeviceParamsWithTimeout(timeout time.Duration) *GetOrganizationFirmwareUpgradesByDeviceParams {
	return &GetOrganizationFirmwareUpgradesByDeviceParams{
		timeout: timeout,
	}
}

// NewGetOrganizationFirmwareUpgradesByDeviceParamsWithContext creates a new GetOrganizationFirmwareUpgradesByDeviceParams object
// with the ability to set a context for a request.
func NewGetOrganizationFirmwareUpgradesByDeviceParamsWithContext(ctx context.Context) *GetOrganizationFirmwareUpgradesByDeviceParams {
	return &GetOrganizationFirmwareUpgradesByDeviceParams{
		Context: ctx,
	}
}

// NewGetOrganizationFirmwareUpgradesByDeviceParamsWithHTTPClient creates a new GetOrganizationFirmwareUpgradesByDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationFirmwareUpgradesByDeviceParamsWithHTTPClient(client *http.Client) *GetOrganizationFirmwareUpgradesByDeviceParams {
	return &GetOrganizationFirmwareUpgradesByDeviceParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationFirmwareUpgradesByDeviceParams contains all the parameters to send to the API endpoint

	for the get organization firmware upgrades by device operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationFirmwareUpgradesByDeviceParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* FirmwareUpgradeBatchIds.

	   Optional parameter to filter by firmware upgrade batch ids.
	*/
	FirmwareUpgradeBatchIds []string

	/* FirmwareUpgradeIds.

	   Optional parameter to filter by firmware upgrade ids.
	*/
	FirmwareUpgradeIds []string

	/* Macs.

	   Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match.
	*/
	Macs []string

	/* NetworkIds.

	   Optional parameter to filter by network
	*/
	NetworkIds []string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 50. Default is 50.
	*/
	PerPage *int64

	/* Serials.

	   Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match.
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

// WithDefaults hydrates default values in the get organization firmware upgrades by device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithDefaults() *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization firmware upgrades by device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithTimeout(timeout time.Duration) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithContext(ctx context.Context) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithHTTPClient(client *http.Client) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithEndingBefore(endingBefore *string) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithFirmwareUpgradeBatchIds adds the firmwareUpgradeBatchIds to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithFirmwareUpgradeBatchIds(firmwareUpgradeBatchIds []string) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetFirmwareUpgradeBatchIds(firmwareUpgradeBatchIds)
	return o
}

// SetFirmwareUpgradeBatchIds adds the firmwareUpgradeBatchIds to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetFirmwareUpgradeBatchIds(firmwareUpgradeBatchIds []string) {
	o.FirmwareUpgradeBatchIds = firmwareUpgradeBatchIds
}

// WithFirmwareUpgradeIds adds the firmwareUpgradeIds to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithFirmwareUpgradeIds(firmwareUpgradeIds []string) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetFirmwareUpgradeIds(firmwareUpgradeIds)
	return o
}

// SetFirmwareUpgradeIds adds the firmwareUpgradeIds to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetFirmwareUpgradeIds(firmwareUpgradeIds []string) {
	o.FirmwareUpgradeIds = firmwareUpgradeIds
}

// WithMacs adds the macs to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithMacs(macs []string) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetMacs(macs []string) {
	o.Macs = macs
}

// WithNetworkIds adds the networkIds to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithNetworkIds(networkIds []string) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetNetworkIds(networkIds)
	return o
}

// SetNetworkIds adds the networkIds to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetNetworkIds(networkIds []string) {
	o.NetworkIds = networkIds
}

// WithOrganizationID adds the organizationID to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithOrganizationID(organizationID string) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithPerPage(perPage *int64) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSerials adds the serials to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithSerials(serials []string) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetSerials(serials)
	return o
}

// SetSerials adds the serials to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetSerials(serials []string) {
	o.Serials = serials
}

// WithStartingAfter adds the startingAfter to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WithStartingAfter(startingAfter *string) *GetOrganizationFirmwareUpgradesByDeviceParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization firmware upgrades by device params
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FirmwareUpgradeBatchIds != nil {

		// binding items for firmwareUpgradeBatchIds
		joinedFirmwareUpgradeBatchIds := o.bindParamFirmwareUpgradeBatchIds(reg)

		// query array param firmwareUpgradeBatchIds
		if err := r.SetQueryParam("firmwareUpgradeBatchIds", joinedFirmwareUpgradeBatchIds...); err != nil {
			return err
		}
	}

	if o.FirmwareUpgradeIds != nil {

		// binding items for firmwareUpgradeIds
		joinedFirmwareUpgradeIds := o.bindParamFirmwareUpgradeIds(reg)

		// query array param firmwareUpgradeIds
		if err := r.SetQueryParam("firmwareUpgradeIds", joinedFirmwareUpgradeIds...); err != nil {
			return err
		}
	}

	if o.Macs != nil {

		// binding items for macs
		joinedMacs := o.bindParamMacs(reg)

		// query array param macs
		if err := r.SetQueryParam("macs", joinedMacs...); err != nil {
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

// bindParamGetOrganizationFirmwareUpgradesByDevice binds the parameter firmwareUpgradeBatchIds
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) bindParamFirmwareUpgradeBatchIds(formats strfmt.Registry) []string {
	firmwareUpgradeBatchIdsIR := o.FirmwareUpgradeBatchIds

	var firmwareUpgradeBatchIdsIC []string
	for _, firmwareUpgradeBatchIdsIIR := range firmwareUpgradeBatchIdsIR { // explode []string

		firmwareUpgradeBatchIdsIIV := firmwareUpgradeBatchIdsIIR // string as string
		firmwareUpgradeBatchIdsIC = append(firmwareUpgradeBatchIdsIC, firmwareUpgradeBatchIdsIIV)
	}

	// items.CollectionFormat: ""
	firmwareUpgradeBatchIdsIS := swag.JoinByFormat(firmwareUpgradeBatchIdsIC, "")

	return firmwareUpgradeBatchIdsIS
}

// bindParamGetOrganizationFirmwareUpgradesByDevice binds the parameter firmwareUpgradeIds
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) bindParamFirmwareUpgradeIds(formats strfmt.Registry) []string {
	firmwareUpgradeIdsIR := o.FirmwareUpgradeIds

	var firmwareUpgradeIdsIC []string
	for _, firmwareUpgradeIdsIIR := range firmwareUpgradeIdsIR { // explode []string

		firmwareUpgradeIdsIIV := firmwareUpgradeIdsIIR // string as string
		firmwareUpgradeIdsIC = append(firmwareUpgradeIdsIC, firmwareUpgradeIdsIIV)
	}

	// items.CollectionFormat: ""
	firmwareUpgradeIdsIS := swag.JoinByFormat(firmwareUpgradeIdsIC, "")

	return firmwareUpgradeIdsIS
}

// bindParamGetOrganizationFirmwareUpgradesByDevice binds the parameter macs
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) bindParamMacs(formats strfmt.Registry) []string {
	macsIR := o.Macs

	var macsIC []string
	for _, macsIIR := range macsIR { // explode []string

		macsIIV := macsIIR // string as string
		macsIC = append(macsIC, macsIIV)
	}

	// items.CollectionFormat: ""
	macsIS := swag.JoinByFormat(macsIC, "")

	return macsIS
}

// bindParamGetOrganizationFirmwareUpgradesByDevice binds the parameter networkIds
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) bindParamNetworkIds(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationFirmwareUpgradesByDevice binds the parameter serials
func (o *GetOrganizationFirmwareUpgradesByDeviceParams) bindParamSerials(formats strfmt.Registry) []string {
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
