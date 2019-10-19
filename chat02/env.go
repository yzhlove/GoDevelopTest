package main

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	S Storage
}

func getEnv() *Env {

	var (
		addr, passwd, db string
	)

	if addr = os.Getenv("APP_REDIS_ADDR"); addr == "" {
		addr = "localhost:6379"
	}

	if passwd = os.Getenv("APP_REDIS_PASSWD"); passwd == "" {
		passwd = ""
	}

	if db = os.Getenv("APP_REDIS_DB"); db == "" {
		db = "0"
	}

	index, err := strconv.Atoi(db)
	if err != nil {
		log.Fatal("connect redis index transform err")
	}

	client := NewRedisCli(addr, passwd, index)
	return &Env{S: client}
}
