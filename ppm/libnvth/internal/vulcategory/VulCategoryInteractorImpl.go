package vulcategory

import (
	"ppm/libnvth/internal/vulcategory/dao"
	"ppm/libnvth/internal/vulcategory/dto"
	"gopkg.in/mgo.v2/bson"
)

// InteractorImpl implements Interactor
type InteractorImpl struct {
	dao dao.VulCategoryDAO
}

// NewInteractorImpl initialize a new InteractorImpl object
func NewInteractorImpl(dao dao.VulCategoryDAO) InteractorImpl {
	return InteractorImpl{dao}
}

var _ Interactor = (*InteractorImpl)(nil)

// List get a list of vulnerability categories
func (interactor InteractorImpl) List(req dto.ReqVulCategoryList) dto.ResVulCategoryList {
	resp := dto.ResVulCategoryList{}
	categories, err := interactor.dao.List()
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Categories = categories
	resp.Total = len(categories)
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Get find a vulnerability category by ID
func (interactor InteractorImpl) Get(req dto.ReqVulCategoryGet) dto.ResVulCategoryGet {
	resp := dto.ResVulCategoryGet{}
	category, err := interactor.dao.Get(req.ID)
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	resp.Category = category
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Create create a new vulnerability category
func (interactor InteractorImpl) Create(req dto.ReqVulCategoryCreate) dto.ResVulCategoryCreate {
	resp := dto.ResVulCategoryCreate{}
	param := dao.VulCategoryDAOParam{
		ID:   req.ID,
		Name: req.Name,
	}
	err := interactor.dao.Insert(param)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Info.SetStatusOK("Created")
	return resp
}

// Update update an existing vulnerability category
func (interactor InteractorImpl) Update(req dto.ReqVulCategoryUpdate) dto.ResVulCategoryUpdate {
	resp := dto.ResVulCategoryUpdate{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	param := dao.VulCategoryDAOParam{
		ID:   req.ID,
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

// Delete remove an vulnerability category
func (interactor InteractorImpl) Delete(req dto.ReqVulCategoryDelete) dto.ResVulCategoryDelete {
	resp := dto.ResVulCategoryDelete{}
	category, err := interactor.dao.Get(req.ID)
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	if len(category.Templates) > 0 {
		resp.Info.SetStatusUnprocessableEntity("Cannot delete because this category is associated with at least one vulnerability template")
		return resp
	}
	err = interactor.dao.Remove(req.ID)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Info.SetStatusOK("Deleted")
	return resp
}
