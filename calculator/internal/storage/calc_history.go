package storage

import (
	"calculator/internal/models"
	"database/sql"
)

func SaveCalcHistory(db *sql.DB, h models.CalcHistory) error {
	_, err := db.Exec(` INSERT INTO calc_history (expression, success, result, error, duration_ms)
						VALUES ($1, $2, $3, $4, $5)`,
		h.Expression,
		h.Success,
		h.Result,
		h.Error,
		h.DurationMs,
	)

	return err
}
