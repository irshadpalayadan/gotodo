package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	todo "github.com/irshadpalayadan/gotodo/package/todo"
	log "github.com/sirupsen/logrus"
)

func getServerStatus(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(res).Encode("Serever running successfully")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/status", getServerStatus).Methods("GET")
	router.HandleFunc("/todos", todo.GetTodo).Methods("GET")
	router.HandleFunc("/todos", todo.AddTodo).Methods("POST")

	http.ListenAndServe(":8000", router)

	defer log.Fatal("Router creation failed")

}
