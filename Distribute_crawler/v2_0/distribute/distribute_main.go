package main

import (
	"crawler_v2.0/distribute/config"
	itemsaverClient "crawler_v2.0/distribute/persist/client"
	"crawler_v2.0/distribute/rpcsupport"
	workerClient "crawler_v2.0/distribute/worker/client"
	"crawler_v2.0/doubanbook/parser"
	"crawler_v2.0/engine"
	"crawler_v2.0/schedular"
	"flag"
	"github.com/rs/zerolog/log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHosts = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts    = flag.String("worker_host", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaverClient.ItemSaver(*itemSaverHosts)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := workerClient.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Schedular:        &schedular.QueuedSchedular{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "https://book.douban.com",
		Parser: engine.NewFuncParser(parser.ParseBookTag, config.ParseBookTag),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Warn().Msgf("error connection to %s : %s", h, err)

		} else {
			clients = append(clients, client)
			log.Warn().Msgf("connected  to %s", h)
		}
	}
	out := make(chan *rpc.Client)
	// 持续纷发客户端
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
