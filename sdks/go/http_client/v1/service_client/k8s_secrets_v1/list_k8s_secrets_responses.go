// Copyright 2019 Polyaxon, Inc.
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

package k8s_secrets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListK8sSecretsReader is a Reader for the ListK8sSecrets structure.
type ListK8sSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListK8sSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListK8sSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListK8sSecretsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListK8sSecretsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListK8sSecretsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListK8sSecretsOK creates a ListK8sSecretsOK with default headers values
func NewListK8sSecretsOK() *ListK8sSecretsOK {
	return &ListK8sSecretsOK{}
}

/*ListK8sSecretsOK handles this case with default header values.

A successful response.
*/
type ListK8sSecretsOK struct {
	Payload *service_model.V1ListK8sResourcesResponse
}

func (o *ListK8sSecretsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets][%d] listK8sSecretsOK  %+v", 200, o.Payload)
}

func (o *ListK8sSecretsOK) GetPayload() *service_model.V1ListK8sResourcesResponse {
	return o.Payload
}

func (o *ListK8sSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListK8sResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8sSecretsNoContent creates a ListK8sSecretsNoContent with default headers values
func NewListK8sSecretsNoContent() *ListK8sSecretsNoContent {
	return &ListK8sSecretsNoContent{}
}

/*ListK8sSecretsNoContent handles this case with default header values.

No content.
*/
type ListK8sSecretsNoContent struct {
	Payload interface{}
}

func (o *ListK8sSecretsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets][%d] listK8sSecretsNoContent  %+v", 204, o.Payload)
}

func (o *ListK8sSecretsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8sSecretsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8sSecretsForbidden creates a ListK8sSecretsForbidden with default headers values
func NewListK8sSecretsForbidden() *ListK8sSecretsForbidden {
	return &ListK8sSecretsForbidden{}
}

/*ListK8sSecretsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListK8sSecretsForbidden struct {
	Payload interface{}
}

func (o *ListK8sSecretsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets][%d] listK8sSecretsForbidden  %+v", 403, o.Payload)
}

func (o *ListK8sSecretsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8sSecretsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8sSecretsNotFound creates a ListK8sSecretsNotFound with default headers values
func NewListK8sSecretsNotFound() *ListK8sSecretsNotFound {
	return &ListK8sSecretsNotFound{}
}

/*ListK8sSecretsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListK8sSecretsNotFound struct {
	Payload interface{}
}

func (o *ListK8sSecretsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets][%d] listK8sSecretsNotFound  %+v", 404, o.Payload)
}

func (o *ListK8sSecretsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8sSecretsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}