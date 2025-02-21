package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulCategoryDelete input for vul category get usecase
type ReqVulCategoryDelete struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"-"`
}
