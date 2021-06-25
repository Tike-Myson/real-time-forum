package sqlite3

import (
	"database/sql"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
)

type DialogsModel struct {
	DB *sql.DB
}

func (m *DialogsModel) CreateDialogsTable() error {
	data, err := m.DB.Prepare(CreateDialogsSQL)
	if err != nil {
		return err
	}
	_, err = data.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (m *DialogsModel) InsertDialogIntoDB(dialog models.Dialog) error {
	insert, err := m.DB.Prepare(InsertDialogSQL)
	if err != nil {
		return err
	}
	_, err = insert.Exec(
		dialog.Id,
		dialog.SenderId,
		dialog.ReceiverId,
		dialog.CreatedAt,
		)
	if err != nil {
		return err
	}
	return nil
}