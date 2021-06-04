package models

import "time"

type User struct {
	Nickname  string `json:"nickname,omitempty"`
	Age       int `json:"age,omitempty"`
	Gender    string `json:"gender,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  []byte `json:"password,omitempty"`
}

type Post struct {
	Id int
	Title string
	Content string
	Author string
	CreatedAt time.Time
	ImageURL string
	Rating int
	Comments []Comment
}

type Comment struct {
	Id int
	PostId int
	Author string
	Content string
	CreatedAt time.Time
	Rating int
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
