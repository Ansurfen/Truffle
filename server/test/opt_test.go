package test

import (
	"fmt"
	"testing"
	"time"
	"truffle/client"
	"truffle/middleware"
	. "truffle/utils"
)

func TestOption(t *testing.T) {
	start := time.Now().Unix()
	opt := LoadOpt(ENV_DEVELOP, EnvOpt{}, DefaultOpt{}, EmailOpt{}, middleware.BreakerOpt{}, client.TracerOpt{})
	fmt.Println(opt.Opt(DEFAULT).(DefaultOpt).Service.Addr)
	end := time.Now().Unix()
	fmt.Println(end - start)
}
