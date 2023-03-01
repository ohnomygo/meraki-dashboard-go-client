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

// NewGetOrganizationNetworksParams creates a new GetOrganizationNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationNetworksParams() *GetOrganizationNetworksParams {
	return &GetOrganizationNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationNetworksParamsWithTimeout creates a new GetOrganizationNetworksParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationNetworksParamsWithTimeout(timeout time.Duration) *GetOrganizationNetworksParams {
	return &GetOrganizationNetworksParams{
		timeout: timeout,
	}
}

// NewGetOrganizationNetworksParamsWithContext creates a new GetOrganizationNetworksParams object
// with the ability to set a context for a request.
func NewGetOrganizationNetworksParamsWithContext(ctx context.Context) *GetOrganizationNetworksParams {
	return &GetOrganizationNetworksParams{
		Context: ctx,
	}
}

// NewGetOrganizationNetworksParamsWithHTTPClient creates a new GetOrganizationNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationNetworksParamsWithHTTPClient(client *http.Client) *GetOrganizationNetworksParams {
	return &GetOrganizationNetworksParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationNetworksParams contains all the parameters to send to the API endpoint

	for the get organization networks operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationNetworksParams struct {

	/* ConfigTemplateID.

	   An optional parameter that is the ID of a config template. Will return all networks bound to that template.
	*/
	ConfigTemplateID *string

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* IsBoundToConfigTemplate.

	   An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false.
	*/
	IsBoundToConfigTemplate *bool

	// OrganizationID.
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000.
	*/
	PerPage *int64

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* Tags.

	   An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).
	*/
	Tags []string

	/* TagsFilterType.

	   An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
	*/
	TagsFilterType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationNetworksParams) WithDefaults() *GetOrganizationNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization networks params
func (o *GetOrganizationNetworksParams) WithTimeout(timeout time.Duration) *GetOrganizationNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization networks params
func (o *GetOrganizationNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization networks params
func (o *GetOrganizationNetworksParams) WithContext(ctx context.Context) *GetOrganizationNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization networks params
func (o *GetOrganizationNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization networks params
func (o *GetOrganizationNetworksParams) WithHTTPClient(client *http.Client) *GetOrganizationNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization networks params
func (o *GetOrganizationNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigTemplateID adds the configTemplateID to the get organization networks params
func (o *GetOrganizationNetworksParams) WithConfigTemplateID(configTemplateID *string) *GetOrganizationNetworksParams {
	o.SetConfigTemplateID(configTemplateID)
	return o
}

// SetConfigTemplateID adds the configTemplateId to the get organization networks params
func (o *GetOrganizationNetworksParams) SetConfigTemplateID(configTemplateID *string) {
	o.ConfigTemplateID = configTemplateID
}

// WithEndingBefore adds the endingBefore to the get organization networks params
func (o *GetOrganizationNetworksParams) WithEndingBefore(endingBefore *string) *GetOrganizationNetworksParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization networks params
func (o *GetOrganizationNetworksParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithIsBoundToConfigTemplate adds the isBoundToConfigTemplate to the get organization networks params
func (o *GetOrganizationNetworksParams) WithIsBoundToConfigTemplate(isBoundToConfigTemplate *bool) *GetOrganizationNetworksParams {
	o.SetIsBoundToConfigTemplate(isBoundToConfigTemplate)
	return o
}

// SetIsBoundToConfigTemplate adds the isBoundToConfigTemplate to the get organization networks params
func (o *GetOrganizationNetworksParams) SetIsBoundToConfigTemplate(isBoundToConfigTemplate *bool) {
	o.IsBoundToConfigTemplate = isBoundToConfigTemplate
}

// WithOrganizationID adds the organizationID to the get organization networks params
func (o *GetOrganizationNetworksParams) WithOrganizationID(organizationID string) *GetOrganizationNetworksParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization networks params
func (o *GetOrganizationNetworksParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization networks params
func (o *GetOrganizationNetworksParams) WithPerPage(perPage *int64) *GetOrganizationNetworksParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization networks params
func (o *GetOrganizationNetworksParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get organization networks params
func (o *GetOrganizationNetworksParams) WithStartingAfter(startingAfter *string) *GetOrganizationNetworksParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization networks params
func (o *GetOrganizationNetworksParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithTags adds the tags to the get organization networks params
func (o *GetOrganizationNetworksParams) WithTags(tags []string) *GetOrganizationNetworksParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get organization networks params
func (o *GetOrganizationNetworksParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithTagsFilterType adds the tagsFilterType to the get organization networks params
func (o *GetOrganizationNetworksParams) WithTagsFilterType(tagsFilterType *string) *GetOrganizationNetworksParams {
	o.SetTagsFilterType(tagsFilterType)
	return o
}

// SetTagsFilterType adds the tagsFilterType to the get organization networks params
func (o *GetOrganizationNetworksParams) SetTagsFilterType(tagsFilterType *string) {
	o.TagsFilterType = tagsFilterType
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConfigTemplateID != nil {

		// query param configTemplateId
		var qrConfigTemplateID string

		if o.ConfigTemplateID != nil {
			qrConfigTemplateID = *o.ConfigTemplateID
		}
		qConfigTemplateID := qrConfigTemplateID
		if qConfigTemplateID != "" {

			if err := r.SetQueryParam("configTemplateId", qConfigTemplateID); err != nil {
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

	if o.IsBoundToConfigTemplate != nil {

		// query param isBoundToConfigTemplate
		var qrIsBoundToConfigTemplate bool

		if o.IsBoundToConfigTemplate != nil {
			qrIsBoundToConfigTemplate = *o.IsBoundToConfigTemplate
		}
		qIsBoundToConfigTemplate := swag.FormatBool(qrIsBoundToConfigTemplate)
		if qIsBoundToConfigTemplate != "" {

			if err := r.SetQueryParam("isBoundToConfigTemplate", qIsBoundToConfigTemplate); err != nil {
				return err
			}
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

// bindParamGetOrganizationNetworks binds the parameter tags
func (o *GetOrganizationNetworksParams) bindParamTags(formats strfmt.Registry) []string {
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
