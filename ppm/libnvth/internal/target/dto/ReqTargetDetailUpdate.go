package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/target/bean"
)

// ReqTargetDetailUpdate input for update target details list usecase
type ReqTargetDetailUpdate struct {
	Info     basedto.ReqInfo         `json:"-"`
	TargetID string                  `json:"-"`
	Details  []bean.TargetDetailBean `json:"details"`
}
