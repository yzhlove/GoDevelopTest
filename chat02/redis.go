package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mattheath/base62"
	"time"
)

const (
	URLIDKEY           = "next.url.id"
	SHORTLINKKEY       = "short_link:%s:url"
	URLHASHKEY         = "url_hash:%s:url"
	SHORTLINKDETAILKEY = "short_link:%s:detail"
)

type RedisCli struct {
	Cli *redis.Client
}

type URLDetail struct {
	URL                 string        `json:"url"`
	CreatedAt           string        `json:"created_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

func NewRedisCli(add string, passwd string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr:     add,
		Password: passwd,
		DB:       db,
	})

	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}
	return &RedisCli{Cli: c}
}

func toSha1(str string) string {
	var (
		sha = sha1.New()
	)
	return string(sha.Sum([]byte(str)))
}

func (cli *RedisCli) Shorten(url string, exp int64) (detail string, err error) {
	var (
		hash       = toSha1(url)
		index      int64
		eid        string
		jsonDetail []byte
		expire     time.Duration
	)
	if detail, err = cli.Cli.Get(fmt.Sprintf(URLHASHKEY, hash)).Result(); err != nil {
		return
	}

	if index, err = cli.Cli.Incr(URLIDKEY).Result(); err != nil {
		return
	}

	eid = base62.EncodeInt64(index)
	expire = time.Minute * time.Duration(exp)

	if err = cli.Cli.Set(fmt.Sprintf(SHORTLINKKEY, eid), url, expire).Err(); err != nil {
		return
	}

	if err = cli.Cli.Set(fmt.Sprintf(URLHASHKEY, hash), eid, expire).Err(); err != nil {
		return
	}

	if jsonDetail, err = json.Marshal(&URLDetail{
		URL:                 url,
		CreatedAt:           time.Now().String(),
		ExpirationInMinutes: time.Duration(exp),
	}); err != nil {
		return
	}
	if err = cli.Cli.Set(fmt.Sprintf(SHORTLINKDETAILKEY, eid), jsonDetail, expire).Err(); err != nil {
		return
	}
	return
}

func (cli *RedisCli) ShortLinkInfo(eid string) (detail interface{}, err error) {
	detail, err = cli.Cli.Get(fmt.Sprintf(SHORTLINKDETAILKEY, eid)).Result()
	return
}

func (cli *RedisCli) UnShorten(eid string) (url string, err error) {
	url, err = cli.Cli.Get(fmt.Sprintf(SHORTLINKKEY, eid)).Result()
	return
}
