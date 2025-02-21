package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqProjectDelete input for vul project get usecase
type ReqProjectDelete struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
