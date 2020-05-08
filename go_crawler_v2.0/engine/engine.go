package engine

import (
	"go_crawler/crawler"
	"log"
)

type Schedular interface {
	Submit(request Request)
	Run()
	WorkReady(chan Request)
	Workchan() chan Request
}

type Processor func(Request) (ParseResult, error)

type ConcurrentEngine struct {
	Scheduler        Schedular
	WorkCount        int
	ItemChan         chan Item
	RequestProcessor Processor
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkCount; i++ {
		CreateWork(e.Scheduler.Workchan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//itemcount := 0
	for {
		result := <-out

		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
			//log.Printf("Got Itme: %d, %v", itemcount, item)
			//itemcount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult, s Schedular) {
	go func() {
		for {
			s.WorkReady(in)
			request := <-in

			result, err := worker(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Crawl url: %s\n", r.Url)

	body, err := crawler.ProxyCrawl(r.Url)

	if err != nil {
		log.Printf("Crawl Error: %s\n", r.Url)
		return ParseResult{}, err
	}

	return r.Parse.Parse(body, r.Url), nil
}
