package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResVulTemplateCreate output for create a new vulnerability template usecase
type ResVulTemplateCreate struct {
	Info basedto.ResInfo `json:"info"`
}
