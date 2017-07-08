package userHandler

import (
	"encoding/json"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleGET(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	log.Println("Request User", username)
	w.Header().Set("Content-Type", "application/json")
	if pgsql.TestUser(username) {
		user := pgsql.GETUser(username)
		json.NewEncoder(w).Encode(user)
	} else {
		status, _ := json.Marshal(http.StatusText(404))
		http.Error(w, string(status), 404)
	}
}
func HandleMOD(w http.ResponseWriter, r *http.Request) {
	log.Println("Mod User", mux.Vars(r)["username"])
	w.Header().Set("Content-Type", "application/json")
}
