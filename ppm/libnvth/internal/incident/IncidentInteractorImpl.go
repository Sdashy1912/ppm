package incident

import (
	"time"

	"ppm/libnvth/internal/incident/dao"
	"ppm/libnvth/internal/incident/dto"
	"gopkg.in/mgo.v2/bson"
)

// InteractorImpl implements Interactor
type InteractorImpl struct {
	dao dao.IncidentDAO
}

// NewInteractorImpl initialize a new InteractorImpl object
func NewInteractorImpl(dao dao.IncidentDAO) InteractorImpl {
	return InteractorImpl{dao}
}

var _ Interactor = (*InteractorImpl)(nil)

// List get a list of vulnerability incidents
func (interactor InteractorImpl) List(req dto.ReqIncidentList) dto.ResIncidentList {
	resp := dto.ResIncidentList{}
	param := dao.IncidentListDAOParam{}
	if req.ProjectID != "" {
		if !bson.IsObjectIdHex(req.ProjectID) {
			resp.Info.SetStatusUnprocessableEntity("parameter project_id is invalid")
			return resp
		}
		bsonID := bson.ObjectIdHex(req.ProjectID)
		param.ProjectID = &bsonID
	} else {
		param.ProjectID = nil
	}
	incidents, err := interactor.dao.List(param)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Incidents = incidents
	resp.Total = len(incidents)
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Get find a vulnerability incident by ID
func (interactor InteractorImpl) Get(req dto.ReqIncidentGet) dto.ResIncidentGet {
	resp := dto.ResIncidentGet{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	incident, err := interactor.dao.Get(bson.ObjectIdHex(req.ID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	resp.Incident = incident
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Create create a new vulnerability incident
func (interactor InteractorImpl) Create(req dto.ReqIncidentCreate) dto.ResIncidentCreate {
	resp := dto.ResIncidentCreate{}
	errs := make(map[string]string)
	if !bson.IsObjectIdHex(req.ProjectID) {
		errs["project_id"] = "invalid project id"
	}
	if !bson.IsObjectIdHex(req.ResponsiblerID) {
		errs["responsibler_id"] = "invalid responsibler id"
	}
	if len(errs) > 0 {
		resp.Info.SetStatusUnprocessableEntity(errs)
		return resp
	}
	now := time.Now()
	param := dao.IncidentDAOParam{
		ID:             bson.NewObjectId(),
		Name:           req.Name,
		CreatedAt:      &now,
		Description:    req.Description,
		Solution:       req.Solution,
		Time:           req.Time,
		Severity:       req.Severity,
		Type:           req.Type,
		ProjectID:      bson.ObjectIdHex(req.ProjectID),
		ResponsiblerID: bson.ObjectIdHex(req.ResponsiblerID),
		Status:         req.Status,
	}
	err := interactor.dao.Insert(param)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.InsertedID = param.ID.Hex()
	resp.Info.SetStatusOK("Created")
	return resp
}

// Update update an existing vulnerability incident
func (interactor InteractorImpl) Update(req dto.ReqIncidentUpdate) dto.ResIncidentUpdate {
	resp := dto.ResIncidentUpdate{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	if !bson.IsObjectIdHex(req.ResponsiblerID) {
		resp.Info.SetStatusUnprocessableEntity(map[string]string{"responsibler_id": "invalid responsibler id"})
		return resp
	}
	now := time.Now()
	param := dao.IncidentDAOParam{
		ID:             bson.ObjectIdHex(req.ID),
		Name:           req.Name,
		LastUpdate:     &now,
		Description:    req.Description,
		Solution:       req.Solution,
		Time:           req.Time,
		Severity:       req.Severity,
		Type:           req.Type,
		ResponsiblerID: bson.ObjectIdHex(req.ResponsiblerID),
		Status:         req.Status,
	}
	err := interactor.dao.Update(param)
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	resp.Info.SetStatusOK("Updated")
	return resp
}

// Delete remove an vulnerability incident
func (interactor InteractorImpl) Delete(req dto.ReqIncidentDelete) dto.ResIncidentDelete {
	resp := dto.ResIncidentDelete{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	_, err := interactor.dao.Get(bson.ObjectIdHex(req.ID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	err = interactor.dao.Remove(bson.ObjectIdHex(req.ID))
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Info.SetStatusOK("Deleted")
	return resp
}
