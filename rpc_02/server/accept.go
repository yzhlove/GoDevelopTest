package server

import (
	"WorkSpace/GoDevelopTest/rpc_02/handle"
	"fmt"
	"log"
	"net"
	"sync"
)

type Accept struct {
	listen net.Listener
	wg     sync.WaitGroup
	Call   func(net.Conn, []byte) bool
}

func NewAccept() *Accept {
	return &Accept{}
}

func (a *Accept) Wait() {
	a.wg.Wait()
}

func (a *Accept) Stop() {
	if err := a.listen.Close(); err != nil {
		log.Println("Listen Close Err:", err)
	}
}

func (a *Accept) Start(address string) {
	go a.monitor(address)
}

func (a *Accept) monitor(address string) {
	a.wg.Add(1)
	defer a.wg.Done()
	var err error
	if a.listen, err = net.Listen("tcp", address); err != nil {
		fmt.Println("NetWork Listen Err:", err)
		return
	}
	for {
		if conn, err := a.listen.Accept(); err != nil {
			fmt.Println("AcceptErr:", err)
		} else {
			go handle.Accept(conn, a.Call)
		}
	}
}
