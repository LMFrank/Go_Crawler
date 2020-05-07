package engine

import (
	"go_crawler/crawler"
	"log"
)

type Schedular interface {
	Submit(Request)
	configureWorkChan(chan Request)
}

type SimpleSchedular struct {
	workerChan chan Request
}

func (s *SimpleSchedular) Submit(r Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleSchedular) configureWorkChan(c chan Request) {
	s.workerChan = c
}

type ConcurrentEngine struct {
	Scheduler Schedular
	WorkCount int
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.configureWorkChan(in)

	for i := 0; i < e.WorkCount; i++ {
		CreateWork(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemcount := 0
	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got Itme: %d, %v", itemcount, item)
			itemcount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult) {
	go func() {
		for {
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

	body, err := crawler.Crawl(r.Url)

	if err != nil {
		log.Printf("Crawl Error: %s\n", r.Url)
		return ParseResult{}, err
	}

	return r.ParseFunc(body), nil
}
