package gee_cache

import "testing"

func TestMap_Add(t *testing.T) {
	m := NewMap(nil, 3)
	m.Add("node1", "node2", "node3")

	t.Log(m.keys)

}

func TestMap_Get(t *testing.T) {
	m := NewMap(nil, 3)
	m.Add("node1", "node2", "node3")

	node := m.Get("kkkkkkk")


	t.Log(node)
}
