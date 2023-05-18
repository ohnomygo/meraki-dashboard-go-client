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

// NewGetOrganizationInventoryDevicesParams creates a new GetOrganizationInventoryDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationInventoryDevicesParams() *GetOrganizationInventoryDevicesParams {
	return &GetOrganizationInventoryDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationInventoryDevicesParamsWithTimeout creates a new GetOrganizationInventoryDevicesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationInventoryDevicesParamsWithTimeout(timeout time.Duration) *GetOrganizationInventoryDevicesParams {
	return &GetOrganizationInventoryDevicesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationInventoryDevicesParamsWithContext creates a new GetOrganizationInventoryDevicesParams object
// with the ability to set a context for a request.
func NewGetOrganizationInventoryDevicesParamsWithContext(ctx context.Context) *GetOrganizationInventoryDevicesParams {
	return &GetOrganizationInventoryDevicesParams{
		Context: ctx,
	}
}

// NewGetOrganizationInventoryDevicesParamsWithHTTPClient creates a new GetOrganizationInventoryDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationInventoryDevicesParamsWithHTTPClient(client *http.Client) *GetOrganizationInventoryDevicesParams {
	return &GetOrganizationInventoryDevicesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationInventoryDevicesParams contains all the parameters to send to the API endpoint

	for the get organization inventory devices operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationInventoryDevicesParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* Macs.

	   Search for devices in inventory based on mac addresses.
	*/
	Macs []string

	/* Models.

	   Search for devices in inventory based on model.
	*/
	Models []string

	/* NetworkIds.

	   Search for devices in inventory based on network ids.
	*/
	NetworkIds []string

	/* OrderNumbers.

	   Search for devices in inventory based on order numbers.
	*/
	OrderNumbers []string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	*/
	PerPage *int64

	/* ProductTypes.

	   Filter devices by product type. Accepted values are appliance, camera, cellularGateway, sensor, switch, systemsManager, and wireless.
	*/
	ProductTypes []string

	/* Search.

	   Search for devices in inventory based on serial number, mac address, or model.
	*/
	Search *string

	/* Serials.

	   Search for devices in inventory based on serials.
	*/
	Serials []string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* Tags.

	   Filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).
	*/
	Tags []string

	/* TagsFilterType.

	   To use with 'tags' parameter, to filter devices which contain ANY or ALL given tags. Accepted values are 'withAnyTags' or 'withAllTags', default is 'withAnyTags'.
	*/
	TagsFilterType *string

	/* UsedState.

	   Filter results by used or unused inventory. Accepted values are 'used' or 'unused'.
	*/
	UsedState *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization inventory devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationInventoryDevicesParams) WithDefaults() *GetOrganizationInventoryDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization inventory devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationInventoryDevicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithTimeout(timeout time.Duration) *GetOrganizationInventoryDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithContext(ctx context.Context) *GetOrganizationInventoryDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithHTTPClient(client *http.Client) *GetOrganizationInventoryDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithEndingBefore(endingBefore *string) *GetOrganizationInventoryDevicesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithMacs adds the macs to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithMacs(macs []string) *GetOrganizationInventoryDevicesParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetMacs(macs []string) {
	o.Macs = macs
}

// WithModels adds the models to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithModels(models []string) *GetOrganizationInventoryDevicesParams {
	o.SetModels(models)
	return o
}

// SetModels adds the models to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetModels(models []string) {
	o.Models = models
}

// WithNetworkIds adds the networkIds to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithNetworkIds(networkIds []string) *GetOrganizationInventoryDevicesParams {
	o.SetNetworkIds(networkIds)
	return o
}

// SetNetworkIds adds the networkIds to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetNetworkIds(networkIds []string) {
	o.NetworkIds = networkIds
}

// WithOrderNumbers adds the orderNumbers to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithOrderNumbers(orderNumbers []string) *GetOrganizationInventoryDevicesParams {
	o.SetOrderNumbers(orderNumbers)
	return o
}

// SetOrderNumbers adds the orderNumbers to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetOrderNumbers(orderNumbers []string) {
	o.OrderNumbers = orderNumbers
}

// WithOrganizationID adds the organizationID to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithOrganizationID(organizationID string) *GetOrganizationInventoryDevicesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithPerPage(perPage *int64) *GetOrganizationInventoryDevicesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithProductTypes adds the productTypes to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithProductTypes(productTypes []string) *GetOrganizationInventoryDevicesParams {
	o.SetProductTypes(productTypes)
	return o
}

// SetProductTypes adds the productTypes to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetProductTypes(productTypes []string) {
	o.ProductTypes = productTypes
}

// WithSearch adds the search to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithSearch(search *string) *GetOrganizationInventoryDevicesParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetSearch(search *string) {
	o.Search = search
}

// WithSerials adds the serials to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithSerials(serials []string) *GetOrganizationInventoryDevicesParams {
	o.SetSerials(serials)
	return o
}

// SetSerials adds the serials to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetSerials(serials []string) {
	o.Serials = serials
}

// WithStartingAfter adds the startingAfter to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithStartingAfter(startingAfter *string) *GetOrganizationInventoryDevicesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithTags adds the tags to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithTags(tags []string) *GetOrganizationInventoryDevicesParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithTagsFilterType adds the tagsFilterType to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithTagsFilterType(tagsFilterType *string) *GetOrganizationInventoryDevicesParams {
	o.SetTagsFilterType(tagsFilterType)
	return o
}

// SetTagsFilterType adds the tagsFilterType to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetTagsFilterType(tagsFilterType *string) {
	o.TagsFilterType = tagsFilterType
}

// WithUsedState adds the usedState to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) WithUsedState(usedState *string) *GetOrganizationInventoryDevicesParams {
	o.SetUsedState(usedState)
	return o
}

// SetUsedState adds the usedState to the get organization inventory devices params
func (o *GetOrganizationInventoryDevicesParams) SetUsedState(usedState *string) {
	o.UsedState = usedState
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationInventoryDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Macs != nil {

		// binding items for macs
		joinedMacs := o.bindParamMacs(reg)

		// query array param macs
		if err := r.SetQueryParam("macs", joinedMacs...); err != nil {
			return err
		}
	}

	if o.Models != nil {

		// binding items for models
		joinedModels := o.bindParamModels(reg)

		// query array param models
		if err := r.SetQueryParam("models", joinedModels...); err != nil {
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

	if o.OrderNumbers != nil {

		// binding items for orderNumbers
		joinedOrderNumbers := o.bindParamOrderNumbers(reg)

		// query array param orderNumbers
		if err := r.SetQueryParam("orderNumbers", joinedOrderNumbers...); err != nil {
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

	if o.ProductTypes != nil {

		// binding items for productTypes
		joinedProductTypes := o.bindParamProductTypes(reg)

		// query array param productTypes
		if err := r.SetQueryParam("productTypes", joinedProductTypes...); err != nil {
			return err
		}
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
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

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// query array param tags
		if err := r.SetQueryParam("tags", joinedTags...); err != nil {
			return err
		}
	}

	if o.TagsFilterType != nil {

		// query param tagsFilterType
		var qrTagsFilterType string

		if o.TagsFilterType != nil {
			qrTagsFilterType = *o.TagsFilterType
		}
		qTagsFilterType := qrTagsFilterType
		if qTagsFilterType != "" {

			if err := r.SetQueryParam("tagsFilterType", qTagsFilterType); err != nil {
				return err
			}
		}
	}

	if o.UsedState != nil {

		// query param usedState
		var qrUsedState string

		if o.UsedState != nil {
			qrUsedState = *o.UsedState
		}
		qUsedState := qrUsedState
		if qUsedState != "" {

			if err := r.SetQueryParam("usedState", qUsedState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrganizationInventoryDevices binds the parameter macs
func (o *GetOrganizationInventoryDevicesParams) bindParamMacs(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationInventoryDevices binds the parameter models
func (o *GetOrganizationInventoryDevicesParams) bindParamModels(formats strfmt.Registry) []string {
	modelsIR := o.Models

	var modelsIC []string
	for _, modelsIIR := range modelsIR { // explode []string

		modelsIIV := modelsIIR // string as string
		modelsIC = append(modelsIC, modelsIIV)
	}

	// items.CollectionFormat: ""
	modelsIS := swag.JoinByFormat(modelsIC, "")

	return modelsIS
}

// bindParamGetOrganizationInventoryDevices binds the parameter networkIds
func (o *GetOrganizationInventoryDevicesParams) bindParamNetworkIds(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationInventoryDevices binds the parameter orderNumbers
func (o *GetOrganizationInventoryDevicesParams) bindParamOrderNumbers(formats strfmt.Registry) []string {
	orderNumbersIR := o.OrderNumbers

	var orderNumbersIC []string
	for _, orderNumbersIIR := range orderNumbersIR { // explode []string

		orderNumbersIIV := orderNumbersIIR // string as string
		orderNumbersIC = append(orderNumbersIC, orderNumbersIIV)
	}

	// items.CollectionFormat: ""
	orderNumbersIS := swag.JoinByFormat(orderNumbersIC, "")

	return orderNumbersIS
}

// bindParamGetOrganizationInventoryDevices binds the parameter productTypes
func (o *GetOrganizationInventoryDevicesParams) bindParamProductTypes(formats strfmt.Registry) []string {
	productTypesIR := o.ProductTypes

	var productTypesIC []string
	for _, productTypesIIR := range productTypesIR { // explode []string

		productTypesIIV := productTypesIIR // string as string
		productTypesIC = append(productTypesIC, productTypesIIV)
	}

	// items.CollectionFormat: ""
	productTypesIS := swag.JoinByFormat(productTypesIC, "")

	return productTypesIS
}

// bindParamGetOrganizationInventoryDevices binds the parameter serials
func (o *GetOrganizationInventoryDevicesParams) bindParamSerials(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationInventoryDevices binds the parameter tags
func (o *GetOrganizationInventoryDevicesParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: ""
	tagsIS := swag.JoinByFormat(tagsIC, "")

	return tagsIS
}
