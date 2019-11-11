package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Receive Data => ", string(buf[0:length]))
		_, _ = conn.Write([]byte(strings.ToUpper(string(buf[0:length]))))
	}
}
