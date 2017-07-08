package rest

import (
	"github.com/gorilla/mux"
	"github.com/Mr-Pi/dos-backend/rest/homeHandler"
	"net/http"
)

func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.Handle).Methods("GET")
	http.ListenAndServe(":8080", r)
}
