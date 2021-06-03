package main

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
	return nil
}
