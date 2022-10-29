package main

import (
	"net/http"

	"example.com/jsonapi/internal/response"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.HandlerFunc("GET", "/status", app.status)
	mux.HandlerFunc("GET", "/hello", app.hello)

	mux.HandlerFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, 200, "Homepage 1.0")
	})

	return app.recoverPanic(mux)
}
