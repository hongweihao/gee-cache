package gee_cache

import (
	"net/http"
	"strings"
)

const BasePath = "/_gee_cache/"
type HttpPool struct {
	// 记录自己的host:port
	self string
	basePath string
}

func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self:     self,
		basePath: BasePath,
	}
}

func (pool *HttpPool) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !strings.HasPrefix(req.URL.Path, BasePath) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	parts := strings.SplitN(req.URL.Path[len(BasePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	name := parts[0]
	key := parts[1]
	if name == "" || key == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	group := GetGroup(name)
	if group == nil {
		http.Error(w, "Not found the group: " + name, http.StatusNotFound)
		return
	}

	byteView, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error() + key, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteView.ByteSlice())
}

type httpGetter struct {
	baseUrl string
}

func NewHttpGetter(baseUrl string) PeerGetter {
	if baseUrl == ""{
		baseUrl = BasePath
	}
	return &httpGetter{baseUrl: baseUrl}
}

func (h httpGetter) Get(key string) ([]byte, error) {
	


}

