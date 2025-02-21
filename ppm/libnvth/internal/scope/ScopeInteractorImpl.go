package scope

import (
	"ppm/libnvth/internal/scope/dao"
	"ppm/libnvth/internal/scope/dto"
	"gopkg.in/mgo.v2/bson"
)

// InteractorImpl implements Interactor
type InteractorImpl struct {
	dao dao.ScopeDAO
}

// NewInteractorImpl initialize a new InteractorImpl object
func NewInteractorImpl(dao dao.ScopeDAO) InteractorImpl {
	return InteractorImpl{dao}
}

var _ Interactor = (*InteractorImpl)(nil)

// List get a list of vulnerability scopes
// func (interactor InteractorImpl) List(req dto.ReqScopeList) dto.ResScopeList {
// 	resp := dto.ResScopeList{}
// 	scopes, err := interactor.dao.List()
// 	if err != nil {
// 		resp.Info.SetStatusInternalServerError()
// 		return resp
// 	}
// 	resp.Scopes = scopes
// 	resp.Total = len(scopes)
// 	resp.Info.SetStatusOK("Ok")
// 	return resp
// }

// Get find a vulnerability scope by ID
func (interactor InteractorImpl) Get(req dto.ReqScopeGet) dto.ResScopeGet {
	resp := dto.ResScopeGet{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	scope, err := interactor.dao.Get(bson.ObjectIdHex(req.ID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	resp.Scope = scope
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Create create a new vulnerability scope
func (interactor InteractorImpl) Create(req dto.ReqScopeCreate) dto.ResScopeCreate {
	resp := dto.ResScopeCreate{}
	errs := make(map[string]string)
	if !bson.IsObjectIdHex(req.ProjectID) {
		errs["project_id"] = "invalid project id"
	}
	if len(errs) > 0 {
		resp.Info.SetStatusUnprocessableEntity(errs)
		return resp
	}
	param := dao.ScopeDAOParam{
		ID:        bson.NewObjectId(),
		Name:      req.Name,
		ProjectID: bson.ObjectIdHex(req.ProjectID),
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

// Update update an existing vulnerability scope
func (interactor InteractorImpl) Update(req dto.ReqScopeUpdate) dto.ResScopeUpdate {
	resp := dto.ResScopeUpdate{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	param := dao.ScopeDAOParam{
		ID:   bson.NewObjectId(),
		Name: req.Name,
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

// Delete remove an vulnerability scope
func (interactor InteractorImpl) Delete(req dto.ReqScopeDelete) dto.ResScopeDelete {
	resp := dto.ResScopeDelete{}
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
		resp.Info.SetStatusUnprocessableEntity("Cannot delete because this scope is associated with at least one targets")
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
