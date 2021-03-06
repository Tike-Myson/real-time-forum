package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	if r.URL.Path != "/api/user/register" {
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
	if r.URL.Path != "/api/user/login" {
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
	if r.URL.Path != "/api/user/logout" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		for _, cookie := range r.Cookies() {
			dCookie := DeleteCookie(cookie.Value)
			http.SetCookie(w, &dCookie)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
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

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/post/create" {
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

		userId, err := app.getUserId(r)
		if err != nil {
			app.serverError(w, err)
			return
		}

		post := models.Post{
			Title: data["title"],
			Content: data["content"],
			UserId: userId,
			CreatedAt: time.Now(),
			ImageURL: data["image_url"],
		}
		err = app.posts.InsertPostIntoDB(post)
		if err != nil {
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(post)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) createDialog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/dialog/create" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		var data map[string]string
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			app.serverError(w, err)
			return
		}

		userId, err := app.getUserId(r)
		if err != nil {
			app.serverError(w, err)
			return
		}

		post := models.Post{
			Title: data["title"],
			Content: data["content"],
			UserId: userId,
			CreatedAt: time.Now(),
			ImageURL: data["image_url"],
		}
		err = app.posts.InsertPostIntoDB(post)
		if err != nil {
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(post)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) createMessage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/post/create" {
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
		userId, err := app.getUserId(r)
		if err != nil {
			app.serverError(w, err)
			return
		}

		post := models.Post{
			Title: data["title"],
			Content: data["content"],
			UserId: userId,
			CreatedAt: time.Now(),
			ImageURL: data["image_url"],
		}
		err = app.posts.InsertPostIntoDB(post)
		if err != nil {
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(post)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) getAllPosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		posts, err := app.posts.Get()
		if err != nil {
			app.serverError(w, err)
			return
		}

		for _, post := range posts {
			post.Comments, err = app.ratings.GetCommentsByPostId(post.Id)
			if err != nil {
				app.serverError(w, err)
				return
			}
			post.Rating, err = app.ratings.GetRatingById(post.Id, "post")
			if err != nil {
				app.serverError(w, err)
				return
			}
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		json.NewEncoder(w).Encode(posts)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) showPost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		u, _ := url.Parse(r.URL.Path)
		id := getFirstParam(u.Path)
		post, err := app.posts.GetPostById(id)
		if err != nil {
			if err == models.ErrNoRecord {
				app.clientError(w, http.StatusNotFound)
				return
			}
			app.serverError(w, err)
			return
		}
		post.Comments, err = app.ratings.GetCommentsByPostId(post.Id)
		if err != nil {
			app.serverError(w, err)
			return
		}
		post.Rating, err = app.ratings.GetRatingById(post.Id, "post")
		if err != nil {
			app.serverError(w, err)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		json.NewEncoder(w).Encode(post)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) showDialog(w http.ResponseWriter, r *http.Request) {

}

func (app *application) serveWs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ws, err := app.Upgrade(w, r)
		if err != nil {
			app.serverError(w, err)
			return
		}
		go Writer(ws)
		Reader(ws)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

}

func (app *application) createComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/comment/create" {
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
		userId, err := app.getUserId(r)
		if err != nil {
			app.serverError(w, err)
			return
		}
		comment := models.Comment{
			PostId: data["post_id"],
			UserId: userId,
			Content: data["content"],
			CreatedAt: time.Now(),
		}

		err = app.comments.InsertCommentIntoDB(comment)
		if err != nil {
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(comment)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) likePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/post/like" {
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
		userId, err := app.getUserId(r)
		if err != nil {
			app.serverError(w, err)
			return
		}
		postId, err := strconv.Atoi(data["post_id"])
		if err != nil {
			app.serverError(w, err)
			return
		}
		err = app.ratings.InsertPostRating(userId, postId, 1)
		if err != nil {
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(data)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) dislikePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/post/dislike" {
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
		userId, err := app.getUserId(r)
		if err != nil {
			app.serverError(w, err)
			return
		}
		postId, err := strconv.Atoi(data["post_id"])
		if err != nil {
			app.serverError(w, err)
			return
		}
		err = app.ratings.InsertPostRating(userId, postId, -1)
		if err != nil {
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(data)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) likeComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/comment/like" {
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
		userId, err := app.getUserId(r)
		if err != nil {
			app.serverError(w, err)
			return
		}
		postId, err := strconv.Atoi(data["post_id"])
		if err != nil {
			app.serverError(w, err)
			return
		}
		err = app.ratings.InsertCommentRating(userId, postId, 1)
		if err != nil {
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(data)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) dislikeComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/comment/dislike" {
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
		userId, err := app.getUserId(r)
		if err != nil {
			app.serverError(w, err)
			return
		}
		postId, err := strconv.Atoi(data["post_id"])
		if err != nil {
			app.serverError(w, err)
			return
		}
		err = app.ratings.InsertCommentRating(userId, postId, -1)
		if err != nil {
			app.serverError(w, err)
			return
		}
		json.NewEncoder(w).Encode(data)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}