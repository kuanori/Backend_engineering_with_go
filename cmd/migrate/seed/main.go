package main

import (
	"app/internal/db"
	"app/internal/env"
	"app/internal/repository"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := repository.NewRepository(conn)

	db.Seed(store, conn)
}
