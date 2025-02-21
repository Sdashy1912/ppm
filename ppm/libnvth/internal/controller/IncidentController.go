package controller

import (
	"net/http"

	"github.com/go-chi/chi"

	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/incident"
	"ppm/libnvth/internal/incident/dao"
	"ppm/libnvth/internal/incident/dto"
	"github.com/go-chi/render"
)

// IncidentController controller
type IncidentController struct {
	interactor incident.Interactor
}

// NewIncidentController initial
func NewIncidentController(session *database.DBSession) IncidentController {
	dao := dao.NewIncidentDAOImpl(session)
	interactor := incident.NewInteractorImpl(dao)
	return IncidentController{interactor}
}

// List list
// func (controller IncidentController) List(w http.ResponseWriter, r *http.Request) {
// 	req := dto.ReqIncidentList{}
// 	resp := controller.interactor.List(req)
// 	render.Status(r, resp.Info.StatusCode)
// 	render.JSON(w, r, resp)
// }

// List list
func (controller IncidentController) List(w http.ResponseWriter, r *http.Request) {
	projectID := r.URL.Query().Get("project")
	req := dto.ReqIncidentList{ProjectID: projectID}
	resp := controller.interactor.List(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Get get
func (controller IncidentController) Get(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "incidentID")
	req := dto.ReqIncidentGet{ID: ID}
	resp := controller.interactor.Get(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Create create
func (controller IncidentController) Create(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqIncidentCreate{}
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
func (controller IncidentController) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "incidentID")
	req := dto.ReqIncidentUpdate{}
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
func (controller IncidentController) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "incidentID")
	req := dto.ReqIncidentDelete{ID: ID}
	resp := controller.interactor.Delete(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}
