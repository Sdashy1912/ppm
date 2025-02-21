package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/scope/bean"
)

// ResScopeList output for vulnerability scope list usecase
type ResScopeList struct {
	Info   basedto.ResInfo  `json:"info"`
	Total  int              `json:"total"`
	Scopes []bean.ScopeBean `json:"scopes"`
}
