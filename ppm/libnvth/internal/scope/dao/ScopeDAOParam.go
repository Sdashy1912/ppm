package dao

import (
	"gopkg.in/mgo.v2/bson"
)

// ScopeDAOParam param for Create, Update
type ScopeDAOParam struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	ProjectID bson.ObjectId `bson:"project_id"`
}
