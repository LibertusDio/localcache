package localcache

import (
	"errors"
	"sync"
)

type slowcache struct {
	Data map[string]interface{}
	L    *sync.Mutex
}

func NewSlowCache() LocalCache {
	cache := slowcache{
		Data: make(map[string]interface{}),
		L:    &sync.Mutex{},
	}
	return &cache
}

func (c *slowcache) Get(key string) interface{} {
	c.L.Lock()
	defer c.L.Unlock()
	return c.Data[key]
}

func (c *slowcache) Set(key string, value interface{}) {
	c.L.Lock()
	defer c.L.Unlock()
	c.Data[key] = value
}

func (c *slowcache) GetLockUnsafe() *sync.Mutex {
	return c.L
}

func (c *slowcache) LockWait(key string) {
	c.L.Lock()
}

func (c *slowcache) Unlock(key string) (err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("invalid unlock")
		}
	}()
	c.L.Unlock()
	return nil
}

func (c *slowcache) Lock(key string) (err error) {
	if !c.L.TryLock() {
		err = errors.New("invalid lock")
		return
	}
	return nil
}
