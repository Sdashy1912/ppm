package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/target/bean"
)

// ReqTargetDetailAdd input for update target details add usecase
type ReqTargetDetailAdd struct {
	Info     basedto.ReqInfo       `json:"-"`
	TargetID string                `json:"-"`
	Detail   bean.TargetDetailBean `json:"detail"`
}
