package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResScopeUpdate output for update a vulnerability scope usecase
type ResScopeUpdate struct {
	Info basedto.ResInfo `json:"info"`
}
