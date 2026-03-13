package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type CalcHistory struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime; not null"`

	Mode    Mode            `gorm:"type:text;not null"`
	Input   json.RawMessage `gorm:"type:jsonb;not null"`
	Success bool            `gorm:"not null"`

	Output     json.RawMessage `gorm:"type:jsonb"`
	Error      sql.NullString  `json:"error"`
	DurationMs int64           `json:"duration_ms"`
	Note       sql.NullString  `json:"note"`
}

// to set the name of the table instead of pluralizing
func (CalcHistory) TableName() string {
	return "calc_history"
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

func (h CalcHistory) String() string {
	var note string
	if h.Note.Valid {
		note = h.Note.String
	}

	var errMsg string
	if h.Error.Valid {
		errMsg = h.Error.String
	}

	return fmt.Sprintf(
		"[#%d] %s\nMode: %s\nSuccess: %t\nDuration: %d ms\nError: %s\nNote: %s\n",
		h.ID,
		h.CreatedAt.Format(time.RFC3339),
		h.Mode,
		h.Success,
		h.DurationMs,
		errMsg,
		note,
	)
}
