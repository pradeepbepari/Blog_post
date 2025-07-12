package database

import "database/sql"

const (
	blogSchema = `
	DROP TABLE IF EXISTS blog_posts;
	CREATE TABLE IF NOT EXISTS blog_posts (
		id UUID PRIMARY KEY,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
)

func DatabaseSchema(db *sql.DB) error {
	_, err := db.Exec(blogSchema)
	if err != nil {
		return err
	}
	return nil
}
