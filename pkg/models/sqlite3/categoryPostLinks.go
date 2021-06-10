package sqlite3

import "database/sql"

type CategoryPostLinkModel struct {
	DB *sql.DB
}

func (m *CategoryPostLinkModel) CreateCategoryPostLinksTable() error {
	categoryPostLinkTable, err := m.DB.Prepare(CreateCategoryPostLinkSQL)
	if err != nil {
		return err
	}
	_, err = categoryPostLinkTable.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (m *CategoryPostLinkModel) InsertCategoryPostLinkIntoDB(postId, categoryName string) error {
	stmt, err := m.DB.Prepare(InsertCategoryPostLinkSQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(postId, categoryName)
	if err != nil {
		return err
	}
	return nil
}