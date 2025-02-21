package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqUserGet input for vul user get usecase
type ReqUserGet struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
