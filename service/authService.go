package service

import (
	"project/config"
	"project/database"
)

type AuthSerivce struct {
	config *config.Config
	db     *database.Database
}

func NewAuthSerivce(
	config *config.Config,
	db *database.Database,
) *AuthSerivce {
	return &AuthSerivce{
		config: config,
		db:     db,
	}
}

func (auth *AuthSerivce) GetAuth() string {
	return auth.db.DB
}
