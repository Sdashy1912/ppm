package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/target/bean"
)

// ResTargetGet output for get vulnerability target output
type ResTargetGet struct {
	Info   basedto.ResInfo `json:"info"`
	Target bean.TargetBean `json:"target"`
}
