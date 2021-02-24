package wireSet

import (
	"project/config"
	"project/database"

	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(newDatabaseSet)

func newDatabaseSet(config *config.Config) *database.Database {
	return database.NewDatabase(config)
}
