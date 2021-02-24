package main

import (
	"project/wireSet"

	"github.com/google/wire"
)

func Initialize(path string) (Application, error) {
	wire.Build(
		newApplication,
		wireSet.ServiceSet,
		wireSet.ConfigSet,
		wireSet.DatabaseSet,
	)
	return Application{}, nil
}
