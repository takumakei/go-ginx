// Package routes provides a way to place route registrations near their implementation.
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route func(r gin.IRouter)

type Routes struct {
	routes []Route
}

func New() *Routes {
	return &Routes{}
}

func (routes *Routes) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	route := func(r gin.IRouter) {
		r.Handle(httpMethod, relativePath, handlers...)
	}
	routes.routes = append(routes.routes, route)
}

func (routes *Routes) POST(relativePath string, handlers ...gin.HandlerFunc) {
	routes.Handle(http.MethodPost, relativePath, handlers...)
}

func (routes *Routes) GET(relativePath string, handlers ...gin.HandlerFunc) {
	routes.Handle(http.MethodGet, relativePath, handlers...)
}

func (routes *Routes) DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	routes.Handle(http.MethodDelete, relativePath, handlers...)
}

func (routes *Routes) PATCH(relativePath string, handlers ...gin.HandlerFunc) {
	routes.Handle(http.MethodPatch, relativePath, handlers...)
}

func (routes *Routes) PUT(relativePath string, handlers ...gin.HandlerFunc) {
	routes.Handle(http.MethodPut, relativePath, handlers...)
}

func (routes *Routes) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) {
	routes.Handle(http.MethodOptions, relativePath, handlers...)
}

func (routes *Routes) HEAD(relativePath string, handlers ...gin.HandlerFunc) {
	routes.Handle(http.MethodHead, relativePath, handlers...)
}

func (routes *Routes) Any(relativePath string, handlers ...gin.HandlerFunc) {
	route := func(r gin.IRouter) {
		r.Any(relativePath, handlers...)
	}
	routes.routes = append(routes.routes, route)
}

func (routes *Routes) Mount(r gin.IRouter) {
	for _, f := range routes.routes {
		f(r)
	}
}
