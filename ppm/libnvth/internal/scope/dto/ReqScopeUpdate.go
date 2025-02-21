package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqScopeUpdate input for vul scope update usecase
type ReqScopeUpdate struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"-"`
	Name string          `json:"name"`
}
