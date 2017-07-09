package handler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/Mr-Pi/dos-backend/types"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

var requiredPermissionsUserInfo = types.UserPermissions{
	ModUser:            true,
	PatchDrinkEveryone: true,
}

var requiredPermissionsUserMod = types.UserPermissions{
	ModUser: true,
}

var requiredPermissionsUserChangePW = types.UserPermissions{
	ModUser:        true,
	SetOwnPassword: true,
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
	if rc != 200 && username != "_self" {
		response.WriteErrorString(rc, "Can't check permissions")
		return
	} else if rc != 200 && username == "_self" {
		response.WriteErrorString(401, "Please login")
		return
	} else if username == "_self" {
		response.WriteEntity(pgsql.GETUser(ownUser))
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

func AddUser(request *restful.Request, response *restful.Response) {
	ownUser, rc := permissions.ReqToUser(request.Request)
	username := request.PathParameter("username")
	if rc != 200 {
		response.WriteErrorString(rc, "Can't check permissions")
		return
	}
	rc = permissions.CheckUserPermissions(ownUser, requiredPermissionsUserMod)
	if rc != 200 {
		response.WriteErrorString(rc, "You need more permissions to create users")
		return
	}
	user := new(types.User)
	err := request.ReadEntity(user)
	user.Username = username
	user.Password, user.Salt = permissions.HashPasswordNew(user.Password)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	err = pgsql.CreateUser(*user)
	if err != nil {
		log.Println("Can't create user", user, err)
		response.WriteErrorString(http.StatusInternalServerError, "User may allready exist")
		return
	}
	response.WriteHeaderAndEntity(http.StatusCreated, user)
}

func UpdateUser(request *restful.Request, response *restful.Response) {
	username := request.PathParameter("username")
	user := pgsql.GETUser(username)
	err := request.ReadEntity(&user)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(permissions.HashPasswordNew("test"))
}
