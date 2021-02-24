package wireSet

import (
	"github.com/google/wire"

	"project/service"
)

var ServiceSet = wire.NewSet(
	service.NewAuthSerivce,
)
