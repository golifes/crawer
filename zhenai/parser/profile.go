package parser

import (
	"crawer/engine"
	"crawer/model"
	"regexp"
)

const ageRe = `([\d])`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	re := regexp.MustCompile(ageRe)
	profile := model.Profile{}
	profile.Name = name
	match := re.FindSubmatch(contents)
	if match != nil {
		age := string(match[1])
		profile.Age = age
	}
	itemList := engine.ItemList{}
	itemList.Items = []interface{}{profile}
	itemList.Category = "profile"
	result := engine.ParseResult{
		Items: itemList,
	}

	return result
}
