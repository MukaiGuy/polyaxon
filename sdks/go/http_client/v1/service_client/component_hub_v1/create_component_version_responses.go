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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateComponentVersionReader is a Reader for the CreateComponentVersion structure.
type CreateComponentVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateComponentVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateComponentVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateComponentVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateComponentVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateComponentVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateComponentVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateComponentVersionOK creates a CreateComponentVersionOK with default headers values
func NewCreateComponentVersionOK() *CreateComponentVersionOK {
	return &CreateComponentVersionOK{}
}

/*CreateComponentVersionOK handles this case with default header values.

A successful response.
*/
type CreateComponentVersionOK struct {
	Payload *service_model.V1ComponentVersion
}

func (o *CreateComponentVersionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{component}/versions][%d] createComponentVersionOK  %+v", 200, o.Payload)
}

func (o *CreateComponentVersionOK) GetPayload() *service_model.V1ComponentVersion {
	return o.Payload
}

func (o *CreateComponentVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ComponentVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateComponentVersionNoContent creates a CreateComponentVersionNoContent with default headers values
func NewCreateComponentVersionNoContent() *CreateComponentVersionNoContent {
	return &CreateComponentVersionNoContent{}
}

/*CreateComponentVersionNoContent handles this case with default header values.

No content.
*/
type CreateComponentVersionNoContent struct {
	Payload interface{}
}

func (o *CreateComponentVersionNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{component}/versions][%d] createComponentVersionNoContent  %+v", 204, o.Payload)
}

func (o *CreateComponentVersionNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateComponentVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateComponentVersionForbidden creates a CreateComponentVersionForbidden with default headers values
func NewCreateComponentVersionForbidden() *CreateComponentVersionForbidden {
	return &CreateComponentVersionForbidden{}
}

/*CreateComponentVersionForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateComponentVersionForbidden struct {
	Payload interface{}
}

func (o *CreateComponentVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{component}/versions][%d] createComponentVersionForbidden  %+v", 403, o.Payload)
}

func (o *CreateComponentVersionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateComponentVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateComponentVersionNotFound creates a CreateComponentVersionNotFound with default headers values
func NewCreateComponentVersionNotFound() *CreateComponentVersionNotFound {
	return &CreateComponentVersionNotFound{}
}

/*CreateComponentVersionNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateComponentVersionNotFound struct {
	Payload interface{}
}

func (o *CreateComponentVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{component}/versions][%d] createComponentVersionNotFound  %+v", 404, o.Payload)
}

func (o *CreateComponentVersionNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateComponentVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateComponentVersionDefault creates a CreateComponentVersionDefault with default headers values
func NewCreateComponentVersionDefault(code int) *CreateComponentVersionDefault {
	return &CreateComponentVersionDefault{
		_statusCode: code,
	}
}

/*CreateComponentVersionDefault handles this case with default header values.

An unexpected error response.
*/
type CreateComponentVersionDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the create component version default response
func (o *CreateComponentVersionDefault) Code() int {
	return o._statusCode
}

func (o *CreateComponentVersionDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{component}/versions][%d] CreateComponentVersion default  %+v", o._statusCode, o.Payload)
}

func (o *CreateComponentVersionDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateComponentVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
