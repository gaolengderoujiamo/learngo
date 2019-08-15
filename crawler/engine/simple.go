package engine

import (
	"learngo/crawler/fecher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	itemCount := 0
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}
	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Feching %s", r.Url)
	body, err := fecher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v",
			r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
