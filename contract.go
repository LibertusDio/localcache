package localcache

type LocalCache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}
