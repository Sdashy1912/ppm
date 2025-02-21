package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResProjectUpdate output for update a vulnerability project usecase
type ResProjectUpdate struct {
	Info basedto.ResInfo `json:"info"`
}
