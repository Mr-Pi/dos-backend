package userListHandler

import (
	"net/http"
	"github.com/Mr-Pi/dos-backend/permissions"
	"encoding/json"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	user, rc := permissions.TokenToUser(r)
	if rc != 0 {
		w.WriteHeader(rc)
		// TODO Request body
		return
	}
	/* FIXME rc = permissions.CheckUserListPermissions(user)
	if rc != 0 {
		w.WriteHeader(rc)
		// TODO Request body
		return
	}*/
	var users []string
	users = append(users, "janne")
	users = append(users, "fritz")
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
