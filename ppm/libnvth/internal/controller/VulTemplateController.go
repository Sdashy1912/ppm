package controller

import (
	"net/http"

	"github.com/go-chi/chi"

	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/vultemplate"
	"ppm/libnvth/internal/vultemplate/dao"
	"ppm/libnvth/internal/vultemplate/dto"
	"github.com/go-chi/render"
)

// VulTemplateController controller
type VulTemplateController struct {
	interactor vultemplate.Interactor
}

// NewVulTemplateController initial
func NewVulTemplateController(session *database.DBSession) VulTemplateController {
	dao := dao.NewVulTemplateDAOImpl(session)
	interactor := vultemplate.NewInteractorImpl(dao)
	return VulTemplateController{interactor}
}

// List list
func (controller VulTemplateController) List(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	orderBy := r.URL.Query().Get("orderBy")
	req := dto.ReqVulTemplateList{CategoryID: category, OrderBy: orderBy}
	resp := controller.interactor.List(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Get get
func (controller VulTemplateController) Get(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "vultemplateID")
	req := dto.ReqVulTemplateGet{ID: ID}
	resp := controller.interactor.Get(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Create create
func (controller VulTemplateController) Create(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqVulTemplateCreate{}
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
func (controller VulTemplateController) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "vultemplateID")
	req := dto.ReqVulTemplateUpdate{}
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
func (controller VulTemplateController) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "vultemplateID")
	req := dto.ReqVulTemplateDelete{ID: ID}
	resp := controller.interactor.Delete(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}
