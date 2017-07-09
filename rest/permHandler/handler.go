package permissionHandler

import (
	"encoding/json"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	permissions := pgsql.ListPermissions()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}

func HandleGET(w http.ResponseWriter, r *http.Request) {
	permType, _ := mux.Vars(r)["permType"]
	log.Println("Request permissions", permType)
	w.Header().Set("Content-Type", "application/json")
	if pgsql.TestPermissions(permType) {
		permissions := pgsql.GETPermissions(permType)
		json.NewEncoder(w).Encode(permissions)
	} else {
		status, _ := json.Marshal(http.StatusText(404))
		http.Error(w, string(status), 404)
	}
}

func HandleMOD(w http.ResponseWriter, r *http.Request) {
	log.Println("Mod permissions", mux.Vars(r)["permType"])
	w.Header().Set("Content-Type", "application/json")
}
