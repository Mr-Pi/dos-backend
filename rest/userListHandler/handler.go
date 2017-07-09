package userListHandler

import (
	"encoding/json"
	//"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"net/http"
	"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/Mr-Pi/dos-backend/types"
)

var requiredPermissions = types.UserPermissions{
	ModUser:            true,
	PatchDrinkEveryone: true,
}

func Handle(w http.ResponseWriter, r *http.Request) {
	username, rc := permissions.ReqToUser(r)
	if rc != 200 {
		w.WriteHeader(rc)
		return
	}
	rc = permissions.CheckUserPermissions(username, requiredPermissions)
	if rc != 200 {
		w.WriteHeader(rc)
		return
	}
	users := pgsql.ListUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
