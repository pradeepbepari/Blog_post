package handler

import (
	"blog-apis/models"
	"blog-apis/repository/mocks"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Post("/posts", h.CreatePost)

	post := &models.BlogPost{
		Title:       "test title",
		Body:        "test content",
		Description: "test description",
	}

	body, _ := json.Marshal(post)

	mockRepo.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	req := httptest.NewRequest(http.MethodPost, "/posts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHandler_GetPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Get("/posts/:id", h.GetPost)

	postID := "f277987f-4bb0-4578-9c1b-3810c5f6df67"
	post := &models.BlogPost{
		// ID:          postID,
		Title:       "test title",
		Body:        "test content",
		Description: "test description",
	}
	mockRepo.EXPECT().GetPost(gomock.Any(), postID).Return(*post, nil).Times(1)

	req := httptest.NewRequest(http.MethodGet, "/posts/"+postID, nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHandler_CreatePost_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Post("/posts", h.CreatePost)

	req := httptest.NewRequest(http.MethodPost, "/posts", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestHandler_GetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Get("/posts", h.GetPosts)

	posts := []models.BlogPost{
		{
			ID:          uuid.New(),
			Title:       "test title 1",
			Body:        "test content 1",
			Description: "test description 1",
		},
		{
			ID:          uuid.New(),
			Title:       "test title 2",
			Body:        "test content 2",
			Description: "test description 2",
		},
	}

	mockRepo.EXPECT().GetBlogPosts(gomock.Any()).Return(posts, nil).Times(1)

	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHandler_UpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Put("/posts/:id", h.UpdatePost)

	postID := "f277987f-4bb0-4578-9c1b-3810c5f6df67"
	post := &models.BlogPost{
		Title:       "updated title",
		Body:        "updated content",
		Description: "updated description",
	}

	body, _ := json.Marshal(post)

	mockRepo.EXPECT().UpdatePost(gomock.Any(), postID, gomock.Any()).Return(nil).Times(1)

	req := httptest.NewRequest(http.MethodPut, "/posts/"+postID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHandler_UpdatePost_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Put("/posts/:id", h.UpdatePost)

	req := httptest.NewRequest(http.MethodPut, "/posts/bad-id", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestHandler_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Delete("/posts/:id", h.DeletePost)

	postID := "f277987f-4bb0-4578-9c1b-3810c5f6df67"

	mockRepo.EXPECT().DeletePost(gomock.Any(), postID).Return(nil).Times(1)

	req := httptest.NewRequest(http.MethodDelete, "/posts/"+postID, nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHandler_DeletePost_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Delete("/posts/:id", h.DeletePost)

	req := httptest.NewRequest(http.MethodDelete, "/posts/bad-id", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
func TestHandler_CreatePost_InTernalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Post("/posts", h.CreatePost)

	post := &models.BlogPost{
		Title:       "test title",
		Body:        "test content",
		Description: "test description",
	}

	body, _ := json.Marshal(post)

	mockRepo.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(fiber.ErrInternalServerError).Times(1)

	req := httptest.NewRequest(http.MethodPost, "/posts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}
func TestHandler_GetPosts_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Get("/posts", h.GetPosts)

	mockRepo.EXPECT().GetBlogPosts(gomock.Any()).Return(nil, fiber.ErrInternalServerError).Times(1)

	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}
func TestHandler_GetPost_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Get("/posts/:id", h.GetPost)

	postID := "f277987f-4bb0-4578-9c1b-3810c5f6df67"

	mockRepo.EXPECT().GetPost(gomock.Any(), postID).Return(models.BlogPost{}, fiber.ErrInternalServerError).Times(1)

	req := httptest.NewRequest(http.MethodGet, "/posts/"+postID, nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}
func TestHandler_UpdatePost_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	h := NewHandler(mockRepo)

	app := fiber.New()
	app.Put("/posts/:id", h.UpdatePost)

	postID := "f277987f-4bb0-4578-9c1b-3810c5f6df67"
	post := &models.BlogPost{
		Title:       "updated title",
		Body:        "updated content",
		Description: "updated description",
	}

	body, _ := json.Marshal(post)

	mockRepo.EXPECT().UpdatePost(gomock.Any(), postID, gomock.Any()).Return(fiber.ErrInternalServerError).Times(1)

	req := httptest.NewRequest(http.MethodPut, "/posts/"+postID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}
