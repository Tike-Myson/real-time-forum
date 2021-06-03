package sqlite3

import "database/sql"

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

