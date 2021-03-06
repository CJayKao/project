// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"project/config"
	"project/wireSet"
)

// Injectors from wire.go:

func Initialize(path string) (*Application, error) {
	configConfig := config.InitConfig(path)
	database := wireSet.NewDatabaseSet(configConfig)
	application := newApplication(configConfig, database)
	return application, nil
}
