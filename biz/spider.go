package spider

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/malisit/kolpa"
)

type Book struct {
	Author        string // 作者
	Press         string // 出版社
	Producer      string //出品方
	translator    string // 译者
	PubDate       string // 出版年
	Pages         string // 页数
	Price         string // 定价
	Binding       string // 装帧
	Series        string /// 丛书
	ISBN          string
	BookNum       string   // 统一书号,几十年前国内出版的书没有ISBN号，当时只有国家统一书号
	Title         string   // 书名
	OriginalTitle string   // 原作名
	Cover         string   // 封面
	Summary       string   // 内容简介
	Tags          []string // 标签
	Score         string   // 评分
	Votes         string   // 评价人数
}

var spiderHeaderMap = map[string]string{
	"Host":                      "movie.douban.com",
	"Connection":                "keep-alive",
	"Cache-Control":             "max-age=0",
	"Upgrade-Insecure-Requests": "1",
	//"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
	"Accept":  "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	"Referer": "https://movie.douban.com/top250",
}

func GetISBNInfo(url string) (book *Book, likeUrlList []string, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	for key, value := range spiderHeaderMap {
		req.Header.Add(key, value)
	}
	c := kolpa.C()
	req.Header.Add("User-Agent", c.UserAgent())

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
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

	book = &Book{
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
			book.translator = v
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
	book.Tags = make([]string, 0)
	for _, v := range tagList {
		if v != "" {
			book.Tags = append(book.Tags, strings.TrimSpace(v))
		}
	}
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
