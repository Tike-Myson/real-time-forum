package main

import "database/sql"

func (app *application) createAllTables() error {
	err := app.posts.CreatePostsTable()
	if err != nil {
		return err
	}
	err = app.categories.CreateCategoriesTable()
	if err != nil {
		return err
	}
	err = app.categoryPostLinks.CreateCategoryPostLinksTable()
	if err != nil {
		return err
	}
	err = app.comments.CreateCommentsTable()
	if err != nil {
		return err
	}
	err = app.ratings.CreateRatingsTable()
	if err != nil {
		return err
	}
	err = app.users.CreateUsersTable()
	if err != nil {
		return err
	}
	err = app.dialogs.CreateDialogsTable()
	if err != nil {
		return err
	}
	err = app.messages.CreateMessagesTable()
	if err != nil {
		return err
	}
	return nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}