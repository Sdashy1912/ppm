package controller

import (
	"net/http"

	"ppm/libnvth/internal/target/bean"

	"github.com/go-chi/chi"

	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/target"
	"ppm/libnvth/internal/target/dao"
	"ppm/libnvth/internal/target/dto"
	"github.com/go-chi/render"
)

// TargetController controller
type TargetController struct {
	interactor target.Interactor
}

// NewTargetController initial
func NewTargetController(session *database.DBSession) TargetController {
	dao := dao.NewTargetDAOImpl(session)
	interactor := target.NewInteractorImpl(dao)
	return TargetController{interactor}
}

// List list
// func (controller TargetController) List(w http.ResponseWriter, r *http.Request) {
// 	req := dto.ReqTargetList{}
// 	resp := controller.interactor.List(req)
// 	render.Status(r, resp.Info.StatusCode)
// 	render.JSON(w, r, resp)
// }

// Get get
func (controller TargetController) Get(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "targetID")
	req := dto.ReqTargetGet{ID: ID}
	resp := controller.interactor.Get(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Create create
func (controller TargetController) Create(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqTargetCreate{}
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
func (controller TargetController) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "targetID")
	req := dto.ReqTargetUpdate{}
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
func (controller TargetController) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "targetID")
	req := dto.ReqTargetDelete{ID: ID}
	resp := controller.interactor.Delete(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// GetDetailList get details list
func (controller TargetController) GetDetailList(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "targetID")
	req := dto.ReqTargetDetailList{TargetID: ID}
	resp := controller.interactor.GetDetailList(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// AddToDetailList add a detail to list
func (controller TargetController) AddToDetailList(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "targetID")
	detail := bean.TargetDetailBean{}
	if err := render.DecodeJSON(r.Body, &detail); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	req := dto.ReqTargetDetailAdd{TargetID: ID, Detail: detail}
	resp := controller.interactor.AddToDetailList(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// UpdateDetailList update details list
func (controller TargetController) UpdateDetailList(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "targetID")
	req := dto.ReqTargetDetailUpdate{}
	details := []bean.TargetDetailBean{}
	if err := render.DecodeJSON(r.Body, &details); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	req.TargetID = ID
	req.Details = details
	resp := controller.interactor.UpdateDetailList(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}
