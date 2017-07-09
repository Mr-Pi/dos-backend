package rest

import (
	"github.com/Mr-Pi/dos-backend/config"
	"github.com/Mr-Pi/dos-backend/rest/handler"
	"github.com/Mr-Pi/dos-backend/rest/tokenHandler"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type SupplierResource struct {
}

func (p SupplierResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/supplier")
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
	ws.Path("/drink")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{ean}").To(handler.GetDrink))
	ws.Route(ws.GET("").To(handler.ListDrinks))
	restful.Add(ws)
}

type UserResource struct {
}

func (p UserResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/user")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{username}").To(handler.GetUser))
	ws.Route(ws.GET("").To(handler.ListUsers))
	restful.Add(ws)
}

type TokenResource struct {
}

func (p TokenResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/token")
	ws.Produces(restful.MIME_JSON)
	// TODO ws.Route(ws.DELETE("/{token}").To(handler.GetUser))
	ws.Route(ws.POST("").To(tokenHandler.RequestToken))
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
