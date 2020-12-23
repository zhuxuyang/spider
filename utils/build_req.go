package utils

import (
	"net/http"

	"github.com/malisit/kolpa"
)

var spiderHeaderMap = map[string]string{
	"Host":                      "movie.douban.com",
	"Connection":                "keep-alive",
	"Cache-Control":             "max-age=0",
	"Upgrade-Insecure-Requests": "1",
	//"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
	"Accept":  "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	"Referer": "https://movie.douban.com/top250",
}

func GetRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range spiderHeaderMap {
		req.Header.Add(key, value)
	}
	c := kolpa.C()
	req.Header.Add("User-Agent", c.UserAgent())
	return req, nil
}
