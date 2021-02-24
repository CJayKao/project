package main

import (
	"project/config"
	"project/database"
)

type application struct {
	DB *database.Database
}

func newApplication(config config.Config) *application {
	return &application{}
}

func (app *application) Start() string {
	return app.DB.DB
}

func main() {
	// app := newApplication()
	// app.Start()
}
