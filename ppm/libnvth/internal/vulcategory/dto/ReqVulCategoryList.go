package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ReqVulCategoryList input for vul category list usecase
type ReqVulCategoryList struct {
	Info    basedto.ReqInfo `json:"-"`
	OrderBy string          `json:"-"`
}
