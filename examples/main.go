package main

import (
	"flag"
)

var db = map[string]interface{}{
	"mkii": "kkk",
}

func main() {
	// 当前节点的端口
	var port int
	// 是否启动api服务
	var api bool

	flag.IntVar(&port, "port", 8090, "server port")
	flag.BoolVar(&api, "api", false, "start a api server?")

	//节点列表
	//当前节点
	nodeMap := map[int]string{
		8090: "localhost:8090",
		8091: "localhost:8091",
		8092: "localhost:8092",
	}




}