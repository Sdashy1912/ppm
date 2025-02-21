package incident

import "ppm/libnvth/internal/incident/dto"

// Interactor interface for vulnerability incidents interactor
type Interactor interface {
	List(req dto.ReqIncidentList) dto.ResIncidentList
	Get(req dto.ReqIncidentGet) dto.ResIncidentGet
	Create(req dto.ReqIncidentCreate) dto.ResIncidentCreate
	Update(req dto.ReqIncidentUpdate) dto.ResIncidentUpdate
	Delete(req dto.ReqIncidentDelete) dto.ResIncidentDelete
}
