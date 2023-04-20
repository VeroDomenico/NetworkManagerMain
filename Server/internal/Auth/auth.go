package auth

import (
	"encoding/json"
	"net/http"
)

//Check session

func authenticateLogin(w http.ResponseWriter, r *http.Request) {
	var LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"Password"`
	}
	//Check request is valid to be decoded in bson for go
	if json.NewDecoder(r.Body).Decode(&LoginRequest) != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	//Authenticate authenticity of the user
	//TODO

	w.WriteHeader(http.StatusOK)

}

// TODO
// func authorizeAccess() {

// }
