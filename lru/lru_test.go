package lru

import (
	"fmt"
	"testing"
)

type String string

func (s String) Len() int64 {
	return int64(len([]byte(s)))
}
func callback(key string, value Value){
	fmt.Println("Deleted:", key, value)
}

func TestGet(t *testing.T) {
	m := NewLRUCache(0, nil)
	m.Set("name", String("mkii"))

	v, ok := m.Get("name")
	if !ok {
		t.Fatal("Failed to get a value by this key", "name")
	}

	value := v.(String)
	if value != "mkii" {
		t.Fatal("got a invalid value", value)
	}

	t.Log(value)
}

func TestSet(t *testing.T) {
	v1 := String("mkii1")
	v2 := String("mkii2")
	v3 := String("mkii3")

	m := NewLRUCache(v1.Len() + v2.Len(), callback)
	t.Log("used:", m.usedBytes)

	m.Set("k1", v1)
	t.Log("used:", m.usedBytes)
	end1 := m.GetEnd().(String)
	t.Log(end1)

	m.Set("k2", v2)
	t.Log("used:", m.usedBytes)
	end2 := m.GetEnd().(String)
	t.Log(end2)

	m.Set("k3", v3)
	t.Log("used:", m.usedBytes)
	end3 := m.GetEnd().(String)
	t.Log(end3)
}

func TestEnd(t *testing.T) {
	cache := NewLRUCache(5000, nil)
	cache.Set("name1", String("mkii1"))
	cache.Set("name2", String("mkii2"))
	cache.Set("name3", String("mkii3"))

	end := cache.GetEnd()
	if end.(String) != "mkii1" {
		t.Fatal("Failed to add an entry to front of list")
	}

	_, ok := cache.Get("name1")
	if !ok {
		t.Fatal("Failed to get a value by this key", "name1")
	}

	end2 := cache.GetEnd()
	if end2.(String) != "mkii2" {
		t.Fatal("Failed to move to front after get a value")
	}

	t.Log(end2)
}

type User struct {
	Name string
	Age  int
}
