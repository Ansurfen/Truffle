package main

import (
	"truffle/ws/engine"
)

func main() {
	engine.NewWsEngine().Setup().Mount().Run()
}
