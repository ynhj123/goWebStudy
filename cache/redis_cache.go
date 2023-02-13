package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx = context.Background()
)

type RedisCache struct {
}

func (RedisCache) setValue(key string, value any) {

}
func (RedisCache) getValue(key string) string {
	return ""
}
func (RedisCache) clearValue(key string) {

}
func (RedisCache) clearAll() {

}
