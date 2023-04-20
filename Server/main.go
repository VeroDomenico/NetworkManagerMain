package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	log.Default().Print("Starting Server")

	//Create new router
	r := mux.NewRouter()

	//Route Handlers
	r.HandleFunc("/login", login).Methods("POST")

	//
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to Serve with Error: " + err.Error())
	}

}
func login(w http.ResponseWriter, r *http.Request) {

	//Call auth to auth to handle authentication
}
