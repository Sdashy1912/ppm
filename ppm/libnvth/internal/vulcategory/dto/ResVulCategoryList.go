package dto

import (
	"ppm/libnvth/internal/basedto"
	"ppm/libnvth/internal/vulcategory/bean"
)

// ResVulCategoryList output for vulnerability category list usecase
type ResVulCategoryList struct {
	Info       basedto.ResInfo            `json:"info"`
	Total      int                        `json:"total"`
	Categories []bean.VulCategoryListBean `json:"categories"`
}
