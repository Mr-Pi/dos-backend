package userHandler

import (
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GETUser(request *restful.Request, response *restful.Response) {
	username := request.PathParameter("username")
	log.Println("Request User", username)
	if pgsql.TestUser(username) {
		user := pgsql.GETUser(username)
		response.WriteEntity(user)
	} else {
	}
}
func HandleMOD(w http.ResponseWriter, r *http.Request) {
	log.Println("Mod User", mux.Vars(r)["username"])
	w.Header().Set("Content-Type", "application/json")
}
