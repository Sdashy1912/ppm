package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/vulcategory/bean"
)

// ResVulCategoryGet output for get vulnerability category output
type ResVulCategoryGet struct {
	Info     basedto.ResInfo         `json:"info"`
	Category bean.VulCategoryGetBean `json:"category"`
}
