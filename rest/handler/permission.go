package handler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/emicklei/go-restful"
	"net/http"
)

func GetPermission(request *restful.Request, response *restful.Response) {
	permissionType := request.PathParameter("type")
	if pgsql.TestPermissions(permissionType) {
		response.WriteEntity(pgsql.GETPermissions(permissionType))
	} else {
		response.WriteErrorString(http.StatusNotFound, "No such permission group")
	}
}

func ListPermissions(request *restful.Request, response *restful.Response) {
	response.WriteEntity(pgsql.ListPermissions())
}
