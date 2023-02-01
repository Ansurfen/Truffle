package test

import (
	"fmt"
	"log"
	"testing"
	"truffle/utils"
)

func TestAuthJWT(t *testing.T) {
	username := "Ansurfen"
	now := utils.NowTimestamp()
	hash := utils.EncodeAESWithKey(utils.ToString(now)+"TRUFFLE", username)
	token, err := utils.ReleaseToken(hash, now)
	if err != nil {
		log.Fatalln(err)
		return
	}
	front_token := "TRUFFLE" + token
	fmt.Printf("Token: %s \n", front_token)
	if utils.AuthJWT(front_token, username) {
		fmt.Println("auth success")
	} else {
		fmt.Println("auth fail")
	}
}
