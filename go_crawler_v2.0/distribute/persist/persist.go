package persist

import (
	"go_crawler/db"
	"go_crawler/engine"
	"gopkg.in/olivere/elastic.v5"
)

type ItemService struct {
	Client *elastic.Client
}

func (s *ItemService) Save(item engine.Item, result *string) error {
	err := db.Save(s.Client, item)

	if err == nil {
		*result = "ok"
	}

	return err

}
