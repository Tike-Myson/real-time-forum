package main

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

var Cookies = make(map[string]string)

func (app *application) GetUserIdByCookie(cookie *http.Cookie) (int, error) {
	ok, login := IsTokenExists(cookie.Value)
	if ok {
		id, err := app.users.GetUserIdByLogin(login)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, nil
}

func MakeCookie(login string) http.Cookie {
	u1 := uuid.NewV4()

	Cookies[u1.String()] = login

	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name: "session",
		Value: u1.String(),
		Expires: expiration,
		HttpOnly: false,
		Path: "/",
	}
	return cookie
}

func DeleteCookie(token string) http.Cookie {
	delete(Cookies, token)
	cookie := http.Cookie{
		Name: "session",
		Path: "/",
		MaxAge: -1,
	}
	return cookie
}

func IsSessionExists(username string) (bool, string) {
	for i, v := range Cookies {
		if username == v {
			return true, i
		}
	}
	return false, ""
}

func IsTokenExists(token string) (bool, string) {
	i, found := Cookies[token]
	if found {
		return true, i
	}
	return false, ""
}

