package handler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/emicklei/go-restful"
	"net/http"
	"strconv"
)

func GetUser(request *restful.Request, response *restful.Response) {
	username := request.PathParameter("username")
	if pgsql.TestUser(username) {
		response.WriteEntity(pgsql.GETUser(username))
	} else {
		response.WriteErrorString(http.StatusNotFound, "User Not Found")
	}
}

func ListUsers(request *restful.Request, response *restful.Response) {
	response.WriteEntity(pgsql.ListUsers())
}

func GetDrink(request *restful.Request, response *restful.Response) {
	drinkEAN := request.PathParameter("ean")
	if pgsql.TestDrink(drinkEAN) {
		response.WriteEntity(pgsql.GETDrink(drinkEAN))
	} else {
		response.WriteErrorString(http.StatusNotFound, "Drink Not Found")
	}
}

func ListDrinks(request *restful.Request, response *restful.Response) {
	response.WriteEntity(pgsql.ListDrinks())
}

func GetSupplier(request *restful.Request, response *restful.Response) {
	supplierID, _ := strconv.ParseInt(request.PathParameter("id"), 10, 64)
	if pgsql.TestSupplier(supplierID) {
		response.WriteEntity(pgsql.GETSupplier(supplierID))
	} else {
		response.WriteErrorString(http.StatusNotFound, "No such supplier")
	}
}

func ListSuppliers(request *restful.Request, response *restful.Response) {
	response.WriteEntity(pgsql.ListSuppliers())
}
