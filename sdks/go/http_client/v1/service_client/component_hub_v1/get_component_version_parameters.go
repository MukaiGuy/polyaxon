// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package component_hub_v1

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

// NewGetComponentVersionParams creates a new GetComponentVersionParams object
// with the default values initialized.
func NewGetComponentVersionParams() *GetComponentVersionParams {
	var ()
	return &GetComponentVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetComponentVersionParamsWithTimeout creates a new GetComponentVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetComponentVersionParamsWithTimeout(timeout time.Duration) *GetComponentVersionParams {
	var ()
	return &GetComponentVersionParams{

		timeout: timeout,
	}
}

// NewGetComponentVersionParamsWithContext creates a new GetComponentVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetComponentVersionParamsWithContext(ctx context.Context) *GetComponentVersionParams {
	var ()
	return &GetComponentVersionParams{

		Context: ctx,
	}
}

// NewGetComponentVersionParamsWithHTTPClient creates a new GetComponentVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetComponentVersionParamsWithHTTPClient(client *http.Client) *GetComponentVersionParams {
	var ()
	return &GetComponentVersionParams{
		HTTPClient: client,
	}
}

/*GetComponentVersionParams contains all the parameters to send to the API endpoint
for the get component version operation typically these are written to a http.Request
*/
type GetComponentVersionParams struct {

	/*Entity
	  Owner of the namespace

	*/
	Entity string
	/*Name
	  Component under namesapce

	*/
	Name string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get component version params
func (o *GetComponentVersionParams) WithTimeout(timeout time.Duration) *GetComponentVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get component version params
func (o *GetComponentVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get component version params
func (o *GetComponentVersionParams) WithContext(ctx context.Context) *GetComponentVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get component version params
func (o *GetComponentVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get component version params
func (o *GetComponentVersionParams) WithHTTPClient(client *http.Client) *GetComponentVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get component version params
func (o *GetComponentVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get component version params
func (o *GetComponentVersionParams) WithEntity(entity string) *GetComponentVersionParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get component version params
func (o *GetComponentVersionParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithName adds the name to the get component version params
func (o *GetComponentVersionParams) WithName(name string) *GetComponentVersionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get component version params
func (o *GetComponentVersionParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the get component version params
func (o *GetComponentVersionParams) WithOwner(owner string) *GetComponentVersionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get component version params
func (o *GetComponentVersionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *GetComponentVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
