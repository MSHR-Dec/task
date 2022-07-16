package adapter

import (
	"github.com/gin-contrib/sessions/redis"
)

func NewRedisConnection() redis.Store {
	store, err := redis.NewStore(10, "tcp", Environment.RedisHost, "", []byte("secret"))
	if err != nil {
		panic("Fail to connect Cache.")
	}

	return store
}
