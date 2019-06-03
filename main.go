package main

import (
	"crawer/engine"
	"crawer/persist"
	"crawer/scheduler"
	"crawer/zhenai/parser"
	"time"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:      "http://www.zhenai.com/zhenghun",
		Retry:    3,
		Timeout:  15 * time.Second,
		Interval: 5,
		Method:   engine.GET,
		Header:   map[string]string{
			//"Content-Type": "application/x-www-form-urlencoded; param=value",
		},
		VerifyProxy: false,
		VerifyTLS:   false,
		ParserFunc:  parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
