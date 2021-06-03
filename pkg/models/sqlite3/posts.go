package sqlite3

import (
	"database/sql"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) CreatePostsTable() error {
	postsTable, err := m.DB.Prepare(CreatePostsTableSQL)
	if err != nil {
		return err
	}
	_, err = postsTable.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (m *PostModel) InsertPostIntoDB(postData models.Post) error {
	insertPost, err := m.DB.Prepare(InsertPostSQL)
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

func (m *PostModel) Get() ([]models.Post, error) {
	var CurrentPost models.Post
	var Posts []models.Post

	rows, err := m.DB.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&CurrentPost.Id, &CurrentPost.Title, &CurrentPost.Content, &CurrentPost.Author, &CurrentPost.CreatedAt, &CurrentPost.ImageURL)
		Posts = append(Posts, CurrentPost)
	}
	return Posts, err
}