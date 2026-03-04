package repo

import (
	"calculator/internal/model"
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("history record not found")

type HistoryRepo struct {
	db *sql.DB
}

func NewHistoryRepo(db *sql.DB) *HistoryRepo {
	return &HistoryRepo{db: db}
}

func (r *HistoryRepo) Save(h model.CalcHistory) error {
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

func (r *HistoryRepo) List(limit int) ([]model.CalcHistory, error) {
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

	history := make([]model.CalcHistory, limit)
	for rows.Next() {
		h := model.CalcHistory{}
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

func (r *HistoryRepo) Delete(id int64) error {
	res, err := r.db.Exec(`DELETE FROM calc_history WHERE id = $1`, id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *HistoryRepo) Clear() error {
	_, err := r.db.Exec(`DELETE FROM calc_history`)
	return err
}

func (r *HistoryRepo) UpdateNote(id int64, note sql.NullString) error {
	res, err := r.db.Exec(`UPDATE calc_history SET note = $1 WHERE id = $2`, note, id)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return ErrNotFound
	}

	return nil
}
