package ip_proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"zhuxuyang/spider/config"
)

var ipList []string
var index int
var mutex sync.Mutex

func GetNextProxyIp() string {
	mutex.Lock()
	next := ""
	if index < len(ipList) {
		next = ipList[index]
		index = index + 1
		if index >= len(ipList) {
			index = 0
		}
	}
	mutex.Unlock()
	return next
}

func InitProxy() {

	list, err := getIpList()
	log.Println(list)
	log.Println(err)
	go func() {
		for {
			list, err := getIpList()
			log.Println(list)
			if err != nil {
				log.Println(err)
				time.Sleep(config.IpListUpdateTime)
				continue
			}
			index = 0
			mutex.Lock()
			ipList = list
			mutex.Unlock()

			time.Sleep(config.IpListUpdateTime)
		}
	}()
}

func getIpList() ([]string, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	httpClient.Transport = &http.Transport{
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	req, err := http.NewRequest("GET", config.IpProxyServerUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body) //把  body 内容读入字符串 s
	if err != nil {
		return nil, err
	}
	result := IpProxy{}
	err = json.Unmarshal(s, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" || result.Message == nil {
		return nil, errors.New(fmt.Sprintf("ip proxy server err %v", result))
	}

	list := make([]string, 0)
	for _, v := range result.Message {
		list = append(list, "http://"+v.Content)
	}
	return list, nil
}

type Message struct {
	ID                    int     `json:"id"`
	IP                    string  `json:"ip"`
	Port                  string  `json:"port"`
	SchemeType            int     `json:"scheme_type"`
	Content               string  `json:"content"`
	AssessTimes           int     `json:"assess_times"`
	SuccessTimes          int     `json:"success_times"`
	AvgResponseTime       float64 `json:"avg_response_time"`
	ContinuousFailedTimes int     `json:"continuous_failed_times"`
	Score                 float64 `json:"score"`
	InsertTime            int     `json:"insert_time"`
	UpdateTime            int     `json:"update_time"`
}

type IpProxy struct {
	Error   string    `json:"error"`
	Message []Message `json:"message"`
}
