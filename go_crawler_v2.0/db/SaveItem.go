package db

import "log"

func SaveItem() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemcount := 0

		for {
			item := <-out
			log.Printf("SaveItem: Got %d, %v\n", itemcount, item)
			itemcount++
		}
	}()

	return out
}
