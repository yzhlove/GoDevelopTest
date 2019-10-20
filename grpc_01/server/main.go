package main

import (
	"WorkSpace/GoDevelopTest/grpc_01/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type UserService struct{}

func (user *UserService) UserLogin(ctx context.Context, userReq *proto.UserReq) (status *proto.UserResp, err error) {

	fmt.Printf("UserLogin username=%s passwd=%s \n", userReq.Name, userReq.Passwd)
	status = new(proto.UserResp)
	if userReq.Name == "yzh" && userReq.Passwd == "123456" {
		status.ErrorCode = 0
		status.Message = "OK"
		status.Info = &proto.UserInfo{
			ID:    1,
			Name:  userReq.Name,
			Token: time.Now().String(),
			Age:   16,
		}
	} else {
		status.ErrorCode = 1
		status.Message = "FAIL"
	}
	return
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("grpc server init err")
		return
	}
	server := grpc.NewServer()
	proto.RegisterUserAccountServer(server, &UserService{})
	if err = server.Serve(listen); err != nil {
		log.Fatal("monitor is err")
		return
	}
}
