package main

import (
	"fmt"

	"github.com/vasiliy-t/gotrain/iface/mocking/redis_client"
	"gopkg.in/redis.v4"
)

func main() {
	opts := &redis.Options{
		Addr: "redis:6379",
	}
	r := redis.NewClient(opts)
	res, err := SetValue(r, "qwerty", "qwerty")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Println(res)

	res, err = GetValue(r, "qwerty")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Println(res)
}

func GetValue(r redis_client.RedisClient, key string) (string, error) {
	res := r.Get(key)
	return res.Result()
}

func SetValue(r redis_client.RedisClient, key string, value interface{}) (string, error) {
	res := r.Set(key, value, 0)
	return res.Result()
}
