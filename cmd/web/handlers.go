package main

import (
	"net/http"
)

func (app *application) homePage(w http.ResponseWriter, r *http.Request) {
	err := app.createAllTables()
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) returnScoreboard(w http.ResponseWriter, r *http.Request){

}

func (app *application) createNewPlayer(w http.ResponseWriter, r *http.Request) {

}
