package sqlite3

import (
	"database/sql"
	"errors"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) GetUserIdByLogin(username string) (int, error) {
	var id int
	stmt := "SELECT id, password FROM users WHERE username = ?"
	row := m.DB.QueryRow(stmt, username)
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoRecord
		} else {
			return 0, err
		}
	}
	return id, nil
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

func (m *UserModel) Authenticate(email string, password []byte) (int,error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, password FROM users WHERE email = ?"
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	return id, nil
}
