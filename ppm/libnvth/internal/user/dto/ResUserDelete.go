package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResUserDelete output for delete a vulnerability user usecase
type ResUserDelete struct {
	Info basedto.ResInfo `json:"info"`
}
