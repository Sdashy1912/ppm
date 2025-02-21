package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqProjectGet input for vul project get usecase
type ReqProjectGet struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
