package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqTargetCreate input for vul scope create usecase
type ReqTargetCreate struct {
	Info        basedto.ReqInfo `json:"-"`
	ID          string          `json:"-"`
	Name        string          `json:"name"`
	Platform    string          `json:"platform"`
	Description string          `json:"description"`
	AssigneeID  string          `json:"assignee_id"`
	ScopeID     string          `json:"scope_id"`
}
