package gorm

import (
	"calculator/internal/model"
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("history record not found")

type HistoryRepo struct {
	db *gorm.DB
}

func NewHistoryRepo(db *gorm.DB) *HistoryRepo {
	return &HistoryRepo{db: db}
}

func (r *HistoryRepo) Save(h model.CalcHistory) error {
	result := r.db.Create(&h)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *HistoryRepo) List(limit int) ([]model.CalcHistory, error) {

	var history []model.CalcHistory
	var result *gorm.DB

	if limit <= 0 {
		result = r.db.Find(&history)
	} else {
		result = r.db.Limit(limit).Find(&history)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return history, nil
}

func (r *HistoryRepo) Get(id int64) (*model.CalcHistory, error) {

	var history model.CalcHistory
	var result *gorm.DB

	result = r.db.First(&history, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &history, nil
}

func (r *HistoryRepo) Delete(id int64) error {
	result := r.db.Delete(&model.CalcHistory{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *HistoryRepo) Clear() error {
	result := r.db.Delete(&model.CalcHistory{})
	return result.Error

}

func (r *HistoryRepo) UpdateNote(id int64, note sql.NullString) error {
	result := r.db.Model(&model.CalcHistory{}).Where("id = ?", id).Update("note", note)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
