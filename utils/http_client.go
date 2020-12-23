package utils

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"zhuxuyang/spider/ip_proxy"

	"github.com/malisit/kolpa"
)

func GetHttpClient() (*http.Client, error) {
	c := &http.Client{
		Timeout: time.Second * 20,
	}
	return c, nil
}

func GetHttpProxyClient() (*http.Client, error) {
	urli := url.URL{}
	//ip := ip_proxy.GetNextProxyIp()
	urlproxy, err := urli.Parse(ip_proxy.GetNextConstIP())
	if err != nil {
		log.Println("GetHttpClient proxy err", err)
		return nil, err
	}
	c := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
		Timeout: time.Second * 20,
	}
	return c, nil
}

func ClientTestFunc(url string) {
	client, _ := GetHttpProxyClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for key, value := range spiderHeaderMap {
		req.Header.Add(key, value)
	}
	c := kolpa.C()
	req.Header.Add("User-Agent", c.UserAgent())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		//if err.Error()==fmt.Sprintf("Get \"%s\": Bad Request",url){
		//panic(err)
		time.Sleep(1 * time.Second)
		ClientTestFunc(url)
		//}

	}
	defer resp.Body.Close()
	s := make([]byte, 2048)
	log.Println(resp.Body.Read(s))
	log.Println(string(s))
}
