package storage

import (
	"calculator/internal/models"
	"database/sql"
)

type HistoryRepo struct {
	db *sql.DB
}

func NewHistoryRepo(db *sql.DB) *HistoryRepo {
	return &HistoryRepo{db: db}
}

func (r *HistoryRepo) SaveHistory(h models.CalcHistory) error {
	_, err := r.db.Exec(`
		INSERT INTO calc_history (mode, input, success, output, error, duration_ms)
		VALUES ($1, $2, $3, $4, $5, $6)
	`,
		h.Mode,
		h.Input,
		h.Success,
		h.Output,
		h.Error,
		h.DurationMs,
	)

	return err
}

func (r *HistoryRepo) GetHistory(limit int) ([]models.CalcHistory, error) {
	var rows *sql.Rows
	var err error
	if limit <= 0 {
		rows, err = r.db.Query(`SELECT * FROM calc_history`)
	} else {
		rows, err = r.db.Query(`SELECT * FROM calc_history LIMIT $1`, limit)
	}

	if err != nil {
		return nil, err
	}

	history := make([]models.CalcHistory, limit)
	for rows.Next() {
		h := models.CalcHistory{}
		err := rows.Scan(
			&h.ID,
			&h.CreatedAt,
			&h.Mode,
			&h.Input,
			&h.Success,
			&h.Output,
			&h.Error,
			&h.DurationMs,
		)
		if err != nil {
			return nil, err
		}
		history = append(history, h)
	}

	return history, nil
}

func (r *HistoryRepo) DeleteHistory(id int) error {
	_, err := r.db.Exec(`DELETE FROM calc_history WHERE id=$1`, id)
	return err
}

func (r *HistoryRepo) ClearHistory() error {
	_, err := r.db.Exec(`DELETE FROM calc_history`)
	return err
}

func (r *HistoryRepo) UpdateNoteHistory(id int64, note string) error {
	_, err := r.db.Exec(`UPDATE calc_history SET note = $1 WHERE id = $2`, note, id)

	return err
}
