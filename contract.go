package localcache

import "sync"

type LocalCache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	GetLockUnsafe() *sync.Mutex
	LockWait(key string)
	Lock(key string) error
	Unlock(key string) error
}
