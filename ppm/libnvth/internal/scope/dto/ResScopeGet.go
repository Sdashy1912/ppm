package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/scope/bean"
)

// ResScopeGet output for get vulnerability scope output
type ResScopeGet struct {
	Info  basedto.ResInfo `json:"info"`
	Scope bean.ScopeBean  `json:"scope"`
}
