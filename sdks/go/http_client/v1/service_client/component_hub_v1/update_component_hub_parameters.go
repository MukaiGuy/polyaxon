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

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewUpdateComponentHubParams creates a new UpdateComponentHubParams object
// with the default values initialized.
func NewUpdateComponentHubParams() *UpdateComponentHubParams {
	var ()
	return &UpdateComponentHubParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateComponentHubParamsWithTimeout creates a new UpdateComponentHubParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateComponentHubParamsWithTimeout(timeout time.Duration) *UpdateComponentHubParams {
	var ()
	return &UpdateComponentHubParams{

		timeout: timeout,
	}
}

// NewUpdateComponentHubParamsWithContext creates a new UpdateComponentHubParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateComponentHubParamsWithContext(ctx context.Context) *UpdateComponentHubParams {
	var ()
	return &UpdateComponentHubParams{

		Context: ctx,
	}
}

// NewUpdateComponentHubParamsWithHTTPClient creates a new UpdateComponentHubParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateComponentHubParamsWithHTTPClient(client *http.Client) *UpdateComponentHubParams {
	var ()
	return &UpdateComponentHubParams{
		HTTPClient: client,
	}
}

/*UpdateComponentHubParams contains all the parameters to send to the API endpoint
for the update component hub operation typically these are written to a http.Request
*/
type UpdateComponentHubParams struct {

	/*Body
	  Component body

	*/
	Body *service_model.V1ComponentHub
	/*ComponentName
	  Optional component name, should be a valid fully qualified value: name[:version]

	*/
	ComponentName string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update component hub params
func (o *UpdateComponentHubParams) WithTimeout(timeout time.Duration) *UpdateComponentHubParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update component hub params
func (o *UpdateComponentHubParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update component hub params
func (o *UpdateComponentHubParams) WithContext(ctx context.Context) *UpdateComponentHubParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update component hub params
func (o *UpdateComponentHubParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update component hub params
func (o *UpdateComponentHubParams) WithHTTPClient(client *http.Client) *UpdateComponentHubParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update component hub params
func (o *UpdateComponentHubParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update component hub params
func (o *UpdateComponentHubParams) WithBody(body *service_model.V1ComponentHub) *UpdateComponentHubParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update component hub params
func (o *UpdateComponentHubParams) SetBody(body *service_model.V1ComponentHub) {
	o.Body = body
}

// WithComponentName adds the componentName to the update component hub params
func (o *UpdateComponentHubParams) WithComponentName(componentName string) *UpdateComponentHubParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the update component hub params
func (o *UpdateComponentHubParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithOwner adds the owner to the update component hub params
func (o *UpdateComponentHubParams) WithOwner(owner string) *UpdateComponentHubParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update component hub params
func (o *UpdateComponentHubParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateComponentHubParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param component.name
	if err := r.SetPathParam("component.name", o.ComponentName); err != nil {
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
