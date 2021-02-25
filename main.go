package main

import (
	"project/config"
	"project/database"
	// "project/database"
)

type Application struct {
	// auth *service.AuthSerivce
	DB     *database.Database
	Config config.Config
}

func newApplication(
	// auth *service.AuthSerivce,
	config config.Config,
	db *database.Database,
) Application {
	return Application{
		Config: config,
		DB:     db,
	}
}

func (app *Application) Start() {

}

func main() {
	// app, _ := Initialize("123")
	// app.Start()
}
