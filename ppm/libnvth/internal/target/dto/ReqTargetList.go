package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqTargetList input for vul scope list usecase
type ReqTargetList struct {
	Info basedto.ReqInfo `json:"-"`
}
