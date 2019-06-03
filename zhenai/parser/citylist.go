package parser

import (
	"crawer/engine"
	"crawer/model"
	"fmt"
	"regexp"
	"time"
)

//历史文章列表
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	cityList := model.CityList{}
	result := engine.ParseResult{}

	itemList := engine.ItemList{}
	itemList.Category = "CityList"
	for _, m := range matches {

		cityList.Name = string(m[2])
		cityList.Url = string(m[1])

		itemList.Items = append(itemList.Items, cityList)
		result.Items = itemList
		//m2城市名字
		//result.Items = append(result.Items.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:      string(m[1]),
			Retry:    3,
			Timeout:  50 * time.Second,
			Interval: 5,
			Method:   engine.GET,
			Header:   map[string]string{
				//"Content-Type": "application/x-www-form-urlencoded; param=value",
			},
			VerifyProxy: false,
			VerifyTLS:   false,
			ParserFunc:  ParseCity,
		})
		fmt.Printf("City: %s,Url : %s\n", m[2], m[1])

	}

	return result
}
