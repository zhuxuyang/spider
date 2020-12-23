package config

import (
	"math/rand"
	"time"
)

const DouBanSpiderSleepTimeMultiple = 2 // 爬虫sleep函数的参数
func SpiderSleep() {
	a := rand.Intn(3)
	time.Sleep(time.Duration(a+5) / DouBanSpiderSleepTimeMultiple * time.Second)
}

const (
	IpProxyServerUrl = "http://47.99.195.197:10000/sql?query=SELECT%20*%20FROM%20PROXY%20WHERE%20SCORE%20%3E%205%20ORDER%20BY%20SCORE%20DESC%20limit%2060"
	//IpProxyServerUrl = "http://47.99.195.197:10000/sql?query=SELECT * FROM PROXY WHERE SCORE > 5 ORDER BY SCORE DESC limit 60"
	IpListUpdateTime = 1 * time.Minute
)
