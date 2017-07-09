package rest

import (
	"github.com/Mr-Pi/dos-backend/config"
	"github.com/Mr-Pi/dos-backend/rest/handler"
	"github.com/Mr-Pi/dos-backend/rest/tokenHandler"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type PermissionResource struct {
}

func (p PermissionResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/permissions")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{type}").To(handler.GetPermission))
	ws.Route(ws.GET("").To(handler.ListPermissions))
	restful.Add(ws)
}

type SupplierResource struct {
}

func (p SupplierResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/suppliers")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{id}").To(handler.GetSupplier))
	ws.Route(ws.GET("").To(handler.ListSuppliers))
	restful.Add(ws)
}

type DrinkResource struct {
}

func (p DrinkResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/drinks")

	ws.Route(ws.POST("/{ean}/order").To(handler.DrinkDrink))
	ws.Route(ws.POST("/{ean}/order/{username}").To(handler.DrinkDrink))

	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("").To(handler.ListDrinks))
	ws.Route(ws.GET("/{ean}").To(handler.GetDrink))
	ws.Route(ws.PUT("/{ean}").To(handler.PutDrink))
	restful.Add(ws)
}

type UserResource struct {
}

func (p UserResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/users")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.PUT("/{username}").To(handler.AddUser))
	ws.Route(ws.GET("/{username}").To(handler.GetUser))
	ws.Route(ws.GET("").To(handler.ListUsers))
	restful.Add(ws)
}

type TokenResource struct {
}

func (p TokenResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/tokens")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.DELETE("").To(tokenHandler.DeleteToken))
	ws.Route(ws.GET("").To(tokenHandler.RequestToken))
	restful.Add(ws)
}

func InitRouter(cfg config.Config) {
	wsContainer := restful.NewContainer()
	userResource := UserResource{}
	userResource.RegisterTo(wsContainer)
	drinkResource := DrinkResource{}
	drinkResource.RegisterTo(wsContainer)
	supplierResource := SupplierResource{}
	supplierResource.RegisterTo(wsContainer)
	permissionResource := PermissionResource{}
	permissionResource.RegisterTo(wsContainer)
	tokenResource := TokenResource{}
	tokenResource.RegisterTo(wsContainer)
	restful.Filter(restful.OPTIONSFilter())
	restful.Filter(globalLogging)
	log.Fatal(http.ListenAndServe(cfg.Listen.Listen, nil))
}

func globalLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[REST] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}
