package routes

import (
	"blog-apis/handler"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(r *fiber.App, h *handler.Handler) {
	router := r.Group("/api/v1")
	router.Post("/posts", h.CreatePost)
	router.Get("/posts", h.GetPosts)
	router.Get("/posts/:id", h.GetPost)
	router.Put("/posts/:id", h.UpdatePost)
	router.Delete("/posts/:id", h.DeletePost)
}
