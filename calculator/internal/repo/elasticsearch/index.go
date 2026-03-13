package elasticsearch

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v9"
)

func EnsureCalcHistoryIndex(es *elasticsearch.Client) error {
	res, err := es.Indices.Exists([]string{"calc-history"})
	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		createCalcHistory(es)
	}

	return nil
}

func createCalcHistory(es *elasticsearch.Client) error {
	mapping := `{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
		"mappings": {
			"dynamic": "strict",
			"properties": {
				"id" : {"type": "long"},
				"created_at" : {"type": "date"},
				"mode" : {"type": "keyword"},
				"input" : {"type": "object", "enabled": false},
				"success" : {"type": "boolean"},
				"output" : {"type": "object", "enabled": false},
				"error" : {"type": "text", "fields" : {
					"keyword" : {"type": "keyword"}
				}},
				"duration_ms" : {"type": "integer"},
				"note" : {"type": "text", "fields" : {
					"keyword" : {"type": "keyword"}
				}}
			}
		}
	}`

	res, err := es.Indices.Create("calc-history", es.Indices.Create.WithBody(strings.NewReader(mapping)))
	defer res.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
