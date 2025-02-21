package target

import "ppm/libnvth/internal/target/dto"

// Interactor interface for vulnerability targets interactor
type Interactor interface {
	// List(req dto.ReqTargetList) dto.ResTargetList
	Get(req dto.ReqTargetGet) dto.ResTargetGet
	Create(req dto.ReqTargetCreate) dto.ResTargetCreate
	Update(req dto.ReqTargetUpdate) dto.ResTargetUpdate
	Delete(req dto.ReqTargetDelete) dto.ResTargetDelete
	GetDetailList(req dto.ReqTargetDetailList) dto.ResTargetDetailList
	AddToDetailList(req dto.ReqTargetDetailAdd) dto.ResTargetDetailAdd
	UpdateDetailList(req dto.ReqTargetDetailUpdate) dto.ResTargetDetailUpdate
}
