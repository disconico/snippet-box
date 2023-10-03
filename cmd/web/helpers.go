package main

import (
	"net/http"
)

func (app *application) logRequests(next http.Handler) http.Handler {
	handlerWithLogs := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			app.logger.Info("Incoming request", "method", r.Method, "uri", r.URL.RequestURI(), "remote_addr", r.RemoteAddr)
			next.ServeHTTP(w, r)
		},
	)

	return handlerWithLogs
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {

	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFoundError(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
