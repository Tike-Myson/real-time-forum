package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homePage)
	return app.logRequest(app.secureHeaders(app.recoverPanic(mux)))

}