package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqTargetDelete input for vul scope get usecase
type ReqTargetDelete struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
