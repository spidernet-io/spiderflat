// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package controller

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

// NewGetIpamStatusParams creates a new GetIpamStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIpamStatusParams() *GetIpamStatusParams {
	return &GetIpamStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIpamStatusParamsWithTimeout creates a new GetIpamStatusParams object
// with the ability to set a timeout on a request.
func NewGetIpamStatusParamsWithTimeout(timeout time.Duration) *GetIpamStatusParams {
	return &GetIpamStatusParams{
		timeout: timeout,
	}
}

// NewGetIpamStatusParamsWithContext creates a new GetIpamStatusParams object
// with the ability to set a context for a request.
func NewGetIpamStatusParamsWithContext(ctx context.Context) *GetIpamStatusParams {
	return &GetIpamStatusParams{
		Context: ctx,
	}
}

// NewGetIpamStatusParamsWithHTTPClient creates a new GetIpamStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIpamStatusParamsWithHTTPClient(client *http.Client) *GetIpamStatusParams {
	return &GetIpamStatusParams{
		HTTPClient: client,
	}
}

/* GetIpamStatusParams contains all the parameters to send to the API endpoint
   for the get ipam status operation.

   Typically these are written to a http.Request.
*/
type GetIpamStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ipam status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIpamStatusParams) WithDefaults() *GetIpamStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ipam status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIpamStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ipam status params
func (o *GetIpamStatusParams) WithTimeout(timeout time.Duration) *GetIpamStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ipam status params
func (o *GetIpamStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ipam status params
func (o *GetIpamStatusParams) WithContext(ctx context.Context) *GetIpamStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ipam status params
func (o *GetIpamStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ipam status params
func (o *GetIpamStatusParams) WithHTTPClient(client *http.Client) *GetIpamStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ipam status params
func (o *GetIpamStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIpamStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}