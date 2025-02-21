package scope

import "ppm/libnvth/internal/scope/dto"

// Interactor interface for vulnerability scopes interactor
type Interactor interface {
	// List(req dto.ReqScopeList) dto.ResScopeList
	Get(req dto.ReqScopeGet) dto.ResScopeGet
	Create(req dto.ReqScopeCreate) dto.ResScopeCreate
	Update(req dto.ReqScopeUpdate) dto.ResScopeUpdate
	Delete(req dto.ReqScopeDelete) dto.ResScopeDelete
}
