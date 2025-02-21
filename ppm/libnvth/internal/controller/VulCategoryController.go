package controller

import (
	"net/http"

	"github.com/go-chi/chi"

	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/vulcategory"
	"ppm/libnvth/internal/vulcategory/dao"
	"ppm/libnvth/internal/vulcategory/dto"
	"github.com/go-chi/render"
)

// VulCategoryController controller
type VulCategoryController struct {
	interactor vulcategory.Interactor
}

// NewVulCategoryController initial
func NewVulCategoryController(session *database.DBSession) VulCategoryController {
	dao := dao.NewVulCategoryDAOImpl(session)
	interactor := vulcategory.NewInteractorImpl(dao)
	return VulCategoryController{interactor}
}

// List list
func (controller VulCategoryController) List(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqVulCategoryList{}
	resp := controller.interactor.List(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Get get
func (controller VulCategoryController) Get(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "vulcategoryID")
	req := dto.ReqVulCategoryGet{ID: ID}
	resp := controller.interactor.Get(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Create create
func (controller VulCategoryController) Create(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqVulCategoryCreate{}
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	resp := controller.interactor.Create(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Update update
func (controller VulCategoryController) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "vulcategoryID")
	req := dto.ReqVulCategoryUpdate{}
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	req.ID = ID
	resp := controller.interactor.Update(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Delete delete
func (controller VulCategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "vulcategoryID")
	req := dto.ReqVulCategoryDelete{ID: ID}
	resp := controller.interactor.Delete(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}
