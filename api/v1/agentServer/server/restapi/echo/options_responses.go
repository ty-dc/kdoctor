// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// OptionsOKCode is the HTTP code returned for type OptionsOK
const OptionsOKCode int = 200

/*
OptionsOK Success

swagger:response optionsOK
*/
type OptionsOK struct {
}

// NewOptionsOK creates OptionsOK with default headers values
func NewOptionsOK() *OptionsOK {

	return &OptionsOK{}
}

// WriteResponse to the client
func (o *OptionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// OptionsInternalServerErrorCode is the HTTP code returned for type OptionsInternalServerError
const OptionsInternalServerErrorCode int = 500

/*
OptionsInternalServerError Failed

swagger:response optionsInternalServerError
*/
type OptionsInternalServerError struct {
}

// NewOptionsInternalServerError creates OptionsInternalServerError with default headers values
func NewOptionsInternalServerError() *OptionsInternalServerError {

	return &OptionsInternalServerError{}
}

// WriteResponse to the client
func (o *OptionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
