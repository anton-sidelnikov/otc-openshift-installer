// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewV2DownloadHostIgnitionParams creates a new V2DownloadHostIgnitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2DownloadHostIgnitionParams() *V2DownloadHostIgnitionParams {
	return &V2DownloadHostIgnitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2DownloadHostIgnitionParamsWithTimeout creates a new V2DownloadHostIgnitionParams object
// with the ability to set a timeout on a request.
func NewV2DownloadHostIgnitionParamsWithTimeout(timeout time.Duration) *V2DownloadHostIgnitionParams {
	return &V2DownloadHostIgnitionParams{
		timeout: timeout,
	}
}

// NewV2DownloadHostIgnitionParamsWithContext creates a new V2DownloadHostIgnitionParams object
// with the ability to set a context for a request.
func NewV2DownloadHostIgnitionParamsWithContext(ctx context.Context) *V2DownloadHostIgnitionParams {
	return &V2DownloadHostIgnitionParams{
		Context: ctx,
	}
}

// NewV2DownloadHostIgnitionParamsWithHTTPClient creates a new V2DownloadHostIgnitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2DownloadHostIgnitionParamsWithHTTPClient(client *http.Client) *V2DownloadHostIgnitionParams {
	return &V2DownloadHostIgnitionParams{
		HTTPClient: client,
	}
}

/*
V2DownloadHostIgnitionParams contains all the parameters to send to the API endpoint

	for the v2 download host ignition operation.

	Typically these are written to a http.Request.
*/
type V2DownloadHostIgnitionParams struct {

	/* HostID.

	   The host whose ignition file should be downloaded.

	   Format: uuid
	*/
	HostID strfmt.UUID

	/* InfraEnvID.

	   The infra-env of the host whose ignition file should be downloaded.

	   Format: uuid
	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 download host ignition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2DownloadHostIgnitionParams) WithDefaults() *V2DownloadHostIgnitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 download host ignition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2DownloadHostIgnitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) WithTimeout(timeout time.Duration) *V2DownloadHostIgnitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) WithContext(ctx context.Context) *V2DownloadHostIgnitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) WithHTTPClient(client *http.Client) *V2DownloadHostIgnitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostID adds the hostID to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) WithHostID(hostID strfmt.UUID) *V2DownloadHostIgnitionParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WithInfraEnvID adds the infraEnvID to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) WithInfraEnvID(infraEnvID strfmt.UUID) *V2DownloadHostIgnitionParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the v2 download host ignition params
func (o *V2DownloadHostIgnitionParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *V2DownloadHostIgnitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
