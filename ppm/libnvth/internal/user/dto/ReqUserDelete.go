package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqUserDelete input for vul user get usecase
type ReqUserDelete struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
