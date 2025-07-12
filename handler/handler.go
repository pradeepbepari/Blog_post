package handler

import (
	"blog-apis/models"
	"blog-apis/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

// CreatePost creates a new blog post
// @Summary Create a new blog post
// @Description Creates a new blog post with the provided details
// @Tags Public
// @Accept json
// @Produce json
// @Param user body models.BlogPost true "Blog post details"
// @Success 200
// @Failure 400 "{\"error\": \"Bad request\"}"
// @Failure 500 "{\"error\": \"Internal server error\"}"
// @Router /posts [post]
func (h *Handler) CreatePost(c *fiber.Ctx) error {
	var reqBody models.BlogPost
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.repo.CreatePost(c, &reqBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "post created successfully", "post": reqBody})
}

// GetPost retrieves a blog post by ID
// @Summary Get a blog post by ID
// @Description Retrieves a blog post by its ID
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} models.BlogPost
// @Failure 400 "{\"error\": \"Post ID is required\"}"
// @Failure 500 "{\"error\": \"Internal server error\"}"
// @Router /posts/{id} [get]
func (h *Handler) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" || len(id) != 36 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Post ID is required"})
	}
	post, err := h.repo.GetPost(c, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"post": post})
}

// GetPosts retrieves all blog posts
// @Summary Get all blog posts
// @Description Retrieves all blog posts
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {array} models.BlogPost
// @Failure 500 "{\"error\": \"Internal server error\"}"
// @Router /posts [get]
func (h *Handler) GetPosts(c *fiber.Ctx) error {
	posts, err := h.repo.GetBlogPosts(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"posts": posts})
}

// UpdatePost updates a blog post by ID
// @Summary Update a blog post by ID
// @Description Updates a blog post with the provided details
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param user body models.BlogPost true "Blog post details"
// @Success 200 "{\"message\": \"post updated successfully\", \"postID\": \"<id>\"}"
// @Failure 400 "{\"error\": \"Post ID is required\"}"
// @Failure 500 "{\"error\": \"Internal server error\"}"
// @Router /posts/{id} [put]
func (h *Handler) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" || len(id) != 36 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Post ID is required"})
	}
	var reqBody models.BlogPost
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.repo.UpdatePost(c, id, &reqBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update post"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "post updated successfully", "postID": id})
}

// DeletePost deletes a blog post by ID
// @Summary Delete a blog post by ID
// @Description Deletes a blog post by its ID
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 "{\"message\": \"post deleted successfully\", \"postID\": \"<id>\"}"
// @Failure 400 "{\"error\": \"Post ID is required\"}"
// @Failure 500 "{\"error\": \"Internal server error\"}"
// @Router /posts/{id} [delete]
func (h *Handler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" || len(id) != 36 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Post ID is required"})
	}
	if err := h.repo.DeletePost(c, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete post"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "post deleted successfully", "postID": id})
}
