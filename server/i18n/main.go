package main

import (
	"truffle/i18n/engine"
)

func main() {
	engine.NewI18NEngine().Setup().Mount().Run()
}
