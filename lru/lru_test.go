package lru

import (
	"testing"
)

type String string

func (s String) Len() int64 {
	return int64(len([]byte(s)))
}

func TestSetAndGet(t *testing.T) {
	m := NewLRUCache(0, nil)
	m.Set("name", String("mkii"))

	valueInterface, ok := m.Get("name")
	if !ok {
		t.Fatal("Failed to get a value by this key", "name")
	}

	value := valueInterface.(String)
	if value != "mkii" {
		t.Fatal("got a invalid value", value)
	}

	t.Log(value)
}

func TestEnd(t *testing.T) {
	cache := NewLRUCache(0, nil)
	cache.Set("name1", String("mkii1"))
	cache.Set("name2", String("mkii2"))
	cache.Set("name3", String("mkii3"))

	end := cache.GetEnd()
	if end.(string) != "mkii1" {
		t.Fatal("Failed to add an entry to front of list")
	}

	_, ok := cache.Get("name1")
	if !ok {
		t.Fatal("Failed to get a value by this key", "name1")
	}

	end2 := cache.GetEnd()
	if end2.(string) != "mkii2" {
		t.Fatal("Failed to move to front after get a value")
	}

	t.Log(end2)
}

type User struct {
	Name string
	Age  int
}
