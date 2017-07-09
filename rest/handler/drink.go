package handler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/Mr-Pi/dos-backend/types"
	"github.com/emicklei/go-restful"
	"net/http"
)

var otherUserPermission = types.UserPermissions{
	PatchDrinkEveryone: true,
}

var requiredPermissionsDrinkMod = types.UserPermissions{
	ModDrink: true,
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

func PutDrink(req *restful.Request, resp *restful.Response) {
	ownUser, rc := permissions.ReqToUser(req.Request)
	ean := req.PathParameter("ean")
	if rc != 200 {
		resp.WriteErrorString(rc, "Can't check permissions")
		return
	}
	rc = permissions.CheckUserPermissions(ownUser, requiredPermissionsDrinkMod)
	if rc != 200 {
		resp.WriteErrorString(rc, "You need more permissions to create drinks")
		return
	}
	drink := new(types.Drink)
	err := req.ReadEntity(user)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	var ok bool
	if pgsql.TestDrink(ean) {
		ok = pgsql.OverwriteDrink(*drink)
	} else {
		ok = pgsql.CreateDrink(*drink)
	}
	if !ok {
		resp.WriteErrorString(http.StatusInternalServerError, "Can not create user")
		return
	}
	resp.WriteHeaderAndEntity(http.StatusCreated, user)
}

func DrinkDrink(req *restful.Request, resp *restful.Response) {
	// Username
	username, rc := permissions.ReqToUser(req.Request)
	if rc != 200 {
		resp.WriteErrorString(rc, "Can't check permissions")
		return
	}
	// Target user
	targetUser := req.PathParameter("username")
	if targetUser != "" {
		rc := permissions.CheckUserPermissions(username, otherUserPermission)
		if rc != 200 {
			resp.WriteErrorString(401, "You are not allowed to do this")
			return
		}
		if !pgsql.TestUser(targetUser) {
			resp.WriteErrorString(404, "Target user not found")
			return
		}
	} else {
		targetUser = username
	}
	// EAN
	drinkEAN := req.PathParameter("ean")
	if !pgsql.TestDrink(drinkEAN) {
		resp.WriteErrorString(http.StatusNotFound, "Drink Not Found")
		return
	}
	// Drink
	drink := pgsql.GETDrink(drinkEAN)
	err := pgsql.DecrementDrinkAmount(drinkEAN, 1)
	if err != nil {
		resp.WriteError(500, err)
		return
	}
	err = pgsql.DecrementUserCredit(targetUser, drink.PriceResell)
	if err != nil {
		resp.WriteError(500, err)
		return
	}
}
