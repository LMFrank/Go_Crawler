package main

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/distribute/rpcsupport"
	"crawler_v2.0/distribute/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
