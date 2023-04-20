package main

import (
	"encoding/json"
	"fmt"
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
	type LoginRequest struct {
		Username string `json: "username"`
		Password string `json: "password"`
	}
	var loginReq LoginRequest

	// Attempt to decode Response
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		fmt.Println("Unable to decode login Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//validate credentials against DB
	valid, err := !validateCredentials(loginReq.Username, loginReq.Password)
	if err != nil {

	}
	token := generateLoginToken(loginReq.Username)

	w.Write([]byte(token))
}

func validateCredentials(username string, password string) (bool err) {

}

func generateLoginToken(username string) string {

}
