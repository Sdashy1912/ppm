package dao

import (
	"ppm/libnvth/internal/customer/dao/bean"
	"ppm/libnvth/internal/customer/dao/param"
	"ppm/libnvth/internal/customer/dto"
	"gopkg.in/mgo.v2/bson"
)

// CustomerDAO data access object for create a new customer usecase
type CustomerDAO interface {
	List() ([]bean.CustomerListBean, error)
	Get(ID bson.ObjectId) (bean.CustomerGetBean, error)
	Insert(payload param.CustomerCreateDAOParam) error
	Update(payload param.CustomerUpdateDAOParam) error
	Delete(ID bson.ObjectId) error
	IndustryStats() ([]dto.ResIndustryStatsData, error)
}
