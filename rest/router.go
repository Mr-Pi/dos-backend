package rest

import (
	"github.com/gorilla/mux"
	"github.com/Mr-Pi/dos-backend/rest/homeHandler"
	"github.com/Mr-Pi/dos-backend/rest/userListHandler"
	"net/http"
)

func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.Handle).Methods("GET")
	r.HandleFunc("/user", userListHandler.Handle).Methods("GET")
	http.ListenAndServe(":8080", r)
}
