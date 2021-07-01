package lru

import "container/list"

type (
	Cache struct {
		// 最大使用内存
		maxBytes int64
		// 当前使用内存
		usedBytes int64

		// map结构，为了提高查找效率
		m map[string]*list.Element
		// 双向链表，为了管理节点的最近使用状态
		list *list.List

		// key 过期回调
		callback func(string, Value)
	}

	// list element 的结构
	entry struct {
		key   string
		value Value
	}
)

type Value interface {
	Len() int64
}

func NewLRUCache(maxBytes int64, callback func(string, Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		m:        make(map[string]*list.Element, 0),
		list:     list.New(),
		callback: callback,
	}
}

// 查找
func (cache *Cache) Get(key string) (Value, bool) {
	e, ok := cache.m[key]
	if !ok {
		return nil, false
	}

	// 找到了，把元素移动到队首
	cache.list.MoveToFront(e)
	item := e.Value.(*entry)
	return item.value, ok
}

// 设置（新增/更新）
func (cache *Cache) Set(key string, value Value) {
	var addByte int64
	if e, ok := cache.m[key]; ok {
		item := e.Value.(*entry)
		addByte = value.Len() - item.value.Len()
		cache.checkBytes(addByte)

		item.value = value
		e.Value = item
		cache.list.MoveToFront(e)
	} else {
		// 没找到，新增
		item := &entry{
			key:   key,
			value: value,
		}
		addByte = item.value.Len()
		cache.checkBytes(addByte)
		front := cache.list.PushFront(item)
		cache.m[key] = front
	}

	cache.usedBytes += addByte
}

// checkBytes 检查添加的容量，判断是否需要淘汰
func (cache *Cache) checkBytes(addByte int64) {
	if addByte <= 0 {
		return
	}

	// 如果容量不足，把队尾的元素淘汰
	for cache.maxBytes > 0 && (cache.maxBytes - cache.usedBytes) < addByte {
		cache.weedOut()
	}
}

// weedOut LRU淘汰，删除 list 最后一个元素
func (cache *Cache) weedOut() {
	e := cache.list.Back()
	item := e.Value.(*entry)
	delete(cache.m, item.key)
	cache.list.Remove(e)
	cache.usedBytes -= item.value.Len()
	if cache.callback != nil{
		cache.callback(item.key, item.value)
	}
}

// 删除
func (cache *Cache) Delete(key string) Value {
	e, ok := cache.m[key]
	if !ok {
		return nil
	}
	item := cache.list.Remove(e).(*entry)
	delete(cache.m, key)
	cache.usedBytes -= item.value.Len()
	if cache.callback != nil{
		cache.callback(item.key, item.value)
	}
	return item.value
}

func (cache *Cache) GetEnd() interface{} {
	element := cache.list.Back()
	elementEntry := element.Value.(*entry)
	return elementEntry.value
}
