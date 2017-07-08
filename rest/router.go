package rest

import (
	"github.com/Mr-Pi/dos-backend/config"
	"github.com/Mr-Pi/dos-backend/rest/homeHandler"
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
	r.HandleFunc("/user/{username}", userHandler.HandleMOD).Methods("POST")
	http.ListenAndServe(cfg.Listen.Listen, r)
}
