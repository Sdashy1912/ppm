package project

import (
	"ppm/libnvth/internal/project/dao"
	"ppm/libnvth/internal/project/dto"
	"gopkg.in/mgo.v2/bson"
)

// InteractorImpl implements Interactor
type InteractorImpl struct {
	dao dao.ProjectDAO
}

// NewInteractorImpl initialize a new InteractorImpl object
func NewInteractorImpl(dao dao.ProjectDAO) InteractorImpl {
	return InteractorImpl{dao}
}

var _ Interactor = (*InteractorImpl)(nil)

// List get a list of vulnerability projects
func (interactor InteractorImpl) List(req dto.ReqProjectList) dto.ResProjectList {
	resp := dto.ResProjectList{}
	projects, err := interactor.dao.List()
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Projects = projects
	resp.Total = len(projects)
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Get find a vulnerability project by ID
func (interactor InteractorImpl) Get(req dto.ReqProjectGet) dto.ResProjectGet {
	resp := dto.ResProjectGet{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	project, err := interactor.dao.Get(bson.ObjectIdHex(req.ID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	resp.Project = project
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Create create a new vulnerability project
func (interactor InteractorImpl) Create(req dto.ReqProjectCreate) dto.ResProjectCreate {
	resp := dto.ResProjectCreate{}
	if !bson.IsObjectIdHex(req.CustomerID) {
		resp.Info.SetStatusUnprocessableEntity(map[string]string{"customer_id": "invalid customer id"})
		return resp
	}
	param := dao.ProjectDAOParam{
		ID:         bson.NewObjectId(),
		Name:       req.Name,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     req.Status,
		CustomerID: bson.ObjectIdHex(req.CustomerID),
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

// Update update an existing vulnerability project
func (interactor InteractorImpl) Update(req dto.ReqProjectUpdate) dto.ResProjectUpdate {
	resp := dto.ResProjectUpdate{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	if !bson.IsObjectIdHex(req.CustomerID) {
		resp.Info.SetStatusUnprocessableEntity(map[string]string{"customer_id": "invalid customer id"})
		return resp
	}
	param := dao.ProjectDAOParam{
		ID:         bson.NewObjectId(),
		Name:       req.Name,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     req.Status,
		CustomerID: bson.ObjectIdHex(req.CustomerID),
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

// Delete remove an vulnerability project
func (interactor InteractorImpl) Delete(req dto.ReqProjectDelete) dto.ResProjectDelete {
	resp := dto.ResProjectDelete{}
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
	ok, err := interactor.dao.IsSafeToDelete(bson.ObjectIdHex(req.ID))
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	if !ok {
		resp.Info.SetStatusUnprocessableEntity("Cannot delete because this project is associated with at least one scopes")
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
