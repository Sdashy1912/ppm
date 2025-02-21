package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqScopeCreate input for vul scope create usecase
type ReqScopeCreate struct {
	Info      basedto.ReqInfo `json:"-"`
	ID        string          `json:"-"`
	Name      string          `bson:"name"`
	ProjectID string          `json:"project_id"`
}
