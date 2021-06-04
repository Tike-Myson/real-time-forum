package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/api/register", app.register)
	mux.HandleFunc("/api/login", app.login)
	mux.HandleFunc("/api/user", app.user)
	mux.HandleFunc("/api/logout", app.logout)

	return app.logRequest(app.secureHeaders(app.recoverPanic(mux)))

}