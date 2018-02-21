package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// RedisCli is a struct to control redis
type RedisCli struct {
	Client *redis.Client
}

func ExampleNewClient() RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := c.Ping().Result()
	fmt.Println(pong, err)

	return RedisCli{c}
}

func (r *RedisCli) Hash(key string) string {
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return strconv.Itoa(h)
}

// Insert inserts data into redis
func (r *RedisCli) Insert(jk string) error {
	k := r.Hash(jk)
	err := r.Client.Set(k, jk, time.Hour*24).Err()

	if err != nil {
		return err
	}

	fmt.Printf("Insertion into redis, KEY: %s", k)
	return nil
}
