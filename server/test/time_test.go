package test

import (
	"fmt"
	"testing"
	"time"
	"truffle/utils"
)

func TestTimestamp(t *testing.T) {
	timer := utils.NewTimer(5, 100, 4, 0.5)
	timer.Add("a")
	ticker1 := time.NewTicker(5 * time.Second)
	for range ticker1.C {
		fmt.Println("加入 b")
		timer.Add("b")
		break
	}
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		fmt.Println("定时")
		timer.Update()
	}
}
