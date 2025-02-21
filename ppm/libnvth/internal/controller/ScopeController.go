package controller

import (
	"net/http"

	"github.com/go-chi/chi"

	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/scope"
	"ppm/libnvth/internal/scope/dao"
	"ppm/libnvth/internal/scope/dto"
	"github.com/go-chi/render"
)

// ScopeController controller
type ScopeController struct {
	interactor scope.Interactor
}

// NewScopeController initial
func NewScopeController(session *database.DBSession) ScopeController {
	dao := dao.NewScopeDAOImpl(session)
	interactor := scope.NewInteractorImpl(dao)
	return ScopeController{interactor}
}

// List list
// func (controller ScopeController) List(w http.ResponseWriter, r *http.Request) {
// 	req := dto.ReqScopeList{}
// 	resp := controller.interactor.List(req)
// 	render.Status(r, resp.Info.StatusCode)
// 	render.JSON(w, r, resp)
// }

// Get get
func (controller ScopeController) Get(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "scopeID")
	req := dto.ReqScopeGet{ID: ID}
	resp := controller.interactor.Get(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Create create
func (controller ScopeController) Create(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqScopeCreate{}
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
func (controller ScopeController) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "scopeID")
	req := dto.ReqScopeUpdate{}
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
func (controller ScopeController) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "scopeID")
	req := dto.ReqScopeDelete{ID: ID}
	resp := controller.interactor.Delete(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}
