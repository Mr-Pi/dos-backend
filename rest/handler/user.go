package handler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/Mr-Pi/dos-backend/types"
	"github.com/emicklei/go-restful"
	"net/http"
)

var requiredPermissionsUserInfo = types.UserPermissions{
	ModUser:            true,
	PatchDrinkEveryone: true,
}

var requiredPermissionsUserMod = types.UserPermissions{
	ModUser: true,
}

func ListUsers(request *restful.Request, response *restful.Response) {
	username, rc := permissions.ReqToUser(request.Request)
	if rc != 200 {
		response.WriteErrorString(rc, "Can't check permissions")
		return
	}
	rc = permissions.CheckUserPermissions(username, requiredPermissionsUserInfo)
	if rc != 200 {
		response.WriteErrorString(rc, "You need more permissions to list all users")
		return
	}
	response.WriteEntity(pgsql.ListUsers())
}

func GetUser(request *restful.Request, response *restful.Response) {
	ownUser, rc := permissions.ReqToUser(request.Request)
	username := request.PathParameter("username")
	if rc != 200 {
		response.WriteErrorString(rc, "Can't check permissions")
		return
	}

	rc = permissions.CheckUserPermissions(ownUser, requiredPermissionsUserInfo)
	if ownUser == username || rc == 200 {
		if pgsql.TestUser(username) {
			response.WriteEntity(pgsql.GETUser(username))
		} else {
			response.WriteErrorString(http.StatusNotFound, "User Not Found")
		}
	} else {
		response.WriteErrorString(rc, "You need more permissions to view user")
	}
}

/*
func AddUser(request *restful.Request, response *restful.Response) {
	ownUser, rc := permissions.ReqToUser(request.Request)
	username := request.PathParameter("username")
	if rc != 200 {
		response.WriteErrorString(rc, "Can't check permissions")
		return
	}
	var user = struct{
	Username    string `json:"uid"`
	FirstName   string `json:"givenName"`
	LastName    string `json:"surname"`
	Credit      float64 `json:"credit"`
	Password    string `json:"-"`
	RFIDTag     []string `json:"rfidTags"`
}
}*/
