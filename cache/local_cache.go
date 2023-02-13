package main

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	c = cache.New(5*time.Minute, 10*time.Minute)
)

type LocalCache struct {
}

func (LocalCache) setValue(key string, value any) {

}
func (LocalCache) getValue(key string) string {
	return ""
}
func (LocalCache) clearValue(key string) {

}
func (LocalCache) clearAll() {

}
