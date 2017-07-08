package rest

import (
	"github.com/gorilla/mux"
	"github.com/Mr-Pi/dos-backend/rest/homeHandler"
	"github.com/Mr-Pi/dos-backend/rest/userListHandler"
	"net/http"
	"github.com/Mr-Pi/dos-backend/config"
)

func InitRouter(cfg config.Config) {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.Handle).Methods("GET")
	r.HandleFunc("/user", userListHandler.Handle).Methods("GET")
	http.ListenAndServe(cfg.Listen.Listen, r)
}
