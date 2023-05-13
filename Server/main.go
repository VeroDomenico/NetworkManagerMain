package main

import (
	"fmt"
	"log"
	"net/http"
	mongodbsetup "networkMangerBackend/internal/MongoDBSetup"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	type Person struct {
		FirstName string `json:"firstName"`
	}
	var doc Person
	var doc2 []Person

	if err := mongodbsetup.FindOne(bson.M{"firstName": "John"}, &doc, "TestCollection"); err != nil {
		// handle the error
		log.Println("Failed to query the database with error: " + err.Error())
		return
	}
	if err := mongodbsetup.FindAll(bson.M{"firstName": "John"}, &doc2, "TestCollection"); err != nil {
		// handle the error
		log.Println("Failed to query the database with error: " + err.Error())
		return
	}

	fmt.Println(doc)
	fmt.Println(doc2)
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
