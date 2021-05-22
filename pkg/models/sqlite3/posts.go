package sqlite3

import (
	"database/sql"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
	"github.com/Tike-Myson/real-time-forum/pkg/models/sqlite3"
)

func createPostsTable(db *sql.DB) error {
	postsTable, err := db.Prepare(sqlite3.CreatePostsTableSQL)
	if err != nil {
		return err
	}
	_, err = postsTable.Exec()
	if err != nil {
		return err
	}
	return nil
}

func insertPostIntoDB(db *sql.DB, postData models.Post) error {
	insertPost, err := db.Prepare(sqlite3.InsertPostSQL)
	if err != nil {
		return err
	}
	_, err = insertPost.Exec(
			postData.Title,
			postData.Content,
			postData.Author,
			postData.CreatedAt,
			postData.ImageURL,
		)

	if err != nil {
		return err
	}
	return nil
}