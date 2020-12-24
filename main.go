package main

import (
	"zhuxuyang/spider/config"
	"zhuxuyang/spider/model"
	"zhuxuyang/spider/resource"

	"github.com/spf13/viper"
)

func main() {
	config.InitViper()
	resource.InitLogger()
	resource.Logger.Info("开始")
	dbConf := viper.GetStringMapString("database")
	resource.InitDB(dbConf["user"], dbConf["password"], dbConf["host"], dbConf["port"], dbConf["name"])
	resource.GetDB().LogMode(true)
	resource.GetDB().AutoMigrate(&model.Proxy{})
	resource.GetDB().Save(&model.Proxy{
		IpPort: "http://47.99.195.197:8080",
	})
	resource.GetDB().Save(&model.Proxy{
		IpPort: "http://127.0.0.1:1087",
	})
	resource.GetDB().Save(&model.Proxy{
		IpPort: "http://47.112.119.49:8081",
	})
	//biz.DouBanSpiderStart(viper.GetInt64("startSubjectID"))
	//"https://book.douban.com/subject/3794471"
	//biz.GetAllTypes()
	//utils.ClientTestFunc("https://book.douban.com/subject/10546125")
}

//2020/12/21 19:17:27 <nil> [] Get "https://book.douban.com/subject/10546125": read tcp 192.168.29.104:54106->188.130.255.5:80: read: connection reset by peer
