package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/project/bean"
)

// ResProjectList output for vulnerability project list usecase
type ResProjectList struct {
	Info     basedto.ResInfo    `json:"info"`
	Total    int                `json:"total"`
	Projects []bean.ProjectBean `json:"projects"`
}
