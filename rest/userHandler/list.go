package userHandler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/Mr-Pi/dos-backend/types"
	"github.com/emicklei/go-restful"
)

var requiredPermissions = types.UserPermissions{
	ModUser:            true,
	PatchDrinkEveryone: true,
}

func ListUsers(req *restful.Request, resp *restful.Response) {
	username, rc := permissions.ReqToUser(req.Request)
	if rc != 200 {
		resp.WriteErrorString(rc, "Can't check permissions")
		return
	}
	rc = permissions.CheckUserPermissions(username, requiredPermissions)
	if rc != 200 {
		resp.WriteErrorString(rc, "You need more permissions to list all users")
		return
	}
	resp.WriteEntity(pgsql.ListUsers())
}
