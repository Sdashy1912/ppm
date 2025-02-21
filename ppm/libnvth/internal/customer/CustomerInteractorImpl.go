package customer

import (
	"log"

	"ppm/libnvth/internal/customer/dao"
	"ppm/libnvth/internal/customer/dao/param"
	"ppm/libnvth/internal/customer/dto"
	"gopkg.in/mgo.v2/bson"
)

// InteractorImpl implements CustomerCreate usecase interface
type InteractorImpl struct {
	dao dao.CustomerDAO
}

// NewCustomerInteractorImpl return a CustomerInteractorImpl which implements CustomerInteractor
func NewCustomerInteractorImpl(dao dao.CustomerDAO) InteractorImpl {
	return InteractorImpl{dao}
}

var _ Interactor = (*InteractorImpl)(nil)

// CreateCustomer persist a new user into database
func (interactor InteractorImpl) CreateCustomer(req dto.ReqCustomerCreate) dto.ResCustomerCreate {
	resp := dto.ResCustomerCreate{}
	newCustomer := param.CustomerCreateDAOParam{
		ID:          bson.NewObjectId(),
		Name:        req.Name,
		Email:       req.Email,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Industry:    req.Industry,
	}
	err := interactor.dao.Insert(newCustomer)
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Info.SetStatusOK("Created")
	resp.InsertedID = newCustomer.ID.Hex()
	return resp
}

// ListCustomer list
func (interactor InteractorImpl) ListCustomer(req dto.ReqCustomerList) dto.ResCustomerList {
	resp := dto.ResCustomerList{}
	customers, err := interactor.dao.List()
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resCustomers := []dto.ResCustomerListBean{}
	for _, c := range customers {
		resCustomers = append(resCustomers, dto.ResCustomerListBean{
			ID:          c.ID.Hex(),
			Name:        c.Name,
			Email:       c.Email,
			Address:     c.Address,
			PhoneNumber: c.PhoneNumber,
			Industry:    c.Industry,
		})
	}
	resp.Info.SetStatusOK("Ok")
	resp.Total = len(resCustomers)
	resp.Customers = resCustomers
	return resp
}

// GetCustomer get customer by ID
func (interactor InteractorImpl) GetCustomer(req dto.ReqCustomerGet) dto.ResCustomerGet {
	resp := dto.ResCustomerGet{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	c, err := interactor.dao.Get(bson.ObjectIdHex(req.ID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	customer := dto.ResCustomerGetBean{
		ID:          c.ID.Hex(),
		Name:        c.Name,
		Email:       c.Email,
		Address:     c.Address,
		PhoneNumber: c.PhoneNumber,
		Industry:    c.Industry,
	}
	projects := []dto.ResCustomerProjectBean{}
	for _, p := range c.Projects {
		projects = append(projects, dto.ResCustomerProjectBean{
			ID:        p.ID.Hex(),
			Name:      p.Name,
			StartTime: p.StartTime,
			EndTime:   p.EndTime,
			Status:    p.Status,
		})
	}
	customer.Projects = projects
	resp.Customer = customer
	resp.Info.SetStatusOK("Ok")
	return resp
}

// UpdateCustomer update
func (interactor InteractorImpl) UpdateCustomer(req dto.ReqCustomerUpdate) dto.ResCustomerUpdate {
	resp := dto.ResCustomerUpdate{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	param := param.CustomerUpdateDAOParam{
		ID:          bson.ObjectIdHex(req.ID),
		Name:        req.Name,
		Email:       req.Email,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Industry:    req.Industry,
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

// DeleteCustomer delete
func (interactor InteractorImpl) DeleteCustomer(req dto.ReqCustomerDelete) dto.ResCustomerDelete {
	resp := dto.ResCustomerDelete{}
	if !bson.IsObjectIdHex(req.ID) {
		resp.Info.SetStatusNotFound()
		return resp
	}
	customer, err := interactor.dao.Get(bson.ObjectIdHex(req.ID))
	if err != nil {
		if err.Error() == "not found" {
			resp.Info.SetStatusNotFound()
		} else {
			resp.Info.SetStatusInternalServerError()
		}
		return resp
	}
	if len(customer.Projects) > 0 {
		resp.Info.SetStatusUnprocessableEntity("Cannot delete because this customer is associated with at least one project")
		return resp
	}
	err = interactor.dao.Delete(bson.ObjectIdHex(req.ID))
	if err != nil {
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Info.SetStatusOK("Deleted")
	return resp
}

// IndustryStats stats
func (interactor InteractorImpl) IndustryStats(req dto.ReqIndustryStats) dto.ResIndustryStats {
	resp := dto.ResIndustryStats{}
	data, err := interactor.dao.IndustryStats()
	if err != nil {
		log.Fatalf("WTF: %s", err)
		resp.Info.SetStatusInternalServerError()
		return resp
	}
	resp.Data = data
	resp.Info.SetStatusOK("OK")
	return resp
}
