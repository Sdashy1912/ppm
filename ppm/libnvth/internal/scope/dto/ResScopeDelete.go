package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResScopeDelete output for delete a vulnerability scope usecase
type ResScopeDelete struct {
	Info basedto.ResInfo `json:"info"`
}
