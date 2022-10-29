package main

import (
	"net/http"

	"example.com/jsonapi/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK mama",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) hello(w http.ResponseWriter, r *http.Request) {
	err := response.JSON(w, 200, "Hello world!")
	if err != nil {
		app.serverError(w, r, err)
	}
}
