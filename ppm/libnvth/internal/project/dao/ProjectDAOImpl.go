package dao

import (
	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/project/bean"
	"gopkg.in/mgo.v2/bson"
)

// ProjectDAOImpl implements ProjectDAO
type ProjectDAOImpl struct {
	session *database.DBSession
}

var _ ProjectDAO = (*ProjectDAOImpl)(nil)

// NewProjectDAOImpl return a new ProjectDAOImpl instance
func NewProjectDAOImpl(session *database.DBSession) ProjectDAOImpl {
	return ProjectDAOImpl{session}
}

// List select all
func (dao ProjectDAOImpl) List() ([]bean.ProjectBean, error) {
	collection := dao.session.Collection("Projects")
	defer collection.Close()
	pipeline := []bson.M{
		{"$lookup": bson.M{
			"from":         "Customers",
			"localField":   "customer_id",
			"foreignField": "_id",
			"as":           "customer",
		}},
		{"$unwind": "$customer"},
	}
	projects := []bean.ProjectBean{}
	err := collection.Pipe(pipeline).All(&projects)
	return projects, err
}

// Get find by id
func (dao ProjectDAOImpl) Get(ID bson.ObjectId) (bean.ProjectBean, error) {
	collection := dao.session.Collection("Projects")
	defer collection.Close()
	project := bean.ProjectBean{}
	pipeline := []bson.M{
		{"$match": bson.M{"_id": ID}},
		{"$lookup": bson.M{
			"from":         "Customers",
			"localField":   "customer_id",
			"foreignField": "_id",
			"as":           "customer",
		}},
		{"$unwind": "$customer"},
		{"$lookup": bson.M{
			"from":         "Scopes",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "scopes",
		}},
		{"$lookup": bson.M{
			"from":         "Incidents",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "incidents",
		}},
	}
	err := collection.Pipe(pipeline).One(&project)
	return project, err
}

// Insert persist a new object
func (dao ProjectDAOImpl) Insert(param ProjectDAOParam) error {
	collection := dao.session.Collection("Projects")
	defer collection.Close()
	return collection.Insert(param)
}

// Update change details
func (dao ProjectDAOImpl) Update(param ProjectDAOParam) error {
	collection := dao.session.Collection("Projects")
	defer collection.Close()
	return collection.UpdateId(param.ID, param)
}

// Remove remove object by id
func (dao ProjectDAOImpl) Remove(ID bson.ObjectId) error {
	collection := dao.session.Collection("Projects")
	defer collection.Close()
	return collection.RemoveId(ID)
}

// IsSafeToDelete check before delete an project
func (dao ProjectDAOImpl) IsSafeToDelete(ID bson.ObjectId) (bool, error) {
	collection := dao.session.Collection("Scopes")
	defer collection.Close()
	result := []bson.M{}
	err := collection.Find(bson.M{"project_id": ID}).All(&result)
	if err != nil {
		return false, err
	}
	if len(result) != 0 {
		return false, nil
	}
	incidentCol := dao.session.Collection("Incidents")
	result = []bson.M{}
	err = incidentCol.Find(bson.M{"project_id": ID}).All(&result)
	if err != nil {
		return false, err
	}
	return len(result) == 0, nil
}
