package database

import (
	"project/config"
)

type Database struct {
	DB string
}

func NewDatabase(config config.Config) *Database {
	// fake db
	// initalize db
	return &Database{
		DB: config.Port,
	}
}
