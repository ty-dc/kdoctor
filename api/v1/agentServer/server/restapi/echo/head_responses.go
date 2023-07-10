// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kdoctor-io/kdoctor/api/v1/agentServer/models"
)

// HeadOKCode is the HTTP code returned for type HeadOK
const HeadOKCode int = 200

/*
HeadOK Success

swagger:response headOK
*/
type HeadOK struct {

	/*
	  In: Body
	*/
	Payload *models.EchoRes `json:"body,omitempty"`
}

// NewHeadOK creates HeadOK with default headers values
func NewHeadOK() *HeadOK {

	return &HeadOK{}
}

// WithPayload adds the payload to the head o k response
func (o *HeadOK) WithPayload(payload *models.EchoRes) *HeadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the head o k response
func (o *HeadOK) SetPayload(payload *models.EchoRes) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HeadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}