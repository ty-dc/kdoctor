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

// New creates a new echo API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for echo API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Get(params *GetParams, opts ...ClientOption) (*GetOK, error)

	GetKdoctoragent(params *GetKdoctoragentParams, opts ...ClientOption) (*GetKdoctoragentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Get echos http request

echo http request
*/
func (a *Client) Get(params *GetParams, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetKdoctoragent echos http request

echo http request
*/
func (a *Client) GetKdoctoragent(params *GetKdoctoragentParams, opts ...ClientOption) (*GetKdoctoragentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKdoctoragentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetKdoctoragent",
		Method:             "GET",
		PathPattern:        "/kdoctoragent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetKdoctoragentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetKdoctoragentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetKdoctoragent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
