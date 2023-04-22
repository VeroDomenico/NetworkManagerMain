package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Check session

func authenticateLogin(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	var LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//Check request is valid to be decoded in bson for go
	if json.NewDecoder(r.Body).Decode(&LoginRequest) != nil {
		fmt.Println("Unable to decode login Request")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}

	//Authenticate authenticity of the user
	//validate credentials against DB
	valid := !validateCredentials(LoginRequest.Username, LoginRequest.Username)

	if valid {
		generateLoginToken(LoginRequest.Username)
	} else {

		w.WriteHeader(http.StatusUnauthorized)
		return w, r
	}

	w.WriteHeader(http.StatusOK)
	return w, r
}

// TODO
// func authorizeAccess() {

// }

// TODO get from mongo db
func validateCredentials(username string, password string) bool {

	// Get the user from the database.
	// user, err := db.GetUser(username)
	// if err != nil {
	// 	// The user does not exist.
	// 	return false
	// }

	// // Compare the user's password with the password that was submitted.
	// if user.Password != password {
	// 	// The password is incorrect.
	// 	return false
	// }

	// // The user's credentials are valid.
	// return true
}

// Used for authorizaiton TODO
func generateLoginToken(username string) string {

	// Generate a random token.
	// token := uuid.New().String()

	// Save the token to the database.
	// err := db.SaveLoginToken(username, token)
	// if err != nil {
	// 	// There was an error saving the token to the database.
	// 	return ""
	// }

	// The token was generated successfully.
	// return token
	return ""
}

//check token
