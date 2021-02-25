package main

import (
	"project/wireSet"

	"github.com/google/wire"
)

func Initialize(path string) (Application, error) {
	wire.Build(
		newApplication,
		wireSet.ConfigSet,
		wireSet.DatabaseSet,
		// wireSet.ServiceSet,
	)
	return Application{}, nil
}
