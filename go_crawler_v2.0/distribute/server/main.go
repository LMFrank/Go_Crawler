package main

import (
	"go_crawler/distribute/persist"
	"go_crawler/distribute/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func main() {

	serveRpc(":1234")
}

func serveRpc(host string) error {

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {

		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemService{
		Client: client,
	})
}
