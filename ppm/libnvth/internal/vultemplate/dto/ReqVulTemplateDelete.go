package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulTemplateDelete input for vul template get usecase
type ReqVulTemplateDelete struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
