package test

import (
	"fmt"
	"testing"
	"time"
	"truffle/breaker"
	"truffle/client"
	. "truffle/utils"
)

func TestOption(t *testing.T) {
	start := time.Now().Unix()
	opt := LoadOpt(ENV_DEVELOP, EnvOpt{}, DefaultOpt{}, EmailOpt{}, breaker.BreakerOpt{}, client.TracerOpt{})
	fmt.Println(opt.Opt(DEFAULT).(DefaultOpt).Service.Addr)
	end := time.Now().Unix()
	fmt.Println(end - start)
}
