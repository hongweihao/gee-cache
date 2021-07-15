package gee_cache

import (
	"fmt"
	"hash/crc32"
	"log"
	"sort"
)

const DefaultReplicas = 3

// hash 函数的形式
type HashFunc func(data []byte) uint32

type Map struct {
	// hash 函数
	hash HashFunc
	// hash环，逻辑上的环，实际上存储的是虚拟节点在环上的 hash 值
	keys []int
	// 实际节点的虚拟节点倍数
	replicas int
	// 虚拟节点和实际节点的映射关系，虚拟节点hash -> 真实节点 name
	hashMap map[uint32]string
}

func NewMap(hash HashFunc, replicas int) *Map {
	m := new(Map)
	m.hash = hash
	if hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	m.replicas = replicas
	m.keys = make([]int, 0)
	m.hashMap = make(map[uint32]string, 0)
	return m
}

// Add 添加实际节点
func (m *Map) Add(nodes ...string) {
	for _, node := range nodes {
		if node == ""{
			log.Println("skip the node: " + node)
			continue
		}
		for i := 0; i < m.replicas; i++ {
			// 创建虚拟节点
			virtualKey := fmt.Sprintf("%v-%v", i, node)
			h := m.hash([]byte(virtualKey))
			m.keys = append(m.keys, int(h))
			m.hashMap[h] = node
		}
	}
	// 从小到大排序
	sort.Ints(m.keys)
}

// Get 获取key实际应该到哪个节点获取，返回节点名称
func (m *Map) Get(key string) string {
	h := m.hash([]byte(key))
	if node, ok:= m.hashMap[h]; ok {
		return node
	}

	hi := int(h)
	l:= len(m.keys)
	idx := sort.Search(l, func(i int) bool {
		// 遍历环的值，如果出现大于或等于key的hash值就找到了虚拟节点
		// 如果没找到会返回 l
		return m.keys[i] >= hi
	})

	return m.hashMap[uint32(m.keys[(idx%l)])]
}
