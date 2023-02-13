package main

type Cache interface {
	setValue(key string, value any)
	getValue(key string) string
	clearValue(key string)
	clearAll()
}

type MultiCache struct {
	localCache LocalCache
	redisCache RedisCache
}

func (c *MultiCache) setValue(key string, value any) {
	c.localCache.setValue(key, value)
	c.redisCache.setValue(key, value)
}
func (c *MultiCache) getValue(key string) string {
	val := c.localCache.getValue(key)
	if val == "" {
		val := c.redisCache.getValue(key)
		if val == "" {
			return ""
		} else {
			return val
		}
	} else {
		return val

	}
}
func (c *MultiCache) clearValue(key string) {
	c.localCache.clearValue(key)
	c.redisCache.clearValue(key)

}
func (c *MultiCache) clearAll() {
	c.redisCache.clearAll()
	c.localCache.clearAll()
}
