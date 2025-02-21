package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqScopeDelete input for vul scope get usecase
type ReqScopeDelete struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
