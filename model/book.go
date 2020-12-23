package model

type Book struct {
	ID            int64
	ISBN          string
	Title         string // 书名
	Author        string // 作者
	Press         string // 出版社
	Producer      string //出品方
	Translator    string // 译者
	PubDate       string // 出版年
	Pages         string // 页数
	Price         string // 定价
	Binding       string // 装帧
	Series        string /// 丛书
	BookNum       string // 统一书号,几十年前国内出版的书没有ISBN号，当时只有国家统一书号
	OriginalTitle string // 原作名
	Cover         string // 封面
	Summary       string // 内容简介
	Tags          string // 标签
	Score         string // 评分
	Votes         string // 评价人数
}
