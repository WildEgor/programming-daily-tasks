package simple_cache

import "sync"

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

type SimpleCache struct {
	mx    sync.RWMutex
	cache map[string]string
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{cache: make(map[string]string)}
}

func (c *SimpleCache) Get(key string) (string, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.cache[key]
	return val, ok
}

func (c *SimpleCache) Set(key, value string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.cache[key] = value
}
