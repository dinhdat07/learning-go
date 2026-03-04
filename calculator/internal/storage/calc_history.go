package storage

import (
	"calculator/internal/models"
	"database/sql"
)

func SaveCalcHistory(db *sql.DB, h models.CalcHistory) error {
	_, err := db.Exec(`
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
