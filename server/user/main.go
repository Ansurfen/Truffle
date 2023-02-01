package main

import (
	"truffle/user/engine"
)

func main() {
	engine.NewUserEngine().Setup().Mount().Run()
}
