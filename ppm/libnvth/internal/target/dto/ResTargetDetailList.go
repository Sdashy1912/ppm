package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/target/bean"
)

// ResTargetDetailList output for get details list
type ResTargetDetailList struct {
	Info    basedto.ResInfo         `json:"info"`
	Total   int                     `json:"total"`
	Details []bean.TargetDetailBean `json:"details"`
}
