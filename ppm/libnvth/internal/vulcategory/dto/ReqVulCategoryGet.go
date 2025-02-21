package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulCategoryGet input for vul category get usecase
type ReqVulCategoryGet struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"-"`
}
