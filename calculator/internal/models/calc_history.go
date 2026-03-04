package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type CalcHistory struct {
	ID        int64
	CreatedAt time.Time

	Mode    string
	Input   json.RawMessage
	Success bool

	Output     json.RawMessage
	Error      sql.NullString
	DurationMs int64
}

func NewHistory(mode string, input any, output any, err error, duration int64) CalcHistory {
	inputJson, _ := json.Marshal(input)
	outputJson, _ := json.Marshal(output)

	return CalcHistory{
		Mode:       mode,
		Input:      inputJson,
		Success:    err == nil,
		DurationMs: duration,
		Output:     outputJson,
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
