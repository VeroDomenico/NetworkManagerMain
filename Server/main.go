package main

import (
	"fmt"
	"log"
	"net/http"
	mongodbsetup "networkMangerBackend/internal/MongoDBSetup"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	firstName string
}

func main() {

	var doc Person

	mongodbsetup.FindOne(bson.M{"firstName": "John"}, &doc, "NetworkManager", "TestCollection")
	fmt.Println(doc.firstName)
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
