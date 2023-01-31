package main

import (
	"truffle/log/engine"
)

func main() {
	engine.NewLogEngine().Setup().Mount().Run()
}
