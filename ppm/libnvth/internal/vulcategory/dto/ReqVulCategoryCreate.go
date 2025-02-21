package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulCategoryCreate input for vul category create usecase
type ReqVulCategoryCreate struct {
	Info basedto.ReqInfo `json:"-"`
	ID   string          `json:"-"`
	Name string          `json:"name"`
}
