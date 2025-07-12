package repository

import (
	"blog-apis/models"

	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	CreatePost(c *fiber.Ctx, post *models.BlogPost) error
	GetBlogPosts(c *fiber.Ctx) ([]models.BlogPost, error)
	GetPost(c *fiber.Ctx, id string) (models.BlogPost, error)
	UpdatePost(c *fiber.Ctx, id string, post *models.BlogPost) error
	DeletePost(c *fiber.Ctx, id string) error
}
