package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/api/user/register", app.register)
	mux.HandleFunc("/api/user/login", app.login)
	mux.HandleFunc("/api/user", app.user)
	mux.HandleFunc("/api/user/logout", app.logout)

	mux.HandleFunc("/api/post/create", app.createPost)
	mux.HandleFunc("/api/post/", app.showPost)
	mux.HandleFunc("/api/post/like", app.likePost)
	mux.HandleFunc("/api/post/dislike", app.dislikePost)

	mux.HandleFunc("/api/comment/create", app.createComment)
	mux.HandleFunc("/api/comment/like", app.likeComment)
	mux.HandleFunc("/api/comment/dislike", app.dislikeComment)

	mux.HandleFunc("/api/dialog/create", app.showPost)
	mux.HandleFunc("/api/message/create", app.showPost)
	mux.HandleFunc("/api/dialog/", app.showDialog)
	mux.HandleFunc("/api/ws/", app.serveWs)

	mux.Handle("/static", http.FileServer(http.Dir("./static")))
	return app.logRequest(app.secureHeaders(app.recoverPanic(mux)))
}