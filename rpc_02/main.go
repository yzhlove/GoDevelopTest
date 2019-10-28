package main

import (
	"WorkSpace/GoDevelopTest/rpc_02/client"
	"WorkSpace/GoDevelopTest/rpc_02/server"
	"fmt"
	"net"
)

func main() {
	const address = "127.0.0.1:8008"
	const testCounter = 100
	var counter int
	accept := server.NewAccept()
	accept.Start(address)
	accept.Call = func(conn net.Conn, bytes []byte) bool {
		str := string(bytes)
		fmt.Println("str => ", str)
		if counter >= testCounter {
			accept.Stop()
			return false
		}
		return true
	}

	client.Send(address)
	accept.Wait()

}
