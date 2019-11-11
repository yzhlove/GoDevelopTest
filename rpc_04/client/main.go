package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	var (
		host   = "www.baidu.com"
		port   = "80"
		remote = host + ":" + port
		data   = make([]uint8, 4096)
		msg    = "GET / HTTP/1.1\n"
		count  = 0
		status = true
	)

	conn, err := net.Dial("tcp", remote)
	if err != nil {
		panic(err)
	}
	_, _ = io.WriteString(conn, msg)
	for status {
		count, err = conn.Read(data)
		status = err == nil
		fmt.Printf(string(data[0:count]))
	}
	_ = conn.Close()
}
