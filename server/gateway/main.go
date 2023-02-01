package main

import (
	"truffle/gateway/server"
)

// @BasePath /
// @title 这是Truffle微服务的网关
// @version 1.0
// @description Truffle是一个综合的平台

func main() {
	server.NewGatewayEngine().Setup().Mount().Run()
}
