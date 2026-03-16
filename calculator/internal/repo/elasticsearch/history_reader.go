package elasticsearch

import (
	"bytes"
	"calculator/internal/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
)

type HistoryReader struct {
	es *elasticsearch.Client
}

func NewHistoryReader(es *elasticsearch.Client) *HistoryReader {
	return &HistoryReader{es: es}
}

func isValidMode(mode string) bool {
	m := model.Mode(mode)
	return m == model.EquationMode || m == model.ExpressionMode || m == model.LinearSystemMode
}

type HistorySearchRequest struct {
	Mode       *string
	Success    *bool
	NoteQuery  string
	ErrorQuery string
	From       *time.Time
	To         *time.Time
}

func Search(es *elasticsearch.Client, req HistorySearchRequest) error {
	must := []interface{}{}
	filter := []interface{}{}

	if req.Mode != nil {
		if isValidMode(*req.Mode) {
			return fmt.Errorf("not valid mode to filter")
		}
		filter = append(filter, map[string]interface{}{
			"term": map[string]interface{}{
				"mode": *req.Mode,
			},
		})
	}

	if req.Success != nil {
		filter = append(filter, map[string]interface{}{
			"term": map[string]interface{}{
				"success": *req.Success,
			},
		})
	}

	if req.From != nil || req.To != nil {
		rangeQuery := map[string]interface{}{}
		if req.From != nil {
			rangeQuery["gte"] = req.From
		}
		if req.To != nil {
			rangeQuery["lte"] = req.To
		}

		filter = append(filter, map[string]interface{}{
			"range": map[string]interface{}{
				"created_at": rangeQuery,
			},
		})
	}

	if req.NoteQuery != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"note": req.NoteQuery,
			},
		})
	}

	if req.ErrorQuery != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"error": req.ErrorQuery,
			},
		})
	}

	// build final query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must":   must,
				"filter": filter,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return err
	}

	res, err := es.Search(es.Search.WithIndex("calc_history"), es.Search.WithBody(&buf))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	//simple println
	fmt.Println(res)

	return nil
}
