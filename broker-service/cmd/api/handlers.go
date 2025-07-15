package main

import (
	"net/http"
)

func (app *Config) GetBroker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{Error: false, Message: "Hit the broker!"}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}
}
