package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	//队列可以存redis
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := work(r)
		//todo 如果异常，添加到队列中
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items.Items {
			log.Printf("Got item %v ", item)

		}
	}

}

func work(r Request) (ParseResult, error) {
	log.Printf("Fetching url %s ", r.Url)

	body, err := r.Fetch()

	if err != nil {
		log.Printf("Fetcher : error fetching  url %s : %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
