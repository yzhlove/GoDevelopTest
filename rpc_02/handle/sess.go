package handle

import (
	"WorkSpace/GoDevelopTest/rpc_02/packet"
	"bufio"
	"fmt"
	"net"
)

func Accept(conn net.Conn, call func(net.Conn, []byte) bool) {
	reader := bufio.NewReader(conn)
	for {
		if pkg, err := packet.UnPacket(reader);
			err != nil || !call(conn, pkg.Body) {
			fmt.Println("UnPackErr:", err)
			_ = conn.Close()
			break
		}
	}
}
