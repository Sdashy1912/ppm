package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulTemplateList input for vul template list usecase
type ReqVulTemplateList struct {
	Info       basedto.ReqInfo `json:"-"`
	CategoryID string          `json:"-"`
	OrderBy    string          `json:"-"`
}
