package repo

import (
	"calculator/internal/model"
	"database/sql"
)

type HistoryRepo interface {
	Save(h model.CalcHistory) error
	List(limit int) ([]model.CalcHistory, error)
	Delete(id int64) error
	Get(id int64) (*model.CalcHistory, error)
	Clear() error
	UpdateNote(id int64, note sql.NullString) error
}
