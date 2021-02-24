package main

import (
	"project/wireSet"

	"github.com/google/wire"
)

func Initialize() application {
	wire.Build(newApplication, wireSet.ConfigSet, wireSet.DatabaseSet)
	return application{}
}
