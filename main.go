package main

import "project/service"

type Application struct {
	auth *service.AuthSerivce
}

func newApplication(
	auth *service.AuthSerivce,
) *Application {
	return &Application{
		auth: auth,
	}
}

func (app *Application) Start() {
	app.Launch()
}

func (app *Application) Launch() {
	app.auth.GetAuth()
}

func main() {
	// app, _ := Initialize("123")
	// app.Start()
}
