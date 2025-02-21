package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqScopeList input for vul scope list usecase
type ReqScopeList struct {
	Info basedto.ReqInfo `json:"-"`
}
