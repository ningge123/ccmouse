package client

import (
	"ccmouse/crawler/engine"
	"ccmouse/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(port string) (chan engine.Item, error) {
	ch := make(chan engine.Item)

	rpc, err := rpcsupport.NewClient(port)

	go func() {
		itemCount := 0
		for {
			item := <- ch
			log.Printf("Item Saver: Got Item #%d: %v", itemCount, item)
			itemCount++

			result := ""
			rpc.Call("ItemSaverService.Save", item, &result)

			if err != nil {
				log.Printf("Item Saver: Save error: %s", err)
			}
		}
	}()

	return ch, err
}
