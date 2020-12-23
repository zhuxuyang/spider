package biz

import (
	"fmt"
	"log"
	"strings"
	"time"

	"zhuxuyang/spider/config"
	"zhuxuyang/spider/model"
	"zhuxuyang/spider/resource"
	"zhuxuyang/spider/utils"

	"github.com/PuerkitoBio/goquery"
)

type SpiderOnc struct {
	ISBN    string
	Title   string
	LikeUrl string
}

func GetISBNInfo(url string) (book *model.Book, likeUrlList []string, err error) {
	client, err := utils.GetHttpProxyClient()
	if err != nil { //此时没有代理，等着
		log.Println("GetISBNInfo sleep 代理还没准备好，2秒后再试")
		time.Sleep(2 * time.Second)
		return GetISBNInfo(url)
	}

	req, err := utils.GetRequest("GET", url)
	if err != nil {
		log.Println("GetISBNInfo  GetRequest  err", err.Error())
		return GetISBNInfo(url)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(fmt.Sprintf("client.Do(req) err %v", err))
		time.Sleep(1 * time.Second)
		return GetISBNInfo(url)
		//return nil, nil, err
	}
	if resp == nil || resp.Body == nil {
		log.Println(fmt.Sprintf("resp==nil||resp.Body==nil %v", resp))
		time.Sleep(1 * time.Second)
		return GetISBNInfo(url)
		//return nil, nil, err
	}
	defer resp.Body.Close()
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	baseMap := make(map[string]string, 0)
	titleNode := dom.Find("#mainpic > a")
	title, _ := titleNode.Attr("title")
	cover, _ := titleNode.Attr("href")

	book = &model.Book{
		Title: title,
		Cover: cover,
	}

	infoNode := dom.Find("#info")
	s := strings.ReplaceAll(infoNode.Text(), " ", "")
	infoStringList := strings.Split(strings.ReplaceAll(strings.TrimSpace(s), ":|：| ", ":"), "\n")

	length := len(infoStringList)
	for i := 0; i < length; i++ {
		s := infoStringList[i]
		if s == "" {
			continue
		}

		if strings.ContainsAny(s, ":") {
			row := strings.SplitN(s, ":", 2)
			if len(row) != 2 {
				continue
			}
			key := strings.TrimSpace(row[0])
			value := strings.TrimSpace(row[1])
			for j := i + 1; j < length; j++ {
				if strings.ContainsAny(infoStringList[j], ":") {
					break
				} else {
					if infoStringList[j] != "" {
						value = value + infoStringList[j]
					}
				}
			}
			baseMap[key] = value
		}
	}

	for k, v := range baseMap {
		switch k {
		case "作者":
			book.Author = v
		case "出版社":
			book.Press = v
		case "出品方":
			book.Producer = v
		case "译者":
			book.Translator = v
		case "出版年":
			book.PubDate = v
		case "页数":
			book.Pages = v
		case "原作名":
			book.OriginalTitle = v
		case "定价":
			book.Price = v
		case "装帧":
			book.Binding = v
		case "丛书":
			book.Series = v
		case "ISBN":
			book.ISBN = v
		case "统一书号":
			book.BookNum = v
		}
	}

	summaryNode := dom.Find("#link-report")
	introNode := summaryNode.Find("div[class=\"intro\"]").Last()
	book.Summary = introNode.Text()
	tagNode := dom.Find("#db-tags-section > div[class=indent]")
	tagList := strings.Split(strings.ReplaceAll(tagNode.Text(), " ", ""), "\n")
	tagList = make([]string, 0)
	for _, v := range tagList {
		if v != "" {
			tagList = append(tagList, strings.TrimSpace(v))
		}
	}
	book.Tags = strings.Join(tagList, ",")
	// 评分区域
	scoreNode := dom.Find("#interest_sectl")
	book.Score = scoreNode.Find("strong[property=\"v:average\"]").Text()
	book.Votes = scoreNode.Find("span[property=\"v:votes\"]").Text()

	// 类似的书
	likeNode := dom.Find("#db-rec-section").Find("a")
	likeUrlMap := make(map[string]string, 0)
	likeNode.Each(func(i int, selection *goquery.Selection) {
		u, _ := selection.Attr("href")
		if u != "" {
			likeUrlMap[u] = u
		}
	})
	// map 是去重的
	for k, _ := range likeUrlMap {
		likeUrlList = append(likeUrlList, k)
	}
	return book, likeUrlList, err
}

func DouBanSpiderStart(startUrl string) {
	//startUrl := "https://book.douban.com/subject/25863515/"

	//startUrl = "https://book.douban.com/subject/10546125/"
	workerChan := make(chan *SpiderOnc, 10000)

	workerChan <- &SpiderOnc{ISBN: "", Title: "", LikeUrl: startUrl}

	i := 0
	for workInfo := range workerChan {
		i++
		config.SpiderSleep()
		book, likeUrls, err := GetISBNInfo(workInfo.LikeUrl)
		if err != nil {
			workerChan <- workInfo
			log.Println(fmt.Sprintf("GetISBNInfo err requeue %v", err))
		}
		if book != nil {
			resource.Logger.Info(fmt.Sprintf("%s 类似的书 ：%s  详细信息：%v", workInfo.Title, book.Title, book))
		}
		if len(likeUrls) > 0 {
			for _, v := range likeUrls {
				workerChan <- &SpiderOnc{ISBN: book.ISBN, Title: book.Title, LikeUrl: v}
			}
		}
	}
}
