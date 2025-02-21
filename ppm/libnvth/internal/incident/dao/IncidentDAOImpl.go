package dao

import (
	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/incident/bean"
	"gopkg.in/mgo.v2/bson"
)

// IncidentDAOImpl implements IncidentDAO
type IncidentDAOImpl struct {
	session *database.DBSession
}

var _ IncidentDAO = (*IncidentDAOImpl)(nil)

// NewIncidentDAOImpl return a new IncidentDAOImpl instance
func NewIncidentDAOImpl(session *database.DBSession) IncidentDAOImpl {
	return IncidentDAOImpl{session}
}

// List get all
func (dao IncidentDAOImpl) List(param IncidentListDAOParam) ([]bean.IncidentListBean, error) {
	collection := dao.session.Collection("Incidents")
	defer collection.Close()
	selector := bson.M{}
	if param.ProjectID != nil {
		selector["project_id"] = param.ProjectID
	}
	incidents := []bean.IncidentListBean{}
	err := collection.Find(selector).All(&incidents)
	return incidents, err
}

// Get find by id
func (dao IncidentDAOImpl) Get(ID bson.ObjectId) (bean.IncidentBean, error) {
	collection := dao.session.Collection("Incidents")
	defer collection.Close()
	incident := bean.IncidentBean{}
	pipeline := []bson.M{
		{"$match": bson.M{"_id": ID}},
		{"$lookup": bson.M{
			"from":         "Projects",
			"localField":   "project_id",
			"foreignField": "_id",
			"as":           "project",
		}},
		{"$unwind": "$project"},
		{"$lookup": bson.M{
			"from":         "Customers",
			"localField":   "project.customer_id",
			"foreignField": "_id",
			"as":           "customer",
		}},
		{"$unwind": "$customer"},
		{"$lookup": bson.M{
			"from":         "Users",
			"localField":   "responsibler_id",
			"foreignField": "_id",
			"as":           "responsibler",
		}},
		{"$unwind": "$responsibler"},
	}
	err := collection.Pipe(pipeline).One(&incident)
	return incident, err
}

// Insert persist a new object
func (dao IncidentDAOImpl) Insert(param IncidentDAOParam) error {
	collection := dao.session.Collection("Incidents")
	defer collection.Close()
	return collection.Insert(param)
}

// Update change details
func (dao IncidentDAOImpl) Update(param IncidentDAOParam) error {
	collection := dao.session.Collection("Incidents")
	defer collection.Close()
	return collection.UpdateId(param.ID, bson.M{"$set": bson.M{
		"name":            param.Name,
		"description":     param.Description,
		"last_update":     param.LastUpdate,
		"solution":        param.Solution,
		"severity":        param.Severity,
		"status":          param.Status,
		"responsibler_id": param.ResponsiblerID,
		"time":            param.Time,
		"type":            param.Type,
	}})
}

// Remove remove object by id
func (dao IncidentDAOImpl) Remove(ID bson.ObjectId) error {
	collection := dao.session.Collection("Incidents")
	defer collection.Close()
	return collection.RemoveId(ID)
}
