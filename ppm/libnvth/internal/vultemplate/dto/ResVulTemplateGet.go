package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/vultemplate/bean"
)

// ResVulTemplateGet output for get vulnerability template output
type ResVulTemplateGet struct {
	Info     basedto.ResInfo      `json:"info"`
	Template bean.VulTemplateBean `json:"template"`
}
