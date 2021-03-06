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

package schemas_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NoOpReader is a Reader for the NoOp structure.
type NoOpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NoOpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNoOpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewNoOpNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNoOpForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNoOpNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewNoOpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNoOpOK creates a NoOpOK with default headers values
func NewNoOpOK() *NoOpOK {
	return &NoOpOK{}
}

/*NoOpOK handles this case with default header values.

A successful response.
*/
type NoOpOK struct {
	Payload *service_model.V1Schemas
}

func (o *NoOpOK) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpOK  %+v", 200, o.Payload)
}

func (o *NoOpOK) GetPayload() *service_model.V1Schemas {
	return o.Payload
}

func (o *NoOpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Schemas)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNoOpNoContent creates a NoOpNoContent with default headers values
func NewNoOpNoContent() *NoOpNoContent {
	return &NoOpNoContent{}
}

/*NoOpNoContent handles this case with default header values.

No content.
*/
type NoOpNoContent struct {
	Payload interface{}
}

func (o *NoOpNoContent) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpNoContent  %+v", 204, o.Payload)
}

func (o *NoOpNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *NoOpNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNoOpForbidden creates a NoOpForbidden with default headers values
func NewNoOpForbidden() *NoOpForbidden {
	return &NoOpForbidden{}
}

/*NoOpForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type NoOpForbidden struct {
	Payload interface{}
}

func (o *NoOpForbidden) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpForbidden  %+v", 403, o.Payload)
}

func (o *NoOpForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *NoOpForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNoOpNotFound creates a NoOpNotFound with default headers values
func NewNoOpNotFound() *NoOpNotFound {
	return &NoOpNotFound{}
}

/*NoOpNotFound handles this case with default header values.

Resource does not exist.
*/
type NoOpNotFound struct {
	Payload interface{}
}

func (o *NoOpNotFound) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpNotFound  %+v", 404, o.Payload)
}

func (o *NoOpNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *NoOpNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNoOpDefault creates a NoOpDefault with default headers values
func NewNoOpDefault(code int) *NoOpDefault {
	return &NoOpDefault{
		_statusCode: code,
	}
}

/*NoOpDefault handles this case with default header values.

An unexpected error response.
*/
type NoOpDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the no op default response
func (o *NoOpDefault) Code() int {
	return o._statusCode
}

func (o *NoOpDefault) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] NoOp default  %+v", o._statusCode, o.Payload)
}

func (o *NoOpDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *NoOpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
