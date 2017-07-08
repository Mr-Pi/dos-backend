package userListHandler

import (
	"encoding/json"
	//"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	//user, rc := permissions.TokenToUser(r)
	/* FIXME rc = permissions.CheckUserListPermissions(user)
	if rc != 0 {
		w.WriteHeader(rc)
		// TODO Request body
		return
	}
	if rc != 0 {
		w.WriteHeader(rc)
		// TODO Request body
		return
	}*/
	users := pgsql.ListUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
