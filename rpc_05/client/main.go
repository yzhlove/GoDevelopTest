package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("what`s your name ?")
	clientName, _ := inputReader.ReadString('\n')

	clientName = strings.Trim(clientName, "\n")
	fmt.Println("client name  => ", clientName)

	for {
		fmt.Println("inout command :")
		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\n")
		if strings.ToUpper(trimInput) == "Q" {
			return
		}
		_, err = conn.Write([]byte(clientName + ":" + trimInput))
		data := make([]byte, 1024)
		length, _ := conn.Read(data)
		fmt.Println("Server => ", string(data[0:length]))
	}
}
