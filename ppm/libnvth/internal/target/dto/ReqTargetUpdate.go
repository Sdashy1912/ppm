package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqTargetUpdate input for vul scope update usecase
type ReqTargetUpdate struct {
	Info        basedto.ReqInfo `json:"-"`
	ID          string          `json:"-"`
	Name        string          `json:"name"`
	Platform    string          `json:"platform"`
	Description string          `json:"description"`
	AssigneeID  string          `json:"assignee_id"`
}
