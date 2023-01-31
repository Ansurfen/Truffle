package main

import (
	"truffle/captcha/engine"
)

func main() {
	engine.NewCaptchaEngine().Setup().Mount().Run()
}
