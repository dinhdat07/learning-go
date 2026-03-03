package models

import (
	"database/sql"
	"time"
)

type CalcHistory struct {
	ID         int64
	CreatedAt  time.Time
	Expression string
	Success    bool
	Result     sql.NullFloat64
	Error      sql.NullString
	DurationMs int64
}

func NewHistory(expr string, ans float64, err error, duration int64) CalcHistory {
	return CalcHistory{
		Expression: expr,
		Success:    err == nil,
		DurationMs: duration,
		Result: sql.NullFloat64{
			Float64: ans,
			Valid:   err == nil,
		},
		Error: sql.NullString{
			String: func() string {
				if err != nil {
					return err.Error()
				}
				return ""
			}(),
			Valid: err != nil,
		},
	}
}
