package main

import "truffle/ws_consumer/engine"

func main() {
	engine.NewWsConsumerEngine().Setup().Mount().Run()
}
