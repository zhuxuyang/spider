package config

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

var configFilePath = "./config/config.yaml"

func InitViper() {
	configFile := flag.String("conf", configFilePath, "path of config file")
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Sprintf("viper read config is failed, err is %v configFile is %v ", err, configFile)
		panic(errStr)
	}
}

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
