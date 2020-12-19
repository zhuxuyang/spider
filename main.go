package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"zhuxuyang/spider/biz"
)

type SpiderOnc struct {
	ISBN    string
	Title   string
	LikeUrl string
}

func main() {
	startUrl := "https://book.douban.com/subject/25863515/"

	workerChan := make(chan *SpiderOnc, 10000)

	baseBook, likeUrls, err := spider.GetISBNInfo(startUrl)
	if err != nil {
		panic(err)
	}
	if baseBook != nil {
		log.Println(baseBook)
	}
	if len(likeUrls) > 0 {
		for _, v := range likeUrls {
			workerChan <- &SpiderOnc{ISBN: baseBook.ISBN, Title: baseBook.Title, LikeUrl: v}
		}
	}

	for workInfo := range workerChan {
		a := rand.Intn(3)
		time.Sleep(time.Duration(a) * time.Second)
		book, likeUrls, err := spider.GetISBNInfo(workInfo.LikeUrl)
		if err != nil {
			panic(err)
		}
		if book != nil {
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ")
			fmt.Println("\n")
			log.Println(workInfo.Title, "类似的书 ：", book.Title, " 详细信息：", book)
		}
		if len(likeUrls) > 0 {
			for _, v := range likeUrls {
				workerChan <- &SpiderOnc{ISBN: book.ISBN, Title: book.Title, LikeUrl: v}
			}
		}
	}
}
