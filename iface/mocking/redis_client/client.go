package redis_client

import (
	"time"

	"gopkg.in/redis.v4"
)

// redis.v4 does not provide client interface, it provides struct
// some day we may want to mock redis client
//
// redis library is used only for presentation purpose
// it could be any package and any task, my was to mock redis client
//
// If we want to mock it for test we could declare an interface, that
// mimics the redis client methods, and use this interface in our code
// then we can just generate mock with gomock or write it by hand
// and use instead of redis.Client struct because it implicitly
// implements our new interface
type RedisClient interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
}
