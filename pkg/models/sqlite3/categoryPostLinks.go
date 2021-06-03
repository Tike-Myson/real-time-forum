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
