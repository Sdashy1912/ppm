package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResTargetUpdate output for update a vulnerability scope usecase
type ResTargetUpdate struct {
	Info basedto.ResInfo `json:"info"`
}
