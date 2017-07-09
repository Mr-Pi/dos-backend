package rest

import (
	"github.com/Mr-Pi/dos-backend/config"
	"github.com/Mr-Pi/dos-backend/rest/drinkHandler"
	"github.com/Mr-Pi/dos-backend/rest/homeHandler"
	"github.com/Mr-Pi/dos-backend/rest/permHandler"
	"github.com/Mr-Pi/dos-backend/rest/supplierHandler"
	"github.com/Mr-Pi/dos-backend/rest/userHandler"
	"github.com/Mr-Pi/dos-backend/rest/userListHandler"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter(cfg config.Config) {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.Handle).Methods("GET")
	r.HandleFunc("/user", userListHandler.Handle).Methods("GET")
	r.HandleFunc("/user/{username}", userHandler.HandleGET).Methods("GET")
	r.HandleFunc("/user/{username}", userHandler.HandleMOD).Methods("PUT", "PATCH")
	r.HandleFunc("/drink", drinkHandler.HandleList).Methods("GET")
	r.HandleFunc("/drink/{drink}", drinkHandler.HandleGET).Methods("GET")
	r.HandleFunc("/drink/{drink}", drinkHandler.HandleMOD).Methods("PUT", "PATCH")
	r.HandleFunc("/supplier", supplierHandler.HandleList).Methods("GET")
	r.HandleFunc("/supplier/{supplier}", supplierHandler.HandleGET).Methods("GET")
	r.HandleFunc("/supplier/{supplier}", supplierHandler.HandleMOD).Methods("PUT", "PATCH")
	r.HandleFunc("/permissions", permissionHandler.HandleList).Methods("GET")
	r.HandleFunc("/permissions/{permType}", permissionHandler.HandleGET).Methods("GET")
	r.HandleFunc("/permissions/{permType}", permissionHandler.HandleMOD).Methods("PUT", "PATCH")
	http.ListenAndServe(cfg.Listen.Listen, r)
}
