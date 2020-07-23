package persist

import (
	"context"
	"log"
)
import "gopkg.in/olivere/elastic.v5"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item Saver: Got Item #%d: %v", itemCount, item)
			itemCount++

			//save(item)
		}
	}()

	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.148.130:9200/"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("douban").
		Type("book").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", nil
	}

	return resp.Id, nil
}
