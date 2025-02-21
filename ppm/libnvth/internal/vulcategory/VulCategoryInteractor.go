package vulcategory

import "ppm/libnvth/internal/vulcategory/dto"

// Interactor interface for vulnerability categories interactor
type Interactor interface {
	List(req dto.ReqVulCategoryList) dto.ResVulCategoryList
	Get(req dto.ReqVulCategoryGet) dto.ResVulCategoryGet
	Create(req dto.ReqVulCategoryCreate) dto.ResVulCategoryCreate
	Update(req dto.ReqVulCategoryUpdate) dto.ResVulCategoryUpdate
	Delete(req dto.ReqVulCategoryDelete) dto.ResVulCategoryDelete
}
