package sqlite3

import (
	"database/sql"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
)

type CommentModel struct {
	DB *sql.DB
}

func (m *CommentModel) CreateCommentsTable() error {
	commentsTable, err := m.DB.Prepare(CreateCommentsTableSQL)
	if err != nil {
		return err
	}
	_, err = commentsTable.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (m *CommentModel) InsertCommentIntoDB(commentData models.Comment) error {
	insertPost, err := m.DB.Prepare(InsertCommentSQL)
	if err != nil {
		return err
	}
	_, err = insertPost.Exec(
		commentData.PostId,
		commentData.Author,
		commentData.Content,
		commentData.CreatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}