package cache

import "container/list"

type (
    Cache struct {
        // 最大使用内存
        //maxBytes int64
        // 当前使用内存
        //nBytes int64

        // map结构，为了提高查找效率
        cache map[string]*list.Element
        // 双向链表，为了管理节点的最近使用状态
        doubleList *list.List

        // key 过期回调
        callback func(string, interface{})
    }

    // list element 的结构
    entry struct {
        key string
        value interface{}
    }
)



func NewCache(maxBytes int64, callback func(string, interface{})) *Cache {
    return &Cache{
        //maxBytes: maxBytes,
        cache: make(map[string]*list.Element, 0),
        doubleList: list.New(),
        //callback: callback,
    }
}

// 查找
func (cache *Cache) Get(key string) (interface{}, bool) {
    element, ok := cache.cache[key]
    if !ok {
        return nil, false
    }

    // 找到了，把元素移动到队首
    cache.doubleList.MoveToFront(element)
    elementEntry := element.Value.(*entry)
    return elementEntry.value, ok
}

// 设置（新增/更新）
func (cache *Cache) Set(key string, value interface{}) {
    element, ok := cache.cache[key]

    // 找到了更新
    if ok {
        elementEntry := element.Value.(*entry)
        elementEntry.value = value
        element.Value = elementEntry
        return
    }

    // 没找到，新增
    elementEntry := &entry{
        key:   key,
        value: value,
    }
    front := cache.doubleList.PushFront(elementEntry)
    cache.cache[key] = front
}

// 删除
func (cache *Cache) Delete(key string) interface{} {
    element, ok := cache.cache[key]
    if !ok {
        return nil
    }
    value := cache.doubleList.Remove(element)
    delete(cache.cache, key)
    return value
}

func (cache *Cache) GetEnd() interface{} {
    element := cache.doubleList.Back()
    elementEntry := element.Value.(*entry)
    return elementEntry.value
}