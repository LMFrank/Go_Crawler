package engine

import (
	"go_crawler/crawler"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, e := range seeds {
		requests = append(requests, e)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Crawling url: %s\n", r.Url)

		body, err := crawler.Crawl(r.Url)
		if err != nil {
			log.Printf("Crawl Error: %s\n", r.Url)
		}

		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item: %s\n", item)
		}
	}
}
