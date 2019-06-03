package model

import "time"

type Profile struct {
	Age  string
	Name string
}

type ArticleList struct {
	Title string //标题
	Url   string
	Ctime time.Duration //创建时间
	Ptime time.Duration //发布时间
}

type Article struct {
	Title   string
	Content string
	Url     string
	Ctime   time.Duration //创建时间
	Ptime   time.Duration //发布时间
}
