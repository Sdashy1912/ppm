package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/project/bean"
)

// ResProjectGet output for get vulnerability project output
type ResProjectGet struct {
	Info    basedto.ResInfo  `json:"info"`
	Project bean.ProjectBean `json:"project"`
}
