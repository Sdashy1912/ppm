package dao

import (
	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/vultemplate/bean"
	"gopkg.in/mgo.v2/bson"
)

// VulTemplateDAOImpl implements VulTemplateDAO
type VulTemplateDAOImpl struct {
	session *database.DBSession
}

var _ VulTemplateDAO = (*VulTemplateDAOImpl)(nil)

// NewVulTemplateDAOImpl return a new VulTemplateDAOImpl instance
func NewVulTemplateDAOImpl(session *database.DBSession) VulTemplateDAOImpl {
	return VulTemplateDAOImpl{session}
}

// List select all
func (dao VulTemplateDAOImpl) List(param VulTemplateListDAOParam) ([]bean.VulTemplateBean, error) {
	collection := dao.session.Collection("VulTemplates")
	defer collection.Close()
	selector := bson.M{}
	if param.CategoryID != "" {
		selector["category_id"] = param.CategoryID
	}
	query := collection.Find(selector)
	if param.OrderBy != "" {
		query.Sort(param.OrderBy)
	}
	templates := []bean.VulTemplateBean{}
	err := query.All(&templates)
	return templates, err
}

// Get find by id
func (dao VulTemplateDAOImpl) Get(ID string) (bean.VulTemplateBean, error) {
	collection := dao.session.Collection("VulTemplates")
	defer collection.Close()
	template := bean.VulTemplateBean{}
	err := collection.FindId(ID).One(&template)
	return template, err
}

// Insert persist a new object
func (dao VulTemplateDAOImpl) Insert(param VulTemplateDAOParam) error {
	collection := dao.session.Collection("VulTemplates")
	defer collection.Close()
	return collection.Insert(param)
}

// Update change details
func (dao VulTemplateDAOImpl) Update(param VulTemplateDAOParam) error {
	collection := dao.session.Collection("VulTemplates")
	defer collection.Close()
	return collection.UpdateId(param.ID, param)
}

// Remove remove object by id
func (dao VulTemplateDAOImpl) Remove(ID string) error {
	collection := dao.session.Collection("VulTemplates")
	defer collection.Close()
	return collection.RemoveId(ID)
}

// RemoveAll remove object by id
func (dao VulTemplateDAOImpl) RemoveAll() error {
	collection := dao.session.Collection("VulTemplates")
	defer collection.Close()
	return collection.DropCollection()
}
