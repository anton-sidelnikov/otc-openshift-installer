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

// NewV2UpdateClusterUISettingsParams creates a new V2UpdateClusterUISettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2UpdateClusterUISettingsParams() *V2UpdateClusterUISettingsParams {
	return &V2UpdateClusterUISettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2UpdateClusterUISettingsParamsWithTimeout creates a new V2UpdateClusterUISettingsParams object
// with the ability to set a timeout on a request.
func NewV2UpdateClusterUISettingsParamsWithTimeout(timeout time.Duration) *V2UpdateClusterUISettingsParams {
	return &V2UpdateClusterUISettingsParams{
		timeout: timeout,
	}
}

// NewV2UpdateClusterUISettingsParamsWithContext creates a new V2UpdateClusterUISettingsParams object
// with the ability to set a context for a request.
func NewV2UpdateClusterUISettingsParamsWithContext(ctx context.Context) *V2UpdateClusterUISettingsParams {
	return &V2UpdateClusterUISettingsParams{
		Context: ctx,
	}
}

// NewV2UpdateClusterUISettingsParamsWithHTTPClient creates a new V2UpdateClusterUISettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2UpdateClusterUISettingsParamsWithHTTPClient(client *http.Client) *V2UpdateClusterUISettingsParams {
	return &V2UpdateClusterUISettingsParams{
		HTTPClient: client,
	}
}

/*
V2UpdateClusterUISettingsParams contains all the parameters to send to the API endpoint

	for the v2 update cluster UI settings operation.

	Typically these are written to a http.Request.
*/
type V2UpdateClusterUISettingsParams struct {

	/* ClusterID.

	   The cluster for which UI settings should be updated.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* UISettings.

	   Settings for the installer UI.
	*/
	UISettings string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 update cluster UI settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2UpdateClusterUISettingsParams) WithDefaults() *V2UpdateClusterUISettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 update cluster UI settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2UpdateClusterUISettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) WithTimeout(timeout time.Duration) *V2UpdateClusterUISettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) WithContext(ctx context.Context) *V2UpdateClusterUISettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) WithHTTPClient(client *http.Client) *V2UpdateClusterUISettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) WithClusterID(clusterID strfmt.UUID) *V2UpdateClusterUISettingsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithUISettings adds the uISettings to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) WithUISettings(uISettings string) *V2UpdateClusterUISettingsParams {
	o.SetUISettings(uISettings)
	return o
}

// SetUISettings adds the uiSettings to the v2 update cluster UI settings params
func (o *V2UpdateClusterUISettingsParams) SetUISettings(uISettings string) {
	o.UISettings = uISettings
}

// WriteToRequest writes these params to a swagger request
func (o *V2UpdateClusterUISettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UISettings); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
