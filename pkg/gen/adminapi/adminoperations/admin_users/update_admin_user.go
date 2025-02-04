// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateAdminUserHandlerFunc turns a function with the right signature into a update admin user handler
type UpdateAdminUserHandlerFunc func(UpdateAdminUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateAdminUserHandlerFunc) Handle(params UpdateAdminUserParams) middleware.Responder {
	return fn(params)
}

// UpdateAdminUserHandler interface for that can handle valid update admin user params
type UpdateAdminUserHandler interface {
	Handle(UpdateAdminUserParams) middleware.Responder
}

// NewUpdateAdminUser creates a new http.Handler for the update admin user operation
func NewUpdateAdminUser(ctx *middleware.Context, handler UpdateAdminUserHandler) *UpdateAdminUser {
	return &UpdateAdminUser{Context: ctx, Handler: handler}
}

/*
	UpdateAdminUser swagger:route PATCH /admin-users/{adminUserId} Admin users updateAdminUser

# Updates an Admin User

This endpoint updates a single Admin User by ID. Do not use this
endpoint directly as it is meant to be used with the Admin UI exclusively.
*/
type UpdateAdminUser struct {
	Context *middleware.Context
	Handler UpdateAdminUserHandler
}

func (o *UpdateAdminUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateAdminUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
