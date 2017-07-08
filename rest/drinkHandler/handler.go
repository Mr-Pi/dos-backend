package drinkHandler

import (
	"encoding/json"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	drinks := pgsql.ListDrinks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drinks)
}

func HandleGET(w http.ResponseWriter, r *http.Request) {
	drinkEAN := mux.Vars(r)["drink"]
	log.Println("Request Drink", drinkEAN)
	w.Header().Set("Content-Type", "application/json")
	if pgsql.TestDrink(drinkEAN) {
		drink := pgsql.GETDrink(drinkEAN)
		json.NewEncoder(w).Encode(drink)
	} else {
		status, _ := json.Marshal(http.StatusText(404))
		http.Error(w, string(status), 404)
	}
}

func HandleMOD(w http.ResponseWriter, r *http.Request) {
	log.Println("Mod Drink", mux.Vars(r)["username"])
	w.Header().Set("Content-Type", "application/json")
}
