package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/target/bean"
)

// ResTargetList output for vulnerability target list usecase
type ResTargetList struct {
	Info    basedto.ResInfo   `json:"info"`
	Total   int               `json:"total"`
	Targets []bean.TargetBean `json:"targets"`
}
