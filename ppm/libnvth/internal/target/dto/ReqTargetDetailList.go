package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqTargetDetailList input for get target details usecase
type ReqTargetDetailList struct {
	Info     basedto.ReqInfo `json:"-"`
	TargetID string          `json:"-"`
}
