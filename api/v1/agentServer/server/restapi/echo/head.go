// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// HeadHandlerFunc turns a function with the right signature into a head handler
type HeadHandlerFunc func(HeadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HeadHandlerFunc) Handle(params HeadParams) middleware.Responder {
	return fn(params)
}

// HeadHandler interface for that can handle valid head params
type HeadHandler interface {
	Handle(HeadParams) middleware.Responder
}

// NewHead creates a new http.Handler for the head operation
func NewHead(ctx *middleware.Context, handler HeadHandler) *Head {
	return &Head{Context: ctx, Handler: handler}
}

/*
	Head swagger:route HEAD / echo head

echo http request

echo http request
*/
type Head struct {
	Context *middleware.Context
	Handler HeadHandler
}

func (o *Head) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewHeadParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}