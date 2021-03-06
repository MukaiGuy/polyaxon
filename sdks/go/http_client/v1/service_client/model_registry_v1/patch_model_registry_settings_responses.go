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

package model_registry_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchModelRegistrySettingsReader is a Reader for the PatchModelRegistrySettings structure.
type PatchModelRegistrySettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchModelRegistrySettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchModelRegistrySettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchModelRegistrySettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchModelRegistrySettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchModelRegistrySettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchModelRegistrySettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchModelRegistrySettingsOK creates a PatchModelRegistrySettingsOK with default headers values
func NewPatchModelRegistrySettingsOK() *PatchModelRegistrySettingsOK {
	return &PatchModelRegistrySettingsOK{}
}

/*PatchModelRegistrySettingsOK handles this case with default header values.

A successful response.
*/
type PatchModelRegistrySettingsOK struct {
	Payload *service_model.V1ModelRegistrySettings
}

func (o *PatchModelRegistrySettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model}/settings][%d] patchModelRegistrySettingsOK  %+v", 200, o.Payload)
}

func (o *PatchModelRegistrySettingsOK) GetPayload() *service_model.V1ModelRegistrySettings {
	return o.Payload
}

func (o *PatchModelRegistrySettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ModelRegistrySettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchModelRegistrySettingsNoContent creates a PatchModelRegistrySettingsNoContent with default headers values
func NewPatchModelRegistrySettingsNoContent() *PatchModelRegistrySettingsNoContent {
	return &PatchModelRegistrySettingsNoContent{}
}

/*PatchModelRegistrySettingsNoContent handles this case with default header values.

No content.
*/
type PatchModelRegistrySettingsNoContent struct {
	Payload interface{}
}

func (o *PatchModelRegistrySettingsNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model}/settings][%d] patchModelRegistrySettingsNoContent  %+v", 204, o.Payload)
}

func (o *PatchModelRegistrySettingsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchModelRegistrySettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchModelRegistrySettingsForbidden creates a PatchModelRegistrySettingsForbidden with default headers values
func NewPatchModelRegistrySettingsForbidden() *PatchModelRegistrySettingsForbidden {
	return &PatchModelRegistrySettingsForbidden{}
}

/*PatchModelRegistrySettingsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchModelRegistrySettingsForbidden struct {
	Payload interface{}
}

func (o *PatchModelRegistrySettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model}/settings][%d] patchModelRegistrySettingsForbidden  %+v", 403, o.Payload)
}

func (o *PatchModelRegistrySettingsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchModelRegistrySettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchModelRegistrySettingsNotFound creates a PatchModelRegistrySettingsNotFound with default headers values
func NewPatchModelRegistrySettingsNotFound() *PatchModelRegistrySettingsNotFound {
	return &PatchModelRegistrySettingsNotFound{}
}

/*PatchModelRegistrySettingsNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchModelRegistrySettingsNotFound struct {
	Payload interface{}
}

func (o *PatchModelRegistrySettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model}/settings][%d] patchModelRegistrySettingsNotFound  %+v", 404, o.Payload)
}

func (o *PatchModelRegistrySettingsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchModelRegistrySettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchModelRegistrySettingsDefault creates a PatchModelRegistrySettingsDefault with default headers values
func NewPatchModelRegistrySettingsDefault(code int) *PatchModelRegistrySettingsDefault {
	return &PatchModelRegistrySettingsDefault{
		_statusCode: code,
	}
}

/*PatchModelRegistrySettingsDefault handles this case with default header values.

An unexpected error response.
*/
type PatchModelRegistrySettingsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch model registry settings default response
func (o *PatchModelRegistrySettingsDefault) Code() int {
	return o._statusCode
}

func (o *PatchModelRegistrySettingsDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model}/settings][%d] PatchModelRegistrySettings default  %+v", o._statusCode, o.Payload)
}

func (o *PatchModelRegistrySettingsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchModelRegistrySettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
