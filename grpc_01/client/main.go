package main

import (
	"WorkSpace/GoDevelopTest/grpc_01/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	client := proto.NewUserAccountClient(conn)
	resp, err := client.UserLogin(context.Background(), &proto.UserReq{
		Name:   "yzh",
		Passwd: "123456",
	})

	log.Println(resp)
}
