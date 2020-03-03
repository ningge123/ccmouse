package server

import (
	"ccmouse/crawler_distributed/persist"
	"ccmouse/crawler_distributed/rpcsupport"
	"github.com/olivere/elastic"
)

func ServeRpc(port, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(port, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}