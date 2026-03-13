package elasticsearch

import (
	"log"

	"github.com/elastic/go-elasticsearch/v9"
)

func Connect() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{Addresses: []string{"http://localhost:9200"}}
	es, _ := elasticsearch.NewClient(cfg)
	res, err := es.Info()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := EnsureCalcHistoryIndex(es); err != nil {
		return nil, err
	}

	log.Println("[ElasticSearch] Connected to search engine succesfully!")
	return es, nil
}
