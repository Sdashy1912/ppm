package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulTemplateGet input for vul template get usecase
type ReqVulTemplateGet struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"id"`
}
