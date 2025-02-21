package controller

import (
	"net/http"

	"github.com/go-chi/chi"

	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/user"
	"ppm/libnvth/internal/user/dao"
	"ppm/libnvth/internal/user/dto"
	"github.com/go-chi/render"
)

// UserController controller
type UserController struct {
	interactor user.Interactor
}

// NewUserController initial
func NewUserController(session *database.DBSession) UserController {
	dao := dao.NewUserDAOImpl(session)
	interactor := user.NewInteractorImpl(dao)
	return UserController{interactor}
}

// List list
func (controller UserController) List(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqUserList{}
	resp := controller.interactor.List(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Get get
func (controller UserController) Get(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "userID")
	req := dto.ReqUserGet{ID: ID}
	resp := controller.interactor.Get(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Create create
func (controller UserController) Create(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqUserCreate{}
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
func (controller UserController) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "userID")
	req := dto.ReqUserUpdate{}
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
func (controller UserController) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "userID")
	req := dto.ReqUserDelete{ID: ID}
	resp := controller.interactor.Delete(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}
