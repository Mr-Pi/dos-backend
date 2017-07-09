package handler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/emicklei/go-restful"
	"net/http"
)

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
