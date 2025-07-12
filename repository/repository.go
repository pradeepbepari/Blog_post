package repository

import (
	"blog-apis/models"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repo {
	return &repo{db: db}
}
func (r *repo) GetBlogPosts(c *fiber.Ctx) ([]models.BlogPost, error) {
	query := `SELECT id, title, body, description, created_at, updated_at FROM blog_posts`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.BlogPost
	for rows.Next() {
		var post models.BlogPost
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Description, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *repo) CreatePost(c *fiber.Ctx, post *models.BlogPost) error {
	post.ID = uuid.New()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	query := `INSERT INTO blog_posts (id, title, body, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, post.ID, post.Title, post.Body, post.Description, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (r *repo) GetPost(c *fiber.Ctx, id string) (models.BlogPost, error) {
	query := `SELECT id, title, body, description, created_at, updated_at FROM blog_posts WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var post models.BlogPost
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.Description, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return models.BlogPost{}, err
	}
	return post, nil
}
func (r *repo) UpdatePost(c *fiber.Ctx, id string, post *models.BlogPost) error {
	post.UpdatedAt = time.Now()

	query := `UPDATE blog_posts SET title = $1, body = $2, description = $3, updated_at = $4 WHERE id = $5`
	_, err := r.db.Exec(query, post.Title, post.Body, post.Description, post.UpdatedAt, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repo) DeletePost(c *fiber.Ctx, id string) error {
	query := `DELETE FROM blog_posts WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
