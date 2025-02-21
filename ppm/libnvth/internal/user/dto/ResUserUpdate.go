package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResUserUpdate output for update a vulnerability user usecase
type ResUserUpdate struct {
	Info basedto.ResInfo `json:"info"`
}
