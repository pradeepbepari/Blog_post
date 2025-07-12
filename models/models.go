package models

import (
	"time"

	"github.com/google/uuid"
)

type Config struct {
	Database   string `envconfig:"DATABASE"`
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	BlogPort   string `envconfig:"BLOG_PORT"`
}
type BlogPost struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
