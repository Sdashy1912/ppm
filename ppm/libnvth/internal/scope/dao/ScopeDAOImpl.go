package dao

import (
	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/scope/bean"
	"gopkg.in/mgo.v2/bson"
)

// ScopeDAOImpl implements ScopeDAO
type ScopeDAOImpl struct {
	session *database.DBSession
}

var _ ScopeDAO = (*ScopeDAOImpl)(nil)

// NewScopeDAOImpl return a new ScopeDAOImpl instance
func NewScopeDAOImpl(session *database.DBSession) ScopeDAOImpl {
	return ScopeDAOImpl{session}
}

// Get find by id
func (dao ScopeDAOImpl) Get(ID bson.ObjectId) (bean.ScopeBean, error) {
	collection := dao.session.Collection("Scopes")
	defer collection.Close()
	scope := bean.ScopeBean{}
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
			"from": "Targets",
			"let":  bson.M{"scope_id": "$_id"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$eq": []string{"$scope_id", "$$scope_id"}}}},
				{"$lookup": bson.M{
					"from":         "Users",
					"localField":   "assignee_id",
					"foreignField": "_id",
					"as":           "assignee"},
				},
				{"$unwind": "$assignee"},
			},
			"as": "targets",
		}},
	}
	err := collection.Pipe(pipeline).One(&scope)
	return scope, err
}

// Insert persist a new object
func (dao ScopeDAOImpl) Insert(param ScopeDAOParam) error {
	collection := dao.session.Collection("Scopes")
	defer collection.Close()
	return collection.Insert(param)
}

// Update change details
func (dao ScopeDAOImpl) Update(param ScopeDAOParam) error {
	collection := dao.session.Collection("Scopes")
	defer collection.Close()
	return collection.UpdateId(param.ID, bson.M{"$set": bson.M{
		"name": param.Name,
	}})
}

// Remove remove object by id
func (dao ScopeDAOImpl) Remove(ID bson.ObjectId) error {
	collection := dao.session.Collection("Scopes")
	defer collection.Close()
	return collection.RemoveId(ID)
}

// IsSafeToDelete check before delete an scope
func (dao ScopeDAOImpl) IsSafeToDelete(ID bson.ObjectId) (bool, error) {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	result := []bson.M{}
	err := collection.Find(bson.M{"scope_id": ID}).All(&result)
	if err != nil {
		return false, err
	}
	return len(result) == 0, nil
}
