package model

import (
	"time"

	"zhuxuyang/spider/resource"
)

type Proxy struct {
	ID        int64
	IpPort    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (m *Proxy) TableName() string {
	return "proxy"
}

func GetProxyList() []string {
	ips := make([]string, 0)
	err := resource.GetDB().Model(&Proxy{}).Pluck("ip_port", &ips).Error
	if err != nil {
		resource.Logger.Error("GetProxyList err " + err.Error())
	}
	return ips
}
