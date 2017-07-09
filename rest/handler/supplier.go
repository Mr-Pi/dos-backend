package handler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/emicklei/go-restful"
	"net/http"
	"strconv"
)




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