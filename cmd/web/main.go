package main

import (
	"flag"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
	"github.com/Tike-Myson/real-time-forum/pkg/models/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	posts interface{
		CreatePostsTable() error
		InsertPostIntoDB(models.Post) error
		Get() ([]models.Post, error)
		GetPostById(string) (models.Post, error)
	}
	categories interface{
		CreateCategoriesTable() error
	}
	categoryPostLinks interface{
		CreateCategoryPostLinksTable() error
	}
	comments interface{
		CreateCommentsTable() error
		InsertCommentIntoDB(models.Comment) error
	}
	ratings interface{
		CreateRatingsTable() error
		InsertPostRating(string, string, int) error
		InsertCommentRating(string, string, int) error
	}
	users interface{
		CreateUsersTable() error
		CreateUser(models.User) error
		Authenticate(string, []byte) (int, error)
	}
	dialogs interface{
		CreateDialogsTable() error
		InsertDialogIntoDB(models.Dialog) error
	}
	messages interface{
		CreateMessagesTable() error
		InsertMessageIntoDB(models.Message) error
	}
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	dsn := flag.String("dsn", "./forum.db", "Sqlite3 data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, Green+"INFO\t"+Reset, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, Red+"ERROR\t"+Reset, log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		errorLog.Println(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		posts: &sqlite3.PostModel{DB: db},
		categories: &sqlite3.CategoryModel{DB: db},
		categoryPostLinks: &sqlite3.CategoryPostLinkModel{DB: db},
		comments: &sqlite3.CommentModel{DB: db},
		ratings: &sqlite3.RatingModel{DB: db},
		users: &sqlite3.UserModel{DB: db},
		dialogs: &sqlite3.DialogsModel{DB: db},
		messages: &sqlite3.MessagesModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Server run on http://127.0.0.1%s\n", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
