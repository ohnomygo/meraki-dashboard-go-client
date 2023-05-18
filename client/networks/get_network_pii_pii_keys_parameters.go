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

// NewGetNetworkPiiPiiKeysParams creates a new GetNetworkPiiPiiKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkPiiPiiKeysParams() *GetNetworkPiiPiiKeysParams {
	return &GetNetworkPiiPiiKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkPiiPiiKeysParamsWithTimeout creates a new GetNetworkPiiPiiKeysParams object
// with the ability to set a timeout on a request.
func NewGetNetworkPiiPiiKeysParamsWithTimeout(timeout time.Duration) *GetNetworkPiiPiiKeysParams {
	return &GetNetworkPiiPiiKeysParams{
		timeout: timeout,
	}
}

// NewGetNetworkPiiPiiKeysParamsWithContext creates a new GetNetworkPiiPiiKeysParams object
// with the ability to set a context for a request.
func NewGetNetworkPiiPiiKeysParamsWithContext(ctx context.Context) *GetNetworkPiiPiiKeysParams {
	return &GetNetworkPiiPiiKeysParams{
		Context: ctx,
	}
}

// NewGetNetworkPiiPiiKeysParamsWithHTTPClient creates a new GetNetworkPiiPiiKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkPiiPiiKeysParamsWithHTTPClient(client *http.Client) *GetNetworkPiiPiiKeysParams {
	return &GetNetworkPiiPiiKeysParams{
		HTTPClient: client,
	}
}

/*
GetNetworkPiiPiiKeysParams contains all the parameters to send to the API endpoint

	for the get network pii pii keys operation.

	Typically these are written to a http.Request.
*/
type GetNetworkPiiPiiKeysParams struct {

	/* BluetoothMac.

	   The MAC of a Bluetooth client
	*/
	BluetoothMac *string

	/* Email.

	   The email of a network user account or a Systems Manager device
	*/
	Email *string

	/* Imei.

	   The IMEI of a Systems Manager device
	*/
	Imei *string

	/* Mac.

	   The MAC of a network client device or a Systems Manager device
	*/
	Mac *string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Serial.

	   The serial of a Systems Manager device
	*/
	Serial *string

	/* Username.

	   The username of a Systems Manager user
	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network pii pii keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkPiiPiiKeysParams) WithDefaults() *GetNetworkPiiPiiKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network pii pii keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkPiiPiiKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithTimeout(timeout time.Duration) *GetNetworkPiiPiiKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithContext(ctx context.Context) *GetNetworkPiiPiiKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithHTTPClient(client *http.Client) *GetNetworkPiiPiiKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBluetoothMac adds the bluetoothMac to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithBluetoothMac(bluetoothMac *string) *GetNetworkPiiPiiKeysParams {
	o.SetBluetoothMac(bluetoothMac)
	return o
}

// SetBluetoothMac adds the bluetoothMac to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetBluetoothMac(bluetoothMac *string) {
	o.BluetoothMac = bluetoothMac
}

// WithEmail adds the email to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithEmail(email *string) *GetNetworkPiiPiiKeysParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetEmail(email *string) {
	o.Email = email
}

// WithImei adds the imei to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithImei(imei *string) *GetNetworkPiiPiiKeysParams {
	o.SetImei(imei)
	return o
}

// SetImei adds the imei to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetImei(imei *string) {
	o.Imei = imei
}

// WithMac adds the mac to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithMac(mac *string) *GetNetworkPiiPiiKeysParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetMac(mac *string) {
	o.Mac = mac
}

// WithNetworkID adds the networkID to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithNetworkID(networkID string) *GetNetworkPiiPiiKeysParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSerial adds the serial to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithSerial(serial *string) *GetNetworkPiiPiiKeysParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithUsername adds the username to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) WithUsername(username *string) *GetNetworkPiiPiiKeysParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get network pii pii keys params
func (o *GetNetworkPiiPiiKeysParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkPiiPiiKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BluetoothMac != nil {

		// query param bluetoothMac
		var qrBluetoothMac string

		if o.BluetoothMac != nil {
			qrBluetoothMac = *o.BluetoothMac
		}
		qBluetoothMac := qrBluetoothMac
		if qBluetoothMac != "" {

			if err := r.SetQueryParam("bluetoothMac", qBluetoothMac); err != nil {
				return err
			}
		}
	}

	if o.Email != nil {

		// query param email
		var qrEmail string

		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {

			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}
	}

	if o.Imei != nil {

		// query param imei
		var qrImei string

		if o.Imei != nil {
			qrImei = *o.Imei
		}
		qImei := qrImei
		if qImei != "" {

			if err := r.SetQueryParam("imei", qImei); err != nil {
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

	if o.Username != nil {

		// query param username
		var qrUsername string

		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
