package dto

import (
	"ppm/libnvth/internal/basedto"
)

// ResVulCategoryDelete output for delete a vulnerability category usecase
type ResVulCategoryDelete struct {
	Info basedto.ResInfo `json:"info"`
}
