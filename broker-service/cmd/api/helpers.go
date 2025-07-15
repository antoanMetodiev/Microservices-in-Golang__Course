package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // 1MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	jsonDecoder := json.NewDecoder(r.Body)

	// при този процес, се пълнят данните реално, но трябва да се подаде pointer към обекта в data параметъра!
	err := jsonDecoder.Decode(data)
	if err != nil {
		return err
	}

	err = jsonDecoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("Body must have only a single JSON value!")
	}

	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, statusCode int, data any, headers ...http.Header) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for _, header := range headers {
		for key, values := range header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
	}

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")

	_, err2 := w.Write(jsonData)
	if err2 != nil {
		return err2
	}

	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := JsonResponse{Error: true, Message: err.Error()}
	return app.writeJSON(w, statusCode, payload)
}
