package main

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/distribute/persist"
	"crawler_v2.0/distribute/rpcsupport"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	err = rpcsupport.ServeRpc(host, &persist.ItemSaverService{Client: client, Index: index})
	return err
}
