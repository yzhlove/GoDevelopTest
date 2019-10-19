package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	pool      *redis.Pool
	redisHost = "127.0.0.1:6379"
)

func init() {
	pool = newRedisPool()
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			if conn, err = redis.Dial("tcp", redisHost); err != nil {
				fmt.Println("redis connect err: ", err)
				return
			}
			return
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			if _, err := c.Do("PING"); err != nil {
				fmt.Println("redis ping err: ", err)
				return err
			}
			return nil
		},
	}
}

func NewRedisClient() *redis.Pool {
	return pool
}
