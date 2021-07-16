package gee_cache

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url "net/url"
	"strings"
	"sync"
)

const (
	BasePath = "/_gee_cache/"
	Replicas = 3
)

type HttpPool struct {
	// 记录自己的host:port
	self     string
	basePath string
	m        sync.Mutex
	// 一致性hash对象
	hashMap *Map
	// 每个实际节点的client对象
	httpGetters map[string]*httpGetter
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
		http.Error(w, "Not found the group: "+name, http.StatusNotFound)
		return
	}

	byteView, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error()+key, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteView.ByteSlice())
}

type httpGetter struct {
	baseUrl string
}

func NewHttpGetter(baseUrl string) PeerGetter {
	return &httpGetter{baseUrl: baseUrl}
}

func (h httpGetter) Get(group, key string) ([]byte, error) {
	remoteUrl := strings.Join([]string{h.baseUrl, url.QueryEscape(group), url.QueryEscape(key)}, "/")
	resp, err := http.Get(remoteUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("server return: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
