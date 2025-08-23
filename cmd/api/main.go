package main

import (
	"app/internal/env"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	} // это как создание обьекта класса в php

	app := &application{
		config: cfg,
	}

	r := app.mount() // Вызывается метод mount у структуры application
	log.Fatal(app.run(r))
}
