package main

import (
	"errors"
	cache "github.com/hongweihao/gee-cache"
	"net/http"
)

var db = map[string]interface{}{
	"mkii": "kkk",
}

func main() {
	self := "127.0.0.1:8089"
	cache.NewGroup("user", 100, cache.GetterFunc(func(key string) ([]byte, error) {
		if v, ok := db[key]; ok {
			return []byte(v.(string)), nil
		}
		return nil, errors.New("not found")
	}))

	pool := cache.NewHttpPool(self)
	http.ListenAndServe(self, pool)
}