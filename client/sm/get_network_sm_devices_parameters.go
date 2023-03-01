// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// NewGetNetworkSmDevicesParams creates a new GetNetworkSmDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSmDevicesParams() *GetNetworkSmDevicesParams {
	return &GetNetworkSmDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmDevicesParamsWithTimeout creates a new GetNetworkSmDevicesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSmDevicesParamsWithTimeout(timeout time.Duration) *GetNetworkSmDevicesParams {
	return &GetNetworkSmDevicesParams{
		timeout: timeout,
	}
}

// NewGetNetworkSmDevicesParamsWithContext creates a new GetNetworkSmDevicesParams object
// with the ability to set a context for a request.
func NewGetNetworkSmDevicesParamsWithContext(ctx context.Context) *GetNetworkSmDevicesParams {
	return &GetNetworkSmDevicesParams{
		Context: ctx,
	}
}

// NewGetNetworkSmDevicesParamsWithHTTPClient creates a new GetNetworkSmDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSmDevicesParamsWithHTTPClient(client *http.Client) *GetNetworkSmDevicesParams {
	return &GetNetworkSmDevicesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSmDevicesParams contains all the parameters to send to the API endpoint

	for the get network sm devices operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSmDevicesParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* Fields.

	    Additional fields that will be displayed for each device.
	   The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,
	   systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,
	   ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,
	   simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,
	   isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,
	   hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, and url.
	*/
	Fields []string

	/* Ids.

	   Filter devices by id(s).
	*/
	Ids []string

	// NetworkID.
	NetworkID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	*/
	PerPage *int64

	/* Scope.

	   Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags.
	*/
	Scope []string

	/* Serials.

	   Filter devices by serial(s).
	*/
	Serials []string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* WifiMacs.

	   Filter devices by wifi mac(s).
	*/
	WifiMacs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network sm devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDevicesParams) WithDefaults() *GetNetworkSmDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sm devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDevicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithTimeout(timeout time.Duration) *GetNetworkSmDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithContext(ctx context.Context) *GetNetworkSmDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithHTTPClient(client *http.Client) *GetNetworkSmDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithEndingBefore(endingBefore *string) *GetNetworkSmDevicesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithFields adds the fields to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithFields(fields []string) *GetNetworkSmDevicesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIds adds the ids to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithIds(ids []string) *GetNetworkSmDevicesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithNetworkID adds the networkID to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithNetworkID(networkID string) *GetNetworkSmDevicesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPerPage adds the perPage to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithPerPage(perPage *int64) *GetNetworkSmDevicesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithScope adds the scope to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithScope(scope []string) *GetNetworkSmDevicesParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetScope(scope []string) {
	o.Scope = scope
}

// WithSerials adds the serials to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithSerials(serials []string) *GetNetworkSmDevicesParams {
	o.SetSerials(serials)
	return o
}

// SetSerials adds the serials to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetSerials(serials []string) {
	o.Serials = serials
}

// WithStartingAfter adds the startingAfter to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithStartingAfter(startingAfter *string) *GetNetworkSmDevicesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithWifiMacs adds the wifiMacs to the get network sm devices params
func (o *GetNetworkSmDevicesParams) WithWifiMacs(wifiMacs []string) *GetNetworkSmDevicesParams {
	o.SetWifiMacs(wifiMacs)
	return o
}

// SetWifiMacs adds the wifiMacs to the get network sm devices params
func (o *GetNetworkSmDevicesParams) SetWifiMacs(wifiMacs []string) {
	o.WifiMacs = wifiMacs
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
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

	if o.Scope != nil {

		// binding items for scope
		joinedScope := o.bindParamScope(reg)

		// query array param scope
		if err := r.SetQueryParam("scope", joinedScope...); err != nil {
			return err
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

	if o.WifiMacs != nil {

		// binding items for wifiMacs
		joinedWifiMacs := o.bindParamWifiMacs(reg)

		// query array param wifiMacs
		if err := r.SetQueryParam("wifiMacs", joinedWifiMacs...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetNetworkSmDevices binds the parameter fields
func (o *GetNetworkSmDevicesParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: ""
	fieldsIS := swag.JoinByFormat(fieldsIC, "")

	return fieldsIS
}

// bindParamGetNetworkSmDevices binds the parameter ids
func (o *GetNetworkSmDevicesParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: ""
	idsIS := swag.JoinByFormat(idsIC, "")

	return idsIS
}

// bindParamGetNetworkSmDevices binds the parameter scope
func (o *GetNetworkSmDevicesParams) bindParamScope(formats strfmt.Registry) []string {
	scopeIR := o.Scope

	var scopeIC []string
	for _, scopeIIR := range scopeIR { // explode []string

		scopeIIV := scopeIIR // string as string
		scopeIC = append(scopeIC, scopeIIV)
	}

	// items.CollectionFormat: ""
	scopeIS := swag.JoinByFormat(scopeIC, "")

	return scopeIS
}

// bindParamGetNetworkSmDevices binds the parameter serials
func (o *GetNetworkSmDevicesParams) bindParamSerials(formats strfmt.Registry) []string {
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

// bindParamGetNetworkSmDevices binds the parameter wifiMacs
func (o *GetNetworkSmDevicesParams) bindParamWifiMacs(formats strfmt.Registry) []string {
	wifiMacsIR := o.WifiMacs

	var wifiMacsIC []string
	for _, wifiMacsIIR := range wifiMacsIR { // explode []string

		wifiMacsIIV := wifiMacsIIR // string as string
		wifiMacsIC = append(wifiMacsIC, wifiMacsIIV)
	}

	// items.CollectionFormat: ""
	wifiMacsIS := swag.JoinByFormat(wifiMacsIC, "")

	return wifiMacsIS
}
