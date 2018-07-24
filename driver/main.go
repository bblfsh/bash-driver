package main

import (
	_ "github.com/bblfsh/bash-driver/driver/impl"
	"github.com/bblfsh/bash-driver/driver/normalizer"

	"gopkg.in/bblfsh/sdk.v2/sdk/driver"
)

func main() {
	driver.Run(normalizer.Transforms)
}
