package sqlite3

import "database/sql"

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