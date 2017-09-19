// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposRepoIDHandlerFunc turns a function with the right signature into a get repos repo ID handler
type GetReposRepoIDHandlerFunc func(GetReposRepoIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposRepoIDHandlerFunc) Handle(params GetReposRepoIDParams) middleware.Responder {
	return fn(params)
}

// GetReposRepoIDHandler interface for that can handle valid get repos repo ID params
type GetReposRepoIDHandler interface {
	Handle(GetReposRepoIDParams) middleware.Responder
}

// NewGetReposRepoID creates a new http.Handler for the get repos repo ID operation
func NewGetReposRepoID(ctx *middleware.Context, handler GetReposRepoIDHandler) *GetReposRepoID {
	return &GetReposRepoID{Context: ctx, Handler: handler}
}

/*GetReposRepoID swagger:route GET /repos/{repoId} repos getReposRepoId

Gets a repo

Returns a single repo by its id

*/
type GetReposRepoID struct {
	Context *middleware.Context
	Handler GetReposRepoIDHandler
}

func (o *GetReposRepoID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposRepoIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
