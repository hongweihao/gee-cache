package gee_cache

// 抽象节点选择接口
type PeerPicker interface {
	// PickPeer
	// @key
	// @return PeerGetter key 应该从此节点获取
	PickPeer(key string) (PeerGetter, bool)
}


// 抽象远程数据获取接口
type PeerGetter interface {
	Get(group, key string) ([]byte, error)
}

// lru cache
// cache
// group
// http pool
// hash map
// peer

// 流程
// 1. register nodes into group
// 2. group set nodes into hash map
// 3. group new current cache and add it into the group list

// 4. set key-value into the current cache
// 5. get value from the current cache by key

// 6. get value from the remote cache by key
// 6.1 get node from consistent hash
// 6.2 get value from node



