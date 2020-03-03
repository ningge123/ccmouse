package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"log"
	"ccmouse/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item, error)  {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for  {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount ++

			err := Save(index, client, item)

			if err != nil {
				log.Printf("Item Saver: error " + "saving item %v: %v", item, err)
			}

		}
	}()

	return out, nil
}

func Save(index string, client *elastic.Client, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	if item.Id == "" {
		return errors.New("must supply Id")
	}

	indexService := client.Index().Index(index).Type(item.Type)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.BodyJson(item).Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}