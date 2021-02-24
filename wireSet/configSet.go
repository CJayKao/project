package wireSet

import (
	"github.com/google/wire"

	"project/config"
)

var ConfigSet = wire.NewSet(NewConfig)

func NewConfig() *config.Config {
	return config.InitConfig()
}
