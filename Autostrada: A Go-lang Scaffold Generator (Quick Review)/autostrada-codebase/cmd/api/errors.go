package main

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/jsonapi/internal/response"
	"example.com/jsonapi/internal/validator"
)

func (app *application) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		app.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err)

	message := "The server encountered a problem and could not process your request"
	app.errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	app.errorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func (app *application) failedValidation(w http.ResponseWriter, r *http.Request, v validator.Validator) {
	err := response.JSON(w, http.StatusUnprocessableEntity, v)
	if err != nil {
		app.serverError(w, r, err)
	}
}
