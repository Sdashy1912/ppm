package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResVulTemplateUpdate output for update a vulnerability template usecase
type ResVulTemplateUpdate struct {
	Info basedto.ResInfo `json:"info"`
}
