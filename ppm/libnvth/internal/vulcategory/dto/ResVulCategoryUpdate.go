package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResVulCategoryUpdate output for update a vulnerability category usecase
type ResVulCategoryUpdate struct {
	Info basedto.ResInfo `json:"info"`
}
