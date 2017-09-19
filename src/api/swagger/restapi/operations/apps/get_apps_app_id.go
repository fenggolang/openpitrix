// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAppsAppIDHandlerFunc turns a function with the right signature into a get apps app ID handler
type GetAppsAppIDHandlerFunc func(GetAppsAppIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAppsAppIDHandlerFunc) Handle(params GetAppsAppIDParams) middleware.Responder {
	return fn(params)
}

// GetAppsAppIDHandler interface for that can handle valid get apps app ID params
type GetAppsAppIDHandler interface {
	Handle(GetAppsAppIDParams) middleware.Responder
}

// NewGetAppsAppID creates a new http.Handler for the get apps app ID operation
func NewGetAppsAppID(ctx *middleware.Context, handler GetAppsAppIDHandler) *GetAppsAppID {
	return &GetAppsAppID{Context: ctx, Handler: handler}
}

/*GetAppsAppID swagger:route GET /apps/{appId} apps getAppsAppId

Gets an app

Returns a single app by its id

*/
type GetAppsAppID struct {
	Context *middleware.Context
	Handler GetAppsAppIDHandler
}

func (o *GetAppsAppID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAppsAppIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
