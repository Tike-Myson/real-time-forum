package sqlite3

import (
	"database/sql"
	"github.com/Tike-Myson/real-time-forum/pkg/models"
)

type MessagesModel struct {
	DB *sql.DB
}

func (m *MessagesModel) CreateMessagesTable() error {
	data, err := m.DB.Prepare(CreateMessagesSQL)
	if err != nil {
		return err
	}
	_, err = data.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (m *MessagesModel) InsertMessageIntoDB(msg models.Message) error {
	insert, err := m.DB.Prepare(InsertMessageSQL)
	if err != nil {
		return err
	}
	_, err = insert.Exec(
		msg.Id,
		msg.DialogId,

		)
	if err != nil {
		return err
	}
	return nil
}


