package wireSet

import (
	"project/config"
	"project/database"

	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(NewDatabaseSet)

func NewDatabaseSet(config *config.Config) *database.Database {
	return database.NewDatabase(config)
}
