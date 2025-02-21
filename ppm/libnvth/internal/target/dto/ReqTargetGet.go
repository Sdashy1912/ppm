package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqTargetGet input for vul scope get usecase
type ReqTargetGet struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
