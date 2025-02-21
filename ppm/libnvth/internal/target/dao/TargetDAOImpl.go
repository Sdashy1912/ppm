package dao

import (
	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/target/bean"
	"gopkg.in/mgo.v2/bson"
)

// TargetDAOImpl implements TargetDAO
type TargetDAOImpl struct {
	session *database.DBSession
}

var _ TargetDAO = (*TargetDAOImpl)(nil)

// NewTargetDAOImpl return a new TargetDAOImpl instance
func NewTargetDAOImpl(session *database.DBSession) TargetDAOImpl {
	return TargetDAOImpl{session}
}

// Get find by id
func (dao TargetDAOImpl) Get(ID bson.ObjectId) (bean.TargetBean, error) {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	target := bean.TargetBean{}
	pipeline := []bson.M{
		{"$match": bson.M{"_id": ID}},
		{"$lookup": bson.M{
			"from":         "Users",
			"localField":   "assignee_id",
			"foreignField": "_id",
			"as":           "assignee",
		}},
		{"$unwind": "$assignee"},
		{"$lookup": bson.M{
			"from":         "Scopes",
			"localField":   "scope_id",
			"foreignField": "_id",
			"as":           "scope",
		}},
		{"$unwind": "$scope"},
		{"$lookup": bson.M{
			"from":         "Projects",
			"localField":   "scope.project_id",
			"foreignField": "_id",
			"as":           "project",
		}},
		{"$unwind": "$project"},
		{"$addFields": bson.M{"requests": bson.M{"$size": bson.M{"$ifNull": []interface{}{"$details", []interface{}{}}}}}},
	}
	err := collection.Pipe(pipeline).One(&target)
	return target, err
}

// Insert persist a new object
func (dao TargetDAOImpl) Insert(param TargetDAOParam) error {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	return collection.Insert(param)
}

// Update change details
func (dao TargetDAOImpl) Update(param TargetDAOParam) error {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	return collection.UpdateId(param.ID, bson.M{"$set": bson.M{
		"name":        param.Name,
		"assignee_id": param.AssigneeID,
		"platform":    param.Platform,
		"description": param.Description,
	}})
}

// Remove remove object by id
func (dao TargetDAOImpl) Remove(ID bson.ObjectId) error {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	return collection.RemoveId(ID)
}

// GetDetailList get all the detail list
func (dao TargetDAOImpl) GetDetailList(ID bson.ObjectId) ([]bean.TargetDetailBean, error) {
	collection := dao.session.Collection("Targets")
	details := []bean.TargetDetailBean{}
	pipeline := []bson.M{
		{"$match": bson.M{"_id": ID}},
		{"$unwind": "$details"},
		{"$replaceRoot": bson.M{"newRoot": "$details"}},
	}
	err := collection.Pipe(pipeline).All(&details)
	return details, err
}

// UpdateDetailList update entire the detail list
func (dao TargetDAOImpl) UpdateDetailList(ID bson.ObjectId, details []bean.TargetDetailBean) error {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	return collection.UpdateId(ID, bson.M{
		"$set": bson.M{"details": details},
	})
}

// AddToDetailList add a single pftarget to the list
func (dao TargetDAOImpl) AddToDetailList(ID bson.ObjectId, detail bean.TargetDetailBean) error {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	return collection.Update(
		bson.M{"_id": ID},
		bson.M{"$push": bson.M{
			"details": detail,
		}},
	)
}
