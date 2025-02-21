package target

import (
	"ppm/libnvth/internal/target/dao"
	"ppm/libnvth/internal/target/dto"
	"gopkg.in/mgo.v2/bson"
    "log"
)

// InteractorImpl implements Interactor
type InteractorImpl struct {
	dao dao.TargetDAO
}

// NewInteractorImpl initialize a new InteractorImpl object
func NewInteractorImpl(dao dao.TargetDAO) InteractorImpl {
	return InteractorImpl{dao}
}

var _ Interactor = (*InteractorImpl)(nil)

// // List get a list of vulnerability targets
// func (interactor InteractorImpl) List(req dto.ReqTargetList) dto.ResTargetList {
// 	resp := dto.ResTargetList{}
// 	targets, err := interactor.dao.List()
// 	if err != nil {
// 		resp.Info.SetStatusInternalServerError()
// 		return resp
// 	}
// 	resp.Targets = targets
// 	resp.Total = len(targets)
// 	resp.Info.SetStatusOK("Ok")
// 	return resp
// }

// Get find a vulnerability target by ID
func (interactor InteractorImpl) Get(req dto.ReqTargetGet) dto.ResTargetGet {
	resp := dto.ResTargetGet{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	target, err := interactor.dao.Get(bson.ObjectIdHex(req.ID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
            log.Printf("WTF %s", err)
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	resp.Target = target
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Create create a new vulnerability target
func (interactor InteractorImpl) Create(req dto.ReqTargetCreate) dto.ResTargetCreate {
	resp := dto.ResTargetCreate{}
	errs := make(map[string]string)
	if !bson.IsObjectIdHex(req.ScopeID) {
		errs["scope_id"] = "invalid scope id"
	}
	if !bson.IsObjectIdHex(req.AssigneeID) {
		errs["assignee_id"] = "invalid assignee id"
	}
	if len(errs) > 0 {
		resp.Info.SetStatusUnprocessableEntity(errs)
		return resp
	}
	param := dao.TargetDAOParam{
		ID:          bson.NewObjectId(),
		Name:        req.Name,
		Description: req.Description,
		Platform:    req.Platform,
		ScopeID:     bson.ObjectIdHex(req.ScopeID),
		AssigneeID:  bson.ObjectIdHex(req.AssigneeID),
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

// Update update an existing vulnerability target
func (interactor InteractorImpl) Update(req dto.ReqTargetUpdate) dto.ResTargetUpdate {
	resp := dto.ResTargetUpdate{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	if !bson.IsObjectIdHex(req.AssigneeID) {
		resp.Info.SetStatusUnprocessableEntity(map[string]string{"assignee_id": "invalid assignee id"})
		return resp
	}
	param := dao.TargetDAOParam{
		ID:          bson.ObjectIdHex(req.ID),
		Name:        req.Name,
		Platform:    req.Platform,
		Description: req.Description,
		AssigneeID:  bson.ObjectIdHex(req.AssigneeID),
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

// Delete remove an vulnerability target
func (interactor InteractorImpl) Delete(req dto.ReqTargetDelete) dto.ResTargetDelete {
	resp := dto.ResTargetDelete{}
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

// UpdateDetailList update an entire details list
func (interactor InteractorImpl) UpdateDetailList(req dto.ReqTargetDetailUpdate) dto.ResTargetDetailUpdate {
	resp := dto.ResTargetDetailUpdate{}
	if !bson.IsObjectIdHex(req.TargetID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	err := interactor.dao.UpdateDetailList(bson.ObjectIdHex(req.TargetID), req.Details)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Inserted = len(req.Details)
	resp.Info.SetStatusOK("Ok")
	return resp
}

// AddToDetailList add a detail to the list
func (interactor InteractorImpl) AddToDetailList(req dto.ReqTargetDetailAdd) dto.ResTargetDetailAdd {
	resp := dto.ResTargetDetailAdd{}
	if !bson.IsObjectIdHex(req.TargetID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	err := interactor.dao.AddToDetailList(bson.ObjectIdHex(req.TargetID), req.Detail)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Inserted = 1
	resp.Info.SetStatusOK("Ok")
	return resp
}

// GetDetailList update an entire details list
func (interactor InteractorImpl) GetDetailList(req dto.ReqTargetDetailList) dto.ResTargetDetailList {
	resp := dto.ResTargetDetailList{}
	if !bson.IsObjectIdHex(req.TargetID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	details, err := interactor.dao.GetDetailList(bson.ObjectIdHex(req.TargetID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
	}
	resp.Details = details
	resp.Total = len(details)
	resp.Info.SetStatusOK("Ok")
	return resp
}
