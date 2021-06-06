package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

const SecretKey = "secret"

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.clientError(w, http.StatusNotFound)
		return
	}
	err := app.createAllTables()
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/register" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

	case "POST":
		var data map[string]string
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			app.serverError(w, err)
			return
		}
		password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
		if err != nil {
			app.serverError(w, err)
			return
		}
		age, err := strconv.Atoi(data["age"])
		if err != nil {
			app.serverError(w, err)
			return
		}
		user := models.User{
			Nickname:  data["nickname"],
			Age:       age,
			Gender:    data["gender"],
			FirstName: data["first_name"],
			LastName:  data["last_name"],
			Email:     data["email"],
			Password: password,
		}
		err = app.users.CreateUser(user)
		if err != nil {
			if err.Error() == "UNIQUE constraint failed: users.email" {
				fmt.Fprintf(w, "Duplicate Email")
				return
			}
			if err.Error() == "UNIQUE constraint failed: users.nickname" {
				fmt.Fprintf(w, "Duplicate Nickname")
				return
			}
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(user)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/login" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

	case "POST":
		var data map[string]string
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			app.serverError(w, err)
			return
		}

		user := models.User{
			Email: data["email"],
			Password: []byte(data["password"]),
		}

		_, err = app.users.Authenticate(user.Email, user.Password)
		if err != nil {
			if errors.Is(err, models.ErrInvalidCredentials) {
				fmt.Fprintf(w, "Invalid credentials")
				return
			} else {
				app.serverError(w, err)
				return
			}
		}
		cookie := MakeCookie(user.Email)
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/logout" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

	case "POST":

	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) user(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/user" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

	case "POST":

	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}
