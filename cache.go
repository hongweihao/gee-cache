package gee_cache

import (
	"github.com/hongweihao/gee-cache/lru"
	"sync"
)

type Cache struct {
	m        sync.Mutex
	lru      *lru.Cache
	maxBytes uint64
}

func NewCache() *Cache {

	return nil
}

// ByteView 用来标识缓存值，只读
type ByteView struct {
	b []byte
}

// Len 计算对象的字节数
func (byteView *ByteView) Len() int {

	return 0
}

// ByteSlice 返回一个拷贝值
func (byteView ByteView) ByteSlice() []byte {

	return nil
}

// String 返回字符串表示
func (byteView ByteView) String() string {

	return ""
}

func (byteView ByteView) cloneBytes(b []byte) []byte {

	return nil
}
