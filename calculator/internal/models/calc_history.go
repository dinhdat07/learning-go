package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type CalcHistory struct {
	ID        int64
	CreatedAt time.Time

	Mode    Mode
	Input   json.RawMessage
	Success bool

	Output     json.RawMessage
	Error      sql.NullString
	DurationMs int64
	Note       sql.NullString
}

func NewHistory(mode Mode, input any, output any, err error, duration int64) CalcHistory {
	in, _ := json.Marshal(input)
	if err != nil {
		in = json.RawMessage(`null`)
	}

	var out json.RawMessage
	if err == nil {
		b, err := json.Marshal(output)
		if err != nil {
			out = json.RawMessage(`null`)
			// if have error when marshal, json null
		} else {
			out = b
		}
	} else {
		// not have valid ans, got NULL
		out = nil
	}

	h := CalcHistory{
		Mode:       mode,
		Input:      in,
		Success:    err == nil,
		Output:     out,
		DurationMs: duration,
	}

	if err != nil {
		h.Error = sql.NullString{String: err.Error(), Valid: true}
	}

	return h
}
