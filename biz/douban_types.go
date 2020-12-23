package biz

import (
	"fmt"
	"log"
	"time"

	"zhuxuyang/spider/utils"

	"github.com/PuerkitoBio/goquery"
)

func GetAllTypes() {
	client, err := utils.GetHttpClient()
	if err != nil { //此时没有代理，等着
		log.Println("GetAllTypes sleep 代理还没准备好，2秒后再试")
		time.Sleep(2 * time.Second)
	}

	req, err := utils.GetRequest("GET", "https://book.douban.com/tag/?view=type")
	if err != nil {
		log.Println("GetAllTypes  GetRequest  err", err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(fmt.Sprintf("client.Do(req) err %v", err))
		time.Sleep(1 * time.Second)
	}
	if resp == nil || resp.Body == nil {
		log.Println(fmt.Sprintf("resp==nil||resp.Body==nil %v", resp))
		time.Sleep(1 * time.Second)
	}

	defer resp.Body.Close()
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}

	dom.Find(".tag-title-wrapper").Each(func(i int, selection *goquery.Selection) {
		log.Println(selection.Text())
		listNode := selection.Find(".tagCol").Find("td")
		listNode.Each(func(i int, item *goquery.Selection) {
			a := item.Find("a")
			url, _ := a.Attr("href")
			log.Println(url, a.Text())
			//.Attr("href")
			log.Println()
		})
	})
}
