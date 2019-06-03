package parser

import (
	"crawer/engine"
	"crawer/model"
	"regexp"
)

const CityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//文章列表
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CityRe)
	matches := re.FindAllSubmatch(contents, -1)
	//result := engine.ParseResult{}
	result := engine.ParseResult{}
	profile := model.Article{}
	for _, m := range matches {
		name := string(m[2])
		profile.Content = name
		itemList := engine.ItemList{}
		itemList.Category = "Article"
		itemList.Items = []interface{}{profile}
		result.Items = itemList
		//m2城市名字
		//result.Items = append(result.Items, " User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc:  ParseProfile,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
		//fmt.Printf("UserName: %s,Url : %s\n", m[2], m[1])
	}
	return result
}
