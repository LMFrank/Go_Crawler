package engine

import (
	"crawler_v2.0/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Schedular   Schedular
	WorkerCount int
	ItemChan    chan Item
}

type Schedular interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Schedular.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Schedular.WorkerChan(), out, e.Schedular)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Schedular.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			item := item
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Schedular.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
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
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.ProxyFetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(body), nil
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
