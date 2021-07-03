package gee_cache

import (
	"errors"
	"testing"
)

func TestGroup(t *testing.T) {
	// 接口型函数的用法，将一个函数强制转换成一个对象GetterFunc(func)
	group := NewGroup("mkii", 20, GetterFunc(func(key string) (ByteView, error) {
		t.Log("Getter was called...............key:" + key)
		if key == "k2" {
			return ByteView{[]byte("getter value")}, nil
		}
		return ByteView{}, errors.New("not found")
	}))

	group.current.Set("k1", ByteView{[]byte("mkii")})

	if v1, err := group.Get("k1"); err == nil {
		t.Log(v1.String())
	} else {
		t.Fatal(err)
	}

	if v2, err := group.Get("k2"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(v2.String())
	}

	if v3, err := group.Get("k3"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(v3.String())
	}
}
