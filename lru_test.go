package cache

import "testing"

func TestSetAndGet(t *testing.T) {
	cache := NewCache(0, nil)
	cache.Set("name", "mkii")

	valueInterface, ok := cache.Get("name")
	if !ok {
		t.Fatal("Failed to get a value by this key", "name")
	}

	value := valueInterface.(string)
	if value != "mkii" {
		t.Fatal("got a invalid value", value)
	}

	t.Log(value)
}

func TestEnd(t *testing.T) {
	cache := NewCache(0, nil)
	cache.Set("name1", "mkii1")
	cache.Set("name2", "mkii2")
	cache.Set("name3", "mkii3")

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
