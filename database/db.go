package database

import (
	"blog-apis/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DBConnection(config models.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	if err = DBHealth(db); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing the database connection: %v", err)
	}
}
func DBHealth(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		log.Printf("Database health check failed: %v", err)
		return err
	}
	log.Println("Database is healthy.")
	return nil
}
