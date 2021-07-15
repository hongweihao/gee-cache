package gee_cache

import "sync"

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (getFunc GetterFunc) Get(key string) ([]byte, error) {
	return getFunc(key)
}

var (
	rw     sync.RWMutex
	groups = make(map[string]*Group, 0)
)

func GetGroup(name string) *Group {
	rw.RLock()
	defer rw.RUnlock()
	return groups[name]
}

func NewGroup(name string, nodes []string, maxBytes int64, getter Getter) *Group {
	group := new(Group)
	group.name = name
	group.getter = getter
	group.current = NewCache(maxBytes)
	group.hashMap = NewMap(nil, DefaultReplicas)

	rw.Lock()
	defer rw.Unlock()
	groups[name] = group
	return group
}

// 相当于一个节点
type Group struct {
	name    string
	current *Cache
	getter  Getter
	hashMap *Map
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
	bv := ByteView{v}
	group.current.Set(key, bv)
	return bv, nil
}

func (group *Group) PickPeer(key string) (PeerGetter, bool) {
	group.hashMap.Get()

}

// loadFromRemotePeer
// get node by key
// get value from node
func (group *Group) loadFromRemotePeer(key string) (ByteView, error) {



}
