package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail = errors.New("models: duplicate email")
	ErrDuplicateUsername = errors.New("models: duplicate login")
)

type User struct {
	Nickname  string `json:"nickname,omitempty"`
	Age       int    `json:"age,omitempty"`
	Gender    string `json:"gender,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  []byte `json:"password,omitempty"`
}

type Post struct {
	Id        int       `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Author    string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ImageURL  string    `json:"image_url,omitempty"`
	Rating    int       `json:"rating,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
}

type Comment struct {
	Id        string    `json:"id,omitempty"`
	PostId    string    `json:"post_id,omitempty"`
	Author    string    `json:"author,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Rating    int       `json:"rating,omitempty"`
}

type Category struct {
	Id int
	Name string
}

type RatingPost struct {
	PostId int
	Author string
	Value int
}

type RatingComment struct {
	CommentId int
	Author string
	Value int
}

type CategoryPostLink struct {
	PostId int
	CategoryId int
}
