package sqlite3

import "database/sql"

type CategoryModel struct {
	DB *sql.DB
}

func (m *CategoryModel) CreateCategoriesTable() error {
	stmt, err := m.DB.Prepare(CreatePostsTableSQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (m *CategoryModel) InsertCategoryIntoDB(categoryName string) error {
	stmt, err := m.DB.Prepare(InsertCategoriesSQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(categoryName)
	if err != nil {
		return err
	}
	return nil
}
