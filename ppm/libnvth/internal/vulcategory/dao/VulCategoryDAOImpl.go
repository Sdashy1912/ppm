package dao

import (
	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/vulcategory/bean"
	"gopkg.in/mgo.v2/bson"
)

// VulCategoryDAOImpl implements VulCategoryDAO
type VulCategoryDAOImpl struct {
	session *database.DBSession
}

var _ VulCategoryDAO = (*VulCategoryDAOImpl)(nil)

// NewVulCategoryDAOImpl return a new VulCategoryDAOImpl instance
func NewVulCategoryDAOImpl(session *database.DBSession) VulCategoryDAOImpl {
	return VulCategoryDAOImpl{session}
}

// List select all
func (dao VulCategoryDAOImpl) List() ([]bean.VulCategoryListBean, error) {
	collection := dao.session.Collection("VulCategories")
	defer collection.Close()
	categories := []bean.VulCategoryListBean{}
	err := collection.Find(nil).All(&categories)
	return categories, err
}

// Get find by id
func (dao VulCategoryDAOImpl) Get(ID string) (bean.VulCategoryGetBean, error) {
	collection := dao.session.Collection("VulCategories")
	defer collection.Close()
	category := bean.VulCategoryGetBean{}
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"_id": ID}},
		bson.M{"$lookup": bson.M{
			"from":         "VulTemplates",
			"localField":   "_id",
			"foreignField": "customer_id",
			"as":           "templates",
		}},
	}
	err := collection.Pipe(pipeline).One(&category)
	return category, err
}

// Insert persist a new object
func (dao VulCategoryDAOImpl) Insert(param VulCategoryDAOParam) error {
	collection := dao.session.Collection("VulCategories")
	defer collection.Close()
	return collection.Insert(param)
}

// Update change details
func (dao VulCategoryDAOImpl) Update(param VulCategoryDAOParam) error {
	collection := dao.session.Collection("VulCategories")
	defer collection.Close()
	return collection.UpdateId(param.ID, param)
}

// Remove remove object by id
func (dao VulCategoryDAOImpl) Remove(ID string) error {
	collection := dao.session.Collection("VulCategories")
	defer collection.Close()
	return collection.RemoveId(ID)
}

// RemoveAll remove object by id
func (dao VulCategoryDAOImpl) RemoveAll() error {
	collection := dao.session.Collection("VulCategories")
	defer collection.Close()
	return collection.DropCollection()
}
