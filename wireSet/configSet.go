package wireSet

import (
	"github.com/google/wire"

	"project/config"
)

var ConfigSet = wire.NewSet(
	config.InitConfig,
)
