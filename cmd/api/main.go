package main

import (
	"app/internal/env"
	"app/internal/repository"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	} // это как создание обьекта класса в php

	repository := repository.NewRepository(nil)

	app := &application{
		config:     cfg,
		repository: repository,
	}

	r := app.mount() // Вызывается метод mount у структуры application
	log.Fatal(app.run(r))
}
