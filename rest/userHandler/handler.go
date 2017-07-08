package userHandler

import (
	"encoding/json"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleGET(w http.ResponseWriter, r *http.Request) {
	log.Println("Request User", mux.Vars(r)["username"])
	w.Header().Set("Content-Type", "application/json")
	user := pgsql.GETUser(mux.Vars(r)["username"])
	json.NewEncoder(w).Encode(user)
}
func HandleMOD(w http.ResponseWriter, r *http.Request) {
	log.Println("Mod User", mux.Vars(r)["username"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(3)
}
