package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/vultemplate/bean"
)

// ResVulTemplateList output for vulnerability template list usecase
type ResVulTemplateList struct {
	Info      basedto.ResInfo        `json:"info"`
	Total     int                    `json:"total"`
	Templates []bean.VulTemplateBean `json:"templates"`
}
