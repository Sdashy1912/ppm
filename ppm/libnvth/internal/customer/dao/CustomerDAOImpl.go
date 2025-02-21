package dao

import (
	"ppm/libnvth/internal/customer/dao/bean"
	"ppm/libnvth/internal/customer/dao/param"
	"ppm/libnvth/internal/customer/dto"
	"ppm/libnvth/internal/database"
	"gopkg.in/mgo.v2/bson"
)

// CustomerDAOImpl implements CustomerCreateDAO
type CustomerDAOImpl struct {
	session *database.DBSession
}

// NewCustomerDAOImpl return an instance of CustomerDAOImpl which implements CustomerDAO
func NewCustomerDAOImpl(session *database.DBSession) CustomerDAOImpl {
	return CustomerDAOImpl{session}
}

var _ CustomerDAO = (*CustomerDAOImpl)(nil)

// Insert persist a new customer into database
func (dao CustomerDAOImpl) Insert(payload param.CustomerCreateDAOParam) error {
	collection := dao.session.Collection("Customers")
	defer collection.Close()
	return collection.Insert(payload)
}

// List get all customers
func (dao CustomerDAOImpl) List() ([]bean.CustomerListBean, error) {
	collection := dao.session.Collection("Customers")
	defer collection.Close()
	customers := []bean.CustomerListBean{}
	err := collection.Find(nil).All(&customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// Get find a customer by id
func (dao CustomerDAOImpl) Get(ID bson.ObjectId) (bean.CustomerGetBean, error) {
	collection := dao.session.Collection("Customers")
	defer collection.Close()
	customer := bean.CustomerGetBean{}
	pipeline := []bson.M{
		{"$match": bson.M{"_id": ID}},
		{
			"$lookup": bson.M{
				"from":         "Projects",
				"localField":   "_id",
				"foreignField": "customer_id",
				"as":           "projects",
			},
		},
	}
	err := collection.Pipe(pipeline).One(&customer)
	return customer, err
}

// Update update a customer
func (dao CustomerDAOImpl) Update(payload param.CustomerUpdateDAOParam) error {
	collection := dao.session.Collection("Customers")
	defer collection.Close()
	return collection.UpdateId(payload.ID, payload)
}

// Delete a customers
func (dao CustomerDAOImpl) Delete(ID bson.ObjectId) error {
	collection := dao.session.Collection("Customers")
	defer collection.Close()
	return collection.RemoveId(ID)
}

// IndustryStats stats
func (dao CustomerDAOImpl) IndustryStats() ([]dto.ResIndustryStatsData, error) {
	collection := dao.session.Collection("Customers")
	defer collection.Close()
	data := []dto.ResIndustryStatsData{}
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":       "$industry",
				"customers": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":       false,
				"industry":  "$_id",
				"customers": true,
			},
		},
	}
	err := collection.Pipe(pipeline).All(&data)
	return data, err
}
