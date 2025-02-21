package project

import "ppm/libnvth/internal/project/dto"

// Interactor interface for vulnerability projects interactor
type Interactor interface {
	List(req dto.ReqProjectList) dto.ResProjectList
	Get(req dto.ReqProjectGet) dto.ResProjectGet
	Create(req dto.ReqProjectCreate) dto.ResProjectCreate
	Update(req dto.ReqProjectUpdate) dto.ResProjectUpdate
	Delete(req dto.ReqProjectDelete) dto.ResProjectDelete
}
