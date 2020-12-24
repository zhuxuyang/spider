package ip_proxy

import (
	"time"

	"zhuxuyang/spider/model"
)

var constIPList = []string{
	"http://47.99.195.197:8080",
	"http://127.0.0.1:1087",
	"http://47.112.119.49:8081",
}
var constIPIndex = 0

var startTime = int64(0)

func GetNextConstIP() string {
	if len(constIPList) == 0 || time.Now().Unix()-startTime > 300 {
		constIPList = model.GetProxyList()
	}

	if constIPIndex >= len(constIPList) {
		constIPIndex = 0
	}
	r := constIPList[index]
	constIPIndex++
	return r
}

func GetCurrentConstIP() string {
	r := constIPList[index]
	return r
}
