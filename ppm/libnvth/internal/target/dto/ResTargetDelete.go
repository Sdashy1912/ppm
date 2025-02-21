package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResTargetDelete output for delete a vulnerability scope usecase
type ResTargetDelete struct {
	Info basedto.ResInfo `json:"info"`
}
