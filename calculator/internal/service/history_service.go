package service

import (
	"calculator/internal/model"
	"calculator/internal/repo"
	"calculator/internal/repo/elasticsearch"
	"database/sql"
	"log"
)

type HistoryService struct {
	historyRepo repo.HistoryRepo
	indexer     elasticsearch.HistoryIndexer
}

func NewHistoryService(r repo.HistoryRepo, i elasticsearch.HistoryIndexer) *HistoryService {
	return &HistoryService{r, i}
}

func (svc *HistoryService) Save(mode model.Mode, input any, output any, err error, duration int64) {
	historyRecord := model.NewHistory(mode, input, output, err, duration)
	if err := svc.historyRepo.Save(historyRecord); err != nil {
		log.Printf("warn: could not save history: %v", err)
		return
	}

	go func() {
		if err := svc.indexer.Index(historyRecord); err != nil {
			log.Printf("warn: failed to index history to es: %v", err)
		}
	}()
}

func (svc *HistoryService) List(limit int) ([]model.CalcHistory, error) {
	return svc.historyRepo.List(limit)
}

func (svc *HistoryService) UpdateNote(id int64, note string) error {
	var noteField sql.NullString

	if note != "" {
		noteField = sql.NullString{
			String: note,
			Valid:  true,
		}
	}

	if err := svc.historyRepo.UpdateNote(id, noteField); err != nil {
		return err
	}

	updated, err := svc.historyRepo.Get(id)
	if err != nil {
		return err
	}

	go func() {
		if err := svc.indexer.Index(*updated); err != nil {
			log.Printf("warn: failed to reindex updated history in es: %v", err)
		}
	}()

	return nil
}

func (svc *HistoryService) Delete(id int) error {
	if err := svc.historyRepo.Delete(int64(id)); err != nil {
		return err
	}

	go func() {
		if err := svc.indexer.Delete(int64(id)); err != nil {
			log.Printf("warn: failed to delete history from es: %v", err)
		}
	}()

	return nil
}

func (svc *HistoryService) Clear() error {
	if err := svc.historyRepo.Clear(); err != nil {
		return err
	}

	go func() {
		if err := svc.indexer.Clear(); err != nil {
			log.Printf("warn: failed to clear history index in es: %v", err)
		}
	}()

	return nil
}
