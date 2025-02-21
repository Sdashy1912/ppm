package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResVulTemplateDelete output for delete a vulnerability template usecase
type ResVulTemplateDelete struct {
	Info basedto.ResInfo `json:"info"`
}
