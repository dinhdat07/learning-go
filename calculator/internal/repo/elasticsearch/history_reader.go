package elasticsearch

import (
	"bytes"
	"calculator/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
)

type HistoryReader struct {
	es *elasticsearch.Client
}

func NewHistoryReader(es *elasticsearch.Client) *HistoryReader {
	return &HistoryReader{es: es}
}

type HistorySearchRequest struct {
	Mode       *string
	Success    *bool
	NoteQuery  string
	ErrorQuery string
	From       *time.Time
	To         *time.Time
}

func (h *HistoryReader) Search(req HistorySearchRequest) error {
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

	return h.execute(query)
}

func (h *HistoryReader) CountBySuccess() error {
	filters := map[string]interface{}{
		"success_true": map[string]interface{}{
			"term": map[string]interface{}{"success": true},
		},
		"success_false": map[string]interface{}{
			"term": map[string]interface{}{"success": false},
		},
	}

	query := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"success_stats": map[string]interface{}{
				"filters": map[string]interface{}{
					"filters": filters,
				},
			},
		},
	}

	return h.execute(query)
}

func (h *HistoryReader) CountByMode() error {
	filters := map[string]interface{}{
		"mode_expression": map[string]interface{}{
			"term": map[string]interface{}{"mode": "expression"},
		},
		"mode_equation": map[string]interface{}{
			"term": map[string]interface{}{"mode": "equation"},
		},
		"mode_linear_system": map[string]interface{}{
			"term": map[string]interface{}{"mode": "linear_system"},
		},
	}

	query := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"mode_stats": map[string]interface{}{
				"filters": map[string]interface{}{
					"filters": filters,
				},
			},
		},
	}

	return h.execute(query)
}

func (h *HistoryReader) CountByError() error {
	query := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"error_stats": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "error.keyword",
				},
			},
		},
	}

	return h.execute(query)
}

func (h *HistoryReader) CountByDate() error {
	query := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"date_stats": map[string]interface{}{
				"date_histogram": map[string]interface{}{
					"field":             "created_at",
					"calendar_interval": "day",
					"format":            "yyyy-MM-dd",
				},
			},
		},
	}

	return h.execute(query)
}

func (h *HistoryReader) execute(query map[string]interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return err
	}

	res, err := h.es.Search(h.es.Search.WithIndex("calc-history"), h.es.Search.WithBody(&buf))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("es error: %s", body)
	}

	io.Copy(os.Stdout, res.Body)
	return nil
}

func isValidMode(mode string) bool {
	m := model.Mode(mode)
	return m == model.EquationMode || m == model.ExpressionMode || m == model.LinearSystemMode
}
