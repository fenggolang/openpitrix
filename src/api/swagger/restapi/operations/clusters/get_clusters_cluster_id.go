// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetClustersClusterIDHandlerFunc turns a function with the right signature into a get clusters cluster ID handler
type GetClustersClusterIDHandlerFunc func(GetClustersClusterIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClustersClusterIDHandlerFunc) Handle(params GetClustersClusterIDParams) middleware.Responder {
	return fn(params)
}

// GetClustersClusterIDHandler interface for that can handle valid get clusters cluster ID params
type GetClustersClusterIDHandler interface {
	Handle(GetClustersClusterIDParams) middleware.Responder
}

// NewGetClustersClusterID creates a new http.Handler for the get clusters cluster ID operation
func NewGetClustersClusterID(ctx *middleware.Context, handler GetClustersClusterIDHandler) *GetClustersClusterID {
	return &GetClustersClusterID{Context: ctx, Handler: handler}
}

/*GetClustersClusterID swagger:route GET /clusters/{clusterId} clusters getClustersClusterId

Gets a cluster

Returns a single cluster by its id

*/
type GetClustersClusterID struct {
	Context *middleware.Context
	Handler GetClustersClusterIDHandler
}

func (o *GetClustersClusterID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetClustersClusterIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
