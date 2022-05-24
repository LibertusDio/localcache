package localcache

import "sync"

type slowcache struct {
	Data map[string]interface{}
	Lock sync.Mutex
}

func NewSlowCache() LocalCache {
	cache := slowcache{
		Data: make(map[string]interface{}),
		Lock: sync.Mutex{},
	}
	return &cache
}

func (c *slowcache) Get(key string) interface{} {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	return c.Data[key]
}

func (c *slowcache) Set(key string, value interface{}) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Data[key] = value
}
