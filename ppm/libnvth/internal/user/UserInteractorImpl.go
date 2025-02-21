package user

import (
	"ppm/libnvth/internal/user/dao"
	"ppm/libnvth/internal/user/dto"
	"ppm/libnvth/internal/util"
	"gopkg.in/mgo.v2/bson"
)

// InteractorImpl implements Interactor
type InteractorImpl struct {
	dao dao.UserDAO
}

// NewInteractorImpl initialize a new InteractorImpl object
func NewInteractorImpl(dao dao.UserDAO) InteractorImpl {
	return InteractorImpl{dao}
}

var _ Interactor = (*InteractorImpl)(nil)

// List get a list of vulnerability users
func (interactor InteractorImpl) List(req dto.ReqUserList) dto.ResUserList {
	resp := dto.ResUserList{}
	users, err := interactor.dao.List()
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Users = users
	resp.Total = len(users)
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Get find a vulnerability user by ID
func (interactor InteractorImpl) Get(req dto.ReqUserGet) dto.ResUserGet {
	resp := dto.ResUserGet{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	user, err := interactor.dao.Get(bson.ObjectIdHex(req.ID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	resp.User = user
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Create create a new vulnerability user
func (interactor InteractorImpl) Create(req dto.ReqUserCreate) dto.ResUserCreate {
	resp := dto.ResUserCreate{}
	param := dao.UserCreateDAOParam{
		ID:          bson.NewObjectId(),
		Email:       req.Email,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		IsActive:    true,
		IsVerified:  false,
		IsAdmin:     false,
		Password:    util.SHA256([]byte("abc@123")),
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

// Update update an existing vulnerability user
func (interactor InteractorImpl) Update(req dto.ReqUserUpdate) dto.ResUserUpdate {
	resp := dto.ResUserUpdate{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	param := dao.UserUpdateDAOParam{
		ID:          bson.ObjectIdHex(req.ID),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
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

// Delete remove an vulnerability user
func (interactor InteractorImpl) Delete(req dto.ReqUserDelete) dto.ResUserDelete {
	resp := dto.ResUserDelete{}
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
		resp.Info.SetStatusUnprocessableEntity("Cannot delete because this user is associated with at least one targets")
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
