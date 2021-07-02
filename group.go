package gee_cache

import "sync"

type Getter interface {
	Get(key string) (ByteView, error)
}

type GetterFunc func(key string) (ByteView, error)

func (getFunc GetterFunc) Get(key string) (ByteView, error) {
	return getFunc(key)
}

var (
	rw     sync.RWMutex
	groups = make(map[string]*Group, 0)
)

func NewGroup(name string, maxBytes int64, getter Getter) *Group {
	group := new(Group)
	group.name = name
	group.getter = getter
	group.current = NewCache(maxBytes)

	rw.Lock()
	defer rw.Unlock()
	groups[name] = group
	return group
}

func GetGroup(name string) *Group {
	rw.RLock()
	defer rw.RUnlock()
	return groups[name]
}

type Group struct {
	name    string
	current *Cache
	getter  Getter
}

// 先从current获取
// 如果获取失败则调用getter获取
// 获取到了需要跟更新到current
func (group *Group) Get(key string) (ByteView, error) {
	if v, ok := group.current.Get(key); ok {
		return v, nil
	}
	return group.load(key)
}

func (group *Group) load(key string) (ByteView, error) {
	return group.loadLocally(key)
}

func (group *Group) loadLocally(key string) (ByteView, error) {
	v, err := group.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	group.current.Set(key, v)
	return v, nil
}
