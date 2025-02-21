package controller

import (
	"net/http"

	"github.com/go-chi/chi"

	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/project"
	"ppm/libnvth/internal/project/dao"
	"ppm/libnvth/internal/project/dto"
	"github.com/go-chi/render"
)

// ProjectController controller
type ProjectController struct {
	interactor project.Interactor
}

// NewProjectController initial
func NewProjectController(session *database.DBSession) ProjectController {
	dao := dao.NewProjectDAOImpl(session)
	interactor := project.NewInteractorImpl(dao)
	return ProjectController{interactor}
}

// List list
func (controller ProjectController) List(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqProjectList{}
	resp := controller.interactor.List(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Get get
func (controller ProjectController) Get(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "projectID")
	req := dto.ReqProjectGet{ID: ID}
	resp := controller.interactor.Get(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Create create
func (controller ProjectController) Create(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqProjectCreate{}
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
func (controller ProjectController) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "projectID")
	req := dto.ReqProjectUpdate{}
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
func (controller ProjectController) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "projectID")
	req := dto.ReqProjectDelete{ID: ID}
	resp := controller.interactor.Delete(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}
