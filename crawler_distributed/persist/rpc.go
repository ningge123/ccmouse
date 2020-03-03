package persist

import (
	"ccmouse/crawler/engine"
	"ccmouse/crawler/persist"
	"github.com/olivere/elastic"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Index, s.Client, item)

	log.Printf("save profile %s", item)

	if err == nil {
		*result = "ok"
	} else {
		log.Printf("item:%s save error: %s", item, err)
	}

	return nil
}
