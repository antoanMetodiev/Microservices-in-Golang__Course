package main

import (
	"log"
	"net/http"
)

type UserAuthRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	user := UserAuthRequest{}

	err := app.readJSON(w, r, user)
	if err != nil {
		log.Panic("******* " + err.Error())
		app.errorJSON(w, err)
		return
	}

	// Validate Email and Password:
	userResponse, err := app.Models.User.GetByEmail(user.Email)
	if err != nil {
		log.Panic("******* " + err.Error())
		app.errorJSON(w, err)
		return
	}

	validUser, err := userResponse.PasswordMatches(user.Password)
	if err != nil {
		log.Panic("******* " + err.Error())
		app.errorJSON(w, err)
		return
	}

	jsonResponse := JsonResponse{
		Error:   false,
		Message: "",
		Data:    validUser,
	}

	app.writeJSON(w, http.StatusOK, jsonResponse)
}
