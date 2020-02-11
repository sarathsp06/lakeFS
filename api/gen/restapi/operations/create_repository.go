// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/treeverse/lakefs/api/gen/models"
)

// CreateRepositoryHandlerFunc turns a function with the right signature into a create repository handler
type CreateRepositoryHandlerFunc func(CreateRepositoryParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateRepositoryHandlerFunc) Handle(params CreateRepositoryParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// CreateRepositoryHandler interface for that can handle valid create repository params
type CreateRepositoryHandler interface {
	Handle(CreateRepositoryParams, *models.User) middleware.Responder
}

// NewCreateRepository creates a new http.Handler for the create repository operation
func NewCreateRepository(ctx *middleware.Context, handler CreateRepositoryHandler) *CreateRepository {
	return &CreateRepository{Context: ctx, Handler: handler}
}

/*CreateRepository swagger:route POST /repositories createRepository

create repository

*/
type CreateRepository struct {
	Context *middleware.Context
	Handler CreateRepositoryHandler
}

func (o *CreateRepository) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateRepositoryParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}