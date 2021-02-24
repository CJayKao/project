package database

import (
	"project/config"
)

type Database struct {
	DB string
}

func NewDatabase(config *config.Config) *Database {
	return &Database{
		DB: config.Port,
	}
}
