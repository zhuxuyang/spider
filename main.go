package main

import (
	"zhuxuyang/spider/biz"
	"zhuxuyang/spider/resource"
)

func main() {

	//ip_proxy.InitProxy()
	//time.Sleep(5 * time.Second)
	//log.Println(ip_proxy.GetNextProxyIp())
	//log.Println(biz.GetHttpClient())
	resource.InitLogger()
	resource.Logger.Info("开始")
	biz.DouBanSpiderStart("https://book.douban.com/subject/3794471")

	//biz.GetAllTypes()
	//utils.ClientTestFunc("https://book.douban.com/subject/10546125")
}

//2020/12/21 19:17:27 <nil> [] Get "https://book.douban.com/subject/10546125": read tcp 192.168.29.104:54106->188.130.255.5:80: read: connection reset by peer
