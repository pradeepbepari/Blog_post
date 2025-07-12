package main

import (
	"blog-apis/database"
	"blog-apis/handler"
	"blog-apis/models"
	"blog-apis/repository"
	"blog-apis/routes"
	"fmt"
	"log"

	_ "blog-apis/docs" // Import the generated swagger package

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/kelseyhightower/envconfig"
)

// @title Blog Posts API
// @version 1.0
// @description This is a sample swagger for Blog Posts API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	var config models.Config
	// Load environment variables into the config struct
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Error processing environment variables: %v", err)
	}
	// Initialize the database connection
	db, err := database.DBConnection(config)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	if err := database.DatabaseSchema(db); err != nil {
		log.Fatalf("Error initializing database schema: %v", err)
	}
	repository := repository.NewRepository(db)
	handler := handler.NewHandler(repository)
	// Initialize the Fiber app
	app := fiber.New()
	routes.ApiRoutes(app, handler)
	app.Get("/swagger/*", swagger.HandlerDefault)
	defer database.CloseDB(db)
	// Start the Fiber app
	if err := app.Listen(fmt.Sprintf(":%s", config.BlogPort)); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
