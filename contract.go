package localcache

import "sync"

type LocalCache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	GetLockUnsafe() *sync.Mutex
}
