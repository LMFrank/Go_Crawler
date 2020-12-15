package main

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/distribute/persist"
	"crawler_v2.0/distribute/rpcsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	err = rpcsupport.ServeRpc(host, &persist.ItemSaverService{Client: client, Index: index})
	return err
}
