package client

import (
	"go_crawler/distribute/rpcsupport"
	"go_crawler/engine"
	"log"
)

func ItemSave(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemcount := 0

		for {
			item := <-out
			log.Printf("Item saver: Got %d,%v", itemcount, item)
			result := ""
			err = client.Call("ItemService.Save", item, &result)

			if err != nil {
				log.Printf("item saver: error saving item %v:%v", item, err)
			}
			itemcount++
		}
	}()
	return out, nil
}
