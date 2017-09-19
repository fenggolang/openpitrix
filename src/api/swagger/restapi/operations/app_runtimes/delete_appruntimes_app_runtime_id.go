// Code generated by go-swagger; DO NOT EDIT.

package app_runtimes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteAppruntimesAppRuntimeIDHandlerFunc turns a function with the right signature into a delete appruntimes app runtime ID handler
type DeleteAppruntimesAppRuntimeIDHandlerFunc func(DeleteAppruntimesAppRuntimeIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAppruntimesAppRuntimeIDHandlerFunc) Handle(params DeleteAppruntimesAppRuntimeIDParams) middleware.Responder {
	return fn(params)
}

// DeleteAppruntimesAppRuntimeIDHandler interface for that can handle valid delete appruntimes app runtime ID params
type DeleteAppruntimesAppRuntimeIDHandler interface {
	Handle(DeleteAppruntimesAppRuntimeIDParams) middleware.Responder
}

// NewDeleteAppruntimesAppRuntimeID creates a new http.Handler for the delete appruntimes app runtime ID operation
func NewDeleteAppruntimesAppRuntimeID(ctx *middleware.Context, handler DeleteAppruntimesAppRuntimeIDHandler) *DeleteAppruntimesAppRuntimeID {
	return &DeleteAppruntimesAppRuntimeID{Context: ctx, Handler: handler}
}

/*DeleteAppruntimesAppRuntimeID swagger:route DELETE /appruntimes/{appRuntimeId} app-runtimes deleteAppruntimesAppRuntimeId

Deletes an app runtime

Delete a single app runtime identified via its id

*/
type DeleteAppruntimesAppRuntimeID struct {
	Context *middleware.Context
	Handler DeleteAppruntimesAppRuntimeIDHandler
}

func (o *DeleteAppruntimesAppRuntimeID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAppruntimesAppRuntimeIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
