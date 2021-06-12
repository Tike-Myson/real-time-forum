package sqlite3

/*------------------------------------------------------*/
/*                                                      */
/*                    POST STATEMENTS                   */
/*                                                      */
/*------------------------------------------------------*/

const CreatePostsTableSQL = `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		created_at TIMESTAMP NOT NULL,
		image_url TEXT
	);
`
const InsertPostSQL = `
	INSERT INTO posts (
		title, content, user_id, created_at, image_url
	) VALUES (?, ?, ?, ?, ?);
`

const GetAllPostsSQL = `
	SELECT * FROM posts
`

/*------------------------------------------------------*/
/*                                                      */
/*                    USER STATEMENTS                   */
/*                                                      */
/*------------------------------------------------------*/

const CreateUsersTableSQL = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		nickname TEXT NOT NULL UNIQUE,
		age INTEGER NOT NULL,
		gender TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
`
const InsertUserSQL = `
	INSERT INTO users (
		nickname, age, gender, first_name, last_name, email, password
	) VALUES (?, ?, ?, ?, ?, ?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                   COMMENT STATEMENTS                 */
/*                                                      */
/*------------------------------------------------------*/

const CreateCommentsTableSQL = `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL
	);
`

const InsertCommentSQL = `
	INSERT INTO comments (
		post_id, user_id, content, created_at
	) VALUES (?, ?, ?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                   CATEGORY STATEMENTS                */
/*                                                      */
/*------------------------------------------------------*/

const CreateCategoryTableSQL = `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
`

const CreateCategoryPostLinkSQL = `
	CREATE TABLE IF NOT EXISTS categoryPostLink (
		post_id INTEGER NOT NULL,
		category_name STRING NOT NULL
	);
`

const InsertCategoriesSQL = `
	INSERT INTO categories (
		name
	) VALUES (?);
`

const InsertCategoryPostLinkSQL = `
	INSERT INTO categoryPostLink (
		post_id, name
	) VALUES (?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                    RATING STATEMENTS                 */
/*                                                      */
/*------------------------------------------------------*/

const CreateRatingPostSQL = `
	CREATE TABLE IF NOT EXISTS ratingPosts (
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		value INTEGER NOT NULL
	);
`

const CreateRatingCommentSQL = `
	CREATE TABLE IF NOT EXISTS ratingComments (
		comment_id INTEGER NOT NULL,
		user_Id INTEGER NOT NULL,
		value INTEGER NOT NULL
	);
`

const InsertRatingPostSQL = `
	INSERT INTO ratingPosts (
		post_id, user_id, value
	) VALUES (?, ?, ?);
`

const InsertRatingCommentSQL = `
	INSERT INTO ratingComments (
		comment_id, user_id, value
	) VALUES (?, ?, ?);
`

const UpdateRatingPostSQL = `
	UPDATE ratingPosts SET value = ?
	WHERE user_id = ? AND post_id = ?;
`

const UpdateRatingCommentSQL = `
	UPDATE ratingComments SET value = ?
	WHERE user_id = ? AND comment_id = ?;
`

const SelectPostRatingByID = `
	SELECT "value" FROM rating user_id = ? AND post_id = ?;
`

const SelectCommentRatingByID = `
	SELECT value FROM rating user_id = ? AND comment_id = ?;
`
