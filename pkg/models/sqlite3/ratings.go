package sqlite3

import (
	"database/sql"
	"fmt"
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

func (m *RatingModel) InsertPostRating(userId, postId string, value int) error {
	var currentRatingValue int
	found, currentRatingValue, err := m.IsRatingExists(userId, postId, "post")
	if err != nil {
		return err
	}
	if !found {
		stmt, err := m.DB.Prepare(InsertRatingPostSQL)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(postId, userId, value)
		if err != nil {
			return err
		}
		return nil
	}
	fmt.Println("GGG 1")
	currentRatingValue += value

	err = m.UpdatePostRating(userId, postId, currentRatingValue)
	if err != nil {
		return err
	}
	return nil
}

func (m *RatingModel) InsertCommentRating(userId, commentId string, value int) error {
	var currentRatingValue int
	found, currentRatingValue, err := m.IsRatingExists(userId, commentId, "comment")
	if err != nil {
		return err
	}
	if !found {
		stmt, err := m.DB.Prepare(InsertRatingCommentSQL)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(commentId, userId, value)
		if err != nil {
			return err
		}
		return nil
	}

	currentRatingValue += value

	err = m.UpdateCommentRating(userId, commentId, currentRatingValue)
	if err != nil {
		return err
	}
	return nil
}

func (m *RatingModel) UpdateCommentRating(userId, commentId string, value int) error {
	_, err := m.DB.Exec(UpdateRatingCommentSQL, value, userId, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (m *RatingModel) UpdatePostRating(userId, postId  string, value int) error {
	_, err := m.DB.Exec(UpdateRatingPostSQL, value, userId, postId)
	if err != nil {
		return err
	}
	return nil
}

func (m *RatingModel) IsRatingExists(userId, id, flag string) (bool, int, error) {
	var value int
	switch flag {
	case "post":
		fmt.Println("GGG")
		rows, err := m.DB.Query(SelectPostRatingByID, userId, id)
		if err != nil {
			return false, 0, err
		}
		if rows.Next() {
			rows.Scan(&value)
			return true, value, err
		}
	case "comment":
		rows, err := m.DB.Query(SelectCommentRatingByID, userId, id)
		if err != nil {
			return false, 0, err
		}
		if rows.Next() {
			rows.Scan(&value)
			return true, value, err
		}
	default:
		return false, 0, nil
	}
	return false, 0, nil
}