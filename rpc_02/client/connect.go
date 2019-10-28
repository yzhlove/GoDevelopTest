package client

import (
	"WorkSpace/GoDevelopTest/rpc_02/packet"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func Send(address string) {

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Net Dial Err:", err)
		return
	}
	for {
		var token string
		for _, value := range rand.Perm(10) {
			token += string(rune(98 + value))
		}
		//fmt.Println("token:", token)
		if err := packet.Pack(conn, []byte(token)); err != nil {
			fmt.Println("Send Err:", err)
			return
		}
		time.Sleep(time.Second)
	}
}
