package main

import (
	"log"

	"zhuxuyang/spider/biz"
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
	log.Println(resource.GetDB().AutoMigrate(&model.Book{}))
	log.Println(resource.GetDB().AutoMigrate(&model.SourceLost{}))

	biz.DouBanSpiderStart(3794471)
	//"https://book.douban.com/subject/3794471"
	//biz.GetAllTypes()
	//utils.ClientTestFunc("https://book.douban.com/subject/10546125")
}

//2020/12/21 19:17:27 <nil> [] Get "https://book.douban.com/subject/10546125": read tcp 192.168.29.104:54106->188.130.255.5:80: read: connection reset by peer
