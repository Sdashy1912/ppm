package customer

import "ppm/libnvth/internal/customer/dto"

// Interactor interface for customer usecases
type Interactor interface {
	ListCustomer(req dto.ReqCustomerList) dto.ResCustomerList
	GetCustomer(req dto.ReqCustomerGet) dto.ResCustomerGet
	CreateCustomer(req dto.ReqCustomerCreate) dto.ResCustomerCreate
	UpdateCustomer(req dto.ReqCustomerUpdate) dto.ResCustomerUpdate
	DeleteCustomer(req dto.ReqCustomerDelete) dto.ResCustomerDelete
	IndustryStats(req dto.ReqIndustryStats) dto.ResIndustryStats
}
