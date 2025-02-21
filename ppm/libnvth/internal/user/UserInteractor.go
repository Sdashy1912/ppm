package user

import "ppm/libnvth/internal/user/dto"

// Interactor interface for vulnerability users interactor
type Interactor interface {
	List(req dto.ReqUserList) dto.ResUserList
	Get(req dto.ReqUserGet) dto.ResUserGet
	Create(req dto.ReqUserCreate) dto.ResUserCreate
	Update(req dto.ReqUserUpdate) dto.ResUserUpdate
	Delete(req dto.ReqUserDelete) dto.ResUserDelete
}
