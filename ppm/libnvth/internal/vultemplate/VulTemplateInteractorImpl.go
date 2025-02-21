package vultemplate

import (
	"ppm/libnvth/internal/vultemplate/dao"
	"ppm/libnvth/internal/vultemplate/dto"
)

// InteractorImpl implements Interactor
type InteractorImpl struct {
	dao dao.VulTemplateDAO
}

// NewInteractorImpl initialize a new InteractorImpl object
func NewInteractorImpl(dao dao.VulTemplateDAO) InteractorImpl {
	return InteractorImpl{dao}
}

var _ Interactor = (*InteractorImpl)(nil)

// List get a list of vulnerability templates
func (interactor InteractorImpl) List(req dto.ReqVulTemplateList) dto.ResVulTemplateList {
	resp := dto.ResVulTemplateList{}
	param := dao.VulTemplateListDAOParam{
		CategoryID: req.CategoryID,
		OrderBy:    req.OrderBy,
	}
	templates, err := interactor.dao.List(param)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Templates = templates
	resp.Total = len(templates)
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Get find a vulnerability template by ID
func (interactor InteractorImpl) Get(req dto.ReqVulTemplateGet) dto.ResVulTemplateGet {
	resp := dto.ResVulTemplateGet{}
	template, err := interactor.dao.Get(req.ID)
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	resp.Template = template
	resp.Info.SetStatusOK("Ok")
	return resp
}

// Create create a new vulnerability template
func (interactor InteractorImpl) Create(req dto.ReqVulTemplateCreate) dto.ResVulTemplateCreate {
	resp := dto.ResVulTemplateCreate{}
	param := dao.VulTemplateDAOParam{
		ID:              req.ID,
		CategoryID:      req.CategoryID,
		Rating:          req.Rating,
		Name:            req.Name,
		IdentifierName:  req.IdentifierName,
		Overview:        req.Overview,
		DetectionMethod: req.DetectionMethod,
		Description:     req.Description,
		Condition:       req.Condition,
		PossibleImpact:  req.PossibleImpact,
		Countermeasure:  req.Countermeasure,
		Remarks:         req.Remarks,
	}
	err := interactor.dao.Insert(param)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Info.SetStatusOK("Created")
	return resp
}

// Update update an existing vulnerability template
func (interactor InteractorImpl) Update(req dto.ReqVulTemplateUpdate) dto.ResVulTemplateUpdate {
	resp := dto.ResVulTemplateUpdate{}
	param := dao.VulTemplateDAOParam{
		ID:              req.ID,
		CategoryID:      req.CategoryID,
		Rating:          req.Rating,
		Name:            req.Name,
		IdentifierName:  req.IdentifierName,
		Overview:        req.Overview,
		DetectionMethod: req.DetectionMethod,
		Description:     req.Description,
		Condition:       req.Condition,
		PossibleImpact:  req.PossibleImpact,
		Countermeasure:  req.Countermeasure,
		Remarks:         req.Remarks,
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

// Delete remove an vulnerability template
func (interactor InteractorImpl) Delete(req dto.ReqVulTemplateDelete) dto.ResVulTemplateDelete {
	resp := dto.ResVulTemplateDelete{}
	_, err := interactor.dao.Get(req.ID)
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
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
