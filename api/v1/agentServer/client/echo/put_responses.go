// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutReader is a Reader for the Put structure.
type PutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOK creates a PutOK with default headers values
func NewPutOK() *PutOK {
	return &PutOK{}
}

/*
PutOK describes a response with status code 200, with default header values.

Success
*/
type PutOK struct {
}

// IsSuccess returns true when this put o k response has a 2xx status code
func (o *PutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put o k response has a 3xx status code
func (o *PutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put o k response has a 4xx status code
func (o *PutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put o k response has a 5xx status code
func (o *PutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put o k response a status code equal to that given
func (o *PutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put o k response
func (o *PutOK) Code() int {
	return 200
}

func (o *PutOK) Error() string {
	return fmt.Sprintf("[PUT /][%d] putOK ", 200)
}

func (o *PutOK) String() string {
	return fmt.Sprintf("[PUT /][%d] putOK ", 200)
}

func (o *PutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutInternalServerError creates a PutInternalServerError with default headers values
func NewPutInternalServerError() *PutInternalServerError {
	return &PutInternalServerError{}
}

/*
PutInternalServerError describes a response with status code 500, with default header values.

Failed
*/
type PutInternalServerError struct {
}

// IsSuccess returns true when this put internal server error response has a 2xx status code
func (o *PutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put internal server error response has a 3xx status code
func (o *PutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put internal server error response has a 4xx status code
func (o *PutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put internal server error response has a 5xx status code
func (o *PutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put internal server error response a status code equal to that given
func (o *PutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put internal server error response
func (o *PutInternalServerError) Code() int {
	return 500
}

func (o *PutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /][%d] putInternalServerError ", 500)
}

func (o *PutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /][%d] putInternalServerError ", 500)
}

func (o *PutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
