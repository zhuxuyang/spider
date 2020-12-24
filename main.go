package main

import (
	"time"

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

	resource.GetDB().AutoMigrate(&model.Work{})
	biz.DealOneWork(viper.GetInt64("startSubjectID"), "")

	ConsumeWork()
}

func ConsumeWork() {
	lastID := viper.GetInt64("lastWorkID")
	for {
		list := model.FindWorkList(lastID)
		for _, work := range list {
			config.SpiderSleep()
			lastID = work.ID
			biz.DealOneWork(work.SourceID, work.BindISBN)
		}
		time.Sleep(time.Microsecond)
	}
}
