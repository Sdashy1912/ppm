package controller

import (
	"net/http"

	"github.com/go-chi/chi"

	"ppm/libnvth/internal/customer"
	"ppm/libnvth/internal/customer/dao"
	"ppm/libnvth/internal/customer/dto"
	"ppm/libnvth/internal/database"
	"github.com/go-chi/render"
)

// CustomerController controller
type CustomerController struct {
	interactor customer.Interactor
}

// NewCustomerController initial
func NewCustomerController(session *database.DBSession) CustomerController {
	dao := dao.NewCustomerDAOImpl(session)
	interactor := customer.NewCustomerInteractorImpl(dao)
	return CustomerController{interactor}
}

// List list
func (controller CustomerController) List(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqCustomerList{}
	resp := controller.interactor.ListCustomer(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Get get
func (controller CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "customerID")
	req := dto.ReqCustomerGet{ID: ID}
	resp := controller.interactor.GetCustomer(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Create create
func (controller CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqCustomerCreate{}
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	resp := controller.interactor.CreateCustomer(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Update update
func (controller CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "customerID")
	req := dto.ReqCustomerUpdate{}
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	req.ID = ID
	resp := controller.interactor.UpdateCustomer(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// Delete delete
func (controller CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "customerID")
	req := dto.ReqCustomerDelete{ID: ID}
	resp := controller.interactor.DeleteCustomer(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}

// IndustryStats stats
func (controller CustomerController) IndustryStats(w http.ResponseWriter, r *http.Request) {
	req := dto.ReqIndustryStats{}
	resp := controller.interactor.IndustryStats(req)
	render.Status(r, resp.Info.StatusCode)
	render.JSON(w, r, resp)
}
