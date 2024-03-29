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

// NewGetOrganizationDevicesParams creates a new GetOrganizationDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationDevicesParams() *GetOrganizationDevicesParams {
	return &GetOrganizationDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationDevicesParamsWithTimeout creates a new GetOrganizationDevicesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationDevicesParamsWithTimeout(timeout time.Duration) *GetOrganizationDevicesParams {
	return &GetOrganizationDevicesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationDevicesParamsWithContext creates a new GetOrganizationDevicesParams object
// with the ability to set a context for a request.
func NewGetOrganizationDevicesParamsWithContext(ctx context.Context) *GetOrganizationDevicesParams {
	return &GetOrganizationDevicesParams{
		Context: ctx,
	}
}

// NewGetOrganizationDevicesParamsWithHTTPClient creates a new GetOrganizationDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationDevicesParamsWithHTTPClient(client *http.Client) *GetOrganizationDevicesParams {
	return &GetOrganizationDevicesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationDevicesParams contains all the parameters to send to the API endpoint

	for the get organization devices operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationDevicesParams struct {

	/* ConfigurationUpdatedAfter.

	   Filter results by whether or not the device's configuration has been updated after the given timestamp
	*/
	ConfigurationUpdatedAfter *string

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* Mac.

	   Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match.
	*/
	Mac *string

	/* Macs.

	   Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match.
	*/
	Macs []string

	/* Model.

	   Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match.
	*/
	Model *string

	/* Models.

	   Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match.
	*/
	Models []string

	/* Name.

	   Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match.
	*/
	Name *string

	/* NetworkIds.

	   Optional parameter to filter devices by network.
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

	/* ProductTypes.

	   Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and sensor.
	*/
	ProductTypes []string

	/* SensorAlertProfileIds.

	   Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices.
	*/
	SensorAlertProfileIds []string

	/* SensorMetrics.

	   Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices.
	*/
	SensorMetrics []string

	/* Serial.

	   Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match.
	*/
	Serial *string

	/* Serials.

	   Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match.
	*/
	Serials []string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* Tags.

	   Optional parameter to filter devices by tags.
	*/
	Tags []string

	/* TagsFilterType.

	   Optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
	*/
	TagsFilterType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationDevicesParams) WithDefaults() *GetOrganizationDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationDevicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization devices params
func (o *GetOrganizationDevicesParams) WithTimeout(timeout time.Duration) *GetOrganizationDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization devices params
func (o *GetOrganizationDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization devices params
func (o *GetOrganizationDevicesParams) WithContext(ctx context.Context) *GetOrganizationDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization devices params
func (o *GetOrganizationDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization devices params
func (o *GetOrganizationDevicesParams) WithHTTPClient(client *http.Client) *GetOrganizationDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization devices params
func (o *GetOrganizationDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigurationUpdatedAfter adds the configurationUpdatedAfter to the get organization devices params
func (o *GetOrganizationDevicesParams) WithConfigurationUpdatedAfter(configurationUpdatedAfter *string) *GetOrganizationDevicesParams {
	o.SetConfigurationUpdatedAfter(configurationUpdatedAfter)
	return o
}

// SetConfigurationUpdatedAfter adds the configurationUpdatedAfter to the get organization devices params
func (o *GetOrganizationDevicesParams) SetConfigurationUpdatedAfter(configurationUpdatedAfter *string) {
	o.ConfigurationUpdatedAfter = configurationUpdatedAfter
}

// WithEndingBefore adds the endingBefore to the get organization devices params
func (o *GetOrganizationDevicesParams) WithEndingBefore(endingBefore *string) *GetOrganizationDevicesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization devices params
func (o *GetOrganizationDevicesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithMac adds the mac to the get organization devices params
func (o *GetOrganizationDevicesParams) WithMac(mac *string) *GetOrganizationDevicesParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the get organization devices params
func (o *GetOrganizationDevicesParams) SetMac(mac *string) {
	o.Mac = mac
}

// WithMacs adds the macs to the get organization devices params
func (o *GetOrganizationDevicesParams) WithMacs(macs []string) *GetOrganizationDevicesParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the get organization devices params
func (o *GetOrganizationDevicesParams) SetMacs(macs []string) {
	o.Macs = macs
}

// WithModel adds the model to the get organization devices params
func (o *GetOrganizationDevicesParams) WithModel(model *string) *GetOrganizationDevicesParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the get organization devices params
func (o *GetOrganizationDevicesParams) SetModel(model *string) {
	o.Model = model
}

// WithModels adds the models to the get organization devices params
func (o *GetOrganizationDevicesParams) WithModels(models []string) *GetOrganizationDevicesParams {
	o.SetModels(models)
	return o
}

// SetModels adds the models to the get organization devices params
func (o *GetOrganizationDevicesParams) SetModels(models []string) {
	o.Models = models
}

// WithName adds the name to the get organization devices params
func (o *GetOrganizationDevicesParams) WithName(name *string) *GetOrganizationDevicesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get organization devices params
func (o *GetOrganizationDevicesParams) SetName(name *string) {
	o.Name = name
}

// WithNetworkIds adds the networkIds to the get organization devices params
func (o *GetOrganizationDevicesParams) WithNetworkIds(networkIds []string) *GetOrganizationDevicesParams {
	o.SetNetworkIds(networkIds)
	return o
}

// SetNetworkIds adds the networkIds to the get organization devices params
func (o *GetOrganizationDevicesParams) SetNetworkIds(networkIds []string) {
	o.NetworkIds = networkIds
}

// WithOrganizationID adds the organizationID to the get organization devices params
func (o *GetOrganizationDevicesParams) WithOrganizationID(organizationID string) *GetOrganizationDevicesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization devices params
func (o *GetOrganizationDevicesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization devices params
func (o *GetOrganizationDevicesParams) WithPerPage(perPage *int64) *GetOrganizationDevicesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization devices params
func (o *GetOrganizationDevicesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithProductTypes adds the productTypes to the get organization devices params
func (o *GetOrganizationDevicesParams) WithProductTypes(productTypes []string) *GetOrganizationDevicesParams {
	o.SetProductTypes(productTypes)
	return o
}

// SetProductTypes adds the productTypes to the get organization devices params
func (o *GetOrganizationDevicesParams) SetProductTypes(productTypes []string) {
	o.ProductTypes = productTypes
}

// WithSensorAlertProfileIds adds the sensorAlertProfileIds to the get organization devices params
func (o *GetOrganizationDevicesParams) WithSensorAlertProfileIds(sensorAlertProfileIds []string) *GetOrganizationDevicesParams {
	o.SetSensorAlertProfileIds(sensorAlertProfileIds)
	return o
}

// SetSensorAlertProfileIds adds the sensorAlertProfileIds to the get organization devices params
func (o *GetOrganizationDevicesParams) SetSensorAlertProfileIds(sensorAlertProfileIds []string) {
	o.SensorAlertProfileIds = sensorAlertProfileIds
}

// WithSensorMetrics adds the sensorMetrics to the get organization devices params
func (o *GetOrganizationDevicesParams) WithSensorMetrics(sensorMetrics []string) *GetOrganizationDevicesParams {
	o.SetSensorMetrics(sensorMetrics)
	return o
}

// SetSensorMetrics adds the sensorMetrics to the get organization devices params
func (o *GetOrganizationDevicesParams) SetSensorMetrics(sensorMetrics []string) {
	o.SensorMetrics = sensorMetrics
}

// WithSerial adds the serial to the get organization devices params
func (o *GetOrganizationDevicesParams) WithSerial(serial *string) *GetOrganizationDevicesParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get organization devices params
func (o *GetOrganizationDevicesParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithSerials adds the serials to the get organization devices params
func (o *GetOrganizationDevicesParams) WithSerials(serials []string) *GetOrganizationDevicesParams {
	o.SetSerials(serials)
	return o
}

// SetSerials adds the serials to the get organization devices params
func (o *GetOrganizationDevicesParams) SetSerials(serials []string) {
	o.Serials = serials
}

// WithStartingAfter adds the startingAfter to the get organization devices params
func (o *GetOrganizationDevicesParams) WithStartingAfter(startingAfter *string) *GetOrganizationDevicesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization devices params
func (o *GetOrganizationDevicesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithTags adds the tags to the get organization devices params
func (o *GetOrganizationDevicesParams) WithTags(tags []string) *GetOrganizationDevicesParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get organization devices params
func (o *GetOrganizationDevicesParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithTagsFilterType adds the tagsFilterType to the get organization devices params
func (o *GetOrganizationDevicesParams) WithTagsFilterType(tagsFilterType *string) *GetOrganizationDevicesParams {
	o.SetTagsFilterType(tagsFilterType)
	return o
}

// SetTagsFilterType adds the tagsFilterType to the get organization devices params
func (o *GetOrganizationDevicesParams) SetTagsFilterType(tagsFilterType *string) {
	o.TagsFilterType = tagsFilterType
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConfigurationUpdatedAfter != nil {

		// query param configurationUpdatedAfter
		var qrConfigurationUpdatedAfter string

		if o.ConfigurationUpdatedAfter != nil {
			qrConfigurationUpdatedAfter = *o.ConfigurationUpdatedAfter
		}
		qConfigurationUpdatedAfter := qrConfigurationUpdatedAfter
		if qConfigurationUpdatedAfter != "" {

			if err := r.SetQueryParam("configurationUpdatedAfter", qConfigurationUpdatedAfter); err != nil {
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

	if o.Macs != nil {

		// binding items for macs
		joinedMacs := o.bindParamMacs(reg)

		// query array param macs
		if err := r.SetQueryParam("macs", joinedMacs...); err != nil {
			return err
		}
	}

	if o.Model != nil {

		// query param model
		var qrModel string

		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {

			if err := r.SetQueryParam("model", qModel); err != nil {
				return err
			}
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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
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

	if o.ProductTypes != nil {

		// binding items for productTypes
		joinedProductTypes := o.bindParamProductTypes(reg)

		// query array param productTypes
		if err := r.SetQueryParam("productTypes", joinedProductTypes...); err != nil {
			return err
		}
	}

	if o.SensorAlertProfileIds != nil {

		// binding items for sensorAlertProfileIds
		joinedSensorAlertProfileIds := o.bindParamSensorAlertProfileIds(reg)

		// query array param sensorAlertProfileIds
		if err := r.SetQueryParam("sensorAlertProfileIds", joinedSensorAlertProfileIds...); err != nil {
			return err
		}
	}

	if o.SensorMetrics != nil {

		// binding items for sensorMetrics
		joinedSensorMetrics := o.bindParamSensorMetrics(reg)

		// query array param sensorMetrics
		if err := r.SetQueryParam("sensorMetrics", joinedSensorMetrics...); err != nil {
			return err
		}
	}

	if o.Serial != nil {

		// query param serial
		var qrSerial string

		if o.Serial != nil {
			qrSerial = *o.Serial
		}
		qSerial := qrSerial
		if qSerial != "" {

			if err := r.SetQueryParam("serial", qSerial); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrganizationDevices binds the parameter macs
func (o *GetOrganizationDevicesParams) bindParamMacs(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationDevices binds the parameter models
func (o *GetOrganizationDevicesParams) bindParamModels(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationDevices binds the parameter networkIds
func (o *GetOrganizationDevicesParams) bindParamNetworkIds(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationDevices binds the parameter productTypes
func (o *GetOrganizationDevicesParams) bindParamProductTypes(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationDevices binds the parameter sensorAlertProfileIds
func (o *GetOrganizationDevicesParams) bindParamSensorAlertProfileIds(formats strfmt.Registry) []string {
	sensorAlertProfileIdsIR := o.SensorAlertProfileIds

	var sensorAlertProfileIdsIC []string
	for _, sensorAlertProfileIdsIIR := range sensorAlertProfileIdsIR { // explode []string

		sensorAlertProfileIdsIIV := sensorAlertProfileIdsIIR // string as string
		sensorAlertProfileIdsIC = append(sensorAlertProfileIdsIC, sensorAlertProfileIdsIIV)
	}

	// items.CollectionFormat: ""
	sensorAlertProfileIdsIS := swag.JoinByFormat(sensorAlertProfileIdsIC, "")

	return sensorAlertProfileIdsIS
}

// bindParamGetOrganizationDevices binds the parameter sensorMetrics
func (o *GetOrganizationDevicesParams) bindParamSensorMetrics(formats strfmt.Registry) []string {
	sensorMetricsIR := o.SensorMetrics

	var sensorMetricsIC []string
	for _, sensorMetricsIIR := range sensorMetricsIR { // explode []string

		sensorMetricsIIV := sensorMetricsIIR // string as string
		sensorMetricsIC = append(sensorMetricsIC, sensorMetricsIIV)
	}

	// items.CollectionFormat: ""
	sensorMetricsIS := swag.JoinByFormat(sensorMetricsIC, "")

	return sensorMetricsIS
}

// bindParamGetOrganizationDevices binds the parameter serials
func (o *GetOrganizationDevicesParams) bindParamSerials(formats strfmt.Registry) []string {
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

// bindParamGetOrganizationDevices binds the parameter tags
func (o *GetOrganizationDevicesParams) bindParamTags(formats strfmt.Registry) []string {
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
