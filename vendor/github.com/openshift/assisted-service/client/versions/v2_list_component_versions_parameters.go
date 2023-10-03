// Code generated by go-swagger; DO NOT EDIT.

package versions

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

// NewV2ListComponentVersionsParams creates a new V2ListComponentVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2ListComponentVersionsParams() *V2ListComponentVersionsParams {
	return &V2ListComponentVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2ListComponentVersionsParamsWithTimeout creates a new V2ListComponentVersionsParams object
// with the ability to set a timeout on a request.
func NewV2ListComponentVersionsParamsWithTimeout(timeout time.Duration) *V2ListComponentVersionsParams {
	return &V2ListComponentVersionsParams{
		timeout: timeout,
	}
}

// NewV2ListComponentVersionsParamsWithContext creates a new V2ListComponentVersionsParams object
// with the ability to set a context for a request.
func NewV2ListComponentVersionsParamsWithContext(ctx context.Context) *V2ListComponentVersionsParams {
	return &V2ListComponentVersionsParams{
		Context: ctx,
	}
}

// NewV2ListComponentVersionsParamsWithHTTPClient creates a new V2ListComponentVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2ListComponentVersionsParamsWithHTTPClient(client *http.Client) *V2ListComponentVersionsParams {
	return &V2ListComponentVersionsParams{
		HTTPClient: client,
	}
}

/*
V2ListComponentVersionsParams contains all the parameters to send to the API endpoint

	for the v2 list component versions operation.

	Typically these are written to a http.Request.
*/
type V2ListComponentVersionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 list component versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2ListComponentVersionsParams) WithDefaults() *V2ListComponentVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 list component versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2ListComponentVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 list component versions params
func (o *V2ListComponentVersionsParams) WithTimeout(timeout time.Duration) *V2ListComponentVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 list component versions params
func (o *V2ListComponentVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 list component versions params
func (o *V2ListComponentVersionsParams) WithContext(ctx context.Context) *V2ListComponentVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 list component versions params
func (o *V2ListComponentVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 list component versions params
func (o *V2ListComponentVersionsParams) WithHTTPClient(client *http.Client) *V2ListComponentVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 list component versions params
func (o *V2ListComponentVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V2ListComponentVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
