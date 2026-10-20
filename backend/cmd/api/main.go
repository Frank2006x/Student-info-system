package main

import (
	"context"
	"log"

	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/config"
	"Frank2006x/Vmate/internals/router"
	"Frank2006x/Vmate/internals/store"
)

func main() {
	settings, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	pools, err := store.ConnectPools(context.Background(), settings.PrimaryDatabaseURL, settings.ReplicaDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pools.Primary.Close()
	defer pools.Replica.Close()

	app := router.New(db.New(pools.Replica), db.New(pools.Primary))
	log.Printf("student information API listening on %s", ":8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
