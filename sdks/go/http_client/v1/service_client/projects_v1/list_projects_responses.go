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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListProjectsReader is a Reader for the ListProjects structure.
type ListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListProjectsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectsOK creates a ListProjectsOK with default headers values
func NewListProjectsOK() *ListProjectsOK {
	return &ListProjectsOK{}
}

/*ListProjectsOK handles this case with default header values.

A successful response.
*/
type ListProjectsOK struct {
	Payload *service_model.V1ListProjectsResponse
}

func (o *ListProjectsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/projects/list][%d] listProjectsOK  %+v", 200, o.Payload)
}

func (o *ListProjectsOK) GetPayload() *service_model.V1ListProjectsResponse {
	return o.Payload
}

func (o *ListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsNoContent creates a ListProjectsNoContent with default headers values
func NewListProjectsNoContent() *ListProjectsNoContent {
	return &ListProjectsNoContent{}
}

/*ListProjectsNoContent handles this case with default header values.

No content.
*/
type ListProjectsNoContent struct {
	Payload interface{}
}

func (o *ListProjectsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/projects/list][%d] listProjectsNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsForbidden creates a ListProjectsForbidden with default headers values
func NewListProjectsForbidden() *ListProjectsForbidden {
	return &ListProjectsForbidden{}
}

/*ListProjectsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListProjectsForbidden struct {
	Payload interface{}
}

func (o *ListProjectsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/projects/list][%d] listProjectsForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsNotFound creates a ListProjectsNotFound with default headers values
func NewListProjectsNotFound() *ListProjectsNotFound {
	return &ListProjectsNotFound{}
}

/*ListProjectsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListProjectsNotFound struct {
	Payload interface{}
}

func (o *ListProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/projects/list][%d] listProjectsNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsDefault creates a ListProjectsDefault with default headers values
func NewListProjectsDefault(code int) *ListProjectsDefault {
	return &ListProjectsDefault{
		_statusCode: code,
	}
}

/*ListProjectsDefault handles this case with default header values.

An unexpected error response.
*/
type ListProjectsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list projects default response
func (o *ListProjectsDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/projects/list][%d] ListProjects default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
