package elasticsearch

import (
	"bytes"
	"calculator/internal/model"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9"
)

type HistoryIndexer struct {
	es *elasticsearch.Client
}

func NewHistoryIndexer(es *elasticsearch.Client) *HistoryIndexer {
	return &HistoryIndexer{es: es}
}

// Basic CRUD operations to sync data with primary DB
func (r *HistoryIndexer) Delete(id int64) error {
	res, err := r.es.Delete("calc-history", fmt.Sprint(id))
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("elasticsearch clear error: %s", res.String())
	}
	return nil
}

func (r *HistoryIndexer) Index(record model.CalcHistory) error {
	doc := toESDoc(record)
	data, _ := json.Marshal(doc)
	res, err := r.es.Index("calc-history", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("elasticsearch clear error: %s", res.String())
	}

	return nil
}

func (r *HistoryIndexer) Clear() error {
	query := `{
		"query": {
			"match_all": {}
		}
	}`
	res, err := r.es.DeleteByQuery([]string{"calc-history"}, bytes.NewReader([]byte(query)))
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("elasticsearch clear error: %s", res.String())
	}

	return nil
}

func toESDoc(h model.CalcHistory) CalcHistoryDoc {
	var errPtr *string
	if h.Error.Valid {
		s := h.Error.String
		errPtr = &s
	}

	var notePtr *string
	if h.Note.Valid {
		s := h.Note.String
		notePtr = &s
	}

	return CalcHistoryDoc{
		ID:         h.ID,
		CreatedAt:  h.CreatedAt,
		Mode:       h.Mode,
		Input:      h.Input,
		Success:    h.Success,
		Output:     h.Output,
		Error:      errPtr,
		DurationMs: h.DurationMs,
		Note:       notePtr,
	}
}
