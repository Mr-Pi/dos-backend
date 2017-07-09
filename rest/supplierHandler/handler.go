package supplierHandler

import (
	"encoding/json"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	suppliers := pgsql.ListSuppliers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suppliers)
}

func HandleGET(w http.ResponseWriter, r *http.Request) {
	supplierID, _ := strconv.ParseInt(mux.Vars(r)["supplier"], 10, 64)
	log.Println("Request Supplier", supplierID)
	w.Header().Set("Content-Type", "application/json")
	if pgsql.TestSupplier(supplierID) {
		supplier := pgsql.GETSupplier(supplierID)
		json.NewEncoder(w).Encode(supplier)
	} else {
		status, _ := json.Marshal(http.StatusText(404))
		http.Error(w, string(status), 404)
	}
}

func HandleMOD(w http.ResponseWriter, r *http.Request) {
	log.Println("Mod Supplier", mux.Vars(r)["username"])
	w.Header().Set("Content-Type", "application/json")
}
