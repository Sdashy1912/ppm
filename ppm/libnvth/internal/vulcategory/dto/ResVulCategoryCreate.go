package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResVulCategoryCreate output for create a new vulnerability category usecase
type ResVulCategoryCreate struct {
	Info basedto.ResInfo `json:"-"`
	ID   string          `json:"id"`
	Name string          `json:"name"`
}
