package sqlite3

import (
	"database/sql"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
	"strconv"
)

type RatingModel struct {
	DB *sql.DB
}

func (m *RatingModel) CreateRatingsTable() error {
	ratingPostsTable, err := m.DB.Prepare(CreateRatingPostSQL)
	if err != nil {
		return err
	}
	_, err = ratingPostsTable.Exec()
	if err != nil {
		return err
	}

	ratingCommentsTable, err := m.DB.Prepare(CreateRatingCommentSQL)
	if err != nil {
		return err
	}
	_, err = ratingCommentsTable.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (m *RatingModel) InsertLikeIntoPost(data models.RatingPost) error {
	insertStmt, err := m.DB.Prepare(InsertRatingPostSQL)
	if err != nil {
		return err
	}
	_, err = insertStmt.Exec(
		data.PostId,
		data.Author,
		data.Value,
	)

	if err != nil {
		return err
	}
	return nil
}

func (m *RatingModel) InsertDislikeIntoPost(data models.RatingPost) error {
	found, err := m.IsRatingExists(data.Author, strconv.Itoa(data.PostId), "post")
	if err != nil {
		return err
	}
	if !found {
		insertStmt, err := m.DB.Prepare(InsertRatingPostSQL)
		if err != nil {
			return err
		}
		_, err = insertStmt.Exec(
			data.PostId,
			data.Author,
			data.Value,
		)

		if err != nil {
			return err
		}
		return nil
	}

}

func (m *RatingModel) InsertLikeIntoComment() {

}

func (m *RatingModel) InsertDislikeIntoComment() {

}

func (m *RatingModel) IsRatingExists(userId, id, flag string) (bool, error) {
	switch flag {
	case "post":
		rows, err := m.DB.Query(SelectPostRatingByID, userId, id)
		if err != nil {
			return false, err
		}
		if rows.Next() {
			return true, err
		}
	case "comment":
		rows, err := m.DB.Query(SelectCommentRatingByID, userId, id)
		if err != nil {
			return false, err
		}
		if rows.Next() {
			return true, err
		}
	default:
		return false, nil
	}
	return false, nil
}