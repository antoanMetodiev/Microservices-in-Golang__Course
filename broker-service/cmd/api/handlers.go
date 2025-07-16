package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type UserAuthRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	authRequest := UserAuthRequest{}

	err := app.readJSON(w, r, authRequest)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	jsonData, _ := json.MarshalIndent(authRequest, "", "\t")

	// Make Request to Auth Microservice:
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a varabiel we'll read response.Body into
	var jsonFromService JsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload JsonResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}
