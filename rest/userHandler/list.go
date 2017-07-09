package userHandler

import (
	"github.com/emicklei/go-restful"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/Mr-Pi/dos-backend/types"
	"errors"
)

var requiredPermissions = types.UserPermissions{
	ModUser:            true,
	PatchDrinkEveryone: true,
}

func ListUsers(req *restful.Request, resp *restful.Response) {
	username, rc := permissions.ReqToUser(req.Request)
	if rc != 200 {
		resp.WriteError(rc, errors.New(""))
		return
	}
	rc = permissions.CheckUserPermissions(username, requiredPermissions)
	if rc != 200 {
		resp.WriteError(rc, errors.New(""))
		return
	}
	resp.WriteEntity(pgsql.ListUsers())
}

