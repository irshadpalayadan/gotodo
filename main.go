package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func getServerStatus(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(res).Encode("Serever running successfully")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/status", getServerStatus).Methods("GET")

	http.ListenAndServe(":8000", router)

	defer log.Fatal("Router creation failed")

}
