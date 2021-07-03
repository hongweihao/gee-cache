package gee_cache

import "testing"

func TestCache(t *testing.T) {
	cache := NewCache(10)
	cache.Set("k1", ByteView{[]byte("值")})
	cache.Set("k2", ByteView{[]byte("value2")})

	if byteView1, ok := cache.Get("k1"); ok {
		t.Log(byteView1.String())
	}
	if byteView2, ok := cache.Get("k2"); ok {
		t.Log(byteView2.String())
	}

	cache.Set("k3", ByteView{[]byte("value3")})
	cache.Set("k4", ByteView{[]byte("value4")})

	if byteView1, ok := cache.Get("k1"); ok {
		t.Log(byteView1.String())
	}
	if byteView2, ok := cache.Get("k2"); ok {
		t.Log(byteView2.String())
	}
}
