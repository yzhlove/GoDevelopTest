package main

import (
	"WorkSpace/GoDevelopTest/file_store_service/account/handle"
	"WorkSpace/GoDevelopTest/file_store_service/account/proto"
	"github.com/micro/go-micro"
	"log"
	"time"
)

func main() {

	//创建一个服务
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	service.Init()
	if err := proto.RegisterUserServiceHandler(service.Server(), new(handle.User)); err != nil {
		log.Fatal("register user service err")
		return
	}
	if err := service.Run(); err != nil {
		log.Fatal("user service run err")
	}
}
