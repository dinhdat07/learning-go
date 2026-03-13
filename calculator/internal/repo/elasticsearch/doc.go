package elasticsearch

import (
	"calculator/internal/model"
	"encoding/json"
	"time"
)

type CalcHistoryDoc struct {
	ID         int64           `json:"id"`
	CreatedAt  time.Time       `json:"created_at"`
	Mode       model.Mode      `json:"mode"`
	Input      json.RawMessage `json:"input"`
	Success    bool            `json:"success"`
	Output     json.RawMessage `json:"output,omitempty"`
	Error      *string         `json:"error,omitempty"`
	DurationMs int64           `json:"duration_ms"`
	Note       *string         `json:"note,omitempty"`
}
