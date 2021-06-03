package sqlite3

import "database/sql"

type CategoryModel struct {
	DB *sql.DB
}

func (m *CategoryModel) CreateCategoriesTable() error {
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
