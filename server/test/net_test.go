package test

import (
	"fmt"
	"net"
	"testing"
)

func TestQueryIp(t *testing.T) {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr().String())
}
