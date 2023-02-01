package test

import (
	"truffle/utils"
	"fmt"
	"testing"
	"time"
)

func TestCrypto(t *testing.T) {
	str := "Hello world"
	// fmt.Println("MD5: ", utils.MD5(str))
	// fmt.Println("SHA256: ", utils.SHA256(str))
	cnt := 0
	start := time.Now().Unix()
	for cnt < 10000 {
		// fmt.Println("EncodeAES: ", utils.EncodeAESWithKey(utils.BackedKey, str))
		// fmt.Println("DecodeAES: ", utils.DecodeAESWithKey(utils.BackedKey, utils.EncodeAESWithKey(utils.BackedKey, str)))
		utils.EncodeAESWithKey(utils.BackedKey, str)
		utils.DecodeAESWithKey(utils.BackedKey, utils.EncodeAESWithKey(utils.BackedKey, str))
		cnt++
	}
	end := time.Now().Unix()
	fmt.Println(end - start)
}
