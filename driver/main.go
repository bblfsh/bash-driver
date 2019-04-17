package main

import (
	_ "github.com/bblfsh/bash-driver/driver/impl"
	"github.com/bblfsh/bash-driver/driver/normalizer"

	"github.com/bblfsh/sdk/v3/driver/server"
)

func main() {
	server.Run(normalizer.Transforms)
}
