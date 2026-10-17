package main

import (
	"context"
	"log"
	"os"

	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/router"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://app:local_dev_password@localhost:5432/students_db?sslmode=disable"
	}
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	app := router.New(db.New(pool))
	log.Printf("student information API listening on %s", ":8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
