package service

import (
	"calculator/internal/model"
	"calculator/internal/repo"
	"database/sql"
	"log"
)

type HistoryService struct {
	historyRepo *repo.HistoryRepo
}

func NewHistoryService(r *repo.HistoryRepo) *HistoryService {
	return &HistoryService{r}
}

func (svc *HistoryService) Save(mode model.Mode, input any, output any, err error, duration int64) {
	historyRecord := model.NewHistory(mode, input, output, err, duration)
	if err := svc.historyRepo.Save(historyRecord); err != nil {
		log.Printf("warn: could not save history: %v", err)
	}
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
	return svc.historyRepo.UpdateNote(id, noteField)
}

func (svc *HistoryService) Delete(id int) error {
	return svc.historyRepo.Delete(int64(id))
}

func (svc *HistoryService) Clear() error {
	return svc.historyRepo.Clear()
}
