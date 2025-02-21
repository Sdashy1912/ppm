package vultemplate

import "ppm/libnvth/internal/vultemplate/dto"

// Interactor interface for vulnerability templates interactor
type Interactor interface {
	List(req dto.ReqVulTemplateList) dto.ResVulTemplateList
	Get(req dto.ReqVulTemplateGet) dto.ResVulTemplateGet
	Create(req dto.ReqVulTemplateCreate) dto.ResVulTemplateCreate
	Update(req dto.ReqVulTemplateUpdate) dto.ResVulTemplateUpdate
	Delete(req dto.ReqVulTemplateDelete) dto.ResVulTemplateDelete
}
