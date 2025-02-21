package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqScopeGet input for vul scope get usecase
type ReqScopeGet struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
