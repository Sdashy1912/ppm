package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulCategoryUpdate input for vul category update usecase
type ReqVulCategoryUpdate struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"-"`
	Name string          `json:"name"`
}
