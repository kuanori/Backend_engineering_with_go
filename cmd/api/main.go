package main

import (
	"app/internal/db"
	"app/internal/env"
	"app/internal/repository"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15mins"),
		},
	} // это как создание обьекта класса в php

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection is established")

	repository := repository.NewRepository(db)

	app := &application{
		config:     cfg,
		repository: repository,
	}

	r := app.mount() // Вызывается метод mount у структуры application
	log.Fatal(app.run(r))
}
