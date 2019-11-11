package client

import (
	"WorkSpace/GoDevelopTest/rpc_02/config"
	"WorkSpace/GoDevelopTest/rpc_03/obj"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", config.Address)
	if err != nil {
		panic(err)
	}

	rpcMsg := &obj.RpcMsg{
		Code:     1001,
		FuncName: "UserSignReq",
	}

}
