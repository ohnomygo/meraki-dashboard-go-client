// Code generated by go-swagger; DO NOT EDIT.

package licensing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new licensing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for licensing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetOrganizationLicensingCotermLicenses(params *GetOrganizationLicensingCotermLicensesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationLicensingCotermLicensesOK, error)

	MoveOrganizationLicensingCotermLicenses(params *MoveOrganizationLicensingCotermLicensesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MoveOrganizationLicensingCotermLicensesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetOrganizationLicensingCotermLicenses lists the licenses in a coterm organization

List the licenses in a coterm organization
*/
func (a *Client) GetOrganizationLicensingCotermLicenses(params *GetOrganizationLicensingCotermLicensesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationLicensingCotermLicensesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationLicensingCotermLicensesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationLicensingCotermLicenses",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationId}/licensing/coterm/licenses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationLicensingCotermLicensesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationLicensingCotermLicensesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationLicensingCotermLicenses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
MoveOrganizationLicensingCotermLicenses moves a license to a different organization coterm only

Moves a license to a different organization (coterm only)
*/
func (a *Client) MoveOrganizationLicensingCotermLicenses(params *MoveOrganizationLicensingCotermLicensesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MoveOrganizationLicensingCotermLicensesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveOrganizationLicensingCotermLicensesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "moveOrganizationLicensingCotermLicenses",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationId}/licensing/coterm/licenses/move",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MoveOrganizationLicensingCotermLicensesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MoveOrganizationLicensingCotermLicensesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for moveOrganizationLicensingCotermLicenses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}