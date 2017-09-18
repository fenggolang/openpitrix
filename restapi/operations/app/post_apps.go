// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostAppsHandlerFunc turns a function with the right signature into a post apps handler
type PostAppsHandlerFunc func(PostAppsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAppsHandlerFunc) Handle(params PostAppsParams) middleware.Responder {
	return fn(params)
}

// PostAppsHandler interface for that can handle valid post apps params
type PostAppsHandler interface {
	Handle(PostAppsParams) middleware.Responder
}

// NewPostApps creates a new http.Handler for the post apps operation
func NewPostApps(ctx *middleware.Context, handler PostAppsHandler) *PostApps {
	return &PostApps{Context: ctx, Handler: handler}
}

/*PostApps swagger:route POST /apps app postApps

Creates an app

Adds a new app to the runtimes list.

*/
type PostApps struct {
	Context *middleware.Context
	Handler PostAppsHandler
}

func (o *PostApps) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAppsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}