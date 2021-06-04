package sqlite3

import (
	"database/sql"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
)

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

func (m *UserModel) CreateUser(user models.User) error {
	insertStmt, err := m.DB.Prepare(InsertUserSQL)
	if err != nil {
		return err
	}
	_, err = insertStmt.Exec(
		user.Nickname,
		user.Age,
		user.Gender,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		)
	if err != nil {
		return err
	}
	return nil
}
