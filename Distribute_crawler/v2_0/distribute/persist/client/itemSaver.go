package client

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/distribute/rpcsupport"
	"crawler_v2.0/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item Saver: Got Item #%d: \n%v", itemCount, item)
			itemCount++

			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: "+"saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}
