package sqlite3

import "database/sql"

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) CreateUsersTable() error {
	usersTable, err := m.DB.Prepare(CreateUsersTableSQL)
	if err != nil {
		return err
	}
	_, err = usersTable.Exec()
	if err != nil {
		return err
	}
	return nil
}
