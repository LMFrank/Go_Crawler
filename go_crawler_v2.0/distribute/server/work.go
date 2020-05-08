package main

import (
	"go_crawler/distribute/rpcsupport"
	"go_crawler/distribute/work/server"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(":1235", &server.CrawlService{}))
}
