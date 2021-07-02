package gee_cache

import (
	"github.com/hongweihao/gee-cache/lru"
	"sync"
)

type Cache struct {
	m        sync.Mutex
	lru      *lru.Cache
	maxBytes int64
}

func NewCache(maxBytes int64) *Cache {
	cache := new(Cache)
	cache.maxBytes = maxBytes
	return cache
}

func (cache *Cache) Set(key string, value ByteView) {
	if cache.lru == nil {
		cache.lru = lru.NewLRUCache(cache.maxBytes, nil)
	}
	cache.m.Lock()
	defer cache.m.Unlock()
	cache.lru.Set(key, value)
}

func (cache *Cache) Get(key string) (ByteView, bool) {
	cache.m.Lock()
	defer cache.m.Unlock()
	if value, ok := cache.lru.Get(key); ok {
		return value.(ByteView), ok
	}
	return ByteView{}, false
}
