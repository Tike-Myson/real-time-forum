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
	UserId    string    `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ImageURL  string    `json:"image_url,omitempty"`
	Rating    int       `json:"rating,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
}

type Comment struct {
	Id        int	    `json:"id,omitempty"`
	PostId    string    `json:"post_id,omitempty"`
	UserId    string    `json:"user_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Rating    int       `json:"rating,omitempty"`
}

type Message struct {
	Id         string    `json:"id,omitempty"`
	Content    string    `json:"content,omitempty"`
	DialogId   string    `json:"dialog_id,omitempty"`
	SenderId   string    `json:"sender_id,omitempty"`
	ReceiverId string    `json:"receiver_id,omitempty"`
	IsRead     bool      `json:"is_read,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type Dialog struct {
	Id         string    `json:"id,omitempty"`
	SenderId   string    `json:"sender_id,omitempty"`
	ReceiverId string    `json:"receiver_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	Messages   []Message `json:"messages"`
}

type Category struct {
	Id int
	Name string
}

type RatingPost struct {
	PostId int
	UserId int
	Value int
}

type RatingComment struct {
	CommentId int
	UserId int
	Value int
}

type CategoryPostLink struct {
	PostId int
	CategoryId int
}
